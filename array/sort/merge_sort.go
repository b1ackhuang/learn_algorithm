package main

func mergeSort(src []int) []int {
	if len(src) <= 1 {
		return src
	}
	var mid = len(src) / 2
	var left = make([]int, 0, mid)
	var right = make([]int, 0, len(src)-mid)
	left = append(left, src[:mid]...)
	right = append(right, src[mid:]...)
	left = mergeSort(left)
	right = mergeSort(right)

	var dst []int
	var i, j = 0, 0
	for k := 0; i < len(left) && j < len(right); k++ {
		if left[i] < right[j] {
			dst = append(dst, left[i])
			i++
		} else {
			dst = append(dst, right[j])
			j++
		}
	}
	dst = append(dst, left[i:]...)
	dst = append(dst, right[j:]...)

	return dst
}
