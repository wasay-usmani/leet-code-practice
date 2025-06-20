package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle returns true if the linked list has a cycle.
func HasCycle(head *ListNode) bool {
	visited := map[*ListNode]struct{}{}
	curr := head
	for curr != nil {
		if _, ok := visited[curr]; ok {
			return true
		}

		visited[curr] = struct{}{}
		curr = curr.Next
	}

	return false
}
