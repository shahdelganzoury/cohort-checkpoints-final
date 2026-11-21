package main

import "testing"

func TestLevel_CompressString_Visible(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{"example_1", "", ""},
		{"example_2", "a", "a1"},
		{"example_3", "aaabbcccc", "a3b2c4"},
		{"example_4", "aabbaa", "a2b2a2"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CompressString(tc.s)
			if got != tc.expected {
				t.Fatalf("CompressString(%q) = %q; want %q", tc.s, got, tc.expected)
			}
		})
	}
}
