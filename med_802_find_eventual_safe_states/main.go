package main

import (
	"fmt"
	"sort"
)

func main() {
	in := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}

	fmt.Println(eventualSafeNodes(in))

	// Time Limit Exceeded
	//99 / 113 testcases passed
}

func eventualSafeNodes(graph [][]int) []int {
	res := make([]int, 0)

	for i := 0; i < len(graph); i++ {
		r, ok := dfs(graph, i, res, make(map[int]struct{}))
		if ok {
			res = append(res, r...)
		}
	}

	sort.Slice(res, func(a, b int) bool {
		return a < b
	})

	return res
}

func dfs(graph [][]int, cur int, res []int, mem map[int]struct{}) ([]int, bool) {
	if _, ok := mem[cur]; ok {
		return nil, false
	}
	mem[cur] = struct{}{}

	for _, out := range graph[cur] {
		r, ok := dfs(graph, out, res, mem)
		if !ok {
			return nil, false
		}
		res = append(res, r...)
	}

	res = append(res, cur)
	delete(mem, cur)

	return nil, true
}
