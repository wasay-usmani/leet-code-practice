package graphs

// numIslands returns the number of islands in the given grid.
func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r >= row || r < 0 ||
			c >= col || c < 0 ||
			grid[r][c] != '1' {
			return
		}

		grid[r][c] = '2'
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r-1, c)
		dfs(r, c-1)
	}
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++

				dfs(i, j)
				// bfs(grid, i, j)
			}
		}
	}

	return count
}
