package customer

type CustomerService interface {
	Get(email string) (Customer, error)
	Put(Customer) error
	SetShippingAddress(string, ShippingAddress) error
}

type Customer struct {
	Email    string
	Name     string
	Password string
}

type ShippingAddress struct {
	Street     string
	Number     string
	Complement string
	District   string
	City       string
	State      string
	Country    string
	PostalCode string
}
