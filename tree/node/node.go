package node

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) Print() {
	fmt.Println(bfs(n))
}

// preOrder 前序遍历
func preOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, root.Val)
	ret = append(ret, preOrder(root.Left)...)
	ret = append(ret, preOrder(root.Right)...)
	return ret
}

// midOrder 中序遍历
func midOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, midOrder(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, midOrder(root.Right)...)
	return ret
}

// postOrder 后序遍历
func postOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, postOrder(root.Left)...)
	ret = append(ret, postOrder(root.Right)...)
	ret = append(ret, root.Val)
	return ret
}

func bfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	var curNodeList, nextNodeList []TreeNode
	curNodeList = append(curNodeList, *root)

	for len(curNodeList) != 0 {
		var curNode = curNodeList[0]
		ret = append(ret, curNode.Val)
		if curNode.Left != nil {
			nextNodeList = append(nextNodeList, *curNode.Left)
		}
		if curNode.Right != nil {
			nextNodeList = append(nextNodeList, *curNode.Right)
		}
		curNodeList = curNodeList[1:]
		if len(curNodeList) == 0 {
			curNodeList = nextNodeList
			nextNodeList = []TreeNode{}
		}
	}

	return ret
}

// FormSlice 从数组中构造二叉树，空节点用-1表示
func FormSlice(data []int) *TreeNode {
	if len(data) == 0 || data[0] == -1 {
		return nil
	}
	var root = &TreeNode{Val: data[0]}

	var curLevelNodeList, nextLevelNodeList []*TreeNode
	curLevelNodeList = append(curLevelNodeList, root)

	for i := 1; i < len(data); i += 2 {
		var curNode = curLevelNodeList[0]
		if curNode == nil { //叶子节点，不会有子树，直接跳过
			nextLevelNodeList = append(nextLevelNodeList, nil)
			nextLevelNodeList = append(nextLevelNodeList, nil)
		} else {
			curNode.Left = newNode(data[i])
			curNode.Right = newNode(data[i+1])
			nextLevelNodeList = append(nextLevelNodeList, curNode.Left)
			nextLevelNodeList = append(nextLevelNodeList, curNode.Right)
		}
		curLevelNodeList = curLevelNodeList[1:]
		if len(curLevelNodeList) == 0 {
			curLevelNodeList = nextLevelNodeList
			nextLevelNodeList = []*TreeNode{}
		}
	}
	return root
}

func newNode(val int) *TreeNode {
	if val == -1 {
		return nil
	}
	return &TreeNode{Val: val}
}
