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

func hiddenRef(amount int) int {
	count := 0
	for _, coin := range []int{25, 10, 5, 1} {
		count += amount / coin
		amount %= coin
	}
	return count
}

// Layer 1: Corner Cases
func TestHidden_CashRegister_CornerCases(t *testing.T) {
	tests := []struct {
		name   string
		amount int
	}{
		{"corner_1", 0},
		{"corner_2", 1},
		{"corner_3", 5},
		{"corner_4", 10},
		{"corner_5", 24},
		{"corner_6", 25},
		{"corner_7", 26},
		{"corner_8", 68},
		{"corner_9", 99},
		{"corner_10", 100},
		{"corner_11", 1000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want := hiddenRef(tc.amount)
			got := hiddenCallWithTimeout(t, func() int { return CashRegister(tc.amount) })
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("corner case mismatch: got %v want %v", got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_CashRegister_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1007))

	for trial := 0; trial < 500; trial++ {
		amount := rng.Intn(10000)
		want := hiddenRef(amount)
		got := hiddenCallWithTimeout(t, func() int { return CashRegister(amount) })
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("trial %d: got %v want %v", trial, got, want)
		}
	}
}
