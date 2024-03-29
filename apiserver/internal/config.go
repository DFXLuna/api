package internal

import (
	"errors"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/ardanlabs/conf/v3/yaml"
)

type Config struct {
	IP   string `conf:"default:0.0.0.0,env:SERVER_IP"`
	Port string `conf:"default:39739,env:SERVER_PORT"`
	//IGDB Config
	IGDBClientID     string `conf:"required,env:IGDB_CLIENT_ID"`
	IGDBClientSecret string `conf:"required,env:IGDB_CLIENT_SECRET"`
	IGDBTokenURL     string `conf:"default:https://id.twitch.tv/oauth2/token,env:IGDB_TOKEN_URL"`
	IGDBBaseURL      string `conf:"default:https://api.igdb.com/v4/,env:IGDB_BASE_URL"`
}

const (
	configPath = "./config.yaml"
)

func NewConfig() (Config, error) {
	var cfg Config
	if bs, err := os.ReadFile(configPath); err == nil {
		_, err := conf.Parse("", &cfg, yaml.WithData(bs))
		if err != nil {
			return Config{}, err
		}
	} else if errors.Is(err, os.ErrNotExist) {
		_, err := conf.Parse("", &cfg)
		if err != nil {
			return Config{}, err
		}
	} else {
		return Config{}, err
	}
	return cfg, nil
}
