package web

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Parse_website(url string, selector string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("url_parse.go[Parse_url]: %v\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Fatalf("url_parse.go[Parse_url]: %v\n", err)
	}

	var text strings.Builder
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		str := strings.TrimSpace(s.Text())
		if str != "" {
			text.WriteString(str)
		}
	})

	return text.String()
}
