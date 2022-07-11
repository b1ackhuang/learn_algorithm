package main

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	var memo = make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		memo = append(memo, make([]int, len(triangle[i])))
	}
	for i := 0; i < len(triangle[0]); i++ {
		memo[len(memo)-1][i] = triangle[len(triangle)-1][i]
	}
	for i := len(triangle) - 2; i >= 0; i++ {
		for j := 0; j <= i; j++ {
			memo[i][j] = triangle[i][j]
			if memo[i+1][j] > memo[i+1][j+1] {
				memo[i][j] += memo[i+1][j]
			} else {
				memo[i][j] += memo[i+1][j+1]
			}
		}
	}
	return memo[0][0]
}
