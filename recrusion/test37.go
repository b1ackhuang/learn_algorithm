package main

var direction = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func solvesSudo(board [][]byte) {
	var rowUsed, colUsed, areaUsed = makeSlice(9, 10), makeSlice(9, 10), makeSlice(3, 3)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] > 0 {
				rowUsed[i][board[i][j]] = true
				colUsed[j][board[i][j]] = true
				areaUsed[j%3+3*(i%3)][board[i][j]] = true
			}
		}
	}
	fillBoard(0, 0, board, 0, rowUsed, colUsed, areaUsed)
}

// 从 (x, y) 开始填充
func fillBoard(x, y int, board [][]byte, filledCnt int, rowUsed, colUsed, areaUsed [][]bool) {
	if filledCnt > 81 {
		return
	}

	if board[x][y] > 0 { //已有数字
		for i := 0; i < len(direction); i++ {

		}
		fillBoard(x+1, y, board, filledCnt+1, rowUsed, colUsed, areaUsed)
		fillBoard(x, y+1, board, filledCnt+1, rowUsed, colUsed, areaUsed)
		return
	}
	for i := 0; i <= 9; i++ {
		if !rowUsed[x][i] {
			continue
		}
		if !colUsed[y][i] {
			continue
		}
		if !areaUsed[y%3+3*(x%3)][i] {
			continue
		}
		rowUsed[x][i] = true
		colUsed[y][i] = true
		areaUsed[y%3+3*(x%3)][i] = true
		board[x][y] = byte(i)
		fillBoard(x+1, y, board, filledCnt+1, rowUsed, colUsed, areaUsed)
		fillBoard(x, y+1, board, filledCnt+1, rowUsed, colUsed, areaUsed)
		rowUsed[x][i] = false
		colUsed[y][i] = false
		areaUsed[y%3+3*(x%3)][i] = false
	}
}

func makeSlice(x, y int) [][]bool {
	var resSlice = make([][]bool, x)
	for i := 0; i < y; i++ {
		resSlice = append(resSlice, make([]bool, y))
	}
	return resSlice
}
