package domain

type ProductOrder struct {
	Id int32 `json:"id"`
	Quantity int32 `json:"quantity"`
	UnitAmount int32 `json:"unit_amount"`
	TotalAmount int32 `json:"total_amount"`
	Discount int32 `json:"discount"`
	Gift bool `json:"is_gift"`
}
