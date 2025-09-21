package main

import (
	"fmt"
	"log"
	"network_modeling/internal/env"
	"network_modeling/internal/tree"
	"network_modeling/internal/web"
)

func main() {
	var url string
	fmt.Scan(&url)
	log.Printf("||\tPARSE WEBSITE\n")
	title := web.Parse_website(url, "t")
	log.Printf("||\tFIND SIMILAR POSTS\n")
	posts := web.Search_google(env.Google_api(), env.Google_ex(), title)
	log.Printf("||\tBUILDING GRAPH FROM POSTS\n")
	graph := tree.Build(posts)

	fmt.Println(graph)
}
