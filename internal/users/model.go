package users

import "errors"

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID        int64    `db:"id"`
	Login     string   `db:"login"`
	Sex       bool     `db:"sex"`
	Age       uint8    `db:"age"`
	Contact   string   `db:"contact"`
	PrimeTime string   `db:"prime_time"`
	Games     []string `db:"usergame"`
}
