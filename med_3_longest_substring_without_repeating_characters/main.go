package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))

}

func lengthOfLongestSubstring(s string) int {
	ln := len(s)
	if ln <= 1 {
		return ln
	}

	// map[char]idx
	m := make(map[byte]int)
	res := 0
	mx := 0 // current max

	var l, r int
	for r = 0; r < ln; r++ {
		// continue to collect uniques
		idx, ok := m[s[r]]
		if !ok {
			m[s[r]] = r
			mx++
			continue
		}

		res = max(mx, res)

		// clear map from left
		for i := l; i < idx; i++ {
			delete(m, s[i])
		}
		mx = len(m)

		// push left to idx+1 (for iteration will add 1)
		l = idx + 1
		m[s[r]] = r
	}

	return max(mx, res)
}
