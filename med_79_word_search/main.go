package main

import "fmt"

func main() {
	// board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	// word := "ABCCED"

	// board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	// word := "ABCB"

	// board := [][]byte{{'a', 'b'}}
	// word := "ba"

	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCESEEEFS"

	fmt.Printf("W - %v\n\n", []byte(word))
	for _, v := range board {
		fmt.Printf("%v\n", v)
	}

	fmt.Printf("%v\n", exist(board, word))
}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	// search for first letter
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if word[0] == board[i][j] {
				// start dfs
				if recoursive_dfs_backtracking(board, word, i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}
