package teststore_test

import (
	"go_restapi/internal/app/model"
	"go_restapi/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u:= model.TestUser(t)
	// проверяем создание пользователя
	assert.NoError(t,  s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "usr@example.com"
	//* 1) ищем несуществующего  пользователя. Должны получить ошибку
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	//* 2) создаём пользователя, а потом ищем в базе по емейлу
	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
