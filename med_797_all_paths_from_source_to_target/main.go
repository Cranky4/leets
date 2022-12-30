package main

import "fmt"

func main() {
	// in := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	// in := [][]int{{1, 2}, {3}, {3}, {}}
	// in := [][]int{{4, 3, 1}, {3, 2, 4}, {}, {4}, {}}

	in := [][]int{{2}, {}, {1}}
	fmt.Printf("%v\n", in)

	fmt.Printf("%v\n", allPathsSourceTarget(in))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	paths, _ := walk(graph, 0)
	return paths
}

func walk(graph [][]int, idx int) ([][]int, bool) {
	if idx == len(graph)-1 {
		return [][]int{{idx}}, true
	}

	if len(graph[idx]) == 0 {
		return [][]int{{idx}}, false
	}
	res := make([][]int, 0)

	for _, child := range graph[idx] {
		path, ok := walk(graph, child)
		if !ok {
			continue
		}
		for _, s := range path {
			rs := make([]int, 0)
			rs = append(rs, idx)
			rs = append(rs, s...)
			res = append(res, rs)
		}
	}

	return res, true
}
