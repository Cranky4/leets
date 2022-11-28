package main

import (
	"fmt"
	"sort"
)

func main() {
	// in := {}{}int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	in := [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}

	fmt.Printf("%v\n", findWinners(in))
}

func findWinners(matches [][]int) [][]int {
	loses := make(map[int]int, len(matches)) // {player}loses

	for i := 0; i < len(matches); i++ {
		win := matches[i][0]
		lose := matches[i][1]

		loses[win] += 0
		loses[lose]++
	}

	nolose := make([]int, 0, len(loses))
	onelose := make([]int, 0, len(loses))

	for p, c := range loses {
		if c == 0 {
			nolose = append(nolose, p)
		}

		if c == 1 {
			onelose = append(onelose, p)
		}
	}

	sort.Slice(nolose, func(i, j int) bool {
		return nolose[i] < nolose[j]
	})

	sort.Slice(onelose, func(i, j int) bool {
		return onelose[i] < onelose[j]
	})

	return [][]int{nolose, onelose}
}
