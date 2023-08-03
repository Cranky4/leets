package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("5646")) // john
	fmt.Println(letterCombinations("5338")) // leet
}

func letterCombinations(digits string) []string {
	btns := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	variants := make([][]rune, 0, len(digits))
	for _, d := range digits {
		v, _ := btns[d]

		variants = append(variants, v)
	}

	vars := getVars(variants)
	res := make([]string, 0, len(vars))

	for _, v := range vars {
		res = append(res, string(v))
	}

	return res
}

func getVars(variants [][]rune) [][]rune {
	if len(variants) == 0 {
		return [][]rune{}
	}

	res := make([][]rune, 0)

	if len(variants) == 1 {
		for _, v := range variants[0] {
			res = append(res, []rune{v})
		}

		return res
	}

	for _, r := range variants[0] {
		for _, sub := range getVars(variants[1:]) {
			item := make([]rune, 0, len(variants))
			item = append(item, r)
			item = append(item, sub...)

			res = append(res, item)
		}
	}

	return res
}
