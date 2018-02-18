package customer

import "errors"

type CustomerStorage interface {
	GetCustomer(email string) (*Customer, error)
	PutCustomer(customer *Customer) error
	UpdateCustomer(customer *Customer) error
}

var CustomerNotFoundErr = errors.New("customer not found")
var CustomerAlreadyExistsErr = errors.New("customer already exists")

type Customer struct {
	Email           string          `json:"email"`
	Name            string          `json:"name"`
	Password        string          `json:"password"`
	ShippingAddress ShippingAddress `json:"shipping"`
}

type ShippingAddress struct {
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}
