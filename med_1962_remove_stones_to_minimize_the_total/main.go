package main

import "fmt"

func main() {
	in := []int{4, 3, 6, 7}
	k := 3

	fmt.Printf("%v\n", minStoneSum(in, k))
}

func minStoneSum(piles []int, k int) int {
	q := NewQueue(len(piles))

	for _, n := range piles {
		q.Add(n)
	}

	for i := 0; i < k; i++ {
		n := q.Get()
		n = n/2 + n%2
		q.Add(n)
	}

	res := 0
	for {
		n := q.Get()
		if n == -1 {
			break
		}

		res += n
	}

	return res
}
