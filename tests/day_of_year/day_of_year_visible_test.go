package main

import "testing"

func TestLevel_DayOfYear_Visible(t *testing.T) {
	tests := []struct {
		name     string
		day      int
		month    int
		year     int
		expected int
	}{
		{"example_1", 1, 3, 2024, 61},
		{"example_2", 1, 3, 2025, 60},
		{"example_3", 29, 2, 2025, -1},
		{"example_4", 29, 2, 2024, 60},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DayOfYear(tc.day, tc.month, tc.year)
			if got != tc.expected {
				t.Fatalf("DayOfYear(%d, %d, %d) = %d; want %d", tc.day, tc.month, tc.year, got, tc.expected)
			}
		})
	}
}
