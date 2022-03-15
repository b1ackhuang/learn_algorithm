package main

import "learn_algorithm/tree/node"

func isSymmetric(root *node.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTowTree(root.Left, root.Right)
}

func isSymmetricTowTree(p *node.TreeNode, q *node.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	return (p.Val == q.Val) && isSymmetricTowTree(p.Left, q.Right) && isSymmetricTowTree(p.Right, q.Left)
}
