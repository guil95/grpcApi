package domain

type Client interface {
	GetDiscount(productId int32) (float32, error)
}