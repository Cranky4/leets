package main

// https://en.wikipedia.org/wiki/Depth-first_search

func recoursive_dfs_backtracking(board [][]byte, word string, i, j, curW int) bool {
	if word[curW] != board[i][j] {
		return false
	}

	rows := len(board)
	cols := len(board[0])

	if curW == len(word)-1 {
		return true
	}

	result := false

	// backtrack
	orig := board[i][j]
	board[i][j] = '\n'

	// show path
	// for _, v := range board {
	// 	fmt.Printf("%v\n", v)
	// }
	// fmt.Printf("\n")

	// up
	if i-1 >= 0 {
		result = result || recoursive_dfs_backtracking(board, word, i-1, j, curW+1)
	}

	// right
	if j+1 < cols {
		result = result || recoursive_dfs_backtracking(board, word, i, j+1, curW+1)
	}

	// down
	if i+1 < rows {
		result = result || recoursive_dfs_backtracking(board, word, i+1, j, curW+1)
	}

	// left
	if j-1 >= 0 {
		result = result || recoursive_dfs_backtracking(board, word, i, j-1, curW+1)
	}

	board[i][j] = orig

	return result
}
