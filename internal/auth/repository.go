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
	ID           int64
	Email        string
	PasswordHash string
}

func (r *Repository) CreateUser(email, passwordHash string) error {
	_, err := r.db.Exec(`
        INSERT INTO users (login, password)
        VALUES (?, ?)
    `, email, passwordHash)
	return err
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	row := r.db.QueryRow(`
        SELECT id, email, password_hash
        FROM users
        WHERE email = ?
    `, email)

	var u User
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
