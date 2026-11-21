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

func hiddenRef(capacity int, stops [][2]int) int {
	if capacity < 0 {
		return -1
	}
	onboard, best := 0, 0
	for _, stop := range stops {
		in, out := stop[0], stop[1]
		if in < 0 || out < 0 || out > onboard {
			return -1
		}
		onboard -= out
		if in > capacity-onboard {
			return -1
		}
		onboard += in
		if onboard > best {
			best = onboard
		}
	}
	return best
}

func hiddenRandomStops(rng *rand.Rand) (int, [][2]int) {
	capacity := rng.Intn(60)
	onboard := 0
	stops := make([][2]int, rng.Intn(65))
	for i := range stops {
		out := rng.Intn(onboard + 2)
		in := rng.Intn(capacity + 3)
		if rng.Intn(7) == 0 {
			out = -1
		}
		if rng.Intn(7) == 0 {
			in = -1
		}
		stops[i] = [2]int{in, out}
		if in < 0 || out < 0 || out > onboard || in > capacity-(onboard-out) {
			onboard = 0
		} else {
			onboard = onboard - out + in
		}
	}
	return capacity, stops
}

// Layer 1: Corner Cases
func TestHidden_BusCapacity_CornerCases(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		stops    [][2]int
	}{
		{"corner_1", 20, [][2]int{{10, 0}, {8, 3}, {2, 5}, {4, 1}}},
		{"corner_2", 10, [][2]int{{10, 0}}},
		{"corner_3", 10, [][2]int{{11, 0}}},
		{"corner_4", 10, [][2]int{{0, 1}}},
		{"corner_5", 0, nil},
		{"corner_6", -1, nil},
		{"corner_7", 10, [][2]int{{8, 0}, {8, 8}}},
		{"corner_8", 10, [][2]int{{1, 0}, {-1, 0}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			call := fmt.Sprintf("%s: BusCapacity(%#v, %#v)", tc.name, tc.capacity, tc.stops)
			want := hiddenRef(tc.capacity, tc.stops)
			got := hiddenCallWithTimeout(t, call, func() int { return BusCapacity(tc.capacity, tc.stops) })
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_BusCapacity_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1004))

	for trial := 0; trial < 500; trial++ {
		capacity, stops := hiddenRandomStops(rng)
		call := fmt.Sprintf("trial %d: BusCapacity(%#v, %#v)", trial, capacity, stops)
		want := hiddenRef(capacity, stops)
		got := hiddenCallWithTimeout(t, call, func() int { return BusCapacity(capacity, stops) })
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
		}
	}
}
