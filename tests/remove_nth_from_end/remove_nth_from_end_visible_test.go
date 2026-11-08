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

func TestLevel9_RemoveNthFromEnd_Visible(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		n        int
		expected []int
	}{
		{"example_1_second_from_end", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"example_2_single_node", []int{1}, 1, []int{}},
		{"example_3_remove_tail", []int{1, 2}, 1, []int{1}},
		{"remove_head", []int{1, 2}, 2, []int{2}},
		{"remove_head_of_five", []int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := visibleToSlice(RemoveNthFromEnd(visibleBuild(tc.vals), tc.n))
			if !visibleEqual(got, tc.expected) {
				t.Fatalf("RemoveNthFromEnd(%v, %d) = %v; want %v", tc.vals, tc.n, got, tc.expected)
			}
		})
	}
}
