package main

import "fmt"

func main() {
	in := []int{2, 0, 2, 1, 1, 0}
	sortColors(in)
	fmt.Println(in)
}

func sortColors(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
