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

// input edge set: e.g. [(1,3), (2,5), (3,2), (3,1)]
// [3][2]int{{1,3}, {3,1}}
func (udg *Graph) create(edges [][]int) {
	_ = filterEdgeArr(edges)
}

func (udg *Graph) setNodeSet() {

}

func (udg *Graph) setEdgeSet() {

}

func (udg *Graph) getNodeSet() {

}

func (udg *Graph) getEdgeSet() {

}

func (udg *Graph) getNeighborNodes() {

}

func (udg *Graph) getNeighborEdges() {

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
	for i, edge := range edges {
		if len(edge) < 2 {
			// delete insignificant data
			edges = append(edges[:i], edges[i+1:]...)
		}
	}
	for i, edge := range edges {
		for j := i + 1; j < len(edges); j++ {
			if (edge[0] == edges[j][0] && edge[1] == edges[j][1]) || (edge[0] == edges[j][1] && edge[1] == edges[j][0]) {
				// delete edges[j], due to the repeated with previous
				edges = append(edges[:j], edges[j+1:]...)
			}
		}
	}
	return edges
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
