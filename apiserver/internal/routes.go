package internal

import (
	"net/http"
	"strconv"
	"time"

	"github.com/DFXLuna/apiserver/internal/igdb"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

type RouteConfig struct {
	IGDBToken  igdb.Token
	IGDBConfig igdb.Config
	Sanitizer  *bluemonday.Policy
	H          *http.Client
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SearchResponse struct {
	Games []igdb.Game `json:"games"`
}

const (
	searchParameter = "q"
	limitParameter  = "l"
	defaultTimeout  = 10 * time.Second
)

func NewRouteConfig(t igdb.Token, c igdb.Config) RouteConfig {
	return RouteConfig{
		IGDBToken:  t,
		IGDBConfig: c,
		Sanitizer:  bluemonday.StrictPolicy(),
		H:          &http.Client{Timeout: defaultTimeout},
	}
}

func WithRouteConfig(f func(*gin.Context, RouteConfig), rc RouteConfig) func(*gin.Context) {
	return func(ctx *gin.Context) {
		f(ctx, rc)
	}
}

func SearchGames(c *gin.Context, rc RouteConfig) {
	searchterm := c.Query(searchParameter)
	if searchterm == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "no search paramater"})
		return
	}

	limitStr := c.Query(limitParameter)
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 0
	}
	searchterm = rc.Sanitizer.Sanitize(searchterm)
	games, err := igdb.SearchGames(c.Request.Context(), rc.IGDBToken, rc.H, rc.IGDBConfig, searchterm, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	status := http.StatusOK
	if len(games) == 0 {
		status = http.StatusNoContent
	}
	c.JSON(status, SearchResponse{Games: games})
}
