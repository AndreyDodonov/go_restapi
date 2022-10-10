package apiserver

import (
	"bytes"
	"encoding/json"
	"go_restapi/internal/app/model"
	"go_restapi/internal/app/store/teststore"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

// тестируем роут создания юзера (шаблон для POST)
func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name       string
		payload    interface{}
		expectCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.com",
				"password": "Pass-word1",
			},
			expectCode: http.StatusCreated, //FIXME test crushed, but function work ok
		},
		{
			name:       "invalid payload",
			payload:    "invalid",
			expectCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectCode: http.StatusUnprocessableEntity, //FIXME test crushed, but function work ok
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectCode, rec.Code)
		})
	}

}

// тестируем главный роут (шаблон для GET)
func TestServer_HandleMain(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Server is working \n Main router")
	// TODO: поменять тест. Наверное должен ожидаться statusOk или что-то из html
}

func TestServerSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	s := newServer(store)
	store.User().Create(u)
	testCases := []struct { // testcases
		name       string
		payload    interface{}
		expectCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectCode: http.StatusOK,
		},
		{
			name:       "invalid payload",
			payload:    "invalid",
			expectCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectCode, rec.Code)
		})
	}
}
