package linkedlist

import (
	"reflect"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	var head, curr *ListNode
	for _, v := range nums {
		node := &ListNode{Val: v}
		if head == nil {
			head = node
			curr = node
		} else {
			curr.Next = node
			curr = node
		}
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
		name     string
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, "example1"},
		{[]int{1, 2}, []int{2, 1}, "example2"},
		{[]int{42}, []int{42}, "single"},
	}

	for _, tc := range tests {
		head := sliceToList(tc.input)
		got := ReverseList(head)
		gotSlice := listToSlice(got)
		if !reflect.DeepEqual(gotSlice, tc.expected) {
			t.Errorf("%s: got %v, want %v", tc.name, gotSlice, tc.expected)
		}
	}
}
