package product

// Product is a buyable product
type Product struct {
	SKU     string   `json:"sku"`
	Name    string   `json:"name"`
	Price   float64  `json:"price"`
	Options []Option `json:"options"`
}

// Option represents a product variant
type Option struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
