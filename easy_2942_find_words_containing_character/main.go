package main

import "fmt"

func main() {
	fmt.Println(findWordsContaining([]string{"leet", "code"}, 'e'))
}

func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0)
	ln := len(words)

	for i := 0; i < ln; i++ {
		for _, j := range words[i] {
			if byte(j) == x {
				res = append(res, i)
				break
			}
		}
	}

	return res
}
