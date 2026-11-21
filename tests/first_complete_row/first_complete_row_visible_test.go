package main

import "testing"

func TestLevel_FirstCompleteRow_Visible(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		called   []int
		expected int
	}{
		{"example_1", [][]int{{4, 1, 7}, {3, 9, 2}, {8, 5, 6}}, []int{9, 4, 1, 3, 7, 2}, 5},
		{"example_2", [][]int{{1}}, []int{1}, 1},
		{"example_3", [][]int{{1}}, []int{2}, -1},
		{"example_4", [][]int{{1, 2}, {3, 4}}, []int{1, 3, 2}, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FirstCompleteRow(tc.grid, tc.called)
			if got != tc.expected {
				t.Fatalf("FirstCompleteRow(%v, %v) = %d; want %d", tc.grid, tc.called, got, tc.expected)
			}
		})
	}
}
