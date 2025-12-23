package games

type Game struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Genre  string `db:"genre"`
	Rating bool   `db:"rating"`
}
