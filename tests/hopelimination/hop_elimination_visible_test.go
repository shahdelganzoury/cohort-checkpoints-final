package main

import "testing"

func visibleRing(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	tail := head
	for _, v := range vals[1:] {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	tail.Next = head
	return head
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

func TestLevel4_HopEliminate_Visible(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		expected []int
	}{
		{"example_1", []int{3, 1, 4, 1, 5}, []int{3, 1, 5, 1, 4}},
		{"example_2_huge_value", []int{1000000000, 7, 8}, []int{1000000000, 8, 7}},
		{"example_3_single_node", []int{9}, []int{9}},
		{"all_equal_values", []int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
		{"empty_circle", []int{}, []int{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := HopEliminate(visibleRing(tc.vals...))
			if !visibleEqual(got, tc.expected) {
				t.Fatalf("HopEliminate(%v) = %v; want %v", tc.vals, got, tc.expected)
			}
		})
	}
}
