package domain

import "errors"

type Chart struct {
	Products []ProductChart `json:"products" validate:"required"`
}

type ProductChart struct {
	Quantity int32 `json:"quantity" validate:"required"`
	ProductId int32 `json:"id" validate:"required"`
}

var ProductGiftError = errors.New("Product gift error")