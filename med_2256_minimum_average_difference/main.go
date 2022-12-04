package main

import (
	"fmt"
	"math"
)

func main() {
	in := []int{2, 5, 3, 9, 5, 3}
	// in := []int{4, 2, 0}

	fmt.Printf("%v\n", minimumAverageDifference(in))
}

func minimumAverageDifference(nums []int) int {
	prefLeft := make([]int, len(nums))
	prefRight := make([]int, len(nums))

	if len(nums) == 0 {
		return 0
	}

	prefLeft[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefLeft[i] = prefLeft[i-1] + nums[i]
	}

	prefRight[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		prefRight[i] = prefRight[i+1] + nums[i]
	}

	min := math.MaxInt
	res := 0
	for i := 0; i < len(prefLeft); i++ {
		l := prefLeft[i] / (i + 1)

		r := 0
		if i < len(prefLeft)-1 {
			r = prefRight[i+1] / (len(prefRight) - i - 1)
		}

		diff := int(math.Abs(float64(l - r)))

		if min > diff {
			min = diff
			res = i
		}
	}

	return res
}
