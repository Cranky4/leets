package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1}}, 0))
	fmt.Println(searchMatrix([][]int{{1}}, 1))
}

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	l := 0
	r := rows - 1

	row := -1

	for l <= r {
		m := (l + r) / 2

		// here!
		if matrix[m][0] <= target && matrix[m][cols-1] >= target {
			row = m
			break
		}

		if matrix[m][0] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if row == -1 {
		return false
	}

	l = 0
	r = cols - 1

	for l <= r {
		m := (l + r) / 2

		// here!
		if matrix[row][m] == target {
			return true
		}

		if matrix[row][m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}
