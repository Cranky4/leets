package main

import (
	"fmt"
	"math"
)

func main() {
	ranks := []int{5, 1, 8}
	cars := 6

	fmt.Println(repairCars(ranks, cars))
}

func repairCars(ranks []int, cars int) int64 {
	var l int64 = 0
	r := int64(math.MaxInt64)
	var mid int64

	for l < r {
		mid = (r + l) / 2

		repCars := ccg(ranks, mid)
		fmt.Printf("l: %d, r: %d, mid: %d, repCars: %d\n", l, r, mid, repCars)

		if repCars >= cars {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func ccg(ranks []int, m int64) int {
	repCars := 0
	for _, rank := range ranks {
		repCars += cc(rank, m)
	}
	return repCars
}

func cc(r int, m int64) int {
	return int(
		math.Floor(
			math.Sqrt(
				float64(m) / float64(r),
			),
		),
	)
}
