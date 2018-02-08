package cookie

/*
import (
	"github.com/gorilla/securecookie"
	"github.com/miaversa/backend/model"
	"net/http"
)

type cartStore struct {
	cookieName string
	sc         *securecookie.SecureCookie
	secure     bool
}

// NewCartStore creates a new Cart Storage
func NewCartStore(cookieName string, sc *securecookie.SecureCookie, secure bool) *cartStore {
	return &cartStore{
		cookieName: cookieName,
		sc:         sc,
		secure:     secure,
	}
}

func (s *cartStore) GetCart(r *http.Request) (model.Cart, error) {
	c := model.Cart{Products: []model.Product{}}
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

func (s *cartStore) SaveCart(w http.ResponseWriter, c model.Cart) error {
	encoded, err := s.sc.Encode(s.cookieName, c)
	if err != nil {
		return err
	}
	cookie := s.createCookie(encoded)
	http.SetCookie(w, cookie)
	return nil
}

func (s *cartStore) DropCart(w http.ResponseWriter) {
	cookie := s.createCookie("")
	http.SetCookie(w, cookie)
}

func (s *cartStore) createCookie(value string) *http.Cookie {
	return &http.Cookie{
		Name:     s.cookieName,
		Value:    value,
		Path:     "/",
		Secure:   s.secure,
		HttpOnly: true,
	}
}
*/
