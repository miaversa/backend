package customer

type CustomerService interface {
	Get(email string) (Customer, error)
	Put(Customer) error
}

type Customer struct {
	Email    string
	Name     string
	Password string
}
