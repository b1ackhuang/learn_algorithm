package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func partation(head *ListNode) (resLHead, resLTail, resMidHead, resMidTail, resRHead *ListNode) {
	var target = head.Val
	var dummyLHead, dummyMidHead, dummyRightHead = &ListNode{}, &ListNode{}, &ListNode{}
	var lTail, mTail, rTail = dummyLHead, dummyMidHead, dummyRightHead
	for cur := head; cur != nil; {
		var next = cur.Next
		if cur.Val == target {
			mTail.Next = cur
			mTail = mTail.Next
		} else if cur.Val < target {
			lTail.Next = cur
			lTail = lTail.Next
		} else {
			rTail.Next = cur
			rTail = rTail.Next
		}
		cur = next
	}
	resLHead = dummyLHead.Next
	resLTail = lTail
	resMidHead = dummyMidHead.Next
	resMidTail = mTail
	resRHead = dummyRightHead.Next
	return
}

func printList(l *ListNode) {
	for cur := l; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val, "->")
	}
}

func copySlice(src []string) []string {
	var dst = make([]string, len(src))
	copy(dst, src)
	return dst
}
