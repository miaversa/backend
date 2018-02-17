package middleware

import (
	"context"
	"github.com/miaversa/backend/uuid"
	"net/http"
	"time"
)

type ctxKeyCartID int

const CartIDKey ctxKeyCartID = 0

func CartID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var id string
		if cookie, err := r.Cookie("cid"); err == nil {
			id = cookie.Value
		} else {
			id = uuid.New()
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, CartIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "cid", Value: id, Expires: expiration}
		http.SetCookie(w, &cookie)
	}
	return http.HandlerFunc(fn)
}

func GetCartID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if cartID, ok := ctx.Value(CartIDKey).(string); ok {
		return cartID
	}
	return ""
}
