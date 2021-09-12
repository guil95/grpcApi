package repositories

import (
	"encoding/json"
	domain2 "github.com/guil95/grpcApi/core/domain"
)

type FileRepository struct {
	db []byte
}

func NewFileRepository(db []byte) *FileRepository {
	return &FileRepository{db: db}
}

func (fr *FileRepository) GetProducts()[]domain2.Product {
	return nil
}

func (fr *FileRepository) GetProductsByChart(chart *domain2.Chart)[]domain2.Product {
	var products []domain2.Product

	_ = json.Unmarshal([]byte(fr.db), &products)

	var productsFilter []domain2.Product
	productsChart := chart.Products

	for _, product := range products {
		for _, productChart := range productsChart {
			if product.Id == productChart.ProductId {
				product.Quantity = productChart.Quantity
				productsFilter = append(productsFilter, product)
			}
		}
	}

	return productsFilter
}
