package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// // hidden reference - do not expose
func hiddenRefCountOccurrences(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		if v == target {
			count++
		}
	}
	return count
}

func hiddenCallLevel9WithTimeout(nums []int, target int, timeout time.Duration) (int, error) {
	type result struct {
		count int
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// panic caught
			}
		}()
		cnt := CountOccurrences(nums, target)
		ch <- result{count: cnt}
	}()

	select {
	case res := <-ch:
		return res.count, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel9_CornerCases(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
	}{
		{"nil_slice", nil, 5},
		{"target_zero_in_zeros", []int{0, 0, 0, 0}, 0},
		{"target_zero_absent", []int{1, 2, 3}, 0},
		{"alternating_targets", []int{1, 2, 1, 2, 1, 2}, 1},
		{"large_spread", []int{-1000, 500, -1000, 200}, -1000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			refVal := hiddenRefCountOccurrences(tc.nums, tc.target)
			gotVal, err := hiddenCallLevel9WithTimeout(tc.nums, tc.target, 1*time.Second)

			if err != nil {
				t.Fatalf("Hidden test %q timed out: %v", tc.name, err)
			}
			if gotVal != refVal {
				t.Fatalf("Hidden test %q mismatch: nums=%v target=%d got=%d want=%d",
					tc.name, tc.nums, tc.target, gotVal, refVal)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel9_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(9009))

	for trial := 0; trial < 500; trial++ {
		length := rng.Intn(60)
		nums := make([]int, length)
		target := rng.Intn(20) - 10

		for i := 0; i < length; i++ {
			nums[i] = rng.Intn(20) - 10
		}

		refVal := hiddenRefCountOccurrences(nums, target)
		gotVal, err := hiddenCallLevel9WithTimeout(nums, target, 500*time.Millisecond)

		if err != nil {
			t.Fatalf("Property trial %d timed out: %v", trial, err)
		}
		if gotVal != refVal {
			t.Fatalf("Property trial %d mismatch! nums=%v target=%d got=%d want=%d",
				trial, nums, target, gotVal, refVal)
		}
	}
}

// Layer 3: Performance Audit (100,000 numbers in < 1 second)
func TestHiddenLevel9_PerformanceAudit(t *testing.T) {
	n := 100000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i % 10
	}

	start := time.Now()
	res, err := hiddenCallLevel9WithTimeout(nums, 5, 2*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Performance audit failed: %v", err)
	}
	if res != n/10 {
		t.Fatalf("Performance audit count mismatch: got %d; want %d (elapsed %v)", res, n/10, elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenCountOccurrences(f *testing.F) {
	f.Add([]byte{1, 4, 2, 4, 5, 4, 3}, byte(4))
	f.Add([]byte{10, 20, 30}, byte(99))
	f.Add([]byte{}, byte(0))

	f.Fuzz(func(t *testing.T, data []byte, targetByte byte) {
		nums := make([]int, len(data))
		for i, b := range data {
			nums[i] = int(b)
		}
		target := int(targetByte)

		refVal := hiddenRefCountOccurrences(nums, target)
		gotVal := CountOccurrences(nums, target)

		if gotVal != refVal {
			t.Fatalf("Fuzz mismatch! nums=%v target=%d -> got %d want %d", nums, target, gotVal, refVal)
		}
	})
}
