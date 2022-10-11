package sqlstore

import (
	"database/sql"
	"fmt"
	"go_restapi/internal/app/model"
	"go_restapi/internal/app/store"
)

// user repository
type UserRepository struct {
	store *Store // ссылка на главное хранилище
}

// Create
func (r *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// Get users
func (r *UserRepository) Get() ([]string, error) {

	u := &model.User{}
	rows, err := r.store.db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Query error (userRepository: 35) : ", err)
		return nil, err
	}
	fmt.Println("Rows : ", rows)
	 rows.Scan(&u.ID, &u.Email, &u.EncryptedPassword)
	 users := make([]string, 0)
	 for rows.Next() {
		// var user string
		if err := rows.Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
			fmt.Println("error in userRepository 47: ", err)
		}
		users = append(users, u.Email)
		fmt.Printf("user is - %s \n", u.Email)
	 }
	return users, nil
}

// Find user by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, encrypted_password FROM users WHERE email =$1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

// Find users by id
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, encrypted_password FROM users WHERE id =$1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
