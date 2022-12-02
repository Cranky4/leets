package main

import "fmt"

func main() {
	// in := "abccccdd"
	in := "a"

	fmt.Printf("%v\n", longestPalindrome(in))
}

func longestPalindrome(s string) int {
	m := make(map[rune]int, len(s)) // [letter]count
	if len(s) == 0 {
		return 0
	}
	res := 0

	for _, l := range s {
		m[l]++

		count := m[l]

		if count == 2 {
			res += 2
			m[l] = 0
		}
	}

	for _, c := range m {
		if c == 1 {
			res++
			break
		}
	}

	return res
}
