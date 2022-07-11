package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var pre, cur, next = (*ListNode)(nil), head, (*ListNode)(nil)
	for ; cur != nil; cur = next {
		next = cur.Next
		cur.Next = pre
		pre = cur
	}
	return pre
}
