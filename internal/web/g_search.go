package web

import (
	"context"
	"log"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func Google_search(key, ex, query string, mx int64) []*customsearch.Result {
	service, err := customsearch.NewService(context.Background(), option.WithAPIKey(key))
	if err != nil {
		log.Fatalf("g_search.go[Search1_google]: %v", err)
	}

	call := service.Cse.List().
		Cx(ex).
		Q(query).
		Num(mx)

	resp, err := call.Do()
	if err != nil {
		log.Fatalf("g_search.go[Search_google]: %v", err)
	}
	return resp.Items
}
