package teststore_test

import (
	"go_restapi/internal/app/model"
	"go_restapi/internal/app/store"
	"go_restapi/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	// проверяем создание пользователя
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "us@examp.com"
	//* 1) ищем несуществующего  пользователя. Должны получить ошибку
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	//* 2) создаём пользователя, а потом ищем в базе по емейлу
	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) { //FIXME в тестах что то не то с подключением к базе, хотя подключение есть - записи делаются

	s := teststore.New()
	// email := "usr@example.com"
	//* 1) ищем несуществующего  пользователя. Должны получить ошибку
	// _, err := s.User().FindByEmail(email)
	// assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	//* 2) создаём пользователя, а потом ищем в базе по емейлу
	u1 := model.TestUser(t)
	// u.Email = email
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
