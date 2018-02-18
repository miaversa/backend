package mem

import (
	"github.com/miaversa/backend/cart"
)

type memCartStorage struct {
	Cart *cart.Cart
}

func NewCartStorage() *memCartStorage {
	return &memCartStorage{}
}

func (s *memCartStorage) GetCart(id string) (*cart.Cart, error) {
	if s.Cart == nil {
		return nil, cart.CartNotFoundErr
	}
	return s.Cart, nil
}

func (s *memCartStorage) SaveCart(c *cart.Cart) error {
	s.Cart = c
	return nil
}

func (s *memCartStorage) DropCart(id string) error {
	s.Cart = nil
	return nil
}
