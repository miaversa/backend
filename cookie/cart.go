package cookie

import (
	"github.com/gorilla/securecookie"
	"github.com/miaversa/backend/model"
	"net/http"
)

type Service interface {
	GetCart(r *http.Request) (model.Cart, error)
	SaveCart(w http.ResponseWriter, c model.Cart) error
}

type service struct {
	sc *securecookie.SecureCookie
}

func New() *service {
	hashKey := []byte("12345")
	blockKey := []byte("1234567890123456")
	return &service{sc: securecookie.New(hashKey, blockKey)}
}

func (s *service) GetCart(r *http.Request) (model.Cart, error) {
	c := model.Cart{Items: []model.CartItem{}}
	cookie, err := r.Cookie("mcart")
	if err != nil {
		return c, nil
	}
	err = s.sc.Decode("mcart", cookie.Value, &c)
	if err != nil {
		return model.Cart{}, err
	}
	return c, nil
}

func (s *service) SaveCart(w http.ResponseWriter, c model.Cart) error {
	encoded, err := s.sc.Encode("mcart", c)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     "mcart",
		Value:    encoded,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return nil
}
