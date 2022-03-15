package main

import (
	"learn_algorithm/tree/node"
	"math"
)

func maxDepth(root *node.TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}
