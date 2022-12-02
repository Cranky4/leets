package main

import (
	"fmt"
	"math"
)

func main() {
	// in := []int{7, 1, 5, 3, 6, 4} // 5
	in := []int{7, 6, 4, 3, 1} // 0

	fmt.Printf("%v\n", maxProfit(in))
}

func maxProfit(prices []int) int {
	min := math.MaxInt
	max := 0
	res := 0

	for _, p := range prices {
		if p < min {
			min = p
			max = 0
		}

		if p > max {
			max = p
			if res < max-min && max != 0 && min != math.MaxInt {
				res = max - min
			}
		}
	}

	return res
}
