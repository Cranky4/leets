package main

import "fmt"

func main() {
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6}))
}

func tupleSameProduct(nums []int) int {
	m := make(map[int]int) // [prod]count
	l := len(nums)

	var res int
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			m[nums[i]*nums[j]]++
		}
	}

	for _, count := range m {
		res += 8 * (count * (count - 1)) / 2
	}

	return res
}
