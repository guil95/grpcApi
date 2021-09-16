package domain

import (
	"errors"
	"log"
)

type Order struct {
	TotalAmount int32             `json:"total_amount"`
	TotalAmountWithDiscount int32 `json:"total_amount_with_discount"`
	TotalDiscount int32           `json:"total_discount"`
	Products []ProductOrder       `json:"products"`
}

func (o *Order) NewOrder() *Order {
	return &Order{}
}

func (o *Order) AddProduct(product *Product, discount int32) {
	log.Println("Include product in order")

	if product == nil {
		return
	}

	if product.Gift == true {
		product.Amount = 0
		discount = 0
		product.Quantity = 1
	}

	o.Products = append(o.Products, ProductOrder{
		Id: product.Id,
		Quantity: product.Quantity,
		UnitAmount: product.Amount,
		TotalAmount: product.Amount*product.Quantity,
		Discount: discount*product.Quantity,
		Gift: product.Gift,
	})
}

func (o *Order) CalcTotals() {
	var totalAmount int32
	var totalAmountWithDiscount int32
	var totalDiscount int32

	for _, product := range o.Products {
		totalAmount += product.UnitAmount * product.Quantity
		totalAmountWithDiscount += product.TotalAmount - product.Discount
		totalDiscount += product.Discount
	}

	o.TotalAmount = totalAmount
	o.TotalAmountWithDiscount = totalAmountWithDiscount
	o.TotalDiscount = totalDiscount
}

var UnprocessableEntity = errors.New("UnprocessableEntity")