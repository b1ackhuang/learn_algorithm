package main

import "fmt"

func main() {
	//var bts = []byte{3, 7, 2, 5}
	//sort.Slice(bts, func(i, j int) bool {
	//	return bts[i] < bts[j]
	//})
	//var inStrArr string
	//_, _ = fmt.Scanln(&inStrArr)
	//var arrStr = strings.Split(inStrArr, ",")
	//var arr []int
	//for i := 0; i < len(arrStr); i++ {
	//	var tmpVal, _ = strconv.ParseInt(arrStr[i], 10, 64)
	//	arr = append(arr, int(tmpVal))
	//}
	//var inStrDst string
	//_, _ = fmt.Scanln(&inStrDst)
	//var dst, _ = strconv.ParseInt(inStrDst, 10, 64)
	//fmt.Println(find(arr, int(dst)))

	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 8, 15},
		{2, 5, 8, 12, 19},
		{6, 6, 9, 16, 22},
		{12, 13, 14, 17, 24},
	}, 22))
}

func find(arr []int, dst int) int {
	if len(arr) <= 0 {
		return -1
	}
	//在循环中保证循环的不变量
	var left, right = 0, len(arr) - 1 // [left, right] 区间内，查找目标元素
	for left <= right {
		var mid = left + (right-left)/2 // (l+r)/2 会溢出
		if arr[mid] == dst {
			return mid
		} else if arr[mid] > dst {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var upRow, downRow = 0, len(matrix) - 1
	var leftCol, rightCol = 0, len(matrix[0]) - 1
	for upRow < downRow && leftCol < rightCol {
		//fmt.Println(upRow, leftCol, "---", downRow, rightCol)
		tmpRightCol, hasTarget := searchRow(matrix, upRow, leftCol, rightCol, target)
		if hasTarget {
			return true
		}
		//fmt.Println("searchRow up:", upRow, leftCol, rightCol, target, ",tmpRightCol:", tmpRightCol)
		tmpLeftCol, hasTarget := searchRow(matrix, downRow, leftCol, rightCol, target)
		if hasTarget {
			return true
		}
		//fmt.Println("searchRow down:", downRow, leftCol, rightCol, target, ",tmpLeftCol:", tmpLeftCol)
		tmpDownRow, hasTarget := searchCol(matrix, leftCol, upRow, downRow, target)
		if hasTarget {
			return true
		}
		//fmt.Println("searchCol left:", leftCol, upRow, downRow, target, ",tmpDownRow:", tmpDownRow)
		tmpUpRow, hasTarget := searchCol(matrix, rightCol, upRow, downRow, target)
		if hasTarget {
			return true
		}
		//fmt.Println("searchCol right:", rightCol, upRow, downRow, target, ",tmpUpRow:", tmpUpRow)
		upRow, downRow, leftCol, rightCol = tmpUpRow+1, tmpDownRow, tmpLeftCol+1, tmpRightCol
	}

	if upRow == downRow {
		_, hasTarget := searchRow(matrix, upRow, leftCol, rightCol, target)
		return hasTarget
	} else {
		_, hasTarget := searchCol(matrix, leftCol, upRow, downRow, target)
		return hasTarget
	}
}

func searchRow(matrix [][]int, row, left, right, target int) (int, bool) {
	var l, r = left, right
	for l <= r {
		var mid = l + (r-l)/2
		if target == matrix[row][mid] {
			return mid, true
		} else if target < matrix[row][mid] {
			if mid == left {
				return mid, false
			}
			if target > matrix[row][mid-1] {
				return mid - 1, false
			}
			r = mid - 1
		} else {
			if mid == right {
				return mid, false
			}
			if target < matrix[row][mid] {
				return mid + 1, false
			}
			l = mid + 1
		}
	}
	return 0, false
}

func searchCol(matrix [][]int, col, up, down, target int) (int, bool) {
	var u, d = up, down
	for u <= d {
		var mid = u + (d-u)/2
		if target == matrix[mid][col] {
			return mid, true
		} else if target < matrix[mid][col] {
			if mid == up {
				return mid, false
			}
			if target > matrix[mid-1][col] {
				return mid - 1, false
			}
			d = mid - 1
		} else {
			if mid == down {
				return mid, false
			}
			if target < matrix[mid][col] {
				return mid + 1, false
			}
			u = mid + 1
		}
	}
	return 0, false
}
