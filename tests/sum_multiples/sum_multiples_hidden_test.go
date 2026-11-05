package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// // hidden reference - do not expose
func hiddenRefSumMultiples(n int, factor int) int {
	if n <= 0 || factor <= 0 {
		return 0
	}
	sum := 0
	for i := 1; i <= n; i++ {
		if i%factor == 0 {
			sum += i
		}
	}
	return sum
}

func hiddenCallLevel5WithTimeout(n int, factor int, timeout time.Duration) (int, error) {
	type result struct {
		sum int
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// panic caught
			}
		}()
		res := SumMultiples(n, factor)
		ch <- result{sum: res}
	}()

	select {
	case res := <-ch:
		return res.sum, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel5_CornerCases(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		factor int
	}{
		{"factor_one_sum_one_to_n", 5, 1}, // 1+2+3+4+5 = 15
		{"both_negative", -5, -2},
		{"factor_equals_n", 10, 10},
		{"large_n", 1000, 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			refVal := hiddenRefSumMultiples(tc.n, tc.factor)
			gotVal, err := hiddenCallLevel5WithTimeout(tc.n, tc.factor, 1*time.Second)

			if err != nil {
				t.Fatalf("Hidden test %q timed out: %v", tc.name, err)
			}
			if gotVal != refVal {
				t.Fatalf("Hidden test %q mismatch: n=%d factor=%d got=%d want=%d",
					tc.name, tc.n, tc.factor, gotVal, refVal)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel5_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(5005))

	for trial := 0; trial < 500; trial++ {
		n := rng.Intn(200) - 50
		factor := rng.Intn(20) - 5

		refVal := hiddenRefSumMultiples(n, factor)
		gotVal, err := hiddenCallLevel5WithTimeout(n, factor, 500*time.Millisecond)

		if err != nil {
			t.Fatalf("Property trial %d timed out: %v", trial, err)
		}
		if gotVal != refVal {
			t.Fatalf("Property trial %d mismatch! n=%d factor=%d got=%d want=%d",
				trial, n, factor, gotVal, refVal)
		}
	}
}

// Layer 3: Performance Audit (n=100,000 in < 1 second)
func TestHiddenLevel5_PerformanceAudit(t *testing.T) {
	start := time.Now()
	res, err := hiddenCallLevel5WithTimeout(100000, 3, 2*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Performance audit failed: %v", err)
	}
	_ = res
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit took too long: %v", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenSumMultiples(f *testing.F) {
	f.Add(10, 3)
	f.Add(5, 2)
	f.Add(0, 5)

	f.Fuzz(func(t *testing.T, n, factor int) {
		if n > 5000 {
			return
		}
		refVal := hiddenRefSumMultiples(n, factor)
		gotVal := SumMultiples(n, factor)

		if gotVal != refVal {
			t.Fatalf("Fuzz mismatch! n=%d factor=%d -> got %d want %d", n, factor, gotVal, refVal)
		}
	})
}
