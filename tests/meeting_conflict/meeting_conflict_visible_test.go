package main

import "testing"

func TestLevel_MeetingConflict_Visible(t *testing.T) {
	tests := []struct {
		name     string
		meetings [][2]int
		expected bool
	}{
		{"example_1", nil, false},
		{"example_2", [][2]int{{540, 600}, {570, 630}, {660, 720}}, true},
		{"example_3", [][2]int{{540, 600}, {600, 660}}, false},
		{"example_4", [][2]int{{0, 1440}, {1, 2}}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MeetingConflict(tc.meetings)
			if got != tc.expected {
				t.Fatalf("%s failed\ncall: MeetingConflict(%#v)\ngot: %#v\nrequired: %#v", tc.name, tc.meetings, got, tc.expected)
			}
		})
	}
}
