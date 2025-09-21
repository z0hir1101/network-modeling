package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type sapling_resp struct {
	Score float64 `json:"score"`
}

func Is_human(key, text string) bool {
	const url = "https://api.sapling.ai/api/v1/aidetect"

	req_body, _ := json.Marshal(map[string]interface{}{
		"key":  key,
		"text": text,
	})
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(req_body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("check.go[Is_human]: %v\n", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Fatalf("check.go[Is_human]: %d\n", resp.StatusCode)
	}

	var result sapling_resp
	json.Unmarshal(body, &result)
	return result.Score < 0.4
}