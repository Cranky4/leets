package main

import "fmt"

func main() {
	// in := []int{2, 3, 1, 1, 4}
	in := []int{3, 2, 1, 0, 4}

	fmt.Printf("%v\n", canJump(in))
}

func canJump(nums []int) bool {
	farthest := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]+i > farthest {
			farthest = nums[i] + i
		}

		if farthest >= len(nums)-1 {
			return true
		}

		if farthest == i {
			return false
		}
	}

	return false
}
