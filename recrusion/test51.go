package main

import "strings"

var res [][]string

func NQueens(n int) [][]string {
	res = nil
	dfs(make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), 0, n, []int{})
	return res
}

func dfs(col, lTilt, rTilt []bool, curLevel, n int, path []int) {
	if len(path) == n {
		res = append(res, drawCheekMap(path))
		return
	}
	for i := 0; i < n; i++ {
		if !col[i] && !lTilt[i-curLevel+n-1] && !rTilt[i+curLevel] {
			col[i] = true
			lTilt[i-curLevel+n-1] = true
			rTilt[i+curLevel] = true
			path = append(path, i)
			dfs(col, lTilt, rTilt, curLevel+1, n, path)
			path = path[0 : len(path)-1]
			col[i] = false
			lTilt[i-curLevel+n-1] = false
			rTilt[i+curLevel] = false
		}
	}
}

func drawCheekMap(points []int) []string {
	var cheekMap = make([]string, len(points))
	for i := 0; i < len(points); i++ {
		var s strings.Builder
		for j := 0; j < len(points); j++ {
			if j == points[i] {
				s.WriteByte('Q')
			} else {
				s.WriteByte('.')
			}
		}
		cheekMap[i] = s.String()
	}
	return cheekMap
}
