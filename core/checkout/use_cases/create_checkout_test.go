package use_cases_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/guil95/grpcApi/core/checkout/domain"
	"github.com/guil95/grpcApi/core/checkout/use_cases"
	"github.com/guil95/grpcApi/core/checkout/use_cases/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckoutWithNonExistsProductsShouldReturnProductNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	chart := &domain.Chart{Products: []domain.ProductChart{{Quantity: 1, ProductId: 1}}}

	c := mock.NewMockClient(ctrl)
	r := mock.NewMockRepository(ctrl)

	r.EXPECT().GetProductsByChart(chart).Return(nil).Times(1)

	service := use_cases.NewService(c, r)

	_,err := service.Checkout(chart)

	assert.NotNil(t, err)
	assert.True(t, err == domain.ProductNotFoundError)
}

func TestCheckoutWithProductGiftShouldReturnProductGiftError(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

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

	service := use_cases.NewService(c, r)

	_,err := service.Checkout(chart)

	assert.NotNil(t, err)
	assert.True(t, err == domain.ProductGiftError)
}

func TestCheckoutWithValidProductShouldReturnOrder(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

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

	service := use_cases.NewService(c, r)

	order,err := service.Checkout(chart)

	assert.Nil(t, err)
	assert.True(t, order.TotalAmountWithDiscount == 5)
	assert.True(t, order.TotalDiscount == 5)
	assert.True(t, order.TotalAmount == 10)
}

func TestCheckoutWithManyEqualsProductsShouldReturnOneProductInOrder(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

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

	service := use_cases.NewService(c, r)

	order,err := service.Checkout(chart)

	assert.Nil(t, err)
	assert.True(t, order.TotalAmountWithDiscount == 15)
	assert.True(t, order.TotalDiscount == 15)
	assert.True(t, order.TotalAmount == 30)
	assert.True(t, len(order.Products) == 1)
}

func TestCheckoutWithValidProductsAndDiscountReturnErrorShouldReturnOrderWithoutDiscount(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

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

	c.EXPECT().GetDiscount(int32(1)).Return(float32(0), errors.New("Qualquer erro")).Times(1)

	service := use_cases.NewService(c, r)

	order,err := service.Checkout(chart)

	assert.Nil(t, err)
	assert.True(t, order.TotalAmountWithDiscount == 30)
	assert.True(t, order.TotalDiscount == int32(0))
	assert.True(t, order.TotalAmount == 30)
	assert.True(t, order.Products[0].Discount == int32(0))
}

