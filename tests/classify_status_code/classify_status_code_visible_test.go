package main

import "testing"

func TestLevel3_ClassifyStatusCode_Visible(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		expected string
	}{
		{"example_1_ok_200", 200, "SUCCESS"},
		{"example_2_created_201", 201, "SUCCESS"},
		{"example_3_not_found_404", 404, "CLIENT_ERROR"},
		{"example_4_server_error_500", 500, "SERVER_ERROR"},
		{"example_5_unknown_redirect_301", 301, "UNKNOWN"},
		{"unknown_negative", -1, "UNKNOWN"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ClassifyStatusCode(tc.code)
			if got != tc.expected {
				t.Fatalf("ClassifyStatusCode(%d) = %q; want %q", tc.code, got, tc.expected)
			}
		})
	}
}
