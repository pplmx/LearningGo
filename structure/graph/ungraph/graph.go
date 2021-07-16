package ungraph

import "fmt"

type EdgeData int   // Edge data, e.g. weight
type VertexData int // Vertex data

type Edge struct {
	isVisited            bool  // mark that whether an edge is visited
	uIdx, vIdx           int   // the index of two vertex of an edge in adj_multi_list
	uNextEdge, vNextEdge *Edge // Respectively point to the next edge attached to these two vertices
	data                 EdgeData
}

type Vertex struct {
	data      VertexData
	firstEdge *Edge
}

// Graph Undirected graph implemented with Adjacency Multi-list
type Graph struct {
	vertexNum, EdgeNum int
	adjMultiList       *Vertex
}

// UnGraph a undirected graph object
type UnGraph struct {
	g         *Graph
	vertexSet []int
	edgeSet   [][]int
	degree    int
}

// Create a function to create a undirected graph
// input edge set:
// e.g.
// [][]int{{1, 3}, {3, 1}}
// [][]int{{1, 3, 2}, {3, 1, 5}}
func (udg UnGraph) Create(edges [][]int) UnGraph {
	udg.setVertexSet(edges)
	return udg
}

func (udg UnGraph) setVertexSet(edges [][]int) UnGraph {
	// Firstly, set edges
	udg.setEdgeSet(edges)
	var vs []int
	for _, edge := range udg.edgeSet {
		vs = append(vs, edge[0])
		vs = append(vs, edge[1])
	}
	filteredVertices := filterVertices(vs)
	udg.vertexSet = filteredVertices
	return udg
}

func (udg UnGraph) setEdgeSet(edges [][]int) UnGraph {
	filteredEdges := filterEdges(edges)
	udg.edgeSet = filteredEdges
	return udg
}

func (udg UnGraph) GetVertexSet() []int {
	return udg.vertexSet
}

func (udg UnGraph) GetEdgeSet() [][]int {
	return udg.edgeSet
}

func (udg UnGraph) GetNeighborVertices() {

}

func (udg UnGraph) GetNeighborEdges() {

}

func (udg UnGraph) GetDegree() int {

}

func filterEdges(edges [][]int) [][]int {
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

func filterVertices(vertices []int) []int {
	// create a map: {int: true}
	existed := make(map[int]bool)
	var filtered []int
	// If the key(values of the slice) is not equal
	// to the already present value in new slice (filtered)
	// then we append it. else we jump on another element.
	for _, vertex := range vertices {
		if _, value := existed[vertex]; !value {
			existed[vertex] = true
			filtered = append(filtered, vertex)
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
	filtered := filterEdges(multiDimensionArr)
	fmt.Printf("4: %v\n", multiDimensionArr)
	fmt.Printf("5: %v\n", filtered)
}
