package main

func insertSort(src []int) []int {
	for i := 1; i < len(src); i++ {
		for j := i; j > 0 && src[j] < src[j-1]; j-- {//[0,i-1]有序，src[i]向左移动到合适的位置是的[0,i]依然有序
			swap(j, j-1, src)
		}
	}
	return src
}
