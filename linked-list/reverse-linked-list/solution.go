package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverses a singly linked list and returns the new head.
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var curr, next *ListNode = head, head.Next
	for next != nil {
		tmp := next.Next
		next.Next = curr
		curr = next
		next = tmp
	}

	head.Next = nil
	return curr
}
