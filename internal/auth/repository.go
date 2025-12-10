package auth

import (
	"database/sql"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type User struct {
	ID       int64
	Login    string
	Password string
}

// Вставка пользователя
func (r *Repository) CreateUser(login, password string) error {
	res, err := r.db.Exec(
		"INSERT INTO Users (Login, Password) VALUES (?, ?)",
		login, password,
	)
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}
	id, _ := res.LastInsertId()
	log.Println("User created with ID:", id)
	return nil
}

// Получаем пользователя по логину
func (r *Repository) GetByLogin(login string) (*User, error) {
	row := r.db.QueryRow(
		"SELECT id, Login, Password FROM Users WHERE Login = ?",
		login,
	)

	var u User
	err := row.Scan(&u.ID, &u.Login, &u.Password)
	if err != nil {
		log.Println("Error fetching user:", err)
		return nil, err
	}
	return &u, nil
}
