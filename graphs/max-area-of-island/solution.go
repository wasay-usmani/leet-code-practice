package graphs

// maxAreaOfIsland returns the area of the largest island in the grid.
func maxAreaOfIsland(grid [][]int) int {
	row, col := len(grid), len(grid[0])

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= row ||
			c < 0 || c >= col ||
			grid[r][c] != 1 {
			return 0
		}

		grid[r][c] = 2
		return 1 + dfs(r+1, c) +
			dfs(r, c+1) +
			dfs(r-1, c) +
			dfs(r, c-1)
	}

	maxArea := 0
	for i, v := range grid {
		for j := range v {
			if grid[i][j] == 1 {
				a := dfs(i, j)
				if a > maxArea {
					maxArea = a
				}
			}
		}
	}

	return maxArea
}
