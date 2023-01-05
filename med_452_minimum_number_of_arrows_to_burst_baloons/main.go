package main

import (
	"fmt"
	"sort"
)

func main() {
	in := [][]int{
		{10, 16},
		{2, 8},
		{1, 6},
		{7, 12},
	}

	fmt.Printf("%v\n", findMinArrowShots(in))
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 1

	end := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			res++
			end = points[i][1]
		}
	}

	return res
}
