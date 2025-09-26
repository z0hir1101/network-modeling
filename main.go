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
	env.Get_flags()
	fmt.Scan(&url)

	log.Printf("||\t* parse website\n")
	_, title := web.Parse_url(url, "h1")

	log.Printf("||\t* find similar sites\n")
	posts := web.Google_search(env.Google_api(), env.Google_ex(), title, 5)

	log.Printf("||\t* building network graph\n")
	graph := tree.Build(posts)
	tree.Graph_view(graph, 1)
}
