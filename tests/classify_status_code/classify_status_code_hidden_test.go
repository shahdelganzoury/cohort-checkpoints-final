package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// // hidden reference - do not expose
func hiddenRefClassifyStatusCode(code int) string {
	if code >= 200 && code <= 299 {
		return "SUCCESS"
	} else if code >= 400 && code <= 499 {
		return "CLIENT_ERROR"
	} else if code >= 500 && code <= 599 {
		return "SERVER_ERROR"
	}
	return "UNKNOWN"
}

func hiddenCallLevel3WithTimeout(code int, timeout time.Duration) (string, error) {
	type result struct {
		category string
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// panic caught
			}
		}()
		res := ClassifyStatusCode(code)
		ch <- result{category: res}
	}()

	select {
	case res := <-ch:
		return res.category, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel3_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		code int
	}{
		{"exact_200_lower_bound", 200},
		{"exact_299_upper_bound", 299},
		{"exact_300_boundary", 300},
		{"exact_399_boundary", 399},
		{"exact_400_lower_bound", 400},
		{"exact_499_upper_bound", 499},
		{"exact_500_lower_bound", 500},
		{"exact_599_upper_bound", 599},
		{"exact_600_boundary", 600},
		{"zero", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			refVal := hiddenRefClassifyStatusCode(tc.code)
			gotVal, err := hiddenCallLevel3WithTimeout(tc.code, 1*time.Second)

			if err != nil {
				t.Fatalf("Hidden test %q timed out: %v", tc.name, err)
			}
			if gotVal != refVal {
				t.Fatalf("Hidden test %q mismatch: code=%d got=%q want=%q", tc.name, tc.code, gotVal, refVal)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel3_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(3003))

	for trial := 0; trial < 500; trial++ {
		code := rng.Intn(1000) - 200

		refVal := hiddenRefClassifyStatusCode(code)
		gotVal, err := hiddenCallLevel3WithTimeout(code, 500*time.Millisecond)

		if err != nil {
			t.Fatalf("Property trial %d timed out: %v", trial, err)
		}
		if gotVal != refVal {
			t.Fatalf("Property trial %d mismatch! code=%d got=%q want=%q", trial, code, gotVal, refVal)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel3_PerformanceAudit(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		ClassifyStatusCode(i % 700)
	}
	elapsed := time.Since(start)

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenClassifyStatusCode(f *testing.F) {
	f.Add(200)
	f.Add(404)
	f.Add(500)
	f.Add(301)

	f.Fuzz(func(t *testing.T, code int) {
		refVal := hiddenRefClassifyStatusCode(code)
		gotVal := ClassifyStatusCode(code)

		if gotVal != refVal {
			t.Fatalf("Fuzz mismatch! code=%d -> got %q want %q", code, gotVal, refVal)
		}
	})
}
