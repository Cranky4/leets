package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "cccaaa"

	m := make(map[rune]int, len(s))

	for _, l := range s {
		m[l]++
	}

	q := NewQueue(len(s))

	for r, n := range m {
		q.Add(item{r: r, n: n})
	}

	var builder strings.Builder

	for {
		n := q.Get()
		if n.n == -1 {
			break
		}

		for i := 0; i < n.n; i++ {
			builder.WriteRune(n.r)
		}
	}

	res := builder.String()

	fmt.Printf("%v\n", res)
}
