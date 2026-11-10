package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
	"unicode"
)

// // hidden reference - do not expose
func hiddenRefWordFlip(s string) string {
	if s == "" {
		return "Invalid Output\n"
	}
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ") + "\n"
}

func hiddenCallWordFlipWithTimeout(s string, timeout time.Duration) (string, error) {
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
		ch <- result{out: WordFlip(s)}
	}()

	select {
	case res := <-ch:
		return res.out, res.err
	case <-time.After(timeout):
		return "", fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel1_WordFlip_CornerCases(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"single_word", "hello", "hello\n"},
		{"single_word_padded", "   hello   ", "hello\n"},
		{"two_words", "a b", "b a\n"},
		{"many_inner_spaces", "a      b      c", "c b a\n"},
		{"leading_spaces", "   one two", "two one\n"},
		{"trailing_spaces", "one two   ", "two one\n"},
		{"single_space", " ", "\n"},
		{"empty", "", "Invalid Output\n"},
		{"punctuation", "Hello, World!", "World! Hello,\n"},
		{"digits", "1 2 3 4 5", "5 4 3 2 1\n"},
		{"palindromic_words", "abc abc abc", "abc abc abc\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if ref := hiddenRefWordFlip(tc.input); ref != tc.want {
				t.Fatalf("internal: reference disagrees with table for %q: %q vs %q", tc.input, ref, tc.want)
			}
			got, err := hiddenCallWordFlipWithTimeout(tc.input, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if got != tc.want {
				t.Fatalf("Hidden test %q mismatch: input=%q got=%q want=%q", tc.name, tc.input, got, tc.want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel1_WordFlip_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1001))
	alphabet := "abcdefghijklmnopqrstuvwxyz0123456789!?,."

	for trial := 0; trial < 500; trial++ {
		var sb strings.Builder
		for i := rng.Intn(4); i > 0; i-- { // leading spaces
			sb.WriteByte(' ')
		}
		nWords := rng.Intn(8)
		for w := 0; w < nWords; w++ {
			l := rng.Intn(8) + 1
			for i := 0; i < l; i++ {
				sb.WriteByte(alphabet[rng.Intn(len(alphabet))])
			}
			for i := rng.Intn(4) + 1; i > 0; i-- {
				sb.WriteByte(' ')
			}
		}
		in := sb.String()

		want := hiddenRefWordFlip(in)
		got, err := hiddenCallWordFlipWithTimeout(in, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v (input=%q)", trial, err, in)
		}
		if got != want {
			t.Fatalf("Property trial %d mismatch! input=%q got=%q want=%q", trial, in, got, want)
		}
	}
}

// Layer 3: Performance Audit (100,000 iterations in < 1s)
func TestHiddenLevel1_WordFlip_PerformanceAudit(t *testing.T) {
	inputs := []string{"First second last", " hello  all  of  you! ", "a b c d e f g h", "     "}
	start := time.Now()
	for i := 0; i < 100000; i++ {
		WordFlip(inputs[i%len(inputs)])
	}
	elapsed := time.Since(start)

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenWordFlip(f *testing.F) {
	f.Add("First second last")
	f.Add("")
	f.Add("     ")
	f.Add(" hello  all  of  you! ")

	f.Fuzz(func(t *testing.T, s string) {
		// only plain spaces are specified as separators
		for _, r := range s {
			if unicode.IsSpace(r) && r != ' ' {
				t.Skip()
			}
		}
		want := hiddenRefWordFlip(s)
		got := WordFlip(s)
		if got != want {
			t.Fatalf("Fuzz mismatch! input=%q got=%q want=%q", s, got, want)
		}
	})
}
