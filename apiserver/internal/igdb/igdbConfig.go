package igdb

type Config struct {
	ClientID      string
	BaseURL       string
	GamesEndpoint string
}

func NewConfig(clientID string) Config {
	return Config{
		ClientID:      clientID,
		BaseURL:       "https://api.igdb.com/v4/",
		GamesEndpoint: "games/",
	}
}
