package pagseguro_test

import (
	"testing"

	"github.com/miaversa/backend/config"
	"github.com/miaversa/backend/pagseguro"
)

func TestSessionID(t *testing.T) {
	config.Load()
	id, err := pagseguro.SessionID()

	if err != nil {
		t.Fatal(err)
	}

	t.Log(id)
}
