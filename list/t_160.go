package main

func getIntersectionNode(ha, hb *ListNode) *ListNode {
	//构造环形链表
	var ta *ListNode
	for cur := ha; cur != nil; cur = cur.Next {
		ta = cur
	}
	ta.Next = hb
	defer func() {
		ta.Next = nil
	}()

	var slow, fast = ha, ha
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			var tmpNode = ha
			for tmpNode != slow {
				tmpNode = tmpNode.Next
				slow = slow.Next
			}
			return slow
		}
	}

	return nil
}
