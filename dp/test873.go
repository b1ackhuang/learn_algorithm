package main

import "fmt"

func lenLongestFibSubseq(arr []int) int {
	var memo = make([][]int, len(arr)) //memo[i][j]: 以arr[i]arr[j]结尾的fb数列
	for i := 0; i < len(arr); i++ {
		memo[i] = make([]int, len(arr))
	}
	var res = 0
	for i := 2; i < len(arr); i++ {
		for j := i - 1; j >= 1; j-- {
			var preSeqNumIndex = binarySearch(arr[:j], arr[i]-arr[j])
			fmt.Println(j, i, arr[:j], arr[i]-arr[j], "-->> ", preSeqNumIndex)
			if preSeqNumIndex >= 0 {
				memo[j][i] = max(memo[preSeqNumIndex][j]+1, 3)
			}
			res = max(res, memo[j][i])
		}
	}
	for i := 0; i < len(memo); i++ {
		fmt.Println(memo[i])
	}
	return res
}

func binarySearch(nums []int, target int) int {
	var l, r = 0, len(nums) - 1
	for l <= r {
		var mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
