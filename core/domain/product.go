package domain

type Product struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Amount int32 `json:"amount"`
	Gift bool `json:"is_gift"`
	Quantity int32 `json:"quantity"`
}
