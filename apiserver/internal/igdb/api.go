package igdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Game struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
}

type SearchResp struct {
	Games []Game
}

func SearchGames(ctx context.Context, t Token, h *http.Client, c Config, name string, limit int) ([]Game, error) {
	query := fmt.Sprintf("search \"%s\"; fields name,id;", name)
	if limit > 0 {
		query = fmt.Sprintf("%s limit %d;", query, limit)
	}
	rdr := strings.NewReader(query)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+c.GamesEndpoint, rdr)
	if err != nil {
		return []Game{}, err
	}
	token, err := t.GetToken(ctx)
	if err != nil {
		return []Game{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Client-ID", c.ClientID)

	resp, err := h.Do(req)
	if err != nil {
		return []Game{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			return []Game{}, fmt.Errorf("%w: error reading body: %s", ErrHTTP, err)
		}
		return []Game{}, fmt.Errorf("%w: statuscode: %v: %s", ErrHTTP, resp.StatusCode, string(bs))
	}
	sr := SearchResp{}
	if err = json.NewDecoder(resp.Body).Decode(&sr.Games); err != nil {
		return []Game{}, nil
	}

	return sr.Games, nil
}
