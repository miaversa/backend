package order

import (
	"github.com/miaversa/backend/cart"
	"github.com/miaversa/backend/customer"
)

type Order struct {
	Cart     *cart.Cart
	Customer *customer.Customer
}
