package array

//给你一个整数数组 nums 和一个整数k ，请你统计并返回 该数组中和为k的子数组的个数。

//示例 1：
//
//输入：nums = [1,1,1], k = 2
//输出：2
//示例 2：
//
//输入：nums = [1,2,3], k = 3
//输出：2

//1 <= nums.length <= 2 * 10^4
//-1000 <= nums[i] <= 1000
//-10^7 <= k <= 10^7

func SubarraySum(nums []int, k int) int {
	//连续问题考虑前缀
	var mp = make(map[int]int)
	mp[0] = 1
	var res = 0
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res += mp[sum-k]
		mp[sum] += 1
	}
	return res
}
