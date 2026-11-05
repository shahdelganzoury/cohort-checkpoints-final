package main

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func TestLevel1_CelsiusToFahrenheit_Visible(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{"example_1_freezing_point", 0.0, 32.0},
		{"example_2_boiling_point", 100.0, 212.0},
		{"example_3_crossover_point", -40.0, -40.0},
		{"body_temperature", 37.0, 98.6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CelsiusToFahrenheit(tc.celsius)
			if !floatEquals(got, tc.expected) {
				t.Fatalf("CelsiusToFahrenheit(%v) = %v; want %v", tc.celsius, got, tc.expected)
			}
		})
	}
}
