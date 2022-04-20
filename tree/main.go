package main

import (
	"fmt"
)

func main() {
	//var root = node.TreeNode{Val: 10}
	//root.Left = &node.TreeNode{Val: 5}
	//root.Right = &node.TreeNode{Val: 15}

	//var root = node.FormSlice([]int{10, 5, 15, 1, -1, 18, -1})
	//root.Print()

	//fmt.Println(strings.Split(" ", ","))
	//sort.Ints()
	var bts = []int{1, 2, 3, 4, 5}
	fmt.Println(bts[0:0])
	fmt.Println(bts[0:1])
	fmt.Println(bts[5:])
	fmt.Println(bts[6:])
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverseList(l1), reverseList(l2)
	var dummyRes = &ListNode{}
	var isCarry = false
	for l1 != nil && l2 != nil {
		var tmp = 0
		if isCarry {
			tmp = 1
		}
		tmp += l1.Val + l2.Val
		if tmp >= 10 {
			isCarry = true
			tmp %= 10
		} else {
			isCarry = false
		}
		dummyRes.Next = &ListNode{Val: tmp}
		l1 = l1.Next
		l2 = l2.Next
	}
	return dummyRes.Next
}

func reverseList(l *ListNode) *ListNode {
	if l == nil {
		return l
	}
	var pre, cur, next = (*ListNode)(nil), l, l.Next
	for ; cur != nil; cur = next {
		next = cur.Next
		cur.Next = pre
		pre = cur
	}
	return pre
}
