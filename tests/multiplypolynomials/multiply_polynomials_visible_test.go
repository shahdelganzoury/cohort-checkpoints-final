package main

import "testing"

// terms are {Coef, Exp} pairs in descending exponent order
func visibleBuild(terms [][2]int) *ListNode {
	var head *ListNode
	for i := len(terms) - 1; i >= 0; i-- {
		head = &ListNode{Coef: terms[i][0], Exp: terms[i][1], Next: head}
	}
	return head
}

func visibleRead(h *ListNode) [][2]int {
	out := [][2]int{}
	for ; h != nil && len(out) < 20000; h = h.Next {
		out = append(out, [2]int{h.Coef, h.Exp})
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

func TestLevel6_MultiplyPoly_Visible(t *testing.T) {
	tests := []struct {
		name     string
		a, b     [][2]int
		expected [][2]int
	}{
		{"example_1", [][2]int{{2, 2}, {3, 0}}, [][2]int{{1, 1}, {1, 0}}, [][2]int{{2, 3}, {2, 2}, {3, 1}, {3, 0}}},
		{"example_2_terms_cancel", [][2]int{{1, 1}, {1, 0}}, [][2]int{{1, 1}, {-1, 0}}, [][2]int{{1, 2}, {-1, 0}}},
		{"example_3_zero_polynomial", nil, [][2]int{{5, 3}}, [][2]int{}},
		{"big_exponents", [][2]int{{3, 5000}}, [][2]int{{-2, 5000}, {4, 0}}, [][2]int{{-6, 10000}, {12, 5000}}},
		{"constants", [][2]int{{7, 0}}, [][2]int{{-1, 0}}, [][2]int{{-7, 0}}},
		{"partial_cancellation", [][2]int{{1, 2}, {1, 1}}, [][2]int{{1, 1}, {-1, 0}}, [][2]int{{1, 3}, {-1, 1}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := visibleRead(MultiplyPoly(visibleBuild(tc.a), visibleBuild(tc.b)))
			if !visibleEqual(got, tc.expected) {
				t.Fatalf("MultiplyPoly(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestLevel6_MultiplyPoly_InputsUnchanged_Visible(t *testing.T) {
	a := [][2]int{{2, 2}, {3, 0}}
	b := [][2]int{{1, 1}, {1, 0}}
	headA, headB := visibleBuild(a), visibleBuild(b)

	MultiplyPoly(headA, headB)

	if got := visibleRead(headA); !visibleEqual(got, a) {
		t.Fatalf("the first input was modified: %v; want %v", got, a)
	}
	if got := visibleRead(headB); !visibleEqual(got, b) {
		t.Fatalf("the second input was modified: %v; want %v", got, b)
	}
}
