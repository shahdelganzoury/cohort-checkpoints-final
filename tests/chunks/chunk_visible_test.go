package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureStdout(fn func()) string {
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

func TestLevel1_Chunk_Visible(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		size     int
		expected string
	}{
		{"example_1_empty_slice", []int{}, 10, "[]\n"},
		{"example_2_size_zero", []int{0, 1, 2, 3, 4, 5, 6, 7}, 0, "\n"},
		{"example_3_size_three", []int{0, 1, 2, 3, 4, 5, 6, 7}, 3, "[[0 1 2] [3 4 5] [6 7]]\n"},
		{"example_4_size_five", []int{0, 1, 2, 3, 4, 5, 6, 7}, 5, "[[0 1 2 3 4] [5 6 7]]\n"},
		{"example_5_size_four", []int{0, 1, 2, 3, 4, 5, 6, 7}, 4, "[[0 1 2 3] [4 5 6 7]]\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := captureStdout(func() { Chunk(tc.slice, tc.size) })
			if got != tc.expected {
				t.Fatalf("Chunk(%v, %d) printed %q; want %q", tc.slice, tc.size, got, tc.expected)
			}
		})
	}
}
