package main

import "testing"

func TestLevel_CashRegister_Visible(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		expected int
	}{
		{"example_1", 0, 0},
		{"example_2", 1, 1},
		{"example_3", 5, 1},
		{"example_4", 10, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CashRegister(tc.amount)
			if got != tc.expected {
				t.Fatalf("CashRegister(%d) = %d; want %d", tc.amount, got, tc.expected)
			}
		})
	}
}
