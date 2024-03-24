package internal

import (
	"github.com/DFXLuna/apiserver/internal/igdb"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

type RouteConfig struct {
	IGDBToken  igdb.Token
	IGDBConfig igdb.Config
	Sanitizer  *bluemonday.Policy
}

func WithRouteConfig(f func(*gin.Context, RouteConfig), rc RouteConfig) func(*gin.Context) {
	return func(ctx *gin.Context) {
		f(ctx, rc)
	}
}

func SearchGames(c *gin.Context, rc RouteConfig) {

}
