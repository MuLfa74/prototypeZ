package games

import (
	"prototypeZ/database"
)

func GetAllGames() ([]Game, error) {
	rows, err := database.DB.Query("SELECT id, title, genre, rating FROM Game")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []Game
	for rows.Next() {
		var g Game
		err := rows.Scan(&g.ID, &g.Title, &g.Genre, &g.Rating)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}
