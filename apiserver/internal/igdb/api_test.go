package igdb_test

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/DFXLuna/apiserver/internal/igdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchGames(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	id := os.Getenv(CLIENT_ID_KEY)
	require.NotEmpty(id, "client id should not be empty")

	secret := os.Getenv(CLIENT_SECRET_KEY)
	require.NotEmpty(secret, "client secret should not be empty")

	tokenurlstr := os.Getenv(TOKEN_URL)
	require.NotEmpty(tokenurlstr, "tokenurl should not be empty")

	tokenurl, err := url.Parse(tokenurlstr)
	require.NoError(err, "should not err parsing url")

	tok := igdb.NewToken(tokenurl, func(req *http.Request) {
		q := req.URL.Query()
		q.Add("client_id", id)
		q.Add("client_secret", secret)
		req.URL.RawQuery = q.Encode()
	})

	c := igdb.NewConfig(id)

	games, err := igdb.SearchGames(context.Background(), tok, http.DefaultClient, c, "Halo")
	require.NoError(err, "should not error searching games")
	assert.NotEmpty(games, "games should not be empty")
}
