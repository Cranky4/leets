package main

import "fmt"

func main() {
	in := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	fmt.Printf("%v\n", search(in, target))
}

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	for min <= max {
		if nums[min] == target {
			return min
		}

		if nums[max] == target {
			return max
		}

		mid := (max + min) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}
