package mem

type memSessionStorage struct {
	ID string
}

func NewSessionStorage() *memSessionStorage {
	return &memSessionStorage{}
}

func (s *memSessionStorage) Get() (string, error) {
	return s.ID, nil
}

func (s *memSessionStorage) Set(id string) error {
	s.ID = id
	return nil
}
