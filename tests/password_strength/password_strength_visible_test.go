package main

import "testing"

func TestLevel_PasswordStrength_Visible(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example_1", "", 0},
		{"example_2", "abc", 0},
		{"example_3", "Go2026!", 3},
		{"example_4", "GoLang2026!", 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := PasswordStrength(tc.s)
			if got != tc.expected {
				t.Fatalf("PasswordStrength(%q) = %d; want %d", tc.s, got, tc.expected)
			}
		})
	}
}
