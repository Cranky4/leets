package main

import "fmt"

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

// 2, 3, 1, 2, 4, 3
// 0, 2, 5, 6, 8, 12, 15
// 7
// 2
func minSubArrayLen(target int, nums []int) int {
	prefs := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefs[i+1] = prefs[i] + nums[i]
	}

	var l, r int

	res := len(nums) + 1
	for r < len(prefs) {
		diff := prefs[r] - prefs[l]

		if diff < target {
			r++
			continue
		}

		if res > r-l {
			res = r - l
		}
		l++
	}

	if res == len(nums)+1 {
		return 0
	}

	return res
}
