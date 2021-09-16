package domain_test

import (
	"github.com/guil95/grpcApi/core/checkout/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder(t *testing.T) {
	t.Run("Test create order and add products", func(t *testing.T) {
		var order domain.Order

		o := order.NewOrder()

		o.AddProduct(&domain.Product{
			Id: 1,
			Amount: 10,
			Quantity: 2,
			Gift: false,
			Description: "Test",
			Title: "Product test",
		},  5)

		o.AddProduct(&domain.Product{
			Id: 2,
			Amount: 10,
			Quantity: 2,
			Gift: false,
			Description: "Test",
			Title: "Product test",
		},  5)

		o.CalcTotals()

		assert.True(t, len(o.Products) == 2)
		assert.True(t, o.TotalAmountWithDiscount == 20)
		assert.True(t, o.TotalAmount == 40)
		assert.True(t, o.TotalDiscount == 20)
	})

	t.Run("Test create order and add gift product its value must be equals 0", func(t *testing.T) {
		var order domain.Order

		o := order.NewOrder()

		o.AddProduct(&domain.Product{
			Id: 1,
			Amount: 10,
			Quantity: 2,
			Gift: false,
			Description: "Test",
			Title: "Product test",
		},  5)

		o.AddProduct(&domain.Product{
			Id: 2,
			Amount: 10,
			Quantity: 2,
			Gift: true,
			Description: "Test",
			Title: "Product test",
		},  5)

		o.CalcTotals()

		assert.True(t, len(o.Products) == 2)
		assert.True(t, o.Products[1].TotalAmount == 0)
		assert.True(t, o.TotalAmountWithDiscount == 10)
		assert.True(t, o.TotalAmount == 20)
		assert.True(t, o.TotalDiscount == 10)
	})

	t.Run("Test create order and add one product and product nil", func(t *testing.T) {
		var order domain.Order

		o := order.NewOrder()

		o.AddProduct(&domain.Product{
			Id: 1,
			Amount: 10,
			Quantity: 2,
			Gift: false,
			Description: "Test",
			Title: "Product test",
		},  5)

		o.AddProduct(nil,  5)

		o.CalcTotals()

		assert.True(t, len(o.Products) == 1)
		assert.True(t, o.TotalAmountWithDiscount == 10)
		assert.True(t, o.TotalAmount == 20)
		assert.True(t, o.TotalDiscount == 10)
	})
}