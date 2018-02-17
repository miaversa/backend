package mem

import (
	"github.com/miaversa/backend/cart"
	"github.com/miaversa/backend/product"
)

type memCartStorage struct {
	Cart cart.Cart
}

func NewCartStorage() *memCartStorage {
	return &memCartStorage{
		Cart: cart.Cart{Products: []product.Product{}},
	}
}

func (s *memCartStorage) GetCart(id string) (cart.Cart, error) {
	return s.Cart, nil
}

func (s *memCartStorage) SaveCart(id string, c cart.Cart) error {
	s.Cart = c
	return nil
}

func (s *memCartStorage) DropCart(id string) error {
	s.Cart.Products = []product.Product{}
	return nil
}
