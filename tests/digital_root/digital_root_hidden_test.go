//go:build hidden
// +build hidden

package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func hiddenCallWithTimeout[T any](t *testing.T, fn func() T) T {
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
		t.Fatalf("student function panicked: %v", p)
		var zero T
		return zero
	case <-time.After(800 * time.Millisecond):
		t.Fatal("student function timed out")
		var zero T
		return zero
	}
}

func hiddenRef(n int) int {
	for n >= 10 {
		total := 0
		for n > 0 {
			total += n % 10
			n /= 10
		}
		n = total
	}
	return n
}

// Layer 1: Corner Cases
func TestHidden_DigitalRoot_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"corner_1", 0},
		{"corner_2", 9},
		{"corner_3", 10},
		{"corner_4", 99},
		{"corner_5", 9875},
		{"corner_6", 493},
		{"corner_7", 1000000000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want := hiddenRef(tc.n)
			got := hiddenCallWithTimeout(t, func() int { return DigitalRoot(tc.n) })
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("corner case mismatch: got %v want %v", got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_DigitalRoot_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1001))

	for trial := 0; trial < 500; trial++ {
		n := rng.Intn(2_000_000_000)
		want := hiddenRef(n)
		got := hiddenCallWithTimeout(t, func() int { return DigitalRoot(n) })
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("trial %d: got %v want %v", trial, got, want)
		}
	}
}
