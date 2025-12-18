package main

import (
	"fmt"
	"log"

	"network_modeling/internal/env"
	"network_modeling/internal/tests"
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
	articles := web.Google_search(env.Google_api(), env.Google_ex(), title)

	log.Printf("||\t* similar sites:\n")
	for _, article := range articles {
		fmt.Println(article.Link)
	}

	log.Printf("||\t* building network graph\n")
	graph := tree.Build(articles)
	tree.Tree_view(graph, "")

	log.Printf("||\t* checking network graph\n")
	t1 := tests.Check_graph(graph) < 5

	log.Printf("||\t* checking for official")
	t2 := tests.Is_official(env.News_api(), title)

	log.Printf("||\t* checking for human")
	_, text := web.Parse_url(url, "p")
	t3 := tests.Is_human(env.Sapling_api(), text)
	fmt.Println("RESULTS:")	

	fmt.Println("NETWORK CHECK:\t\t", t1)
	fmt.Println("FOR OFFICIAL CHECK:\t", t2)
	fmt.Println("FOR HUMAN CHECK:\t", t3)
}