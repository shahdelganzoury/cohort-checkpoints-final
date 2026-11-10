package main

import "testing"

func TestLevel1_NotDecimal_Visible(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"example_1_small_decimal", "0.1", "1\n"},
		{"example_2_integer_part", "174.2", "1742\n"},
		{"example_3_four_decimals", "0.1255", "1255\n"},
		{"example_4_many_decimals", "1.20525856", "120525856\n"},
		{"example_5_not_a_number", "-0.0f00d00", "-0.0f00d00\n"},
		{"example_6_empty", "", "\n"},
		{"example_7_negative", "-19.525856", "-19525856\n"},
		{"example_8_no_decimal_point", "1952", "1952\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NotDecimal(tc.input)
			if got != tc.expected {
				t.Fatalf("NotDecimal(%q) = %q; want %q", tc.input, got, tc.expected)
			}
		})
	}
}
