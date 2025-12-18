package web

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Parse_url(url string, selector string) (int, string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("url_parse.go[Parse_url]: %v\n", err)
		return 1, ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("url_parse.go[Parse_url]: %d\n", resp.StatusCode)
		return 1, ""
	}

	var text strings.Builder
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		str := strings.TrimSpace(s.Text())
		if str != "" {
			text.WriteString(str)
		}
	})

	return 0, text.String()
}
