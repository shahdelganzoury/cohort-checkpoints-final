package main

import "testing"

func TestLevel9_CountOccurrences_Visible(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"example_1_multiple_matches", []int{1, 4, 2, 4, 5, 4, 3}, 4, 3},
		{"example_2_target_absent", []int{10, 20, 30}, 99, 0},
		{"example_3_empty_slice", []int{}, 5, 0},
		{"single_match", []int{42}, 42, 1},
		{"single_non_match", []int{42}, 10, 0},
		{"all_matching", []int{7, 7, 7, 7}, 7, 4},
		{"negative_target", []int{-5, 0, -5, 5, -5}, -5, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountOccurrences(tc.nums, tc.target)
			if got != tc.expected {
				t.Fatalf("CountOccurrences(%v, %d) = %d; want %d", tc.nums, tc.target, got, tc.expected)
			}
		})
	}
}
