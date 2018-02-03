package login

type dummyAuth struct {
	email    string
	password string
}

func NewDummyAuth(email, password string) *dummyAuth {
	return &dummyAuth{email: email, password: password}
}

func (d *dummyAuth) Validate(email, password string) bool {
	if d.email == email && d.password == password {
		return true
	}
	return false
}
