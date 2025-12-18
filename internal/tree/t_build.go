package tree

import (
	"fmt"
	"network_modeling/internal/web"
	"strconv"

	"google.golang.org/api/customsearch/v1"
)

func Build(results []*customsearch.Result) []*T_node {
	t_nodes := make([]*T_node, 0, len(results))
	for _, result := range results {
		err, post := web.Parse_url(result.Link, "p")
		if err > 0 { continue }
		fmt.Println(result.Link)
		t_nodes = append(t_nodes, &T_node{
			URL: 		result.Link,
			Post:		post,
			ORG:     	1,
			Is_used: 	false,
		})
	}
	
	tree := make([]*T_node, 0, len(results))
	for i, tn := range t_nodes {
		if !tn.Is_used {
			tree = append(tree, tn)
		}
		for _, item := range t_nodes[i + 1:] {
			if item.Is_used { continue }
			for range 20 {
				fmt.Printf("_")
			}
			fmt.Printf("\n")
			
			fmt.Println(tn.URL)
			fmt.Println(item.URL)
			parse_tnodes(tn, item)
		}
	}
	for range 20 {
		fmt.Printf("_")
	}	
	fmt.Println()
	return tree
}

func Tree_view(graph []*T_node, id string) {
	for i, t_node := range graph {
		fmt.Printf("|%s|=", id + strconv.Itoa(i + 1))
		if len(t_node.Items) > 0 {
			Tree_view(t_node.Items, (id + strconv.Itoa(i + 1)))
		}
	}
}
