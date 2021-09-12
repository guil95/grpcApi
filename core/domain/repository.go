package domain

type Repository interface {
	GetProducts()[]Product
	GetProductsByChart(chart *Chart)[]Product
}
