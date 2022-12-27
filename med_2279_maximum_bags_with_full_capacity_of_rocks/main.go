package main

import (
	"fmt"
	"sort"
)

func main() {
	caps := []int{2, 3, 4, 5}
	rocks := []int{1, 2, 4, 4}
	avl := 2

	fmt.Printf("%v\n", maximumBags(caps, rocks, avl))
}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	avails := make([]int, len(capacity))
	for i, c := range capacity {
		avails[i] = c - rocks[i]
	}

	sort.Slice(avails, func(i, j int) bool {
		return avails[i] <= avails[j]
	})

	res := 0
	for _, avail := range avails {
		// already full
		if avail == 0 {
			res++
			continue
		}

		// not enough rocks
		if additionalRocks == 0 || additionalRocks < avail {
			break
		}

		additionalRocks -= avail
		res++
	}

	return res
}
