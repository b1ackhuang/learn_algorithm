package main

func waysToSplitArray(nums []int) int {
	var lSum = make([]int, len(nums))
	lSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		lSum[i] = lSum[i-1] + nums[i]
	}
	var rSum = make([]int, len(nums))
	rSum[len(rSum)-1] = 0
	for i := len(rSum) - 2; i >= 0; i-- {
		rSum[i] = rSum[i+1] + nums[i+1]
	}
	var res = 0
	for i := 0; i < len(nums)-1; i++ {
		if lSum[i] >= rSum[i] {
			res++
		}
	}
	return res
}
