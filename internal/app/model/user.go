package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

// Валидируем структуру User
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		 validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(3,10)),
		)
}

// проверяем не пустой ли пароль и зашифровываем пароль
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

// шифруем пароль. MinCost - слабое шифрование.
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}
