package middleware

import (
	"context"
	"github.com/miaversa/backend/cart"
	"net/http"
)

type ctxKeyCart int

const CartKey ctxKeyCart = 0

type cartMiddleware struct {
	db cart.CartStorage
}

func NewCartMiddleware(db cart.CartStorage) *cartMiddleware {
	return &cartMiddleware{db: db}
}

func (m *cartMiddleware) Cart(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cid := GetCartID(r.Context())
		if cid == "" {
			panic("cid em branco")
		}
		ca, err := m.db.GetCart(cid)
		if err != nil && err.Error() != "item not found" {
			panic(err)
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, CartKey, ca)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		ca2, ok := GetCart(r.Context())
		err = m.db.SaveCart(ca2)
		if err != nil {
			panic(err)
		}
	}
	return http.HandlerFunc(fn)
}

func GetCart(ctx context.Context) (cart.Cart, bool) {
	if ctx == nil {
		return cart.New("x"), false
	}
	cart, ok := ctx.Value(CartKey).(cart.Cart)
	return cart, ok
}
