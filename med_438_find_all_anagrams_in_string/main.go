package main

import "fmt"

func main() {
	in := "cbaebabacd"
	p := "abc"

	fmt.Printf("%v\n", findAnagrams(in, p))
}

func findAnagrams(s string, p string) []int {
	sl := len(s)
	pl := len(p)
	mp := make(map[byte]int, pl)
	for _, i := range p {
		mp[byte(i)] += 1
	}

	res := make([]int, 0, sl/pl)

	var w []byte
	var mw map[byte]int
L:
	for i := 0; i <= sl-pl; i++ {

		if i == 0 {
			w = []byte(s[i : i+pl])
			mw = make(map[byte]int, pl)

			for _, j := range w {
				mw[j]++
			}
		} else {
			// cut off first
			f := w[0]
			w = w[1:]

			// fmt.Printf("- %v\n", f)
			mw[f]--
			// add new
			w = append(w, s[i+pl-1])

			// fmt.Printf("+ %v\n", s[i+pl-1])
			mw[s[i+pl-1]]++
		}

		// fmt.Printf("%v\n", mw)

		for j, count := range mw {
			if mp[j] != count {
				continue L
			}
		}

		res = append(res, i)
	}

	return res
}

// abc
// cbaebabacd

// cba = 0
// c <- ba <- e
