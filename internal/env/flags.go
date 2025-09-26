package env

import "flag"

var (
	ga *string // 'google api' key
	ex *string // 'google search engine' id
	na *string // 'news api' key
	sa *string // 'sapling ai' key
)

func Get_flags() {
	ga = flag.String("ga", " ", "google api-key")
	ex = flag.String("ex", " ", "google search engine")
	na = flag.String("na", " ", "news api key")
	sa = flag.String("sa", " ", "sapling ai api-key")

	flag.Parse()
}

func Google_api() string {
	return *ga
}

func Google_ex() string {
	return *ex
}

func News_api() string {
	return *na
}

func Sapling_api() string {
	return *sa
}