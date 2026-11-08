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

func TestLevel3_DeleteNode_Visible(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		val      int
		expected []int
	}{
		{"example_1_delete_head", []int{1, 2, 3}, 1, []int{2, 3}},
		{"example_2_delete_middle", []int{1, 2, 3}, 2, []int{1, 3}},
		{"example_3_not_found", []int{1, 2, 3}, 9, []int{1, 2, 3}},
		{"delete_tail", []int{1, 2, 3}, 3, []int{1, 2}},
		{"only_first_match_removed", []int{1, 2, 2}, 2, []int{1, 2}},
		{"empty_list", []int{}, 1, []int{}},
		{"single_node_deleted", []int{1}, 1, []int{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := visibleToSlice(DeleteNode(visibleBuild(tc.vals), tc.val))
			if !visibleEqual(got, tc.expected) {
				t.Fatalf("DeleteNode(%v, %d) = %v; want %v", tc.vals, tc.val, got, tc.expected)
			}
		})
	}
}
