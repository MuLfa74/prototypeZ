package requests

func GetRequestsForGame(gameID int) ([]Request, error) {
	return GetRequestsByGame(gameID)
}

func GetGameTitleByID(gameID int) (string, error) {
	return GetGameTitle(gameID)
}
