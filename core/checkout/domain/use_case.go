package domain

type UseCase interface {
	Checkout(chart *Chart)(*Order, error)
}
