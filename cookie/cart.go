package cookie

import (
	"github.com/gorilla/securecookie"
	"github.com/miaversa/backend/model"
	"net/http"
)

type cartService struct {
	cookieName string
	sc         *securecookie.SecureCookie
	secure     bool
}

func New(cookieName, hashKey, blockKey string, secure bool) *cartService {
	return &cartService{
		cookieName: cookieName,
		sc:         securecookie.New([]byte(hashKey), []byte(blockKey)),
		secure:     secure,
	}
}

func (s *cartService) GetCart(r *http.Request) (model.Cart, error) {
	c := model.Cart{Items: []model.CartItem{}}
	cookie, err := r.Cookie(s.cookieName)
	if err != nil {
		return c, nil
	}
	err = s.sc.Decode(s.cookieName, cookie.Value, &c)
	if err != nil {
		return model.Cart{}, err
	}
	return c, nil
}

func (s *cartService) SaveCart(w http.ResponseWriter, c model.Cart) error {
	encoded, err := s.sc.Encode(s.cookieName, c)
	if err != nil {
		return err
	}
	cookie := s.createCookie(encoded)
	http.SetCookie(w, cookie)
	return nil
}

func (s *cartService) DropCart(w http.ResponseWriter) {
	cookie := s.createCookie("")
	http.SetCookie(w, cookie)
}

func (s *cartService) createCookie(value string) *http.Cookie {
	return &http.Cookie{
		Name:     s.cookieName,
		Value:    value,
		Path:     "/",
		Secure:   s.secure,
		HttpOnly: true,
	}
}
