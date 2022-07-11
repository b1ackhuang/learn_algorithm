package array

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func majorityElement(nums []int) int {
	return partitionElements(nums, len(nums)/2)
}

func partitionElements(nums []int, halfLen int) int {
	var l, r = partition(nums)
	if r-l+1 > halfLen {
		return nums[l]
	}
	if len(nums[:l]) > halfLen {
		return partitionElements(nums[:l], halfLen)
	}
	if len(nums[r+1:]) > halfLen {
		return partitionElements(nums[r+1:], halfLen)
	}
	return -1
}

func partition(nums []int) (int, int) {
	var target = nums[rand.Intn(len(nums))]
	var l, cur, r = -1, 0, len(nums)
	for cur < r {
		if nums[cur] < target {
			l++
			swap(nums, l, cur)
			cur++
		} else if nums[cur] > target {
			r--
			swap(nums, cur, r)
		} else {
			cur++
		}
	}
	return l + 1, r - 1
}
