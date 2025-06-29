package linkedlists

// ListNode is a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList reorders the linked list as per the problem statement.
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid := getMid(head)
	reversed := reverse(mid)

	c1 := head
	c2 := reversed
	var f1, f2 *ListNode

	for c1 != nil && c2 != nil {
		// Backup
		f1 = c1.Next
		f2 = c2.Next

		// Linking
		c1.Next = c2
		c2.Next = f1

		// Move
		c1 = f1
		c2 = f2
	}
}

func getMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
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
