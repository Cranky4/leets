package main

import "fmt"

func main() {
	s := "abc"
	shifts := [][]int{
		{0, 1, 0},
		{1, 2, 1},
		{0, 2, 1},
	}

	fmt.Println(shiftingLetters(s, shifts))
}

func shiftingLetters(s string, shifts [][]int) string {
	const abcLen = 26

	pref := make([]int32, len(s)+1)
	for _, shift := range shifts {
		if shift[2] == 1 {
			pref[shift[0]]++
			pref[shift[1]+1]--
		} else {
			pref[shift[0]]--
			pref[shift[1]+1]++
		}
	}

	for i := 1; i < len(pref); i++ {
		pref[i] += pref[i-1]
	}

	mn := 'a'
	mx := 'z'
	res := []rune(s)

	for i := 0; i < len(res); i++ {
		sh := pref[i]
		if sh > abcLen {
			sh = sh % abcLen
		}
		if sh < -abcLen {
			sh = sh % -abcLen
		}

		newVal := res[i] + sh

		if newVal > mx {
			res[i] = mn + newVal - mx - 1
			continue
		}

		if newVal < mn {
			res[i] = mx - mn + newVal + 1
			continue
		}

		res[i] = newVal
	}

	return string(res)
}
