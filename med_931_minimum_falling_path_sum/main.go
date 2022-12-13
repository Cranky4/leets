package main

import (
	"fmt"
	"math"
)

func main() {
	in := [][]int{
		{100, -42, -46, -41},
		{31, 97, 10, -10},
		{-58, -51, 82, 89},
		{51, 81, 69, -51},
	}

	fmt.Printf("%v\n", minFallingPathSum(in))
}

func minFallingPathSum(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	res := math.MaxInt

	if rows == 1 {
		for _, i := range matrix[0] {
			if res > i {
				res = i
			}
		}

		return res
	}

	for row := 1; row < rows; row++ {
		for col := 0; col < cols; col++ {
			min := math.MaxInt

			// left
			if row-1 >= 0 && col-1 >= 0 && min > matrix[row-1][col-1] {
				min = matrix[row-1][col-1]
			}

			// top
			if row-1 >= 0 && min > matrix[row-1][col] {
				min = matrix[row-1][col]
			}

			// right
			if row-1 >= 0 && col+1 < cols && min > matrix[row-1][col+1] {
				min = matrix[row-1][col+1]
			}

			matrix[row][col] += min

			if row == rows-1 && matrix[row][col] < res {
				res = matrix[row][col]
			}
		}
	}

	return res
}
