package mem_test

import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/mem"
	"testing"
)

func TestCustomerService(t *testing.T) {
	storage := mem.NewCustomerStorage()
	email, name, password := "maria@gmail.com", "Maria Madalena", "123"
	c := customer.Customer{
		Email:    email,
		Name:     name,
		Password: password,
	}
	err := storage.Put(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	c, err = storage.Get(email)
	if err != nil {
		t.Fatal(err.Error())
	}
	if email != c.Email {
		t.Fatal("erro ao recuperar o email do cliente")
	}
	if name != c.Name {
		t.Fatal("erro ao recuperar o nome do cliente")
	}
	if password != c.Password {
		t.Fatal("erro ao recuperar a senha do cliente")
	}
}

func TestCustomerServiceNotFound(t *testing.T) {
	storage := mem.NewCustomerStorage()
	_, err := storage.Get("joao@gmail.com")
	if err == nil {
		t.Fatal("deveria retornar um erro")
	}
}

func TestCustomerServiceAlreadyExists(t *testing.T) {
	storage := mem.NewCustomerStorage()
	email, name, password := "maria@gmail.com", "Maria Madalena", "123"
	c := customer.Customer{
		Email:    email,
		Name:     name,
		Password: password,
	}
	err := storage.Put(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = storage.Put(c)
	if err == nil {
		t.Fatal("deveria retornar um erro de usuário já cadastrado")
	}
}

func TestShipping(t *testing.T) {
	storage := mem.NewCustomerStorage()
	email := "maria@gmail.com"
	city := "Sao Paulo"
	addr := customer.ShippingAddress{City: city}
	err := storage.SetShippingAddress(email, addr)
	if err != nil {
		t.Fatal(err.Error())
	}
	addr, err = storage.GetShippingAddress(email)
	if err != nil {
		t.Fatal(err.Error())
	}
	if city != addr.City {
		t.Fatal("erro ao recuperar o endereço")
	}
}
func TestShippingNotFound(t *testing.T) {
	storage := mem.NewCustomerStorage()
	_, err := storage.GetShippingAddress("joao@gmail.com")
	if err == nil {
		t.Fatal("deveria retornar um erro")
	}
}
