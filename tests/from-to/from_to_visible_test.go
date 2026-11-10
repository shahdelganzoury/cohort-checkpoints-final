package main

import "testing"

func TestLevel1_FromTo_Visible(t *testing.T) {
	tests := []struct {
		name     string
		from, to int
		expected string
	}{
		{"example_1_ascending", 1, 10, "01, 02, 03, 04, 05, 06, 07, 08, 09, 10\n"},
		{"example_2_descending", 10, 1, "10, 09, 08, 07, 06, 05, 04, 03, 02, 01\n"},
		{"example_3_same_number", 10, 10, "10\n"},
		{"example_4_invalid_too_big", 100, 10, "Invalid\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FromTo(tc.from, tc.to)
			if got != tc.expected {
				t.Fatalf("FromTo(%d, %d) = %q; want %q", tc.from, tc.to, got, tc.expected)
			}
		})
	}
}
