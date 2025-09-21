package env

import (
	"flag"
	"log"
)

func News_api() string {
	api := *flag.String("ga", "", "google api-key")
	if api == "" {
		log.Fatalf("google.go[Google_api]: google api-key\n")
	}
	return api
}
