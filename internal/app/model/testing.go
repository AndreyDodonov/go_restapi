package model

import "testing"

// хелпер для теста. Для того, чтобы не инициализировать структуру каждый раз
func TestUser(t *testing.T) *User {
	return &User{
		Email: "us@examp.com",
		Password: "Pass-word1",
	}
}