package middleware_test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	bkdb "github.com/miaversa/backend/dynamodb"
	"github.com/miaversa/backend/middleware"
	"github.com/miaversa/backend/product"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestCli(debug bool) *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("sa-east-1"), Endpoint: aws.String("http://localhost:8000")})
	if err != nil {
		panic(err)
	}
	if debug {
		return dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	}

	return dynamodb.New(sess, aws.NewConfig())
}

type testCartMiddlewareHandler struct {
}

func (h *testCartMiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cart := middleware.GetCart(r.Context())
	cart.AddProduct(product.Product{Name: "Um produto"})
}

func TestCartMiddleware(t *testing.T) {
	cli := newTestCli(false)
	tcmh := &testCartMiddlewareHandler{}

	dynamo, err := bkdb.NewCartStorage(cli)
	if err != nil {
		t.Fatal(err)
	}
	mCart := middleware.NewCartMiddleware(dynamo)
	mCartWithCartID := middleware.CartID(mCart.Cart(tcmh))

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	mCartWithCartID.ServeHTTP(rr, req)

	t.Log(req.Context())
}
