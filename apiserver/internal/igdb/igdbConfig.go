package igdb

type Config struct {
	ClientID      string
	BaseURL       string
	GamesEndpoint string
}

func NewConfig(clientID string, BaseURL string) Config {
	return Config{
		ClientID:      clientID,
		BaseURL:       BaseURL,
		GamesEndpoint: "games/",
	}
}
