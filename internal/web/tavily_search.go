package web

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type TavilyRequest struct {
	APIKey      string `json:"api_key"`
	Query       string `json:"query"`
	StartDate   string `json:"start_date"`   // e.g., "2024-01-01"
	EndDate     string `json:"end_date"`     // e.g., "2025-11-17"
	SearchDepth string `json:"search_depth"` // "basic" or "advanced"
}

type TavilyResponse struct {
	Query        string         `json:"query"`
	Answer       string         `json:"answer,omitempty"`
	Results      []*TavilyResult `json:"results"`
	ResponseTime float64        `json:"response_time"`
}

type TavilyResult struct {
	Title      string  `json:"title"`
	URL        string  `json:"url"`
	Content    string  `json:"content"`
	Score      float64 `json:"score"`
	RawContent *string `json:"raw_content,omitempty"`
	Favicon    string  `json:"favicon,omitempty"`
}

func Tavily_search(key, query string) []*TavilyResult {
	reqBody := TavilyRequest{
		APIKey:      key,
		Query:       query,
		SearchDepth: "basic",
	}
	jsonBody, _ := json.Marshal(reqBody)

	resp, err := http.Post("https://api.tavily.com/search", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("tavily_search.go[Tavily_search]: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var tavily_resp TavilyResponse
	err = json.Unmarshal(body, &tavily_resp)

	return tavily_resp.Results
}
