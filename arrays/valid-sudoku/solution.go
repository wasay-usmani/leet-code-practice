package validsudoku

import "unicode"

// IsValidSudoku returns true if the given 9x9 board is a valid Sudoku.
func IsValidSudoku(board [][]byte) bool {
	transposed := make([][]byte, len(board))
	for i := range transposed {
		transposed[i] = make([]byte, len(board[0]))
	}

	box := make([][]byte, len(board))
	for i := range box {
		box[i] = make([]byte, len(board[0]))
	}

	for i, row := range board {
		visited := map[byte]struct{}{}
		for j, val := range row {
			if _, ok := visited[val]; ok && unicode.IsDigit(rune(val)) {
				return false
			}

			visited[val] = struct{}{}
			transposed[j][i] = val
			boxNumber := (i/3)*3 + (j / 3)
			position := (i%3)*3 + (j % 3)
			box[boxNumber][position] = val
		}
	}

	for i := 0; i < len(board[0]); i++ {
		visited := map[byte]struct{}{}
		visitedCol := map[byte]struct{}{}
		for j := 0; j < len(board[0]); j++ {
			if _, ok := visited[transposed[i][j]]; ok && unicode.IsDigit(rune(transposed[i][j])) {
				return false
			}

			if _, ok := visitedCol[box[i][j]]; ok && unicode.IsDigit(rune(box[i][j])) {
				return false
			}

			visited[transposed[i][j]] = struct{}{}
			visitedCol[box[i][j]] = struct{}{}
		}
	}
	
	return true
}
