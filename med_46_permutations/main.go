package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}

func permute(nums []int) [][]int {
	l := len(nums)
	if l == 1 {
		return [][]int{nums}
	}

	res := make([][]int, 0)

	for i := 0; i < l; i++ {
		cur := nums[i]
		newNums := make([]int, 0, l-1)
		newNums = append(newNums, nums[0:i]...)
		newNums = append(newNums, nums[i+1:]...)

		submutations := permute(newNums)

		for _, m := range submutations {
			item := make([]int, 0, len(newNums)+1)
			item = append(item, cur)
			item = append(item, m...)

			res = append(res, item)
		}
	}

	return res
}
