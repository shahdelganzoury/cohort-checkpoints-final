package main

import "testing"

func visibleBuild(s string) *ListNode {
	var head, tail *ListNode
	for i := 0; i < len(s); i++ {
		n := &ListNode{Val: s[i], Prev: tail}
		if tail == nil {
			head = n
		} else {
			tail.Next = n
		}
		tail = n
	}
	return head
}

// visibleRead returns the forward and backward text of a list.
func visibleRead(h *ListNode) (string, string) {
	fwd, back := []byte{}, []byte{}
	var last *ListNode
	for ; h != nil && len(fwd) < 100000; h = h.Next {
		fwd = append(fwd, h.Val)
		last = h
	}
	for ; last != nil && len(back) < 100000; last = last.Prev {
		back = append(back, last.Val)
	}
	return string(fwd), string(back)
}

func visibleReverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func TestLevel2_Crossover_Visible(t *testing.T) {
	tests := []struct {
		name         string
		a, b         string
		i, j         int
		wantA, wantB string
	}{
		{"example_1_swap_tails", "AAAA", "bbbbb", 2, 1, "AAbbbb", "bAA"},
		{"example_2_a_becomes_empty", "ABC", "xyz", 0, 3, "", "xyzABC"},
		{"example_3_i_past_the_end", "ABC", "xy", 10, 1, "ABCy", "x"},
		{"a_is_nil", "", "xy", 0, 1, "y", "x"},
		{"b_is_nil", "Q", "", 1, 0, "Q", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotA, gotB := Crossover(visibleBuild(tc.a), visibleBuild(tc.b), tc.i, tc.j)
			for _, side := range []struct {
				name string
				head *ListNode
				want string
			}{{"a", gotA, tc.wantA}, {"b", gotB, tc.wantB}} {
				fwd, back := visibleRead(side.head)
				if fwd != side.want {
					t.Fatalf("%s forward = %q; want %q", side.name, fwd, side.want)
				}
				if back != visibleReverse(side.want) {
					t.Fatalf("%s backward = %q; want %q (check the Prev pointers)", side.name, back, visibleReverse(side.want))
				}
			}
		})
	}
}
