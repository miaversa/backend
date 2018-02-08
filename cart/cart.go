package cart

import (
	"github.com/miaversa/backend/product"
)

// Cart stores the cart
type Cart struct {
	Shipping float64           `json:"shipping"`
	Products []product.Product `json:"products"`
}

func New() Cart {
	return Cart{Products: []product.Product{}}
}

// Total returns the total
func (c Cart) Total() float64 {
	sum := c.Shipping
	for _, p := range c.Products {
		sum += p.Price
	}
	return sum
}

// Quantity returns the item quantity in the cart
func (c Cart) Quantity() int {
	return len(c.Products)
}

// AddItem adds a new item in the cart
func (c *Cart) AddProduct(product product.Product) {
	c.Products = append(c.Products, product)
}

// RemoveItem remove an item from the cart
func (c *Cart) RemoveProduct(index int) {
	list := []product.Product{}
	for k, v := range c.Products {
		if index != k {
			list = append(list, v)
		}
	}
	c.Products = list
}
