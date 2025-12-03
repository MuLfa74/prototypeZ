package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, u *User) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetByID(ctx context.Context, id int64) (*User, error) {
	user := &User{}
	var gamesJSON []byte

	query := `
        SELECT ID, Login, Sex, Age, Contact, Primetime, UserGame
        FROM Users
        WHERE ID = ?
    `

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Login, &user.Sex, &user.Age,
		&user.Contact, &user.PrimeTime, &gamesJSON,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	if len(gamesJSON) > 0 {
		json.Unmarshal(gamesJSON, &user.Games)
	}

	return user, nil
}

func (r *repository) Update(ctx context.Context, u *User) error {
	gamesJSON, _ := json.Marshal(u.Games)

	query := `
        UPDATE Users
        SET Sex = ?, Age = ?, Contact = ?, Prime-time = ?, UserGame = ?
        WHERE ID = ?
    `

	_, err := r.db.ExecContext(ctx, query,
		u.Sex, u.Age, u.Contact, u.PrimeTime, gamesJSON, u.ID,
	)

	return err
}
