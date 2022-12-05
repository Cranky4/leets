package main

import (
	"fmt"
	"math"
)

func main() {
	in := 5

	fmt.Printf("%v\n", firstBadVersion(in))
}

func firstBadVersion(n int) int {
	min := 0
	max := n
	res := math.MaxInt

	for min <= max {
		med := (max + min) / 2

		if isBadVersion(med) {
			res = med
			max = med - 1
		} else {
			min = med + 1
		}
	}

	return res
}

func isBadVersion(n int) bool {
	return n == 4
}
