package cookie

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

type sessionService struct {
	cookieName string
	sc         *securecookie.SecureCookie
	secure     bool
}

func NewSessionService(cookieName, hashKey, blockKey string, secure bool) *sessionService {
	return &sessionService{
		cookieName: cookieName,
		sc:         securecookie.New([]byte(hashKey), []byte(blockKey)),
		secure:     secure,
	}
}

func (s *sessionService) Get(r *http.Request) (session string) {
	if cookie, err := r.Cookie(s.cookieName); err == nil {
		s.sc.Decode(s.cookieName, cookie.Value, &session)
	}
	return
}

func (s *sessionService) Set(w http.ResponseWriter, session string) error {
	encoded, err := s.sc.Encode(s.cookieName, session)
	if err != nil {
		return err
	}
	cookie := s.createCookie(encoded)
	http.SetCookie(w, cookie)
	return nil
}

func (s *sessionService) createCookie(value string) *http.Cookie {
	return &http.Cookie{
		Name:     s.cookieName,
		Value:    value,
		Path:     "/",
		Secure:   s.secure,
		HttpOnly: true,
	}
}
