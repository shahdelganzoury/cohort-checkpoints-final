//go:build hidden
// +build hidden

package main

import (
	"fmt"
	"math/rand"
	"reflect"
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

func hiddenRef(nums []int) [2]int {
	n := len(nums) + 2
	seen := make([]bool, n+1)
	for _, v := range nums {
		seen[v] = true
	}
	result := [2]int{}
	j := 0
	for i := 1; i <= n; i++ {
		if !seen[i] {
			result[j] = i
			j++
		}
	}
	return result
}

func hiddenRandomMissing(rng *rand.Rand) []int {
	n := rng.Intn(120) + 2
	a := rng.Intn(n) + 1
	b := rng.Intn(n-1) + 1
	if b >= a {
		b++
	}
	nums := make([]int, 0, n-2)
	for i := 1; i <= n; i++ {
		if i != a && i != b {
			nums = append(nums, i)
		}
	}
	rng.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

// Layer 1: Corner Cases
func TestHidden_TwoMissingNumbers_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{"corner_1", nil},
		{"corner_2", []int{1, 3, 6, 2}},
		{"corner_3", []int{2, 4, 5, 6}},
		{"corner_4", []int{2, 3}},
		{"corner_5", []int{1, 4}},
		{"corner_6", []int{1, 2, 3, 4}},
		{"corner_7", []int{3, 4, 5, 6}},
		{"corner_8", []int{1, 2, 4, 6, 7, 8}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			call := fmt.Sprintf("%s: TwoMissingNumbers(%#v)", tc.name, tc.nums)
			want := hiddenRef(tc.nums)
			got := hiddenCallWithTimeout(t, call, func() [2]int {
				a, b := TwoMissingNumbers(tc.nums)
				return [2]int{a, b}
			})
			if got != want {
				t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_TwoMissingNumbers_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1010))

	for trial := 0; trial < 500; trial++ {
		nums := hiddenRandomMissing(rng)
		call := fmt.Sprintf("trial %d: TwoMissingNumbers(%#v)", trial, nums)
		want := hiddenRef(nums)
		got := hiddenCallWithTimeout(t, call, func() [2]int {
			a, b := TwoMissingNumbers(nums)
			return [2]int{a, b}
		})
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("%s\ngot: %#v\nrequired: %#v", call, got, want)
		}
	}
}
