package ungraph

import "fmt"

type EdgeInfo int           // EdgeInfo, e.g. weight
type VertexType interface{} // VertexType, e.g. 1, 2, 3 or A, B, C

type Edge struct {
	isVisited            bool  // mark that whether the current edge is visited
	uIdx, vIdx           int   // the index of two vertices of the current edge in adj_multi_list
	uNextEdge, vNextEdge *Edge // Respectively point to the next edge attached to these two vertices
	info                 EdgeInfo
}

type Vertex struct {
	data      VertexType
	firstEdge *Edge
}

// Graph Undirected graph implemented with Adjacency Multi-list
type Graph struct {
	vertexNum, edgeNum int
	adjMultiList       []Vertex
}

// UnGraph an undirected graph "class"
type UnGraph struct {
	g         *Graph
	vertexSet []VertexType
	edgeSet   [][]VertexType
	degree    int
}

// Create a function to create a undirected graph
// input edge set:
// e.g.
// [][]VertexType{{1, 3}, {3, 1}}
// [][]VertexType{{1, 3, 2}, {3, 1, 5}}
func (udg UnGraph) Create(edges [][]VertexType) UnGraph {
	udg.setVertexSet(edges)

	// init a undirected graph
	graph := Graph{
		vertexNum:    len(udg.vertexSet),
		edgeNum:      len(udg.edgeSet),
		adjMultiList: nil,
	}

	// init all vertices in adjacency multi-list
	graph.adjMultiList = make([]Vertex, graph.vertexNum)
	for i, vertex := range udg.vertexSet {
		graph.adjMultiList[i].data = vertex
		graph.adjMultiList[i].firstEdge = nil
	}

	// init all edges
	for _, edge := range udg.edgeSet {
		var e Edge
		e.uIdx = udg.locateVertex(edge[0])
		e.vIdx = udg.locateVertex(edge[1])
		e.isVisited = false
		e.uNextEdge = udg.g.adjMultiList[e.uIdx].firstEdge
		e.vNextEdge = udg.g.adjMultiList[e.vIdx].firstEdge
		udg.g.adjMultiList[e.uIdx].firstEdge = &e
		udg.g.adjMultiList[e.vIdx].firstEdge = &e
	}

	udg.g = &graph
	return udg
}

func (udg UnGraph) GetNeighborVertices(v VertexType) []VertexType {
	idx := udg.locateVertex(v)
	tmp := udg.g.adjMultiList[idx].firstEdge
	for tmp != nil {
		if tmp.uIdx == idx {

		}
	}
	return nil
}

func (udg UnGraph) GetNeighborEdges(v VertexType) [][]VertexType {
	return nil
}

// BFS Breadth-First-Search
// Queue: enqueue, dequeue
func (udg UnGraph) BFS() []VertexType {
	return nil
}

// DFS Breadth-First-Search
// Stack: push, pop
func (udg UnGraph) DFS(v VertexType) {
	idx := udg.locateVertex(v)
	visited := make(map[VertexType]bool)
	visited[v] = true
	tmp := udg.g.adjMultiList[idx].firstEdge
	if tmp != nil {
		udg.DFS(udg.g.adjMultiList[tmp.vIdx].data)
	}
}

func (udg UnGraph) GetDegree(node VertexType) int {
	return 0
}

func (udg UnGraph) GetVertexSet() []VertexType {
	return udg.vertexSet
}

func (udg UnGraph) GetEdgeSet() [][]VertexType {
	return udg.edgeSet
}

func (udg UnGraph) locateVertex(v VertexType) int {
	for i := 0; i < udg.g.vertexNum; i++ {
		if udg.g.adjMultiList[i].data == v {
			return i
		}
	}
	return -1
}

func (udg UnGraph) setVertexSet(edges [][]VertexType) UnGraph {
	// Firstly, set edges
	udg.setEdgeSet(edges)
	var vs []VertexType
	for _, edge := range udg.edgeSet {
		vs = append(vs, edge[0])
		vs = append(vs, edge[1])
	}
	filteredVertices := filterVertices(vs)
	udg.vertexSet = filteredVertices
	return udg
}

func (udg UnGraph) setEdgeSet(edges [][]VertexType) UnGraph {
	filteredEdges := filterEdges(edges)
	udg.edgeSet = filteredEdges
	return udg
}

func filterEdges(edges [][]VertexType) [][]VertexType {
	//edges := [][]VertexType{
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
	var filtered [][]VertexType
	for _, edge := range edges {
		if len(edge) > 1 && len(edge) < 4 {
			// filter insignificant info
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

func filterVertices(vertices []VertexType) []VertexType {
	// create a map: {VertexType: true}
	existed := make(map[VertexType]bool)
	var filtered []VertexType
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

func PrintEdges(edges [][]VertexType) {
	for _, edge := range edges {
		fmt.Printf("(%v, %v)\n", edge[0], edge[1])
	}
}

func TestUndirectedGraph() {
	/*    multiDimensionArr := [][]VertexType{
	      {1, 2, 3},
	      {2, 3, 10},
	      {1, 3, 5},
	      {4, 3, 10},
	      {3, 4, 15},
	      {4, 5},
	      {1, 2},
	      {1},
	  }*/
	multiDimensionArr := [][]VertexType{
		{"A", "B", 3},
		{"B", "C", 10},
		{"A", "C", 5},
		{"D", "C", 10},
		{"C", "D", 15},
		{"D", "E"},
		{"A", "B"},
		{"A"},
	}
	fmt.Printf("1: %v\n", multiDimensionArr)
	filtered := filterEdges(multiDimensionArr)
	PrintEdges(filtered)
	fmt.Printf("4: %v\n", multiDimensionArr)
	fmt.Printf("5: %v\n", filtered)
}
