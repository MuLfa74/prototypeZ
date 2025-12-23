package games

// Service слой может содержать бизнес-логику.
// Пока простая передача данных из репозитория.
func GetGamesList() ([]Game, error) {
	return GetAllGames()
}
