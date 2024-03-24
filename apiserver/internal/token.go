package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Token struct {
	h         *http.Client
	tokenUrl  *url.URL
	token     string
	expiresAt time.Time
	// Used to add authentication to the HTTP request used to fulfill the token
	af AuthFunction
	m  *sync.Mutex
}

const (
	defaultTimeout       = 10 * time.Second
	defaultRefreshMargin = 5 * time.Minute
)

var (
	ErrHTTP = fmt.Errorf("Error making http call to token service")
)

type AuthFunction func(req *http.Request)

type tokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// Autorefreshing token
// Sets grant_type=client_credentials
func NewToken(tokenURL *url.URL, af AuthFunction) Token {
	h := http.Client{
		Timeout: defaultTimeout,
	}

	q := tokenURL.Query()
	q.Add("grant_type", "client_credentials")
	tokenURL.RawQuery = q.Encode()

	return Token{
		h:         &h,
		expiresAt: time.Now(),
		tokenUrl:  tokenURL,
		af:        af,
		m:         &sync.Mutex{},
	}
}

func (t *Token) GetToken(ctx context.Context) (string, error) {
	t.m.Lock()
	defer t.m.Unlock()
	if time.Now().After(t.expiresAt.Add(-defaultRefreshMargin)) {
		err := t.refreshToken(ctx)
		if err != nil {
			return "", err
		}
	}
	return t.token, nil
}

func (t *Token) refreshToken(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.tokenUrl.String(), nil)
	if err != nil {
		return err
	}
	t.af(req)

	resp, err := t.h.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("%w: error reading body: %s", ErrHTTP, err)
		}
		return fmt.Errorf("%w: statuscode: %v: %s", ErrHTTP, resp.StatusCode, string(bs))
	}

	tr := tokenResp{}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		return err
	}
	// not 100% accurate, should read HTTP date header to find correct expiration
	newExp := time.Now().Add(time.Duration(time.Duration(tr.ExpiresIn*1000) * time.Second))
	t.expiresAt = newExp
	t.token = tr.AccessToken

	return nil
}
