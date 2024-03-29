package internal

import (
	"context"
	"fmt"
	"net/url"

	"github.com/DFXLuna/apiserver/internal/igdb"
	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context, errch chan error) {
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		errch <- err
		return
	}
	conf, err := NewConfig()
	if err != nil {
		errch <- err
		return
	}
	tokenURL, err := url.Parse(conf.IGDBTokenURL)
	if err != nil {
		errch <- err
		return
	}
	t := igdb.NewToken(tokenURL, igdb.DefaultAuthFunction(conf.IGDBClientID, conf.IGDBClientSecret))
	iconf := igdb.NewConfig(conf.IGDBClientID, conf.IGDBBaseURL)
	rc := NewRouteConfig(t, iconf)

	// Routes
	router.GET("/v1/search", WithRouteConfig(SearchGames, rc))

	errch <- router.Run(fmt.Sprintf("%s:%s", conf.IP, conf.Port))
}
