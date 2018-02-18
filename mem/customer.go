package mem

import (
	"github.com/miaversa/backend/customer"
)

type memCustomerStorage struct {
	customers map[string]*customer.Customer
}

func NewCustomerStorage() *memCustomerStorage {
	return &memCustomerStorage{
		customers: map[string]*customer.Customer{},
	}
}

func (s *memCustomerStorage) GetCustomer(email string) (*customer.Customer, error) {
	if _, ok := s.customers[email]; !ok {
		return nil, customer.CustomerNotFoundErr
	}
	return s.customers[email], nil
}

func (s *memCustomerStorage) PutCustomer(c *customer.Customer) error {
	if _, ok := s.customers[c.Email]; ok {
		return customer.CustomerAlreadyExistsErr
	}
	s.customers[c.Email] = c
	return nil
}

func (s *memCustomerStorage) UpdateCustomer(c *customer.Customer) error {
	if _, ok := s.customers[c.Email]; !ok {
		return customer.CustomerNotFoundErr
	}
	s.customers[c.Email] = c
	return nil
}
