package main

import (
	"learn_algorithm/tree/node"
)

func main() {
	//var root = node.TreeNode{Val: 10}
	//root.Left = &node.TreeNode{Val: 5}
	//root.Right = &node.TreeNode{Val: 15}

	var root = node.FormSlice([]int{10, 5, 15, 1, -1, 18, -1})
	root.Print()
}
