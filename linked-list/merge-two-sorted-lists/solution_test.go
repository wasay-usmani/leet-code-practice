package mergetwosortedlists

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

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
		name     string
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}, "example1"},
		{[]int{}, []int{0}, []int{0}, "one empty"},
		{[]int{2}, []int{1}, []int{1, 2}, "single elements"},
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}, "alternating"},
	}

	for _, tc := range tests {
		l1 := sliceToList(tc.list1)
		l2 := sliceToList(tc.list2)
		got := MergeTwoLists(l1, l2)
		gotSlice := listToSlice(got)
		if !reflect.DeepEqual(gotSlice, tc.expected) {
			t.Errorf("%s: got %v, want %v", tc.name, gotSlice, tc.expected)
		}
	}
}
