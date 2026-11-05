package main

import "testing"

func TestLevel5_SumMultiples_Visible(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		factor   int
		expected int
	}{
		{"example_1_multiples_of_3", 10, 3, 18},
		{"example_2_multiples_of_2", 5, 2, 6},
		{"example_3_zero_n", 0, 5, 0},
		{"negative_n", -10, 3, 0},
		{"zero_factor", 10, 0, 0},
		{"factor_larger_than_n", 5, 10, 0},
		{"exact_multiple", 6, 6, 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SumMultiples(tc.n, tc.factor)
			if got != tc.expected {
				t.Fatalf("SumMultiples(%d, %d) = %d; want %d", tc.n, tc.factor, got, tc.expected)
			}
		})
	}
}
