package utils

import "fmt"

func printEdges(edges [][]int) {
	for _, edge := range edges {
		fmt.Printf("(%d, %d)\n", edge[0], edge[1])
	}
}

func printArcs(arcs [][]int) {
	for _, arc := range arcs {
		fmt.Printf("(%d -> %d)\n", arc[0], arc[1])
	}
}
