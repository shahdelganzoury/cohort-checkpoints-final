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

func hiddenRef(d, m, y int) int {
	if y < 1 || m < 1 || m > 12 || d < 1 {
		return -1
	}
	days := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
		days[1] = 29
	}
	if d > days[m-1] {
		return -1
	}
	for i := 0; i < m-1; i++ {
		d += days[i]
	}
	return d
}

// Layer 1: Corner Cases
func TestHidden_DayOfYear_CornerCases(t *testing.T) {
	tests := []struct {
		name  string
		day   int
		month int
		year  int
	}{
		{"corner_1", 1, 3, 2024},
		{"corner_2", 1, 3, 2025},
		{"corner_3", 29, 2, 2025},
		{"corner_4", 29, 2, 2024},
		{"corner_5", 31, 12, 2000},
		{"corner_6", 1, 1, 1},
		{"corner_7", 31, 4, 2026},
		{"corner_8", 29, 2, 1900},
		{"corner_9", 29, 2, 2000},
		{"corner_10", 1, 13, 2026},
		{"corner_11", 1, 1, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want := hiddenRef(tc.day, tc.month, tc.year)
			got := hiddenCallWithTimeout(t, func() int { return DayOfYear(tc.day, tc.month, tc.year) })
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("corner case mismatch: got %v want %v", got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_DayOfYear_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1002))

	for trial := 0; trial < 500; trial++ {
		day, month, year := rng.Intn(36)-2, rng.Intn(16)-2, rng.Intn(2500)-5
		want := hiddenRef(day, month, year)
		got := hiddenCallWithTimeout(t, func() int { return DayOfYear(day, month, year) })
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("trial %d: got %v want %v", trial, got, want)
		}
	}
}
