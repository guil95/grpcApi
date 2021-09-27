package use_cases_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/guil95/grpcApi/config"
	"github.com/guil95/grpcApi/core/checkout/domain"
	"github.com/guil95/grpcApi/core/checkout/use_cases"
	"github.com/guil95/grpcApi/core/checkout/use_cases/mock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)
func TestCheckout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	os.Setenv("API_PORT", "api")
	os.Setenv("DB_FILE", "teste")
	os.Setenv("CLIENT_HOST", "host")
	os.Setenv("CLIENT_PORT", "port")

	t.Run("Test checkout with valid product should return order", func(t *testing.T) {
		os.Setenv("BLACK_FRIDAY_DATE", "2020-02-02")

		chart := &domain.Chart{Products: []domain.ProductChart{{Quantity: 1, ProductId: 1}}}
		c := mock.NewMockClient(ctrl)
		r := mock.NewMockRepository(ctrl)

		r.EXPECT().GetProductsByChart(chart).Return([]domain.Product{{
			Id: 1,
			Title: "Produto teste",
			Description: "Produto teste",
			Amount: 10,
			Gift: false,
			Quantity: 1,
		}}).Times(1)

		c.EXPECT().GetDiscount(int32(1)).Return(float32(0.5), nil).Times(1)

		conf, _ := config.RetrieveConfig()

		service := use_cases.NewCreateCheckoutUseCase(c, r, conf)

		order,err := service.Checkout(chart)

		assert.Nil(t, err)
		assert.True(t, order.TotalAmountWithDiscount == 5)
		assert.True(t, order.TotalDiscount == 5)
		assert.True(t, order.TotalAmount == 10)
	})

	t.Run("Test checkout in black friday with valid product should return order with gift", func(t *testing.T) {
		os.Setenv("BLACK_FRIDAY_DATE", time.Now().Format("2006-01-02"))

		chart := &domain.Chart{Products: []domain.ProductChart{{Quantity: 1, ProductId: 1}}}

		c := mock.NewMockClient(ctrl)
		r := mock.NewMockRepository(ctrl)

		r.EXPECT().GetGiftProducts().Return([]domain.Product{{
			Id: 2,
			Title: "Produto teste",
			Description: "Produto teste",
			Amount: 20,
			Gift: true,
			Quantity: 1,
		}}).Times(1)

		r.EXPECT().GetProductsByChart(chart).Return([]domain.Product{{
			Id: 1,
			Title: "Produto teste",
			Description: "Produto teste",
			Amount: 10,
			Gift: false,
			Quantity: 1,
		}}).Times(1)

		c.EXPECT().GetDiscount(int32(1)).Return(float32(0.5), nil).Times(1)

		conf, _ := config.RetrieveConfig()

		service := use_cases.NewCreateCheckoutUseCase(c, r, conf)

		order,err := service.Checkout(chart)

		assert.Nil(t, err)
		assert.True(t, order.TotalAmountWithDiscount == 5)
		assert.True(t, order.TotalDiscount == 5)
		assert.True(t, order.TotalAmount == 10)
		assert.True(t, hasGiftProduct(order.Products))
	})

	t.Run("Test checkout in black friday with valid product should return order without gift", func(t *testing.T) {
		os.Setenv("BLACK_FRIDAY_DATE", time.Now().Format("2006-01-02"))

		chart := &domain.Chart{Products: []domain.ProductChart{{Quantity: 1, ProductId: 1}}}

		c := mock.NewMockClient(ctrl)
		r := mock.NewMockRepository(ctrl)

		r.EXPECT().GetGiftProducts().Return(nil).Times(1)

		r.EXPECT().GetProductsByChart(chart).Return([]domain.Product{{
			Id: 1,
			Title: "Produto teste",
			Description: "Produto teste",
			Amount: 10,
			Gift: false,
			Quantity: 1,
		}}).Times(1)

		c.EXPECT().GetDiscount(int32(1)).Return(float32(0.5), nil).Times(1)

		conf, _ := config.RetrieveConfig()

		service := use_cases.NewCreateCheckoutUseCase(c, r, conf)

		order,err := service.Checkout(chart)

		assert.Nil(t, err)
		assert.True(t, order.TotalAmountWithDiscount == 5)
		assert.True(t, order.TotalDiscount == 5)
		assert.True(t, order.TotalAmount == 10)
		assert.False(t, hasGiftProduct(order.Products))
	})

	t.Run("Test checkout with non exists products should return ProductNotFoundError", func(t *testing.T) {
		os.Setenv("BLACK_FRIDAY_DATE", "2020-02-02")

		chart := &domain.Chart{Products: []domain.ProductChart{{Quantity: 1, ProductId: 1}}}

		c := mock.NewMockClient(ctrl)
		r := mock.NewMockRepository(ctrl)

		r.EXPECT().GetProductsByChart(chart).Return(nil).Times(1)

		conf, _ := config.RetrieveConfig()

		service := use_cases.NewCreateCheckoutUseCase(c, r, conf)

		_,err := service.Checkout(chart)

		assert.NotNil(t, err)
		assert.True(t, err == domain.ProductNotFoundError)
	})

	t.Run("Test checkout with gift product should return ProductGiftError", func(t *testing.T) {
		os.Setenv("BLACK_FRIDAY_DATE", "2020-02-02")

		chart := &domain.Chart{Products: []domain.ProductChart{{Quantity: 1, ProductId: 1}}}

		c := mock.NewMockClient(ctrl)
		r := mock.NewMockRepository(ctrl)

		r.EXPECT().GetProductsByChart(chart).Return([]domain.Product{{
			Id: 1,
			Title: "Produto teste",
			Description: "Produto teste",
			Amount: 10,
			Gift: true,
			Quantity: 1,
		}}).Times(1)

		conf, _ := config.RetrieveConfig()

		service := use_cases.NewCreateCheckoutUseCase(c, r, conf)

		_,err := service.Checkout(chart)

		assert.NotNil(t, err)
		assert.True(t, err == domain.ProductGiftError)
	})

	t.Run("Test checkout with many equals products should return one product in order", func(t *testing.T) {
		os.Setenv("BLACK_FRIDAY_DATE", "2020-02-02")

		chart := &domain.Chart{Products: []domain.ProductChart{
			{Quantity: 1, ProductId: 1},
			{Quantity: 1, ProductId: 1},
			{Quantity: 1, ProductId: 1},
		}}

		c := mock.NewMockClient(ctrl)
		r := mock.NewMockRepository(ctrl)

		r.EXPECT().GetProductsByChart(chart).Return([]domain.Product{{
			Id: 1,
			Title: "Produto teste",
			Description: "Produto teste",
			Amount: 10,
			Gift: false,
			Quantity: 3,
		}})

		c.EXPECT().GetDiscount(int32(1)).Return(float32(0.5), nil).Times(1)

		conf, _ := config.RetrieveConfig()

		service := use_cases.NewCreateCheckoutUseCase(c, r, conf)

		order,err := service.Checkout(chart)

		assert.Nil(t, err)
		assert.True(t, order.TotalAmountWithDiscount == 15)
		assert.True(t, order.TotalDiscount == 15)
		assert.True(t, order.TotalAmount == 30)
		assert.True(t, len(order.Products) == 1)
	})

	t.Run("Test checkout with valid products and discount return error should return order without discount", func(t *testing.T) {
		os.Setenv("BLACK_FRIDAY_DATE", "2020-02-02")

		chart := &domain.Chart{Products: []domain.ProductChart{
			{Quantity: 1, ProductId: 1},
		}}

		c := mock.NewMockClient(ctrl)
		r := mock.NewMockRepository(ctrl)

		r.EXPECT().GetProductsByChart(chart).Return([]domain.Product{{
			Id: 1,
			Title: "Produto teste",
			Description: "Produto teste",
			Amount: 10,
			Gift: false,
			Quantity: 1,
		}})

		c.EXPECT().GetDiscount(int32(1)).Return(float32(0), errors.New("Qualquer erro")).Times(1)

		conf, _ := config.RetrieveConfig()

		service := use_cases.NewCreateCheckoutUseCase(c, r, conf)

		order,err := service.Checkout(chart)

		assert.Nil(t, err)
		assert.True(t, order.TotalAmountWithDiscount == 10)
		assert.True(t, order.TotalDiscount == int32(0))
		assert.True(t, order.TotalAmount == 10)
		assert.True(t, order.Products[0].Discount == int32(0))
	})
}

func hasGiftProduct(products []domain.ProductOrder) bool {
	for _,item := range products {
		if item.Gift == true {
			return true
		}
	}
	return false
}


