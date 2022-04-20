package main

import "fmt"

const (
	MaskHistoryVisited  = 1 << 2
	MaskConnectPacific  = 1 << 1
	MaskConnectAtlantic = 1 << 0
	MaskAll             = MaskConnectPacific | MaskConnectAtlantic
)

func pacificAtlantic(height [][]int) [][]int {
	if len(height) == 0 || len(height[0]) == 0 {
		return nil
	}
	var history [][]int
	for i := 0; i < len(height); i++ {
		history = append(history, make([]int, len(height[0])))
	}

	var resPoint [][]int
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height[0]); j++ {
			var tmpRes = judgeConnect(height, i, j, history)
			if tmpRes&MaskAll == MaskAll {
				resPoint = append(resPoint, []int{i, j})
			}
			history[i][j] = tmpRes
		}
	}
	fmt.Println(history)
	return resPoint
}

// 判断height[x][y]是满足条件
func judgeConnect(height [][]int, x, y int, history [][]int) int {
	var visited [][]bool
	for i := 0; i < len(height); i++ {
		visited = append(visited, make([]bool, len(height[0])))
	}

	var connectMask = MaskHistoryVisited //0x01：pacific， 0x10：atlantic
	var points [][]int
	points = append(points, []int{x, y})
	for len(points) != 0 {
		var curX, curY = points[0][0], points[0][1]
		visited[curX][curY] = true
		if history[curX][curY]&MaskHistoryVisited == MaskHistoryVisited { //曾经遍历过,使用先前内容加速
			connectMask |= history[curX][curY]
		} else {
			if curX == 0 || curY == 0 { //左上角，太平洋
				connectMask |= MaskConnectPacific
			}
			if curX == len(height)-1 || curY == len(height[0])-1 { //右下角，大西洋
				connectMask |= MaskConnectAtlantic
			}
			for i := 0; i < len(direction); i++ {
				var nextX, nextY = curX + direction[i][0], curY + direction[i][1]
				if InBoard(nextX, nextY, height) && //在地图中
					!visited[nextX][nextY] && //未被遍历
					height[curX][curY] >= height[nextX][nextY] { //连通区域
					points = append(points, []int{nextX, nextY})
				}
			}
		}
		if connectMask&MaskAll == MaskAll { //提前返回
			return connectMask
		}
		points = points[1:]
	}
	return connectMask
}

func InBoard(x, y int, board [][]int) bool {
	return x >= 0 && y >= 0 && x < len(board) && y < len(board[0])
}
