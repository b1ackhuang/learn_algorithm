package array

func moveZeroes(nums []int) {
	var l, r = 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
		r++
	}
	l++
	for l < len(nums) {
		nums[l] = 0
		l++
	}
}

func swap(nums []int, x, y int) {
	var tmp = nums[x]
	nums[x] = nums[y]
	nums[y] = tmp
}
