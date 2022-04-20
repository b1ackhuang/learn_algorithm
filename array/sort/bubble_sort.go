package main

func bubbleSort(src []int) []int {
	if len(src) == 0 {
		return src
	}
	for i := 1; i < len(src); i++ {
		for j := 0; j < len(src)-i; j++ { //[0,len(src)-i) 从左到右将最大值移动到最右端
			if src[j] > src[j+1] {
				swap(j, j+1, src)
			}
		}
	}
	return src
}
