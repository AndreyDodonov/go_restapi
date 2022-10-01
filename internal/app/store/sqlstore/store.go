package sqlstore

import (
	"database/sql"
	"go_restapi/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}


// Создаём пользователя. Пример:
// example: store.User().Create()
func (s *Store) User() store.UserRepository {

	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
