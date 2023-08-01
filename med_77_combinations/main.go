package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
	fmt.Println(combine(3, 3))
}

func combine(n int, k int) [][]int {
	return app(1, n, k)
}

func app(from, to, cap int) [][]int {
	res := make([][]int, 0)

	for i := from; i <= to; i++ {
		if cap > 1 {
			for _, subitem := range app(i+1, to, cap-1) {
				item := make([]int, 0, cap-1)
				item = append(item, i)
				item = append(item, subitem...)

				res = append(res, item)
			}
		} else {
			item := make([]int, 0, cap)
			item = append(item, i)
			res = append(res, item)
		}
	}

	return res
}
