package session

import (
	"net/http"
)

// SessionService defines the session api
type SessionService interface {
	Get(r *http.Request) (email string)
	Set(w http.ResponseWriter, email string) error
}
