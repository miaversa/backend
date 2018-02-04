package customer

import (
	"errors"
)

type dummyCustomerService struct {
	customers map[string]Customer
}

func NewDummyCustomerService() *dummyCustomerService {
	return &dummyCustomerService{customers: map[string]Customer{}}
}

func (s *dummyCustomerService) Get(email string) (Customer, error) {
	return Customer{}, nil
}

func (s *dummyCustomerService) Put(c Customer) error {
	if _, ok := s.customers[c.Email]; ok {
		return errors.New("email exists")
	}
	s.customers[c.Email] = c
	return nil
}
