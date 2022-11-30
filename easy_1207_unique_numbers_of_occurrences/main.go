package main

import "fmt"

func main() {
	// in := []int{1, 2, 2, 1, 1, 3}
	// in := []int{1, 2}
	in := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}

	fmt.Printf("%v\n", uniqueOccurrences(in))
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int, len(arr))

	for _, n := range arr {
		m[n]++
	}

	mm := make(map[int]struct{}, len(m))
	for _, n := range m {
		_, ok := mm[n]
		if ok {
			return false
		}

		mm[n] = struct{}{}
	}

	return true
}
