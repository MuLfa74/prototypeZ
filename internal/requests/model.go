package requests

import "time"

type Request struct {
	RequestID int       `db:"request_id"`
	GameID    int       `db:"game_id"`
	UserID    int       `db:"user_id"`
	Type      bool      `db:"type"` // false — короткая, true — долгая
	Purpose   string    `db:"purpose"`
	Sex       bool      `db:"sex"` // false — муж, true — жен
	Age       uint8     `db:"age"`
	Contact   string    `db:"contact"`
	PrimeTime string    `db:"primetime"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
}
