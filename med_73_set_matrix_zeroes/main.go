package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	in := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setZeroes(in)

	fmt.Println(in)

}

func setZeroes(matrix [][]int) {
	trueZeros := make(map[string]struct{})

	for row, cols := range matrix {
		for col, val := range cols {
			if val == 0 {
				trueZeros[fmt.Sprintf("%d-%d", row, col)] = struct{}{}
			}
		}
	}

	totalRows := len(matrix) - 1
	totalCols := len(matrix[0]) - 1

	for coordStr := range trueZeros {
		coords := strings.SplitN(coordStr, "-", 2)
		row, _ := strconv.Atoi(coords[0])
		col, _ := strconv.Atoi(coords[1])

		// row
		for i := 0; i <= totalCols; i++ {
			matrix[row][i] = 0
		}

		// col
		for i := 0; i <= totalRows; i++ {
			matrix[i][col] = 0
		}
	}
}
