package mem

import (
	"github.com/miaversa/backend/handler"
	"github.com/miaversa/backend/model"
)

type memCartStorage struct {
	Cart model.Cart
}

func NewCartStorage() handler.CartStorage {
	return &memCartStorage{
		Cart: model.Cart{Products: []model.Product{}},
	}
}

func (s *memCartStorage) GetCart() (model.Cart, error) {
	return s.Cart, nil
}

func (s *memCartStorage) SaveCart(c model.Cart) error {
	s.Cart = c
	return nil
}

func (s *memCartStorage) DropCart() error {
	s.Cart.Products = []model.Product{}
	return nil
}
