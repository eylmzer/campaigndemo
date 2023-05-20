package product

import "errors"

type Product struct {
	Code  string
	Price float64
	Stock int
}

func NewProduct(code string, price float64, stock int) *Product {
	return &Product{
		Code:  code,
		Price: price,
		Stock: stock,
	}
}

func (p *Product) DecreaseStock(quantity int) error {
	if p.Stock < quantity {
		return errors.New("insufficient stock")
	}
	p.Stock -= quantity
	return nil
}
