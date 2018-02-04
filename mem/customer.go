package mem

import (
	"errors"
	"github.com/miaversa/backend/customer"
)

type memCustomerService struct {
	customers map[string]customer.Customer
	address   map[string]customer.ShippingAddress
}

func NewCustomerService() *memCustomerService {
	return &memCustomerService{customers: map[string]customer.Customer{}}
}

func (s *memCustomerService) Get(email string) (customer.Customer, error) {
	return customer.Customer{}, nil
}

func (s *memCustomerService) Put(c customer.Customer) error {
	if _, ok := s.customers[c.Email]; ok {
		return errors.New("email exists")
	}
	s.customers[c.Email] = c
	return nil
}

func (s *memCustomerService) SetShippingAddress(email string, sa customer.ShippingAddress) error {
	if _, ok := s.customers[email]; !ok {
		return errors.New("customer not found")
	}
	s.address[email] = sa
	return nil
}
