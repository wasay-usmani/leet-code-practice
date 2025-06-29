package linkedlists

import (
	"reflect"
	"testing"
)

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestReorderList(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"even", []int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{"odd", []int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		{"single", []int{1}, []int{1}},
		{"two", []int{1, 2}, []int{1, 2}},
		{"empty", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			reorderList(head)
			got := listToSlice(head)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("reorderList() = %v, want %v", got, tt.expect)
			}
		})
	}
}
