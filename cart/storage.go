package cart

type CartStorage interface {
	GetCart() (Cart, error)
	SaveCart(c Cart) error
	DropCart() error
}
