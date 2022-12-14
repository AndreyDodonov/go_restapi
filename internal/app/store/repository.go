/**
 * Интерефейсы репозиториев
 */

package store

import "go_restapi/internal/app/model"

// user repository
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	Get() ([]string, error)
	Find(int) (*model.User, error)
}

