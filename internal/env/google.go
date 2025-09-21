package env

import (
	"flag"
	"log"
)

func Google_api() string {
	api := *flag.String("ga", "", "google api-key")
	if api == "" {
		log.Fatalf("google.go[Google_api]: no google api-key\n")
	}
	return api
}

func Google_ex() string {
	api := *flag.String("ex", "", "google api-key")
	if api == "" {
		log.Fatalf("google.go[Google_api]: no google search engine\n")
	}
	return api
}
