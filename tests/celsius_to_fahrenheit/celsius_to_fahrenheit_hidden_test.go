package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// // hidden reference - do not expose
func hiddenRefCelsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func hiddenCallLevel1WithTimeout(c float64, timeout time.Duration) (float64, error) {
	type result struct {
		f float64
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// panic caught
			}
		}()
		res := CelsiusToFahrenheit(c)
		ch <- result{f: res}
	}()

	select {
	case res := <-ch:
		return res.f, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel1_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		c    float64
	}{
		{"absolute_zero", -273.15},
		{"negative_decimal", -17.777778},
		{"large_positive", 1000.0},
		{"zero", 0.0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			refVal := hiddenRefCelsiusToFahrenheit(tc.c)
			gotVal, err := hiddenCallLevel1WithTimeout(tc.c, 1*time.Second)

			if err != nil {
				t.Fatalf("Hidden test %q timed out: %v", tc.name, err)
			}
			if !floatEquals(gotVal, refVal) {
				t.Fatalf("Hidden test %q mismatch: c=%v got=%v want=%v", tc.name, tc.c, gotVal, refVal)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel1_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1001))

	for trial := 0; trial < 500; trial++ {
		c := (rng.Float64() * 1000.0) - 500.0

		refVal := hiddenRefCelsiusToFahrenheit(c)
		gotVal, err := hiddenCallLevel1WithTimeout(c, 500*time.Millisecond)

		if err != nil {
			t.Fatalf("Property trial %d timed out: %v", trial, err)
		}
		if !floatEquals(gotVal, refVal) {
			t.Fatalf("Property trial %d mismatch! c=%v got=%v want=%v", trial, c, gotVal, refVal)
		}
	}
}

// Layer 3: Performance Audit (100,000 iterations in < 500ms)
func TestHiddenLevel1_PerformanceAudit(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		CelsiusToFahrenheit(float64(i % 100))
	}
	elapsed := time.Since(start)

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenCelsiusToFahrenheit(f *testing.F) {
	f.Add(0.0)
	f.Add(100.0)
	f.Add(-40.0)
	f.Add(37.0)

	f.Fuzz(func(t *testing.T, c float64) {
		refVal := hiddenRefCelsiusToFahrenheit(c)
		gotVal := CelsiusToFahrenheit(c)

		if !floatEquals(gotVal, refVal) {
			t.Fatalf("Fuzz mismatch! c=%v -> got %v want %v", c, gotVal, refVal)
		}
	})
}
