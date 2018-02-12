package customer

type CustomerService interface {
	Get(email string) (Customer, error)
	Put(Customer) error
	GetShippingAddress(string) (ShippingAddress, error)
	SetShippingAddress(string, ShippingAddress) error
}

type Customer struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ShippingAddress struct {
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}
