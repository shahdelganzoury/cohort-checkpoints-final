package main

import "testing"

func TestLevel_DigitPositionMatch_Visible(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{"example_1", "58342", "57382", 3},
		{"example_2", "", "", 0},
		{"example_3", "000", "000", 3},
		{"example_4", "123", "456", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DigitPositionMatch(tc.a, tc.b)
			if got != tc.expected {
				t.Fatalf("DigitPositionMatch(%q, %q) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
