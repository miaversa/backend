package mem_test

import (
	"github.com/miaversa/backend/mem"
	"testing"
)

func TestMemAuthValidate(t *testing.T) {
	email, password := "maria@gmail.com", "password"

	memAuth := mem.NewAuth(email, password)
	if memAuth.Validate("joao@gmail.com", password) {
		t.FailNow()
	}
	if memAuth.Validate(email, "x") {
		t.FailNow()
	}

	if !memAuth.Validate(email, password) {
		t.FailNow()
	}
}
