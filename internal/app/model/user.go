package model

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

// Валидируем структуру User.
// Проверяем наличие емейл и пароля, длину пароля и валидность емейла
func (u *User) Validate() error {
	fmt.Println("validation is working, u = ", u) //TODO debug
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(3, 10)),
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

// метод который будет затирать данные, которые должны быть недоступны
func (u *User) Sanitize()  {
	fmt.Println("Sanitize is working, u = ", u) //TODO debug
	u.Password =""
}

// шифруем пароль. MinCost - слабое шифрование.
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}
