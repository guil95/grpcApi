package use_cases

import (
	"fmt"
	"github.com/guil95/grpcApi/core/checkout/domain"
	"log"
	"sync"
)

type CreateCheckoutUseCase struct {
	client     domain.Client
	repository domain.Repository
}

func NewCreateCheckoutUseCase(client domain.Client, repository domain.Repository) *CreateCheckoutUseCase {
	return &CreateCheckoutUseCase{client: client, repository: repository}
}
var order domain.Order

func (s *CreateCheckoutUseCase) Checkout(chart *domain.Chart) (*domain.Order, error){
	var wg sync.WaitGroup

	s.mergeProducts(chart)

	o := order.NewOrder()

	products := s.retrieveProductsByChart(chart)

	if products == nil {
		return nil, domain.ProductNotFoundError
	}

	err := s.verifyGiftProducts(products)

	if err != nil {
		return o, err
	}

	for _, product := range products {
		wg.Add(1)

		go func(product domain.Product, order *domain.Order) {
			product, discountValue := s.applyDiscount(product)
			o.AddProduct(product, discountValue)

			defer wg.Done()
		}(product, o)
	}

	wg.Wait()

	o.CalcTotals()

	return o, nil
}

func (s *CreateCheckoutUseCase) verifyGiftProducts(products []domain.Product) error {
	for _, product := range products {
		if product.Gift == true {
			return domain.ProductGiftError
		}
	}

	return nil
}

func (s *CreateCheckoutUseCase) retrieveProductsByChart(chart *domain.Chart) []domain.Product {
	products := s.repository.GetProductsByChart(chart)

	return products
}

func (s *CreateCheckoutUseCase) applyDiscount(product domain.Product) (domain.Product, int32){
	discount, _ := s.client.GetDiscount(product.Id)

	log.Println(fmt.Sprintf("Discount applied: %f", discount))

	productFloatAmount := float32(product.Amount) / 100

	discountValue := int32((productFloatAmount * discount) * 100)

	return product, discountValue
}

func (s *CreateCheckoutUseCase) mergeProducts(chart *domain.Chart) {
	productExists := make(map[int32]domain.ProductChart)
	var products []domain.ProductChart

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