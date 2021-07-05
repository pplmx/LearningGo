package structure

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
func (udg *Graph) create(edges [...][2]int) {

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

func filterEdges(edges [...][2]int) {

}