package graph

import "fmt"

type EdgeData int // Edge data, e.g. weight
type NodeData int // Node data

type Edge struct {
	isVisited            bool  // mark that whether an edge is visited
	uIdx, vIdx           int   // the index of two node of an edge in adj_multi_list
	uNextEdge, vNextEdge *Edge // Respectively point to the next edge attached to these two nodes
	data                 EdgeData
}

type Node struct {
	data      NodeData
	firstEdge *Edge
}

// Graph Undirected graph implemented with Adjacency Multi-list
type Graph struct {
	nodeNum, EdgeNum int
	adjMultiList     *Node
}

// input edge set:
// e.g.
// [][]int{{1, 3}, {3, 1}}
// [][]int{{1, 3, 2}, {3, 1, 5}}
func (udg Graph) create(edges [][]int) Graph {
	filtered := filterEdgeArr(edges)
	for _, edge := range filtered {

	}
	return udg
}

func (udg Graph) setNodeSet(edges [][]int) {

}

func (udg Graph) setEdgeSet() {

}

func (udg Graph) getNodeSet() {

}

func (udg Graph) getEdgeSet() {

}

func (udg Graph) getNeighborNodes() {

}

func (udg Graph) getNeighborEdges() {

}

func filterEdgeArr(edges [][]int) [][]int {
	//edges := [][]int{
	//    {1, 2, 3},
	//    {2, 3, 10},
	//    {1, 3, 5},
	//    {4, 3, 10},
	//    {3, 4, 15}, // will be deleted
	//    {1, 2, 15}, // will be deleted
	//    {4, 5},
	//    {1, 2}, // will be deleted
	//    {1},    // will be deleted
	//}
	var filtered [][]int
	for _, edge := range edges {
		if len(edge) > 1 && len(edge) < 4 {
			// filter insignificant data
			filtered = append(filtered, edge)
		}
	}
	for i, edge := range filtered {
		for j := i + 1; j < len(filtered); j++ {
			if (edge[0] == filtered[j][0] && edge[1] == filtered[j][1]) ||
				(edge[0] == filtered[j][1] && edge[1] == filtered[j][0]) {
				// delete filtered[j], due to the repeated with previous
				filtered = append(filtered[:j], filtered[j+1:]...)
			}
		}
	}
	return filtered
}

func TestUndirectedGraph() {
	multiDimensionArr := [][]int{
		{1, 2, 3},
		{2, 3, 10},
		{1, 3, 5},
		{4, 3, 10},
		{3, 4, 15},
		{4, 5},
		{1, 2},
		{1},
	}
	fmt.Printf("1: %v\n", multiDimensionArr)
	filtered := filterEdgeArr(multiDimensionArr)
	fmt.Printf("4: %v\n", multiDimensionArr)
	fmt.Printf("5: %v\n", filtered)
}
