package dynamodb_test

import (
	"github.com/miaversa/backend/cart"
	bkdb "github.com/miaversa/backend/dynamodb"
	"github.com/miaversa/backend/product"
	"github.com/miaversa/backend/testutil"
	"github.com/miaversa/backend/uuid"
	"testing"
)

func TestEnsureTable(t *testing.T) {
	cli := testutil.NewDynamoDB()
	_, err := bkdb.NewCartStorage(cli)
	if err != nil {
		t.Fatal(err)
	}
	_, err = bkdb.NewCartStorage(cli)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCart(t *testing.T) {
	cli := testutil.NewDynamoDB()
	dynamo, err := bkdb.NewCartStorage(cli)
	if err != nil {
		t.Fatal(err)
	}
	cid := uuid.New()
	carrinho := cart.New(cid)
	produto1 := product.Product{
		SKU:     "sku1",
		Name:    "name1",
		Price:   101,
		Options: []product.Option{product.Option{Name: "opt", Value: "val"}},
	}
	produto2 := product.Product{
		SKU:     "sku2",
		Name:    "name2",
		Price:   102,
		Options: []product.Option{product.Option{Name: "opt", Value: "val"}},
	}
	carrinho.AddProduct(produto1)
	carrinho.AddProduct(produto2)
	err = dynamo.SaveCart(carrinho)
	if err != nil {
		t.Fatal(err)
	}
	carrinho2, err := dynamo.GetCart(cid)
	if err != nil {
		t.Fatal(err)
	}
	if carrinho.Key != carrinho2.Key {
		t.Fatal("erro na chave")
	}
	if carrinho.Quantity() != carrinho2.Quantity() {
		t.Fatal("erro na quantidade")
	}
	if carrinho.Total() != carrinho2.Total() {
		t.Fatal("Erro no total")
	}
	if carrinho.Products[0].Name != carrinho2.Products[0].Name {
		t.Fatal("erro no nome de um produto")
	}
}
