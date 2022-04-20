package main

func swap(i, j int, src []int) {
	var tmp = src[i]
	src[i] = src[j]
	src[j] = tmp
}
