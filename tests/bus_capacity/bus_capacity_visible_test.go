package main

import "testing"

func TestLevel_BusCapacity_Visible(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		stops    [][2]int
		expected int
	}{
		{"example_1", 20, [][2]int{{10, 0}, {8, 3}, {2, 5}, {4, 1}}, 15},
		{"example_2", 10, [][2]int{{10, 0}}, 10},
		{"example_3", 10, [][2]int{{11, 0}}, -1},
		{"example_4", 10, [][2]int{{0, 1}}, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := BusCapacity(tc.capacity, tc.stops)
			if got != tc.expected {
				t.Fatalf("%s failed\ncall: BusCapacity(%#v, %#v)\ngot: %#v\nrequired: %#v", tc.name, tc.capacity, tc.stops, got, tc.expected)
			}
		})
	}
}
