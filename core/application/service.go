package application

import (
	"fmt"
	domain2 "github.com/guil95/grpcApi/core/domain"
	"log"
	"sync"
)

type Service struct {
	client     domain2.Client
	repository domain2.Repository
}

func NewService(client domain2.Client, repository domain2.Repository) *Service {
	return &Service{client: client, repository: repository}
}
var order domain2.Order

func (s *Service) Checkout(chart *domain2.Chart) (*domain2.Order, error){
	var wg sync.WaitGroup

	s.mergeProducts(chart)

	o := order.NewOrder()

	products := s.retrieveProductsByChart(chart)

	err := s.verifyGiftProducts(products)

	if err != nil {
		return o, err
	}

	for _, product := range products {
		wg.Add(1)

		go func(product domain2.Product, order *domain2.Order) {
			product, discountValue := s.applyDiscount(product)
			o.AddProduct(product, product.Quantity, discountValue)

			defer wg.Done()
		}(product, o)
	}

	wg.Wait()

	o.CalcTotals()

	return o, nil
}

func (s *Service) verifyGiftProducts(products []domain2.Product) error {
	for _, product := range products {
		if product.Gift == true {
			return domain2.ProductGiftError
		}
	}

	return nil
}

func (s *Service) retrieveProductsByChart(chart *domain2.Chart) []domain2.Product {
	products := s.repository.GetProductsByChart(chart)

	return products
}

func (s *Service) applyDiscount(product domain2.Product) (domain2.Product, int32){
	discount, _ := s.client.GetDiscount(product.Id)

	log.Println(fmt.Sprintf("Discount applied: %f", discount))

	productFloatAmount := float32(product.Amount) / 100

	discountValue := int32((productFloatAmount * discount) * 100)

	return product, discountValue
}

func (s *Service) mergeProducts(chart *domain2.Chart) {
	productExists := make(map[int32]domain2.ProductChart)
	var products []domain2.ProductChart

	for _,item := range chart.Products {
		if _, ok := productExists[item.ProductId]; ok {
			product := productExists[item.ProductId]
			product.Quantity += item.Quantity
			productExists[item.ProductId] = product

			continue
		}

		productExists[item.ProductId] = item
	}

	for _, i := range productExists {
		products = append(products, i)
	}

	chart.Products = products
}