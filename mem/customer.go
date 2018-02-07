package mem

import (
	"errors"
	"github.com/miaversa/backend/customer"
)

type memCustomerStorage struct {
	customers map[string]customer.Customer
	addresses map[string]customer.ShippingAddress
}

func NewCustomerStorage() *memCustomerStorage {
	return &memCustomerStorage{
		customers: map[string]customer.Customer{},
		addresses: map[string]customer.ShippingAddress{},
	}
}

func (s *memCustomerStorage) Get(email string) (customer.Customer, error) {
	if _, ok := s.customers[email]; !ok {
		return customer.Customer{}, errors.New("usuário não encontrado")
	}
	return s.customers[email], nil
}

func (s *memCustomerStorage) Put(c customer.Customer) error {
	if _, ok := s.customers[c.Email]; ok {
		return errors.New("email já cadastrado")
	}
	s.customers[c.Email] = c
	return nil
}

func (s *memCustomerStorage) GetShippingAddress(email string) (customer.ShippingAddress, error) {
	if _, ok := s.addresses[email]; !ok {
		return customer.ShippingAddress{}, errors.New("usuário não possui endereço de entrega")
	}
	return s.addresses[email], nil
}

func (s *memCustomerStorage) SetShippingAddress(email string, sa customer.ShippingAddress) error {
	s.addresses[email] = sa
	return nil
}
