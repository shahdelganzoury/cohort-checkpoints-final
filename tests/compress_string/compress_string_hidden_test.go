//go:build hidden
// +build hidden

package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
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

func hiddenRef(s string) string {
	if s == "" {
		return ""
	}
	var b strings.Builder
	start := 0
	for i := 1; i <= len(s); i++ {
		if i == len(s) || s[i] != s[start] {
			b.WriteByte(s[start])
			b.WriteString(strconv.Itoa(i - start))
			start = i
		}
	}
	return b.String()
}

func hiddenRandomString(rng *rand.Rand, chars string, n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte(chars[rng.Intn(len(chars))])
	}
	return b.String()
}

// Layer 1: Corner Cases
func TestHidden_CompressString_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		{"corner_1", ""},
		{"corner_2", "a"},
		{"corner_3", "aaabbcccc"},
		{"corner_4", "aabbaa"},
		{"corner_5", "abcd"},
		{"corner_6", "aaaaaaaaaaaa"},
		{"corner_7", "ababa"},
		{"corner_8", "zzzzzzzzzzzzzzzzz"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			call := fmt.Sprintf("%s: CompressString(%#v)", tc.name, tc.s)
			want := hiddenRef(tc.s)
			got := hiddenCallWithTimeout(t, call, func() string { return CompressString(tc.s) })
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_CompressString_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1005))

	for trial := 0; trial < 500; trial++ {
		s := hiddenRandomString(rng, "aabbccddef", rng.Intn(101))
		call := fmt.Sprintf("trial %d: CompressString(%#v)", trial, s)
		want := hiddenRef(s)
		got := hiddenCallWithTimeout(t, call, func() string { return CompressString(s) })
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
		}
	}
}
