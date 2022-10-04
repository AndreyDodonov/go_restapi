package apiserver

import (
	"database/sql"
	"fmt"
	"go_restapi/internal/app/store/sqlstore"
	"net/http"
)

// Start
func Start(config *Config) error  {
	fmt.Println(config.DatabaseURL) //TODO debug
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		fmt.Println("error in new db") //TODO debug
		return err
	}
	defer db.Close()
	fmt.Println("new sql store") //TODO debug
	store := sqlstore.New(db)
	srv := newServer(store)

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