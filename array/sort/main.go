package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var ins = ""
	_, _ = fmt.Scanln(&ins)
	var inputSlice = strings.Split(ins, ",")
	var src = make([]int, 0, len(inputSlice))
	for i := 0; i < len(inputSlice); i++ {
		var tmpVal, _ = strconv.ParseInt(inputSlice[i], 10, 64)
		src = append(src, int(tmpVal))
	}
	//var dst bubbleSort(src)
	//var dst selectSort(src)
	//var dst = insertSort(src)
	//math.Min(float64())
	//var dst = mergeSort(src)
	var dst = HeapSort(src)
	fmt.Println(dst)
}

const InvalidMaxVal = int(10e9 + 1)

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var minRes = InvalidMaxVal
	var l, r = 0, len(nums) - 1
	var tmpSum = nums[0]
	for r < len(nums) {
		if tmpSum == target {
			minRes = int(math.Min(float64(minRes), float64(r-l+1)))
			tmpSum -= nums[l]
			l++
		} else if tmpSum < target {
			r++
			tmpSum += nums[r]
		} else { //tmpSum > target
			if l < r {
				tmpSum -= nums[l]
				l++
			} else {
				r++
				tmpSum += nums[r]
			}
		}
	}
	return minRes
}
