package mem

import (
	"net/http"
)

type memSessionService struct {
	cookieName string
	Email      string
}

func NewSessionService(name string) *memSessionService {
	return &memSessionService{
		cookieName: name,
	}
}

func (s *memSessionService) Get(r *http.Request) string {
	return s.Email
}

func (s *memSessionService) Set(w http.ResponseWriter, session string) error {
	c := &http.Cookie{Name: s.cookieName, Value: session}
	http.SetCookie(w, c)
	return nil
}
