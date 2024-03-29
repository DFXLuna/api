package igdb_test

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/DFXLuna/apiserver/internal/igdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetToken(t *testing.T) {
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

	tok := igdb.NewToken(tokenurl, igdb.DefaultAuthFunction(id, secret))

	str, err := tok.GetToken(context.Background())
	require.NoError(err, "should not error getting token")
	assert.NotEmpty(str, "token should not be empty")
}
