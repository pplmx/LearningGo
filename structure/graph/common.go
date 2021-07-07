package graph

import "fmt"

type UndirectedEdge struct {
	one, another int // the two nodes of an edge
}

type DirectedArc struct {
	from, to int // the two nodes of an arc
}

func filterEdges(edges []UndirectedEdge) []UndirectedEdge {
	keys := make(map[UndirectedEdge]bool)
	var filtered []UndirectedEdge
	for _, edge := range edges {
		if _, value := keys[edge]; !value {
			keys[edge] = true
			filtered = append(filtered, edge)
		}
	}
	return filtered
}

func filterArcs(arcs []DirectedArc) []DirectedArc {
	keys := make(map[DirectedArc]bool)
	var filtered []DirectedArc
	// If the key(values of the slice) is not equal
	// to the already present value in new slice (filtered)
	// then we append it. else we jump on another element.
	for _, arc := range arcs {
		if _, value := keys[arc]; !value {
			keys[arc] = true
			filtered = append(filtered, arc)
		}
	}
	return filtered
}

func printEdges(edges []UndirectedEdge) {
	for _, edge := range edges {
		fmt.Printf("(%d, %d)\n", edge.one, edge.another)
	}
}

func printArcs(arcs []DirectedArc) {
	for _, arc := range arcs {
		fmt.Printf("(%d -> %d)\n", arc.from, arc.to)
	}
}
