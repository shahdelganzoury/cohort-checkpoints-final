//go:build hidden
// +build hidden

package main

import (
	"math/rand"
	"reflect"
	"strings"
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

func hiddenRef(s string) int {
	score := 0
	if len(s) >= 8 {
		score++
	}
	upper, digit, special := false, false, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			upper = true
		}
		if c >= '0' && c <= '9' {
			digit = true
		}
		if strings.ContainsRune("!@#$%", rune(c)) {
			special = true
		}
	}
	if upper {
		score++
	}
	if digit {
		score++
	}
	if special {
		score++
	}
	return score
}

func hiddenRandomString(rng *rand.Rand, chars string, n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte(chars[rng.Intn(len(chars))])
	}
	return b.String()
}

// Layer 1: Corner Cases
func TestHidden_PasswordStrength_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		{"corner_1", ""},
		{"corner_2", "abc"},
		{"corner_3", "Go2026!"},
		{"corner_4", "GoLang2026!"},
		{"corner_5", "ABCDEFGH"},
		{"corner_6", "12345678"},
		{"corner_7", "abcdefgh"},
		{"corner_8", "!@#$%"},
		{"corner_9", "a?bC12345"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want := hiddenRef(tc.s)
			got := hiddenCallWithTimeout(t, func() int { return PasswordStrength(tc.s) })
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("corner case mismatch: got %v want %v", got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_PasswordStrength_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1003))

	for trial := 0; trial < 500; trial++ {
		password := hiddenRandomString(rng, "abcxyzABCXYZ019!@#$%?_-", rng.Intn(25))
		want := hiddenRef(password)
		got := hiddenCallWithTimeout(t, func() int { return PasswordStrength(password) })
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("trial %d: got %v want %v", trial, got, want)
		}
	}
}
