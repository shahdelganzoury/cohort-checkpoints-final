package main

import "testing"

func TestLevel1_WordFlip_Visible(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"example_1_basic", "First second last", "last second First\n"},
		{"example_2_empty", "", "Invalid Output\n"},
		{"example_3_only_spaces", "     ", "\n"},
		{"example_4_extra_spaces", " hello  all  of  you! ", "you! of all hello\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := WordFlip(tc.input)
			if got != tc.expected {
				t.Fatalf("WordFlip(%q) = %q; want %q", tc.input, got, tc.expected)
			}
		})
	}
}
