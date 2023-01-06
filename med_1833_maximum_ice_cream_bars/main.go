package main

import (
	"fmt"
	"sort"
)

func main() {
	costs := []int{1, 3, 2, 4, 1}
	coins := 7

	fmt.Printf("%v\n", maxIceCream(costs, coins))
}

func maxIceCream(costs []int, coins int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})

	res := 0
	for i := 0; i < len(costs); i++ {
		if coins < costs[i] {
			break
		}

		res++
		coins -= costs[i]
	}

	return res
}
