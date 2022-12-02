package main

import (
	"fmt"
	"sort"
)

func main() {
	// true
	// w1 := "abc"
	// w2 := "bca"

	// false
	// w1 := "a"
	// w2 := "aa"

	// true
	// w1 := "cabbba"
	// w2 := "abbccc"

	// false
	w1 := "uau"
	w2 := "ssx"

	fmt.Printf("%v\n", closeStrings(w1, w2))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	m1 := make(map[rune]int)
	for _, s := range word1 {
		m1[s]++
	}
	a1 := make([]int, 0, len(m1))
	for _, n := range m1 {
		a1 = append(a1, n)
	}

	m2 := make(map[rune]int)
	for _, s := range word2 {
		m2[s]++
	}
	a2 := make([]int, 0, len(m2))
	for _, n := range m2 {
		a2 = append(a2, n)
	}

	// check alphabet
	for l := range m2 {
		_, ok := m1[l]
		if !ok {
			return false
		}
	}
	for l := range m1 {
		_, ok := m2[l]
		if !ok {
			return false
		}
	}

	// sort counts
	sort.Slice(a1, func(i, j int) bool {
		return a1[i] > a1[j]
	})

	sort.Slice(a2, func(i, j int) bool {
		return a2[i] > a2[j]
	})

	// compare count pairs
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}
