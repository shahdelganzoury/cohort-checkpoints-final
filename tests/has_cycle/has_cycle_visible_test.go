package main

import (
	"testing"
	"time"
)

func visibleBuildCycle(vals []int, pos int) *ListNode {
	var head, tail *ListNode
	nodes := []*ListNode{}
	for _, v := range vals {
		n := &ListNode{Val: v}
		nodes = append(nodes, n)
		if head == nil {
			head = n
		} else {
			tail.Next = n
		}
		tail = n
	}
	if pos >= 0 && pos < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return head
}

func TestLevel8_HasCycle_Visible(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		pos      int
		expected bool
	}{
		{"example_1_cycle_to_index_1", []int{3, 2, 0, -4}, 1, true},
		{"example_2_cycle_to_head", []int{1, 2}, 0, true},
		{"example_3_single_no_cycle", []int{1}, -1, false},
		{"empty_list", []int{}, -1, false},
		{"no_cycle_long", []int{1, 2, 3, 4, 5}, -1, false},
		{"equal_values_no_cycle", []int{1, 1, 1}, -1, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := visibleBuildCycle(tc.vals, tc.pos)
			done := make(chan bool, 1)
			go func() { done <- HasCycle(head) }()

			select {
			case got := <-done:
				if got != tc.expected {
					t.Fatalf("HasCycle(%v, cycle at %d) = %v; want %v", tc.vals, tc.pos, got, tc.expected)
				}
			case <-time.After(2 * time.Second):
				t.Fatalf("HasCycle(%v, cycle at %d) timed out (infinite loop?)", tc.vals, tc.pos)
			}
		})
	}
}
