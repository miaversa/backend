package model

// Cart stores the cart
type Cart struct {
	Shipping float64    `json:"shipping"`
	Items    []CartItem `json:"items"`
}

// Total returns the total
func (c Cart) Total() float64 {
	sum := c.Shipping
	for _, v := range c.Items {
		sum += v.Product.Price * float64(v.Quantity)
	}
	return sum
}

// Quantity returns the item quantity in the cart
func (c Cart) Quantity() int {
	sum := 0
	for _, v := range c.Items {
		sum += v.Quantity
	}
	return sum
}

// AddItem adds a new item in the cart
func (c *Cart) AddItem(item CartItem) {
	newItem := true
	for k, v := range c.Items {
		if v.Product.SKU == item.Product.SKU {
			c.Items[k].Quantity += 1
			newItem = false
		}
	}
	if newItem {
		c.Items = append(c.Items, item)
	}
}

// RemoveItem remove an item from the cart
func (c *Cart) RemoveItem(sku string) {
	var index int = -1
	for k, v := range c.Items {
		if v.Product.SKU == sku {
			index = k
		}
	}
	if index > -1 {
		c.Items = append(c.Items[:index], c.Items[index+1:]...)
	}
}

// CartItem represents one item type in the cart
type CartItem struct {
	Product  Product `json:"product"`
	Quantity int
}

// Total returns the Item total (price * quantity)
func (ci CartItem) Total() float64 {
	return ci.Product.Price * float64(ci.Quantity)
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
