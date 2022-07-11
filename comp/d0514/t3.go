package main

import (
	"sort"
)

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][1]
	})
	//fmt.Println(tiles)
	var worth []int
	var worthLen []int
	for i := 0; i < len(tiles); i++ {
		worth = append(worth, tiles[i][1]-tiles[i][0]+1)
		worthLen = append(worthLen, tiles[i][1]-tiles[i][0]+1)
		if i < len(tiles)-1 {
			worth = append(worth, 0)
			worthLen = append(worthLen, tiles[i+1][0]-tiles[i][1]-1)
		}
	}
	var res = 0
	var l, r = 0, 0
	var tmpLen, tmpWorth = 0, 0
	tmpLen = worthLen[0]
	tmpWorth = min(carpetLen, worth[0])
	for true {
		if tmpLen < carpetLen {
			r++
			if r == len(worthLen) {
				break
			}
			tmpLen += worthLen[r]
			tmpWorth += worth[r]
		} else {
			tmpLen -= worthLen[l]
			tmpWorth -= worth[l]
			l++
		}

		var realWorth = tmpWorth
		if tmpLen > carpetLen && worth[r] > 0 {
			realWorth -= tmpLen - carpetLen
		}
		if res < realWorth {
			res = realWorth
		}
	}
	//for i := 0; i < len(worthLen); i += 2 {
	//	var tmpLen, tmpW = 0, 0
	//	for j := i; j < len(worthLen) && tmpLen < carpetLen; j++ {
	//		tmpW += min(worth[j], carpetLen-tmpLen)
	//		tmpLen += worthLen[j]
	//	}
	//	if tmpW > res {
	//		res = tmpW
	//	}
	//}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
