package main

import "fmt"

func main() {
	//["MinStack","push","push","push","getMin","pop","top","getMin"]
	//[[],[-2],[0],[-3],[],[],[],[]]

	var minStack = Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())

}
