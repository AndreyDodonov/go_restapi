package apiserver

import (
	"database/sql"
	"go_restapi/internal/app/store/sqlstore"
	"net/http"

	"github.com/gorilla/sessions"
)

// Start
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddress, srv)
}

// стартуем базу данных
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	// проверяем методом Ping подключение к базе. Потому что подключение по необходимости - lazy
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
