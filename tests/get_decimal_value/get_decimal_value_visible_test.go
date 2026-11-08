package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

func visibleBuildBits(bits []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, b := range bits {
		tail.Next = &ListNode{Val: b}
		tail = tail.Next
	}
	return dummy.Next
}

func TestLevel1_GetDecimalValue_Visible(t *testing.T) {
	tests := []struct {
		name     string
		bits     []int
		expected int
	}{
		{"example_1_101", []int{1, 0, 1}, 5},
		{"example_2_single_zero", []int{0}, 0},
		{"single_one", []int{1}, 1},
		{"all_ones_three", []int{1, 1, 1}, 7},
		{"leading_zero", []int{0, 1, 0}, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetDecimalValue(visibleBuildBits(tc.bits))
			if got != tc.expected {
				t.Fatalf("GetDecimalValue(%v) = %d; want %d", tc.bits, got, tc.expected)
			}
		})
	}
}
