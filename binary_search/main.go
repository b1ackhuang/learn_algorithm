package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inStrArr string
	_, _ = fmt.Scanln(&inStrArr)
	var arrStr = strings.Split(inStrArr, ",")
	var arr []int
	for i := 0; i < len(arrStr); i++ {
		var tmpVal, _ = strconv.ParseInt(arrStr[i], 10, 64)
		arr = append(arr, int(tmpVal))
	}
	var inStrDst string
	_, _ = fmt.Scanln(&inStrDst)
	var dst, _ = strconv.ParseInt(inStrDst, 10, 64)
	fmt.Println(find(arr, int(dst)))
}

func find(arr []int, dst int) int {
	if len(arr) <= 0 {
		return -1
	}
	//在循环中保证循环的不变量
	var left, right = 0, len(arr) - 1 // [left, right] 区间内，查找目标元素
	for left <= right {
		var mid = (left + right) / 2
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
