package main

import "fmt"

func numSquares(n int) int {
	if n <= 1 {
		return 1
	}
	//strconv.FormatBool()
	var memo = make([]int, n+1) //对n进行拆分的最大乘积
	memo[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			memo[i] = max3(memo[i], (i-j)*j, (i-j)*memo[j])
		}
	}
	return memo[n]
}

func max3(x, y, z int) int {
	return max2(x, max2(y, z))
}

func max2(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func canPartition(nums []int) bool {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	var target = sum / 2
	// 考虑nums[i],[0,i]中是否有组合的和等于j
	var memo [][]bool
	for i := 0; i < len(nums); i++ {
		memo = append(memo, make([]bool, target+1))
	}
	memo[0][nums[0]] = true

	for i := 1; i < len(memo); i++ {
		for j := 0; j <= target; j++ {
			memo[i][j] = memo[i-1][j]
			if j >= nums[i] {
				memo[i][j] = memo[i][j] || memo[i-1][j-nums[i]]
			}
		}
	}
	fmt.Println(memo)
	return memo[len(nums)-1][target]
}
