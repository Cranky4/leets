package main

import "fmt"

func main() {
	n := 3
	// t := [][]int{{1, 3}, {2, 3}}
	t := [][]int{{1, 3}, {2, 3}, {3, 1}}

	fmt.Printf("%v\n", findJudge(n, t))
}

func findJudge(n int, trust [][]int) int {
	tr1 := make([]int, n) // [idx]count
	tr2 := make([]int, n)

	for i := 0; i < len(trust); i++ {
		tr1[trust[i][1]-1]++
		tr2[trust[i][0]-1]++
	}

	for idx, tr := range tr2 {
		if tr == 0 && tr1[idx] == n-1 {
			return idx + 1
		}
	}

	return -1
}
