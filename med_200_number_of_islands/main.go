package main

import "fmt"

func main() {
	in := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	fmt.Printf("%v\n", numIslands(in))
}

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	res := 0

	for row, r := range grid {
		for col := range r {
			if grid[row][col] == 49 {
				grid = walk(grid, row, col, rows, cols)

				res++
			}
		}
	}

	return res
}

func walk(grid [][]byte, row, col, rows, cols int) [][]byte {
	if grid[row][col] != 49 {
		return grid
	}

	grid[row][col] = 50

	// up
	if row-1 >= 0 {
		grid = walk(grid, row-1, col, rows, cols)
	}

	// right
	if col+1 < cols {
		grid = walk(grid, row, col+1, rows, cols)
	}

	// down
	if row+1 < rows {
		grid = walk(grid, row+1, col, rows, cols)
	}

	// left
	if col-1 >= 0 {
		grid = walk(grid, row, col-1, rows, cols)
	}

	return grid
}
