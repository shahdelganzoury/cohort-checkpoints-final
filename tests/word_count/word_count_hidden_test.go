package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var hiddenLevel6Vocab = []string{"go", "Go", "GO", "a", "b", "hello", "world", ""}

// // hidden reference - do not expose
func hiddenRefWordCount(words []string) map[string]int {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	return counts
}

func hiddenCallLevel6WithTimeout(words []string, timeout time.Duration) (map[string]int, error) {
	type result struct {
		counts map[string]int
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// panic caught
			}
		}()
		res := WordCount(words)
		ch <- result{counts: res}
	}()

	select {
	case res := <-ch:
		return res.counts, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel6_CornerCases(t *testing.T) {
	tests := []struct {
		name  string
		words []string
	}{
		{"nil_slice", nil},
		{"empty_slice", []string{}},
		{"single_word", []string{"solo"}},
		{"case_sensitive", []string{"Go", "go", "GO", "go"}},
		{"empty_string_word", []string{"", "", "a"}},
		{"all_same_word", []string{"x", "x", "x", "x", "x"}},
		{"all_unique", []string{"a", "b", "c", "d", "e"}},
		{"words_with_spaces_and_punctuation", []string{"hi there", "hi there", "hi,", "hi"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			refVal := hiddenRefWordCount(tc.words)
			gotVal, err := hiddenCallLevel6WithTimeout(tc.words, 1*time.Second)

			if err != nil {
				t.Fatalf("Hidden test %q timed out: %v", tc.name, err)
			}
			if gotVal == nil {
				t.Fatalf("Hidden test %q returned a nil map; want a non-nil map", tc.name)
			}
			if !reflect.DeepEqual(gotVal, refVal) {
				t.Fatalf("Hidden test %q mismatch: words=%v got=%v want=%v", tc.name, tc.words, gotVal, refVal)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel6_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(6006))

	for trial := 0; trial < 500; trial++ {
		length := rng.Intn(60)
		words := make([]string, length)
		for i := 0; i < length; i++ {
			words[i] = hiddenLevel6Vocab[rng.Intn(len(hiddenLevel6Vocab))]
		}

		refVal := hiddenRefWordCount(words)
		gotVal, err := hiddenCallLevel6WithTimeout(words, 500*time.Millisecond)

		if err != nil {
			t.Fatalf("Property trial %d timed out: %v", trial, err)
		}
		if gotVal == nil {
			t.Fatalf("Property trial %d returned a nil map; want a non-nil map", trial)
		}
		if !reflect.DeepEqual(gotVal, refVal) {
			t.Fatalf("Property trial %d mismatch! words=%v got=%v want=%v", trial, words, gotVal, refVal)
		}
	}
}

// Layer 3: Performance Audit (100,000 words in < 1 second)
func TestHiddenLevel6_PerformanceAudit(t *testing.T) {
	n := 100000
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = fmt.Sprintf("w%d", i%10)
	}

	start := time.Now()
	res, err := hiddenCallLevel6WithTimeout(words, 2*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Performance audit failed: %v", err)
	}
	if len(res) != 10 {
		t.Fatalf("Performance audit key count mismatch: got %d; want 10 (elapsed %v)", len(res), elapsed)
	}
	for k, v := range res {
		if v != n/10 {
			t.Fatalf("Performance audit count mismatch for %q: got %d; want %d", k, v, n/10)
		}
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit took too long: %v", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenWordCount(f *testing.F) {
	f.Add([]byte{1, 2, 1, 3})
	f.Add([]byte{7, 7, 7})
	f.Add([]byte{})

	f.Fuzz(func(t *testing.T, data []byte) {
		words := make([]string, len(data))
		for i, b := range data {
			words[i] = hiddenLevel6Vocab[int(b)%len(hiddenLevel6Vocab)]
		}

		refVal := hiddenRefWordCount(words)
		gotVal := WordCount(words)

		if !reflect.DeepEqual(gotVal, refVal) {
			t.Fatalf("Fuzz mismatch! words=%v -> got %v want %v", words, gotVal, refVal)
		}
	})
}
