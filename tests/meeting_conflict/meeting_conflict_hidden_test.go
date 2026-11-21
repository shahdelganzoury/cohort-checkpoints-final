//go:build hidden
// +build hidden

package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func hiddenCallWithTimeout[T any](t *testing.T, label string, fn func() T) T {
	t.Helper()
	ch := make(chan T, 1)
	panicCh := make(chan any, 1)
	go func() {
		defer func() {
			if v := recover(); v != nil {
				panicCh <- v
			}
		}()
		ch <- fn()
	}()
	select {
	case got := <-ch:
		return got
	case p := <-panicCh:
		t.Fatalf("%s\nstudent function panicked: %v", label, p)
		var zero T
		return zero
	case <-time.After(800 * time.Millisecond):
		t.Fatalf("%s\nstudent function timed out", label)
		var zero T
		return zero
	}
}

func hiddenRef(meetings [][2]int) bool {
	for i := 0; i < len(meetings); i++ {
		for j := i + 1; j < len(meetings); j++ {
			a, b := meetings[i], meetings[j]
			if a[0] < b[1] && b[0] < a[1] {
				return true
			}
		}
	}
	return false
}

func hiddenRandomMeetings(rng *rand.Rand) [][2]int {
	n := rng.Intn(35)
	meetings := make([][2]int, n)
	for i := range meetings {
		s := rng.Intn(1440)
		meetings[i] = [2]int{s, s + 1 + rng.Intn(1440-s)}
	}
	return meetings
}

// Layer 1: Corner Cases
func TestHidden_MeetingConflict_CornerCases(t *testing.T) {
	tests := []struct {
		name     string
		meetings [][2]int
	}{
		{"corner_1", nil},
		{"corner_2", [][2]int{{540, 600}, {570, 630}, {660, 720}}},
		{"corner_3", [][2]int{{540, 600}, {600, 660}}},
		{"corner_4", [][2]int{{0, 1440}, {1, 2}}},
		{"corner_5", [][2]int{{900, 1000}, {100, 200}}},
		{"corner_6", [][2]int{{0, 1}, {1, 2}, {2, 3}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			call := fmt.Sprintf("%s: MeetingConflict(%#v)", tc.name, tc.meetings)
			want := hiddenRef(tc.meetings)
			got := hiddenCallWithTimeout(t, call, func() bool { return MeetingConflict(tc.meetings) })
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_MeetingConflict_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1008))

	for trial := 0; trial < 500; trial++ {
		meetings := hiddenRandomMeetings(rng)
		call := fmt.Sprintf("trial %d: MeetingConflict(%#v)", trial, meetings)
		want := hiddenRef(meetings)
		got := hiddenCallWithTimeout(t, call, func() bool { return MeetingConflict(meetings) })
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
		}
	}
}
