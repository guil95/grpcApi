package domain

type CreateCheckoutUseCase interface {
	Checkout(chart *Chart)(*Order, error)
}
