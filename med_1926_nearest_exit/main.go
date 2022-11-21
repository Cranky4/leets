package main

import "fmt"

func main() {
	// 46 .
	// 43 +

	// field := [][]byte{{43, 43, 46, 43}, {46, 46, 46, 43}, {43, 43, 43, 46}}
	// start := [2]int{1, 2}
	// 1

	// field := [][]byte{{43, 43, 43}, {46, 46, 46}, {43, 43, 43}}
	// start := [2]int{1, 0}
	// 2

	field := [][]byte{{46, 43}}
	start := [2]int{0, 0}
	// -1

	for _, v := range field {
		fmt.Printf("%v\n", v)
	}

	out := bfs(start, field)

	fmt.Printf("%v", out)
}
