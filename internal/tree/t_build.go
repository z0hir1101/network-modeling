package tree

import (
	"network_modeling/internal/web"

	"google.golang.org/api/customsearch/v1"
)

func Build(results []*customsearch.Result) []*t_node {
	t_nodes := make([]*t_node, 0, len(results))
	for _, result := range results {
		t_nodes = append(t_nodes, &t_node{
			url: 		result.DisplayLink,
			post:		web.Parse_website(result.DisplayLink, "d"),
			org:     1,
			is_used: false,
		})
	}

	graph := make([]*t_node, 0, len(results))
	for i, tn := range t_nodes {
		if !tn.is_used {
			graph = append(graph, tn)
		}
		parse_slice(tn, t_nodes[i:])
	}

	return graph
}

func parse_slice(tn *t_node, slice []*t_node) {
	for _, item := range slice {
		if item.is_used {	continue }
		parse_tnodes(tn, item)
	}
}
