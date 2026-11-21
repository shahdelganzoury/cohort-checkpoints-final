//go:build hidden
// +build hidden

package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
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

func hiddenRef(a, b string) int {
	count := 0
	for i := range a {
		if a[i] == b[i] {
			count++
		}
	}
	return count
}

func hiddenRandomString(rng *rand.Rand, chars string, n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte(chars[rng.Intn(len(chars))])
	}
	return b.String()
}

func hiddenRandomDigitPair(rng *rand.Rand) (string, string) {
	n := rng.Intn(80)
	return hiddenRandomString(rng, "0123456789", n), hiddenRandomString(rng, "0123456789", n)
}

// Layer 1: Corner Cases
func TestHidden_DigitPositionMatch_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
	}{
		{"corner_1", "58342", "57382"},
		{"corner_2", "", ""},
		{"corner_3", "000", "000"},
		{"corner_4", "123", "456"},
		{"corner_5", "090", "009"},
		{"corner_6", "1", "1"},
		{"corner_7", "0", "9"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			call := fmt.Sprintf("%s: DigitPositionMatch(%#v, %#v)", tc.name, tc.a, tc.b)
			want := hiddenRef(tc.a, tc.b)
			got := hiddenCallWithTimeout(t, call, func() int { return DigitPositionMatch(tc.a, tc.b) })
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_DigitPositionMatch_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1006))

	for trial := 0; trial < 500; trial++ {
		a, b := hiddenRandomDigitPair(rng)
		call := fmt.Sprintf("trial %d: DigitPositionMatch(%#v, %#v)", trial, a, b)
		want := hiddenRef(a, b)
		got := hiddenCallWithTimeout(t, call, func() int { return DigitPositionMatch(a, b) })
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
		}
	}
}
