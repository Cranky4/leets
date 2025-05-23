package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findEvenNumbers([]int{2, 1, 3, 0}))
}

func findEvenNumbers(digits []int) []int {
	m := make(map[int]int, 0)
	for _, d := range digits {
		m[d]++
	}

	res := make([]int, 0)

	// d..
	for i := range m {
		if i == 0 {
			continue
		}

		m[i]--

		// .d.
		for j, jpar := range m {
			if jpar == 0 {
				continue
			}

			m[j]--

			// ..d
			for k, kpar := range m {
				if kpar == 0 || k%2 > 0 {
					continue
				}

				res = append(res, i*100+j*10+k)
			}

			m[j]++
		}

		m[i]++
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}
