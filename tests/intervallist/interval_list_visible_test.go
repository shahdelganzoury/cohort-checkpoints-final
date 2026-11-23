package main

import "testing"

func visibleBuild(iv [][2]int) *ListNode {
	var head *ListNode
	for i := len(iv) - 1; i >= 0; i-- {
		head = &ListNode{Start: iv[i][0], End: iv[i][1], Next: head}
	}
	return head
}

func visibleRead(h *ListNode) [][2]int {
	out := [][2]int{}
	for ; h != nil && len(out) < 10000; h = h.Next {
		out = append(out, [2]int{h.Start, h.End})
	}
	return out
}

func visibleEqual(a, b [][2]int) bool {
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

func TestLevel5_AddInterval_Visible(t *testing.T) {
	tests := []struct {
		name       string
		iv         [][2]int
		start, end int
		expected   [][2]int
	}{
		{"example_1_touching_merge", [][2]int{{1, 3}, {5, 6}, {10, 12}}, 3, 4, [][2]int{{1, 4}, {5, 6}, {10, 12}}},
		{"example_2_swallow_everything", [][2]int{{1, 4}, {5, 6}, {7, 7}, {8, 9}, {10, 12}}, 2, 11, [][2]int{{1, 12}}},
		{"example_3_empty_list", [][2]int{}, 4, 4, [][2]int{{4, 4}}},
		{"new_interval_at_the_front", [][2]int{{5, 6}}, 1, 2, [][2]int{{1, 2}, {5, 6}}},
		{"new_interval_at_the_back", [][2]int{{1, 2}}, 8, 9, [][2]int{{1, 2}, {8, 9}}},
		{"neighbours_that_do_not_touch_stay_apart", [][2]int{{1, 2}}, 3, 4, [][2]int{{1, 2}, {3, 4}}},
		{"start_after_end_changes_nothing", [][2]int{{1, 12}, {13, 13}}, 9, 1, [][2]int{{1, 12}, {13, 13}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := visibleRead(AddInterval(visibleBuild(tc.iv), tc.start, tc.end))
			if !visibleEqual(got, tc.expected) {
				t.Fatalf("AddInterval(%v, [%d,%d]) = %v; want %v", tc.iv, tc.start, tc.end, got, tc.expected)
			}
		})
	}
}
