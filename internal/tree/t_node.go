package tree

import (
	"fmt"
	"network_modeling/internal/alg"
)

type T_node struct {
	URL     string
	Post    string
	Items   []*T_node
	Is_used bool
	ORG     float64
}

func parse_tnodes(tn1, tn2 *T_node) {
	simil := alg.Shingle_alg(tn1.Post, tn2.Post, 3)
	fmt.Println(simil)
	if simil >= 0.6 {
		tn1.Items = append(tn1.Items, tn2)
		tn2.ORG = tn1.ORG * simil
		tn2.Is_used = true
	}
}
