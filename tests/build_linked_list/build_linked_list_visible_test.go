package main

import "testing"

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

func TestLevel5_BuildLinkedList_Visible(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"example_1_three_values", []int{3, 7, 9}, []int{3, 7, 9}},
		{"example_2_empty_slice", []int{}, []int{}},
		{"nil_slice", nil, []int{}},
		{"single_value", []int{42}, []int{42}},
		{"negative_values", []int{-1, -2}, []int{-1, -2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := BuildLinkedList(tc.nums)
			if len(tc.nums) == 0 && head != nil {
				t.Fatalf("BuildLinkedList(%v) must return nil for an empty input", tc.nums)
			}
			got := visibleToSlice(head)
			if !visibleEqual(got, tc.expected) {
				t.Fatalf("BuildLinkedList(%v) = %v; want %v", tc.nums, got, tc.expected)
			}
		})
	}
}
