package search

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//
//
//
//示例 1：
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//示例 2：
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 300
//grid[i][j] 的值为 '0' 或 '1'

//0 0 [[118 118 118 118 48] [49 118 48 118 48] [49 118 48 48 48] [48 48 48 48 48]]
//1 0 [[118 118 118 118 48] [118 118 48 118 48] [49 118 48 48 48] [48 48 48 48 48]]
//2 0 [[118 118 118 118 48] [118 118 48 118 48] [118 118 48 48 48] [48 48 48 48 48]]

const Visited = 'v'

var (
	direction = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var res = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				mark(grid, i, j)
			}
		}
	}
	return res
}

func mark(grid [][]byte, x, y int) {
	grid[x][y] = Visited
	for i := 0; i < len(direction); i++ {
		var nextX, nextY = x + direction[i][0], y + direction[i][1]
		if inRange(grid, nextX, nextY) &&
			grid[nextX][nextY] == '1' {
			mark(grid, nextX, nextY)
		}
	}
}

func inRange(grid [][]byte, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}
