package main

import "testing"

func TestLevel_DigitalRoot_Visible(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"example_1", 0, 0},
		{"example_2", 9, 9},
		{"example_3", 10, 1},
		{"example_4", 99, 9},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DigitalRoot(tc.n)
			if got != tc.expected {
				t.Fatalf("%s failed\ncall: DigitalRoot(%#v)\ngot: %#v\nrequired: %#v", tc.name, tc.n, got, tc.expected)
			}
		})
	}
}
