package order

import (
	"errors"

	"github.com/eylmzer/campaingdemo/pkg/product"
)

type Order struct {
	Product  *product.Product
	Quantity int
}

func NewOrder(product *product.Product, quantity int) (*Order, error) {
	if quantity <= 0 {
		return nil, errors.New("invalid quantity")
	}

	err := product.DecreaseStock(quantity)
	if err != nil {
		return nil, err
	}

	return &Order{
		Product:  product,
		Quantity: quantity,
	}, nil
}
