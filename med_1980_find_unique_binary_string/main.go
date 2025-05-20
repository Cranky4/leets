package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(findDifferentBinaryString([]string{"000", "111", "011", "001", "100"}))
}

func findDifferentBinaryString(nums []string) string {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[bToD([]rune(n))] = struct{}{}
	}

	mx := int(math.Pow(2, float64(len(nums))) - 1)
	for i := 0; i <= mx; i++ {
		if _, ok := m[i]; !ok {
			return dToB(i, len(nums[0]))
		}
	}

	return ""
}

func bToD(s []rune) int {
	res := 0

	for i := len(s) - 1; i >= 0; i-- {
		d, _ := strconv.Atoi(string(s[i]))

		res += d * int(math.Pow(2, float64(i)))
	}

	return res
}

func dToB(n int, ln int) string {
	res := make([]string, ln)
	for i := range res {
		res[i] = "0"
	}

	i := 0
	for n > 0 {
		res[i] = strconv.Itoa(n % 2)
		n /= 2
		i++
	}

	return strings.Join(res, "")
}
