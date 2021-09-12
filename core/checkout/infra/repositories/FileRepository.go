package repositories

import (
	"encoding/json"
	"github.com/guil95/grpcApi/core/checkout/domain"
)

type FileRepository struct {
	db []byte
}

func NewFileRepository(db []byte) *FileRepository {
	return &FileRepository{db: db}
}

func (fr *FileRepository) GetProducts()[]domain.Product {
	return nil
}

func (fr *FileRepository) GetProductsByChart(chart *domain.Chart)[]domain.Product {
	var products []domain.Product

	_ = json.Unmarshal([]byte(fr.db), &products)

	var productsFilter []domain.Product
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
