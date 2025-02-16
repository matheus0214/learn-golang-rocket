package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Result struct {
	Search       []SearchResult `json:"search"`
	TotalResults string         `json:"total_results"`
	Response     string         `json:"response"`
}

type SearchResult struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

func Search(apiKey, title string) (Result, error) {
	v := url.Values{}

	v.Set("apikey", apiKey)
	v.Set("s", title)

	resp, err := http.Get("http://www.omdbapi.com/?" + v.Encode())
	if err != nil {
		return Result{}, fmt.Errorf("failed to make omdb request: %w", err)
	}

	defer resp.Body.Close()

	var res Result

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return Result{}, fmt.Errorf("failed to decode omdb response: %w", err)
	}

	return res, nil
}
