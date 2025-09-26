package tree

import (
	"network_modeling/internal/alg"
)

type t_node struct {
	url     string
	post    string
	items   []*t_node
	is_used bool
	org     float32
}

func parse_tnodes(tn1, tn2 *t_node) {
	simil := alg.Shingle_alg(tn1.post, tn2.post, 3)

	if simil >= 0.0 {
		tn1.items = append(tn1.items, tn2)
		tn2.org = tn1.org * simil
		tn2.is_used = true
	}
}
