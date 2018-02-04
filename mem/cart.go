package mem

import (
	"encoding/base64"
	"encoding/json"
	"github.com/miaversa/backend/model"
	"net/http"
)

type memCartStore struct {
	cookieName string
	Cart       model.Cart
}

func NewCartStore(name string) *memCartStore {
	return &memCartStore{
		cookieName: name,
		Cart:       model.Cart{Items: []model.CartItem{}},
	}
}

func (s *memCartStore) GetCart(r *http.Request) (model.Cart, error) {
	return s.Cart, nil
}

func (s *memCartStore) SaveCart(w http.ResponseWriter, c model.Cart) error {
	s.Cart = c
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	str := base64.StdEncoding.EncodeToString(b)

	cartCookie := &http.Cookie{Name: s.cookieName, Value: str}
	http.SetCookie(w, cartCookie)
	return nil
}

func (s *memCartStore) DropCart(w http.ResponseWriter) {
}
