package model

// Cart stores the cart
type Cart struct {
	Shipping float64   `json:"shipping"`
	Products []Product `json:"products"`
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
	sum := 0
	for _, _ = range c.Products {
		sum++
	}
	return sum
}

// AddItem adds a new item in the cart
func (c *Cart) AddProduct(product Product) {
	c.Products = append(c.Products, product)
}

// RemoveItem remove an item from the cart
func (c *Cart) RemoveProduct(index int) {
	c.Products = append(c.Products[:index], c.Products[index+1:]...)
}

// ProductOption represents a product variant
type ProductOption struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

// Product is a buyable product
type Product struct {
	SKU     string          `json:"sku"`
	Name    string          `json:"name"`
	Price   float64         `json:"price"`
	Options []ProductOption `json:"options"`
}
