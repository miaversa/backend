package mem_test

import (
	"github.com/miaversa/backend/mem"
	"github.com/miaversa/backend/product"
	"testing"
)

func TestCartStorage(t *testing.T) {
	storage := mem.NewCartStorage()
	c, err := storage.GetCart()
	if err != nil {
		t.Fatal("erro ao pegar o carrinho")
	}
	sku, name, price := "sku", "name", 10.5
	optKey, optVal := "size", "15"
	p := product.Product{
		SKU:   sku,
		Name:  name,
		Price: price,
		Options: []product.Option{
			product.Option{Name: optKey, Value: optVal},
		},
	}
	c.AddProduct(p)
	storage.SaveCart(c)
	c, err = storage.GetCart()
	if err != nil {
		t.Fatal("erro ao pegar o carrinho")
	}
	if sku != c.Products[0].SKU {
		t.Fatal("erro ao recuperar o codigo do produto")
	}
	if name != c.Products[0].Name {
		t.Fatal("erro ao recuperar o nome do produto")
	}
	if price != c.Products[0].Price {
		t.Fatal("erro ao recuperar o preço do produto")
	}
	if optKey != c.Products[0].Options[0].Name {
		t.Fatal("erro ao recuperar o nome da primeira opção")
	}
	if optVal != c.Products[0].Options[0].Value {
		t.Fatal("erro ao recuperar o valor da primeira opção")
	}
	err = storage.DropCart()
	if err != nil {
		t.Fatal("erro ao dropar o carrinho")
	}
	if len(storage.Cart.Products) > 0 {
		t.Fatal("erro ao dropar o carrinho")
	}
}
