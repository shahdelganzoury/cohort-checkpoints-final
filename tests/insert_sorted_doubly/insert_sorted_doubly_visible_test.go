package main

import "testing"

func visibleBuildDoubly(vals []int) *DoublyListNode {
	var head, tail *DoublyListNode
	for _, v := range vals {
		n := &DoublyListNode{Val: v}
		if head == nil {
			head = n
		} else {
			tail.Next = n
			n.Prev = tail
		}
		tail = n
	}
	return head
}

func visibleForwardBackward(head *DoublyListNode) ([]int, []int) {
	forward, backward := []int{}, []int{}
	var tail *DoublyListNode
	for curr := head; curr != nil && len(forward) < 10000; curr = curr.Next {
		forward = append(forward, curr.Val)
		tail = curr
	}
	for curr := tail; curr != nil && len(backward) < 10000; curr = curr.Prev {
		backward = append(backward, curr.Val)
	}
	return forward, backward
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

func TestLevel10_InsertSortedDoubly_Visible(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		val      int
		expected []int
	}{
		{"example_1_new_head", []int{3, 7}, 1, []int{1, 3, 7}},
		{"example_2_middle", []int{3, 7}, 5, []int{3, 5, 7}},
		{"example_3_new_tail", []int{3, 7}, 9, []int{3, 7, 9}},
		{"example_4_empty_list", []int{}, 5, []int{5}},
		{"single_node_before", []int{5}, 1, []int{1, 5}},
		{"single_node_after", []int{5}, 9, []int{5, 9}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := InsertSortedDoubly(visibleBuildDoubly(tc.vals), tc.val)
			forward, backward := visibleForwardBackward(head)
			if !visibleEqual(forward, tc.expected) {
				t.Fatalf("forward = %v; want %v", forward, tc.expected)
			}
			reversed := make([]int, len(tc.expected))
			for i, v := range tc.expected {
				reversed[len(tc.expected)-1-i] = v
			}
			if !visibleEqual(backward, reversed) {
				t.Fatalf("backward = %v; want %v (check the Prev pointers)", backward, reversed)
			}
		})
	}
}
