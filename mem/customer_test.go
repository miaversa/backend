package mem_test

import (
	"testing"

	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/mem"
)

func TestCustomerService(t *testing.T) {
	storage := mem.NewCustomerStorage()
	email, name, password := "maria@gmail.com", "Maria Madalena", "123"
	c := &customer.Customer{
		Email:    email,
		Name:     name,
		Password: password,
	}
	err := storage.PutCustomer(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	c, err = storage.GetCustomer(email)
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
	_, err := storage.GetCustomer("joao@gmail.com")
	if err == nil {
		t.Fatal("deveria retornar um erro")
	}
}

func TestCustomerServiceAlreadyExists(t *testing.T) {
	storage := mem.NewCustomerStorage()
	email, name, password := "maria@gmail.com", "Maria Madalena", "123"
	c := &customer.Customer{
		Email:    email,
		Name:     name,
		Password: password,
	}
	err := storage.PutCustomer(c)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = storage.PutCustomer(c)
	if err == nil {
		t.Fatal("deveria retornar um erro de usuário já cadastrado")
	}
}
