package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	var mergedArr = mergeSortedArrays(nums1, nums2)
	if len(mergedArr) == 0 {
		return 0
	}
	if len(mergedArr)%2 == 1 {
		return float64(mergedArr[len(mergedArr)/2])
	}
	return (float64(mergedArr[len(mergedArr)/2] + mergedArr[len(mergedArr)/2-1])) / 2.0
}

func mergeSortedArrays(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	var dstArr = make([]int, len(nums1)+len(nums2), len(nums1)+len(nums2))
	for i, cur1, cur2 := 0, 0, 0; i < len(dstArr); i++ {
		if cur1 == len(nums1) {
			dstArr[i] = nums2[cur2]
			cur2 += 1
			continue
		}
		if cur2 == len(nums2) {
			dstArr[i] = nums1[cur1]
			cur1 += 1
			continue
		}
		if nums1[cur1] <= nums2[cur2] {
			dstArr[i] = nums1[cur1]
			cur1 += 1
		} else {
			dstArr[i] = nums2[cur2]
			cur2 += 1
		}
	}
	return dstArr
}
