package main

func HeapSort(nums []int) []int {
	var minHeap MinHeap
	minHeap.Init(nums)
	var ret = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = minHeap.Top()
	}
	return ret
}
