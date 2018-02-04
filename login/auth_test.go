package login_test

import (
	"github.com/miaversa/backend/login"
	"testing"
)

func TestDummyAuthValidate(t *testing.T) {
	email, password := "maria@gmail.com", "password"

	dummyAuth := login.NewDummyAuth(email, password)
	if dummyAuth.Validate("joao@gmail.com", password) {
		t.FailNow()
	}
	if dummyAuth.Validate(email, "x") {
		t.FailNow()
	}

	if !dummyAuth.Validate(email, password) {
		t.FailNow()
	}
}
