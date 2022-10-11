package sqlstore_test

import (
	"go_restapi/internal/app/model"
	"go_restapi/internal/app/store"
	"go_restapi/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u:= model.TestUser(t)
	// проверяем создание пользователя
	assert.NoError(t,  s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) { //FIXME в тестах что то не то с подключением к базе, хотя подключение есть - записи делаются
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	u1 := model.TestUser(t)
	s := sqlstore.New(db)
	email := "usr@example.com"
	//* 1) ищем несуществующего  пользователя. Должны получить ошибку
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	//* 2) создаём пользователя, а потом ищем в базе по емейлу

	//u.Email = email
	s.User().Create(u1)
	u2, err := s.User().FindByEmail(u1.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_Find(t *testing.T) { //FIXME в тестах что то не то с подключением к базе, хотя подключение есть - записи делаются
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
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