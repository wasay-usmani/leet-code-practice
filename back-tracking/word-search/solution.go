package backtracking

// exist returns true if the word exists in the board.
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(i, j, 0, board, word) {
				return true
			}
		}
	}

	return false
}

func dfs(r, c, count int, board [][]byte, word string) bool {
	if count == len(word) {
		return true
	}

	if r >= len(board) || r < 0 ||
		c >= len(board[0]) || c < 0 ||
		board[r][c] != word[count] {
		return false
	}

	temp := board[r][c]
	board[r][c] = '*'
	found := dfs(r+1, c, count+1, board, word) ||
		dfs(r-1, c, count+1, board, word) ||
		dfs(r, c+1, count+1, board, word) ||
		dfs(r, c-1, count+1, board, word)
	board[r][c] = temp
	return found
}
