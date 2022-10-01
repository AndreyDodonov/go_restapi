package apiserver

import (
	"database/sql"
	"go_restapi/internal/app/store/sqlstore"
	"net/http"
)

// Start
func Start(config *Config) error  {
	db, err := newDB(config.databaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddress, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}