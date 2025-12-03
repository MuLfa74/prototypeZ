package requests

import "prototypeZ/database"

func GetRequestsByGame(gameID int) ([]Request, error) {
	rows, err := database.DB.Query(`
		SELECT requestid, gameid, userid, type, purpose, sex, age, contact, primetime, created_at, expires_at
		FROM Request
		WHERE gameid = ?`, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reqs []Request
	for rows.Next() {
		var r Request
		err := rows.Scan(&r.RequestID, &r.GameID, &r.UserID, &r.Type, &r.Purpose, &r.Sex, &r.Age, &r.Contact, &r.PrimeTime, &r.CreatedAt, &r.ExpiresAt)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, r)
	}
	return reqs, nil
}

func GetGameTitle(gameID int) (string, error) {
	var title string
	err := database.DB.QueryRow("SELECT title FROM Game WHERE id = ?", gameID).Scan(&title)
	if err != nil {
		return "", err
	}
	return title, nil
}
