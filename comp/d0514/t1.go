package main

func divisorSubstrings(num int, k int) int {
	var tmpNum = num
	var nums = make([]int, 12)
	var base = 10
	var endIndex = 0
	for ; endIndex < 12 && num > 0; endIndex++ {
		nums[endIndex] = num % base
		num /= base
	}
	nums = nums[:endIndex]
	var wNum = 0
	var mul = 1
	for i := 0; i < k; i++ {
		wNum += mul * nums[i]
		mul *= base
	}
	mul /= 10

	var res = 0
	for i := k; i <= len(nums); i++ {
		if wNum != 0 && tmpNum%wNum == 0 {
			res++
		}
		if i < len(nums) {
			wNum /= base
			wNum += nums[i] * mul
		}
	}
	return res
}
