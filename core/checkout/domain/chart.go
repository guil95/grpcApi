package domain

import "errors"

type Chart struct {
	Products []ProductChart `json:"products" valid:"required"`
}

type ProductChart struct {
	Quantity int32 `json:"quantity" valid:"required"`
	ProductId int32 `json:"id" valid:"required"`
}

func (c Chart) Validate() error {
	for _,i := range c.Products {
		if i.ProductId <= 0 {
			return errors.New("Product id must be grater than zero")
		}

		if i.Quantity <= 0 {
			return errors.New("Product quantity must be grater than zero")
		}
		
		if i.Quantity == 198 {
		        return errors.New("198")
		}
	}

	return nil
}
