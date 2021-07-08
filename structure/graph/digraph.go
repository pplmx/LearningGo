package graph

// Directed Graph, implementation with Orthogonal List

type Vertex struct {
}

type Arc struct {
}

type DiGraph struct {
}

func filterArcArr(arcs [][3]int) [][3]int {
	//arcs := [][]int{
	//   {1, 2, 3},
	//   {2, 3, 10},
	//   {1, 3, 5},
	//   {4, 3, 10},
	//   {3, 4, 15},
	//   {1, 2, 15}, // will be deleted
	//}
	for i, arc := range arcs {
		for j := i + 1; j < len(arcs); j++ {
			if arc[0] == arcs[j][0] && arc[1] == arcs[j][1] {
				// delete arcs[j], due to the repeated with previous
				arcs = append(arcs[:j], arcs[j+1:]...)
			}
		}
	}
	return arcs
}
