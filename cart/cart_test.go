package cart_test

import (
	"github.com/miaversa/backend/cart"
	"github.com/miaversa/backend/product"
	"testing"
)

func TestCart(t *testing.T) {
	c := cart.New()
	if c.Products == nil {
		t.Fatal("lista de produtos não inicializada")
	}
	if 0 != c.Total() {
		t.Fatal("um carrinho inicializado deve ser zerado de valor")
	}
	if 0 != c.Quantity() {
		t.Fatal("um carrinho inicializado deve ser zerado de quantidades")
	}
}

func TestCartAddProduct(t *testing.T) {
	c := cart.New()
	p := product.Product{
		SKU:     "sku",
		Name:    "name",
		Price:   100,
		Options: []product.Option{product.Option{Name: "opt", Value: "val"}},
	}
	c.AddProduct(p)
	if 100 != c.Total() {
		t.Fatal("o total do carrinho não atualizou")
	}
	if 1 != c.Quantity() {
		t.Fatal("a quantidade no carrinho não atualizou")
	}
	if p.Name != c.Products[0].Name {
		t.Fatal("produto não inserido no carrinho")
	}
}

func TestCartRemoveProduct(t *testing.T) {
	c := cart.New()
	p := product.Product{
		SKU:     "sku",
		Name:    "name",
		Price:   100,
		Options: []product.Option{product.Option{Name: "opt", Value: "val"}},
	}
	p2 := product.Product{
		SKU:     "sku2",
		Name:    "name2",
		Price:   100,
		Options: []product.Option{product.Option{Name: "opt2", Value: "val2"}},
	}
	c.AddProduct(p)
	c.RemoveProduct(0)
	if len(c.Products) > 0 {
		t.Fatal("deveria remover todos os produtos do carrinho")
	}
	c.AddProduct(p)
	c.AddProduct(p2)
	c.RemoveProduct(0)
}
