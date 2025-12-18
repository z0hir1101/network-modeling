package tests

import (
	"network_modeling/internal/tree"
	"slices"
)

var l []float64

func Check_graph(graph []*tree.T_node) float64 {
	read_tree(graph, 1)
	mx := slices.Max(l)
	mn := slices.Min(l)
	
	sum := 0.0
	for _, n := range l {
		sum += n
	}
	mean := sum / float64(len(l))

	return ((mx - mn) / 2) / mean * 100
}

func read_tree(graph []*tree.T_node, org float64) {
	for _, tn := range graph {
		l = append(l, tn.ORG / org)
		if len(tn.Items) > 0 {
			read_tree(tn.Items, tn.ORG)
		}
	}
}