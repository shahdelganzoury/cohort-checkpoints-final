package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// // hidden reference - do not expose
func hiddenRefFromTo(from, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}
	var parts []string
	if from <= to {
		for i := from; i <= to; i++ {
			parts = append(parts, fmt.Sprintf("%02d", i))
		}
	} else {
		for i := from; i >= to; i-- {
			parts = append(parts, fmt.Sprintf("%02d", i))
		}
	}
	return strings.Join(parts, ", ") + "\n"
}

func hiddenCallFromToWithTimeout(from, to int, timeout time.Duration) (string, error) {
	type result struct {
		out string
		err error
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{err: fmt.Errorf("panic: %v", r)}
			}
		}()
		ch <- result{out: FromTo(from, to)}
	}()

	select {
	case res := <-ch:
		return res.out, res.err
	case <-time.After(timeout):
		return "", fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel1_FromTo_CornerCases(t *testing.T) {
	tests := []struct {
		name     string
		from, to int
	}{
		{"zero_zero", 0, 0},
		{"max_max", 99, 99},
		{"full_ascending", 0, 99},
		{"full_descending", 99, 0},
		{"negative_from", -1, 5},
		{"negative_to", 5, -1},
		{"both_negative", -1, -1},
		{"to_too_big", 5, 100},
		{"from_too_big", 100, 5},
		{"both_too_big", 100, 100},
		{"single_digit_same", 5, 5},
		{"cross_ten_ascending", 9, 10},
		{"cross_ten_descending", 10, 9},
		{"huge_values", 1000000, -1000000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want := hiddenRefFromTo(tc.from, tc.to)
			got, err := hiddenCallFromToWithTimeout(tc.from, tc.to, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if got != want {
				t.Fatalf("Hidden test %q mismatch: from=%d to=%d got=%q want=%q", tc.name, tc.from, tc.to, got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel1_FromTo_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1001))

	for trial := 0; trial < 500; trial++ {
		from := rng.Intn(141) - 20 // -20..120
		to := rng.Intn(141) - 20

		want := hiddenRefFromTo(from, to)
		got, err := hiddenCallFromToWithTimeout(from, to, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if got != want {
			t.Fatalf("Property trial %d mismatch! from=%d to=%d got=%q want=%q", trial, from, to, got, want)
		}
	}
}

// Layer 3: Performance Audit (10,000 iterations in < 1s)
func TestHiddenLevel1_FromTo_PerformanceAudit(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		FromTo(i%100, (i*7)%100)
	}
	elapsed := time.Since(start)

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenFromTo(f *testing.F) {
	f.Add(1, 10)
	f.Add(10, 1)
	f.Add(10, 10)
	f.Add(100, 10)

	f.Fuzz(func(t *testing.T, from, to int) {
		want := hiddenRefFromTo(from, to)
		got := FromTo(from, to)
		if got != want {
			t.Fatalf("Fuzz mismatch! from=%d to=%d got=%q want=%q", from, to, got, want)
		}
	})
}
