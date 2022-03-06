package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var ins = ""
	_, _ = fmt.Scanln(&ins)
	var inputSlice = strings.Split(ins, ",")
	var src = make([]int, 0, len(inputSlice))
	for i := 0; i < len(inputSlice); i++ {
		var tmpVal, _ = strconv.ParseInt(inputSlice[i], 10, 64)
		src = append(src, int(tmpVal))
	}
	//var dst bubbleSort(src)
	//var dst selectSort(src)
	//var dst = insertSort(src)
	var dst = mergeSort(src)
	fmt.Println(dst)
}
