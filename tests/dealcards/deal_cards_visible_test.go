package main

import "testing"

func visibleBuild(vals ...int) *ListNode {
	var head *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		head = &ListNode{Val: vals[i], Next: head}
	}
	return head
}

func visibleHands(hands []*ListNode) [][]int {
	out := make([][]int, len(hands))
	for i, h := range hands {
		out[i] = []int{}
		for ; h != nil && len(out[i]) < 10000; h = h.Next {
			out[i] = append(out[i], h.Val)
		}
	}
	return out
}

func visibleEqualHands(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestLevel3_DealCards_Visible(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		k        int
		expected [][]int
	}{
		{"example_1_three_hands", []int{1, 2, 3, 4, 5, 6, 7}, 3, [][]int{{1, 4, 7}, {2, 5}, {3, 6}}},
		{"example_2_more_hands_than_nodes", []int{10, 20}, 4, [][]int{{10}, {20}, {}, {}}},
		{"example_3_empty_list", []int{}, 2, [][]int{{}, {}}},
		{"single_hand_keeps_everything", []int{5, 6, 7}, 1, [][]int{{5, 6, 7}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := visibleHands(DealCards(visibleBuild(tc.vals...), tc.k))
			if !visibleEqualHands(got, tc.expected) {
				t.Fatalf("DealCards(%v, %d) = %v; want %v", tc.vals, tc.k, got, tc.expected)
			}
		})
	}
}

func TestLevel3_DealCards_NonPositiveK_Visible(t *testing.T) {
	for _, k := range []int{0, -1} {
		if got := DealCards(visibleBuild(1, 2), k); got != nil {
			t.Fatalf("DealCards(list, %d) must return nil; got %v", k, got)
		}
	}
}
