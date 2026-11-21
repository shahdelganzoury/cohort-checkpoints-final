package main

import "testing"

func TestLevel_TwoMissingNumbers_Visible(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [2]int
	}{
		{"example_1", nil, [2]int{1, 2}},
		{"example_2", []int{1, 3, 6, 2}, [2]int{4, 5}},
		{"example_3", []int{2, 4, 5, 6}, [2]int{1, 3}},
		{"example_4", []int{2, 3}, [2]int{1, 4}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, b := TwoMissingNumbers(tc.nums)
			got := [2]int{a, b}
			if got != tc.expected {
				t.Fatalf("TwoMissingNumbers(%v) = %v; want %v", tc.nums, got, tc.expected)
			}
		})
	}
}
