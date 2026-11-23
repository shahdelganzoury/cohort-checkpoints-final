package main

import "testing"

// spec: [value, index of the Random target or -1]
func visibleBuild(spec [][2]int) *ListNode {
	nodes := make([]*ListNode, len(spec))
	for i := range spec {
		nodes[i] = &ListNode{Val: spec[i][0]}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	for i, s := range spec {
		if s[1] >= 0 {
			nodes[i].Random = nodes[s[1]]
		}
	}
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}

// visibleEncode returns the shape of a list as value/random-index pairs.
func visibleEncode(head *ListNode) [][2]int {
	index := map[*ListNode]int{}
	i := 0
	for c := head; c != nil && i < 10000; c = c.Next {
		index[c] = i
		i++
	}
	out := [][2]int{}
	for c := head; c != nil && len(out) < 10000; c = c.Next {
		r := -1
		if c.Random != nil {
			if idx, ok := index[c.Random]; ok {
				r = idx
			} else {
				r = -2 // points outside the list
			}
		}
		out = append(out, [2]int{c.Val, r})
	}
	return out
}

func visibleSameShape(a, b [][2]int) bool {
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

func visibleShares(orig, cp *ListNode) bool {
	seen := map[*ListNode]bool{}
	for c := orig; c != nil; c = c.Next {
		seen[c] = true
	}
	for c := cp; c != nil; c = c.Next {
		if seen[c] || (c.Random != nil && seen[c.Random]) {
			return true
		}
	}
	return false
}

func TestLevel1_CopyRandomList_Visible(t *testing.T) {
	tests := []struct {
		name string
		spec [][2]int
	}{
		{"example_1_mixed_randoms", [][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}},
		{"example_2_equal_values", [][2]int{{3, -1}, {3, 0}, {3, -1}}},
		{"example_3_empty_list", [][2]int{}},
		{"two_nodes_random_to_last", [][2]int{{1, 1}, {2, 1}}},
		{"single_node_self_random", [][2]int{{5, 0}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			orig := visibleBuild(tc.spec)
			cp := CopyRandomList(orig)

			if len(tc.spec) == 0 {
				if cp != nil {
					t.Fatalf("an empty list must copy to nil")
				}
				return
			}
			if got := visibleEncode(cp); !visibleSameShape(got, tc.spec) {
				t.Fatalf("copy shape = %v; want %v", got, tc.spec)
			}
			if got := visibleEncode(orig); !visibleSameShape(got, tc.spec) {
				t.Fatalf("the original list was changed: %v; want %v", got, tc.spec)
			}
			if visibleShares(orig, cp) {
				t.Fatalf("the copy must not share any node (or Random target) with the original")
			}
		})
	}
}
