package teststore

import (
	"go_restapi/internal/app/model"
	"go_restapi/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

// New
func New() *Store {
	return &Store{}
}


// Создаём пользователя. Пример:
// example: store.User().Create()
func (s *Store) User() store.UserRepository {

	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string] *model.User),
	}

	return s.userRepository
}
