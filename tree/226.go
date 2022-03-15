package main

import "learn_algorithm/tree/node"

func invertTree(root *node.TreeNode) *node.TreeNode {
	if root == nil {
		return root
	}
	var tmpNode = root.Left
	root.Left = root.Right
	root.Right = tmpNode
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
