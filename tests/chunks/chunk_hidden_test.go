package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"testing"
	"time"
)

// // hidden reference - do not expose
func hiddenRefChunk(slice []int, size int) string {
	if size == 0 {
		return "\n"
	}
	res := [][]int{}
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		res = append(res, slice[i:end])
	}
	return fmt.Sprintln(res)
}

func hiddenCaptureStdout(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	done := make(chan string, 1)
	go func() {
		var b bytes.Buffer
		io.Copy(&b, r)
		done <- b.String()
	}()
	func() {
		defer func() {
			os.Stdout = old
			w.Close()
		}()
		fn()
	}()
	return <-done
}

func hiddenCallChunkWithTimeout(slice []int, size int, timeout time.Duration) (string, error) {
	type result struct {
		out string
		err error
	}
	ch := make(chan result, 1)
	in := append([]int(nil), slice...) // protect the reference input

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{err: fmt.Errorf("panic: %v", r)}
			}
		}()
		out := hiddenCaptureStdout(func() { Chunk(in, size) })
		ch <- result{out: out}
	}()

	select {
	case res := <-ch:
		return res.out, res.err
	case <-time.After(timeout):
		return "", fmt.Errorf("execution timed out after %v", timeout)
	}
}

func hiddenRange(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}

// Layer 1: Corner Cases
func TestHiddenLevel1_Chunk_CornerCases(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		size  int
	}{
		{"empty_size_zero", []int{}, 0},
		{"empty_size_one", []int{}, 1},
		{"size_one", hiddenRange(5), 1},
		{"size_equals_len", hiddenRange(6), 6},
		{"size_bigger_than_len", hiddenRange(4), 100},
		{"exact_multiple", hiddenRange(9), 3},
		{"one_leftover", hiddenRange(7), 3},
		{"single_element", []int{42}, 1},
		{"single_element_big_size", []int{42}, 5},
		{"negative_values", []int{-3, -2, -1, 0, 1, 2}, 4},
		{"large_slice", hiddenRange(1000), 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want := hiddenRefChunk(tc.slice, tc.size)
			got, err := hiddenCallChunkWithTimeout(tc.slice, tc.size, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if got != want {
				t.Fatalf("Hidden test %q mismatch: size=%d\ngot =%q\nwant=%q", tc.name, tc.size, got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel1_Chunk_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1001))

	for trial := 0; trial < 500; trial++ {
		n := rng.Intn(60)
		slice := make([]int, n)
		for i := range slice {
			slice[i] = rng.Intn(2001) - 1000
		}
		size := rng.Intn(25) // includes 0

		want := hiddenRefChunk(slice, size)
		got, err := hiddenCallChunkWithTimeout(slice, size, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if got != want {
			t.Fatalf("Property trial %d mismatch! slice=%v size=%d\ngot =%q\nwant=%q", trial, slice, size, got, want)
		}
	}
}

// Layer 3: Performance Audit (20,000 iterations in < 1s)
func TestHiddenLevel1_Chunk_PerformanceAudit(t *testing.T) {
	devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Skip("cannot open os.DevNull")
	}
	defer devNull.Close()

	slice := hiddenRange(100)
	old := os.Stdout
	os.Stdout = devNull
	start := time.Now()
	for i := 0; i < 20000; i++ {
		Chunk(slice, i%20+1)
	}
	elapsed := time.Since(start)
	os.Stdout = old

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenChunk(f *testing.F) {
	f.Add(8, 3)
	f.Add(8, 0)
	f.Add(0, 10)
	f.Add(8, 5)

	f.Fuzz(func(t *testing.T, n int, size int) {
		if n < 0 {
			n = -n
		}
		if size < 0 {
			size = -size
		}
		if n < 0 || size < 0 { // MinInt overflow
			t.Skip()
		}
		n %= 200
		size %= 50

		slice := hiddenRange(n)
		want := hiddenRefChunk(slice, size)
		got := hiddenCaptureStdout(func() { Chunk(append([]int(nil), slice...), size) })
		if got != want {
			t.Fatalf("Fuzz mismatch! n=%d size=%d\ngot =%q\nwant=%q", n, size, got, want)
		}
	})
}
