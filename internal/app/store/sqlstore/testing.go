package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
	"fmt"
)

// Test store
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	//fmt.Print("db is: ", databaseURL)
	if err != nil {
		fmt.Println("error in sql open: ", err)
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("error in db Ping: ", err)
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}
		db.Close() //TODO add defer?
	}
}
