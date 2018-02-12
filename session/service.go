package session

// SessionService defines the session api
type SessionService interface {
	Get() (string, error)
	Set(email string) error
}
