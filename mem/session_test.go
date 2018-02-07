package mem_test

import (
	"github.com/miaversa/backend/mem"
	"testing"
)

func TestSessionStorage(t *testing.T) {
	storage := mem.NewSessionStorage()
	email := "maria@gmail.com"
	err := storage.Set(email)
	if err != nil {
		t.Fatal("erro ao definir a sessão")
	}
	session, err := storage.Get()
	if err != nil {
		t.Fatal("erro ao pegar a sessão")
	}
	if email != session {
		t.Fatal("a sessão é diferente")
	}
}
