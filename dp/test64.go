package main

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var memo [][]int
	for i := 0; i < len(grid); i++ {
		memo = append(memo, make([]int, len(grid[0])))
	}
	for i := 0; i < len(grid); i++ {
		memo[i][0] = grid[i][0]
	}
	for i := 0; i < len(grid[0]); i++ {
		memo[0][i] = grid[0][i]
	}
	for i := 1; i < len(memo); i++ {
		for j := 1; j < len(memo[0]); j++ {
			if memo[i][j-1] < memo[i-1][j] {
				memo[i][j] = grid[i][j] + memo[i][j-1]
			} else {
				memo[i][j] = grid[i][j] + memo[i-1][j]
			}
		}
	}
	return memo[len(memo)-1][len(memo[0])-1]
}
