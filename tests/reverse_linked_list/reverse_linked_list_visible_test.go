package main

import "testing"

func visibleBuild(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

func visibleToSlice(head *ListNode) []int {
	out := []int{}
	for curr := head; curr != nil && len(out) < 10000; curr = curr.Next {
		out = append(out, curr.Val)
	}
	return out
}

func visibleEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestLevel6_ReverseLinkedList_Visible(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		expected []int
	}{
		{"example_1_five_nodes", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"example_2_two_nodes", []int{1, 2}, []int{2, 1}},
		{"example_3_empty_list", []int{}, []int{}},
		{"single_node", []int{7}, []int{7}},
		{"duplicates", []int{1, 1, 2}, []int{2, 1, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := visibleToSlice(ReverseLinkedList(visibleBuild(tc.vals)))
			if !visibleEqual(got, tc.expected) {
				t.Fatalf("ReverseLinkedList(%v) = %v; want %v", tc.vals, got, tc.expected)
			}
		})
	}
}
