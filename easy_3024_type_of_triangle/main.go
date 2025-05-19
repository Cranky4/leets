package main

import "fmt"

func main() {
	fmt.Println(triangleType([]int{2, 2, 2}))
}

func triangleType(nums []int) string {
	longest, x, y := -1, -1, -1
	a, b, c := nums[0], nums[1], nums[2]
	if a >= b && a >= c {
		longest = a
		x, y = b, c
	} else if b >= a && b >= c {
		longest = b
		x, y = a, c
	} else {
		longest = c
		x, y = a, b
	}

	if longest >= x+y {
		return "none"
	} else if a == b && a == c && b == c {
		return "equilateral"
	} else if a == b || b == c || c == a {
		return "isosceles"
	}

	return "scalene"
}
