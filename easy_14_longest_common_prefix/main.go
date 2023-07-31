package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	pref := ""
	l := 1

	for {
		if len(strs[0]) < l {
			return pref
		}

		p := strs[0][0:l]

		for _, s := range strs {
			if len(s) < l {
				return pref
			}

			if p != s[0:l] {
				return pref
			}
		}
		l++
		pref = p
	}

	return pref
}
