package memory

import "github.com/authorhealth/go-exercise/domain"

type Store struct {
	userStore *UserStore
}

var _ domain.Storer = (*Store)(nil)

func NewStore() *Store {
	return &Store{
		userStore: NewUserStore(),
	}
}

func (s *Store) Users() domain.UserStorer {
	return s.userStore
}
