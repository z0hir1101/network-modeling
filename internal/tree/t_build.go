package tree

import (
	"fmt"
	"network_modeling/internal/web"

	"google.golang.org/api/customsearch/v1"
)

func Build(results []*customsearch.Result) []*t_node {
	t_nodes := make([]*t_node, 0, len(results))
	for _, result := range results {
		err, post := web.Parse_url(result.Link, "p")
		if err > 0 { continue }
		t_nodes = append(t_nodes, &t_node{
			url: 		result.DisplayLink,
			post:		post,
			org:     1,
			is_used: false,
		})
	}

	graph := make([]*t_node, 0, len(results))
	for i, tn := range t_nodes {
		if !tn.is_used {
			graph = append(graph, tn)
		}
		for _, item := range graph[i + 1:] {
			if item.is_used {	continue }
			parse_tnodes(tn, item)
		}
	}
	return graph
}

func Graph_view(graph []*t_node, id int) {
	fmt.Printf("|%d|->", id)
	for i, t_node := range graph {
		if i == 0 { continue }
		Graph_view(t_node.items, (i + 1) * 10)
	}
}
