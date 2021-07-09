package graph

// Directed Graph, implementation with Orthogonal List

type Vertex struct {
}

type Arc struct {
}

type DiGraph struct {
}

func filterArcArr(arcs [][]int) [][]int {
	//arcs := [][]int{
	//   {1, 2, 3},
	//   {2, 3, 10},
	//   {1, 3, 5},
	//   {4, 3, 10},
	//   {3, 4, 15},
	//   {1, 2, 15}, // will be deleted
	//}
	var filtered [][]int
	for _, arc := range arcs {
		if len(arc) >= 2 {
			// filter insignificant data
			filtered = append(filtered, arc)
		}
	}
	for i, arc := range filtered {
		for j := i + 1; j < len(filtered); j++ {
			if arc[0] == filtered[j][0] && arc[1] == filtered[j][1] {
				// delete arcs[j], due to the repeated with previous
				filtered = append(filtered[:j], filtered[j+1:]...)
			}
		}
	}
	return filtered
}
