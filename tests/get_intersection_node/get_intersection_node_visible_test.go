package main

import "testing"

func visibleBuildOnto(vals []int, tail *ListNode) *ListNode {
	head := tail
	for i := len(vals) - 1; i >= 0; i-- {
		head = &ListNode{Val: vals[i], Next: head}
	}
	return head
}

func TestLevel7_GetIntersectionNode_Visible(t *testing.T) {
	shared := visibleBuildOnto([]int{8, 4, 5}, nil)
	a := visibleBuildOnto([]int{4, 1}, shared)
	b := visibleBuildOnto([]int{5, 6, 1}, shared)

	same1 := visibleBuildOnto([]int{1, 2, 3}, nil)
	same2 := visibleBuildOnto([]int{1, 2, 3}, nil)

	whole := visibleBuildOnto([]int{1, 2, 3}, nil)

	tests := []struct {
		name     string
		a, b     *ListNode
		expected *ListNode
	}{
		{"example_1_intersect_at_8", a, b, shared},
		{"example_2_no_intersection", visibleBuildOnto([]int{2, 6, 4}, nil), visibleBuildOnto([]int{1, 5}, nil), nil},
		{"equal_values_different_nodes", same1, same2, nil},
		{"same_list", whole, whole, whole},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetIntersectionNode(tc.a, tc.b)
			if got != tc.expected {
				t.Fatalf("GetIntersectionNode returned %p; want %p", got, tc.expected)
			}
		})
	}
}
