package main

import "fmt"

func main() {
	// in := "USA"
	// in := "FlaG"
	in := "flag"

	fmt.Printf("%v\n", detectCapitalUse(in))
}

func detectCapitalUse(word string) bool {
	bytes := []byte(word)
	allLower := false
	allUpper := false

	if isLower(bytes[0]) {
		allLower = true
	}

	for i := 1; i < len(bytes); i++ {
		if isLower(bytes[i]) && allUpper {
			return false
		} else if isLower(bytes[i]) {
			allLower = true
		} else if isUpper(bytes[i]) && allLower {
			return false
		} else {
			allUpper = true
		}
	}

	return true
}

func isLower(b byte) bool {
	return b >= byte('a')
}

func isUpper(b byte) bool {
	return b >= byte('A') && b < byte('a')
}
