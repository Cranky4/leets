package main

import "fmt"

func main() {
	n := 7
	in := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	apples := []bool{false, false, true, false, true, true, false}

	fmt.Printf("%v\n", minTime(n, in, apples))
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	tree := make(map[int][]int, n)

	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	res, _ := walk(0, 0, tree, hasApple)

	return res
}

func walk(node, prev int, tree map[int][]int, apples []bool) (int, bool) {
	res := 0
	ex := false

	if apples[node] {
		ex = true
	}

	childs, ok := tree[node]

	if !ok {
		return res, ex
	}

	for _, c := range childs {
		if c == prev {
			continue
		}

		r, e := walk(c, node, tree, apples)
		if e {
			res += r + 2
			ex = e
		}
	}

	return res, ex
}
