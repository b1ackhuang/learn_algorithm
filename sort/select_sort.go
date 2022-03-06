package main

func selectSort(src []int) []int {
	if len(src) == 0 {
		return src
	}
	for i := 0; i < len(src); i++ {
		for j := i + 1; j < len(src); j++ {
			if src[i] > src[j] {
				swap(i, j, src)
			}
		}
	}
	return src
}
