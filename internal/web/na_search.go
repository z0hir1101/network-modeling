package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Article struct {
    Author      string   `json:"author"`
    Title       string    `json:"title"`
    Description string   `json:"description"`
    URL         string    `json:"url"`
    PublishedAt time.Time `json:"publishedAt"`
}

type NewsAPIResponse struct {
    Status       string    `json:"status"`
    TotalResults int       `json:"totalResults"`
    Articles     []Article `json:"articles"`
}

func Newsapi_search (key, query string) []Article { 
	baseURL := "https://newsapi.org/v2/everything/?"
	u, _ := url.Parse(baseURL)
	q := u.Query()
	q.Set("q", query)
	q.Set("apiKey", key)
	q.Set("sortBy", "publishedAt")
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
	return result.Articles
}
