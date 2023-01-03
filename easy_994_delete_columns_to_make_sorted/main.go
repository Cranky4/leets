package main

import "fmt"

func main() {
	in := []string{
		"abc",
		"bce",
		"cae",
	}

	fmt.Printf("%v\n", minDeletionSize(in))
}

func minDeletionSize(strs []string) int {
	l := len(strs[0])

	res := 0
	for i := 0; i < l; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				res++
				break
			}
		}
	}

	return res
}
