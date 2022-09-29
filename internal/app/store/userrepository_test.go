package store_test

import (
	"go_restapi/internal/app/model"
	"go_restapi/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
 // проверяем создание пользователя
	u, err := s.User().Create(&model.User{
		Email:             "user@example.org",
		EncryptedPassword: "password", //!TODO будет присваиваться автоматически, потом убрать
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	email := "user@example.com"
	//* 1) ищем несуществующего  пользователя. Должны получить ошибку
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	//* 2) создаём пользователя, а потом смотрим в базе есть ли он
	s.User().Create(&model.User{
		Email: "example@mail.com",
		EncryptedPassword: "password",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t ,err)
	assert.NotNil(t, u)
}
