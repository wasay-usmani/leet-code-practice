package linkedlistcycle

import "testing"

func makeListWithCycle(vals []int, pos int) *ListNode {
	var head, curr, cycleNode *ListNode
	for i, v := range vals {
		node := &ListNode{Val: v}
		if head == nil {
			head = node
		} else {
			curr.Next = node
		}
		curr = node
		if i == pos {
			cycleNode = node
		}
	}
	if curr != nil && cycleNode != nil {
		curr.Next = cycleNode
	}
	return head
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		vals     []int
		pos      int
		hasCycle bool
		name     string
	}{
		{[]int{3, 2, 0, -4}, 1, true, "example1"},
		{[]int{1, 2}, 0, true, "example2"},
		{[]int{1}, -1, false, "example3"},
		{[]int{}, -1, false, "empty"},
		{[]int{1, 2, 3, 4, 5}, -1, false, "no cycle"},
		{[]int{1, 2, 3, 4, 5}, 2, true, "cycle at 2"},
	}

	for _, tc := range tests {
		head := makeListWithCycle(tc.vals, tc.pos)
		got := HasCycle(head)
		if got != tc.hasCycle {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.hasCycle)
		}
	}
}
