package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome([]string{"ab", "ac", "cc", "cc", "cc", "ca", "ba", "bb"}))
}

func longestPalindrome(words []string) int {
	m := make(map[string]int)
	res := 0

	for _, w := range words {
		m[w]++

		b := []byte(w)
		rev := string([]byte{b[1], b[0]})

		if b[0] == b[1] && m[w] == 2 {
			res += 4
			m[w] -= 2

			continue
		}

		count, ok := m[rev]
		if ok && count > 0 && b[0] != b[1] {
			m[w]--
			m[rev]--

			res += 4

			continue
		}
	}

	for w, count := range m {
		if count == 1 {
			b := []byte(w)
			if b[0] == b[1] {
				res += 2
				break
			}
		}
	}

	return res
}
