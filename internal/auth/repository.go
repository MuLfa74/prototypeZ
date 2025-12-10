package auth

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type User struct {
	ID       int64
	login    string
	password string
}

func (r *Repository) CreateUser(login, password string) error {
	_, err := r.db.Exec(`
        INSERT INTO users (login, password)
        VALUES (?, ?)
    `, login, password)
	return err
}

func (r *Repository) GetByEmail(login string) (*User, error) {
	row := r.db.QueryRow(`
        SELECT id, login, password
        FROM users
        WHERE login = ?
    `, login)

	var u User
	err := row.Scan(&u.ID, &u.login, &u.password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
