package prb0142

type ListNode struct {
	Val  int
	Next *ListNode
}

// 3, 2, 0, -4
func detectCycle(head *ListNode) *ListNode {

	slow, fast := head, head

	for {
		fast = stepTwo(fast)
		if fast == nil {
			return nil
		}
		slow = slow.Next

		if slow == fast {
			break
		}
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func stepTwo(node *ListNode) *ListNode {
	if node != nil && node.Next != nil && node.Next.Next != nil {
		return node.Next.Next
	} else {
		return nil
	}
}
