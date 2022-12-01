package main

import "fmt"

func main() {
	// in := "book"
	in := "textbook"

	fmt.Printf("%v\n", halvesAreAlike(in))
}

func halvesAreAlike(s string) bool {
	l := len(s)
	h := int(l / 2)

	a := []byte(s[:h])
	b := []byte(s[h:])

	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}

	n1 := 0
	n2 := 0

	for _, s := range a {
		_, ok := vowels[s]
		if ok {
			n1++
		}
	}

	for _, s := range b {
		_, ok := vowels[s]
		if ok {
			n2++
		}

		if n2 > n1 {
			return false
		}
	}

	return n1 == n2
}
