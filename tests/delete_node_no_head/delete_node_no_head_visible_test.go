package main

import "testing"

func visibleBuild(vals []int) (*ListNode, []*ListNode) {
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
	return head, nodes
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

func TestLevel4_DeleteNodeNoHead_Visible(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		idx      int
		expected []int
	}{
		{"example_1_delete_5", []int{4, 5, 1, 9}, 1, []int{4, 1, 9}},
		{"example_2_delete_1", []int{4, 5, 1, 9}, 2, []int{4, 5, 9}},
		{"delete_first_node", []int{4, 5, 1, 9}, 0, []int{5, 1, 9}},
		{"two_nodes", []int{1, 2}, 0, []int{2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := visibleBuild(tc.vals)
			DeleteNodeNoHead(nodes[tc.idx])
			got := visibleToSlice(head)
			if !visibleEqual(got, tc.expected) {
				t.Fatalf("after deleting index %d of %v got %v; want %v", tc.idx, tc.vals, got, tc.expected)
			}
		})
	}
}
