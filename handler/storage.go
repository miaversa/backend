package handler

import (
	"github.com/miaversa/backend/model"
)

type CartStorage interface {
	GetCart() (model.Cart, error)
	SaveCart(c model.Cart) error
	DropCart() error
}
