package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

/*
Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/
func productExceptSelf(nums []int) []int {
	//         1  2  3  4
	// pref    1  1  2  6 24
	// suff    24 24 12 4  1

	prefs := make([]int, len(nums)+1)
	prefs[0] = 1

	for i := 0; i < len(nums); i++ {
		prefs[i+1] = prefs[i] * nums[i]
	}

	suffs := make([]int, len(nums)+1)
	suffs[len(nums)] = 1

	for i := len(nums) - 1; i >= 0; i-- {
		suffs[i] = suffs[i+1] * nums[i]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = prefs[i] * suffs[i+1]
	}

	return nums
}
