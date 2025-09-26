package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type newsapi_resp struct {
	TotalResults int `json:"totalResults"`
}

func Is_official(key, q string) bool { // check request in some official domains
	q = strings.ReplaceAll(q, " ", "-")
	const url = "https://newsapi.org/v2/everything/?"
	params := fmt.Sprintf("?q=%s&domains=bbc.com&language=en&apiKey=%s", q, key)
	resp, err := http.Get(url + params)
	if err != nil {
		log.Fatalf("check.go[Is_official]: %v\n", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Fatalf("check.go[Is_official]: %d\n", resp.StatusCode)
	}

	var result newsapi_resp
	json.Unmarshal(body, &result)
	return result.TotalResults > 0
}
