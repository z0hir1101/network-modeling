package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type NewsAPIResponse struct {
	TotalResults int `json:"totalResults"`
}

func Is_official(key, query string) bool { // check in official websites
	baseURL := "https://newsapi.org/v2/everything/?"
	u, _ := url.Parse(baseURL)
	q := u.Query()
	q.Set("q", query)
	q.Set("apiKey", key)
	q.Set("language", "en")

	u.RawQuery = q.Encode()

	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalf("na_search.go[Newsapi_search]: %v\n", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("na_search.go[Newsapi_search]: %d\n", resp.StatusCode)
	}

	var result NewsAPIResponse
	json.Unmarshal(body, &result)
	return result.TotalResults > 0
}
