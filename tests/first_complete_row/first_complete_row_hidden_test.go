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

func hiddenRef(grid [][]int, called []int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	seen := make(map[int]bool, len(called))
	for i, v := range called {
		seen[v] = true
		for _, row := range grid {
			ok := true
			for _, n := range row {
				if !seen[n] {
					ok = false
					break
				}
			}
			if ok {
				return i + 1
			}
		}
	}
	return -1
}

func hiddenRandomGrid(rng *rand.Rand) ([][]int, []int) {
	rows, cols := rng.Intn(4)+1, rng.Intn(5)+1
	pool := rng.Perm(300)
	grid := make([][]int, rows)
	for r := 0; r < rows; r++ {
		grid[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			grid[r][c] = pool[r*cols+c] + 1
		}
	}
	calls := append([]int{}, pool[:rows*cols]...)
	for i := range calls {
		calls[i]++
	}
	for i := 0; i < rng.Intn(5); i++ {
		calls = append(calls, 500+i)
	}
	rng.Shuffle(len(calls), func(i, j int) {
		calls[i], calls[j] = calls[j], calls[i]
	})
	calls = calls[:rng.Intn(len(calls)+1)]
	return grid, calls
}

// Layer 1: Corner Cases
func TestHidden_FirstCompleteRow_CornerCases(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		called []int
	}{
		{"corner_1", [][]int{{4, 1, 7}, {3, 9, 2}, {8, 5, 6}}, []int{9, 4, 1, 3, 7, 2}},
		{"corner_2", [][]int{{1}}, []int{1}},
		{"corner_3", [][]int{{1}}, []int{2}},
		{"corner_4", [][]int{{1, 2}, {3, 4}}, []int{1, 3, 2}},
		{"corner_5", [][]int{{1, 2}, {3, 4}}, nil},
		{"corner_6", [][]int{{1}, {2}}, []int{2}},
		{"corner_7", [][]int{{1, 2, 3}}, []int{99, 2, 1, 3}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want := hiddenRef(tc.grid, tc.called)
			got := hiddenCallWithTimeout(t, func() int { return FirstCompleteRow(tc.grid, tc.called) })
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("corner case mismatch: got %v want %v", got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials)
func TestHidden_FirstCompleteRow_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1009))

	for trial := 0; trial < 500; trial++ {
		grid, called := hiddenRandomGrid(rng)
		want := hiddenRef(grid, called)
		got := hiddenCallWithTimeout(t, func() int { return FirstCompleteRow(grid, called) })
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("trial %d: got %v want %v", trial, got, want)
		}
	}
}
