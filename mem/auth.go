package mem

type memAuth struct {
	email    string
	password string
}

func NewAuth(email, password string) *memAuth {
	return &memAuth{email: email, password: password}
}

func (d *memAuth) Validate(email, password string) bool {
	if d.email == email && d.password == password {
		return true
	}
	return false
}
