package structure

import "fmt"

type UndirectedEdge struct {
	one, another int // the two nodes of an edge
}

type DirectedArc struct {
	from, to int // the two nodes of an arc
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
