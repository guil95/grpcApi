package domain

import "errors"

type Product struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Amount int32 `json:"amount"`
	Gift bool `json:"is_gift"`
	Quantity int32 `json:"quantity"`
}

var ProductGiftError = errors.New("Product gift error")
var ProductNotFoundError = errors.New("Product not found error")