package cart

type CartStorage interface {
	GetCart(id string) (Cart, error)
	SaveCart(c Cart) error
	DropCart(id string) error
}
