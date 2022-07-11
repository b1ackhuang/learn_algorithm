package main

import (
	"fmt"
	"sort"
)

func CombinationSum4(nums []int, target int) int {
	// [i,j]表示考虑第[0,i]个数时，组合数和=target的组合数
	var memo = make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = make([]int, target+1)
	}

	sort.Ints(nums)
	for i := 1; i*nums[0] <= target; i++ {
		memo[0][i*nums[0]] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			//memo[i][j] = memo[i-1][j]
			for k := 0; k*nums[i] < j; k++ {
				memo[i][j] += (k + 1) * memo[i-1][j-k*nums[i]]
			}
			if j%nums[i] == 0 {
				memo[i][j] += 1
			}
		}
	}

	fmt.Println(memo)
	return memo[len(nums)-1][target]
}
