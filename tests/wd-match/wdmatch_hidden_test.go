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
func hiddenRefWdmatch(args []string) string {
	if len(args) != 2 {
		return ""
	}
	a, b := args[0], args[1]
	j := 0
	for i := 0; i < len(b) && j < len(a); i++ {
		if b[i] == a[j] {
			j++
		}
	}
	if j == len(a) {
		return a + "\n"
	}
	return ""
}

func hiddenRunMain(args []string) string {
	oldArgs := os.Args
	oldOut := os.Stdout
	os.Args = append([]string{"wdmatch"}, args...)
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
			os.Args = oldArgs
			os.Stdout = oldOut
			w.Close()
		}()
		main()
	}()
	return <-done
}

func hiddenCallMainWithTimeout(args []string, timeout time.Duration) (string, error) {
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
		ch <- result{out: hiddenRunMain(args)}
	}()

	select {
	case res := <-ch:
		return res.out, res.err
	case <-time.After(timeout):
		return "", fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel1_Wdmatch_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"no_args", []string{}},
		{"one_arg", []string{"abc"}},
		{"three_args", []string{"a", "b", "c"}},
		{"identical", []string{"abc", "abc"}},
		{"single_char_match", []string{"a", "xyzabc"}},
		{"single_char_no_match", []string{"a", "xyzbc"}},
		{"first_longer_than_second", []string{"abcdef", "abc"}},
		{"order_matters", []string{"abc", "cba"}},
		{"order_ok_with_gaps", []string{"abc", "a1b2c3"}},
		{"repeated_letters_enough", []string{"aab", "a1a2b"}},
		{"repeated_letters_not_enough", []string{"aab", "ab"}},
		{"spaces_in_both", []string{"a b", "a  b"}},
		{"space_missing", []string{"a b", "ab"}},
		{"case_sensitive", []string{"abc", "ABC"}},
		{"digits", []string{"123", "1x2y3z"}},
		{"last_char_missing", []string{"hello", "hell"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want := hiddenRefWdmatch(tc.args)
			got, err := hiddenCallMainWithTimeout(tc.args, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if got != want {
				t.Fatalf("Hidden test %q mismatch: args=%q got=%q want=%q", tc.name, tc.args, got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel1_Wdmatch_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1001))
	alphabet := "abcde"

	for trial := 0; trial < 500; trial++ {
		la := rng.Intn(6) + 1
		a := make([]byte, la)
		for i := range a {
			a[i] = alphabet[rng.Intn(len(alphabet))]
		}

		var b []byte
		if rng.Intn(2) == 0 {
			// guaranteed subsequence: interleave noise
			for _, c := range a {
				for k := rng.Intn(3); k > 0; k-- {
					b = append(b, alphabet[rng.Intn(len(alphabet))])
				}
				b = append(b, c)
			}
			for k := rng.Intn(3); k > 0; k-- {
				b = append(b, alphabet[rng.Intn(len(alphabet))])
			}
		} else {
			lb := rng.Intn(12)
			b = make([]byte, lb)
			for i := range b {
				b[i] = alphabet[rng.Intn(len(alphabet))]
			}
		}

		args := []string{string(a), string(b)}
		want := hiddenRefWdmatch(args)
		got, err := hiddenCallMainWithTimeout(args, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v (args=%q)", trial, err, args)
		}
		if got != want {
			t.Fatalf("Property trial %d mismatch! args=%q got=%q want=%q", trial, args, got, want)
		}
	}
}

// Layer 3: Performance Audit (20,000 runs in < 1s)
func TestHiddenLevel1_Wdmatch_PerformanceAudit(t *testing.T) {
	devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		t.Skip("cannot open os.DevNull")
	}
	defer devNull.Close()

	oldArgs := os.Args
	oldOut := os.Stdout
	os.Stdout = devNull
	start := time.Now()
	for i := 0; i < 20000; i++ {
		os.Args = []string{"wdmatch", "faya", "fgvvfdxcacpolhyghbreda"}
		main()
	}
	elapsed := time.Since(start)
	os.Stdout = oldOut
	os.Args = oldArgs

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenWdmatch(f *testing.F) {
	f.Add("123", "123")
	f.Add("faya", "fgvvfdxcacpolhyghbreda")
	f.Add("faya", "fgvvfdxcacpolhyghbred")
	f.Add("error", "rrerrrfiiljdfxjyuifrrvcoojh")

	f.Fuzz(func(t *testing.T, a, b string) {
		if a == "" {
			t.Skip()
		}
		for i := 0; i < len(a); i++ {
			if a[i] >= 128 {
				t.Skip()
			}
		}
		for i := 0; i < len(b); i++ {
			if b[i] >= 128 {
				t.Skip()
			}
		}
		args := []string{a, b}
		want := hiddenRefWdmatch(args)
		got := hiddenRunMain(args)
		if got != want {
			t.Fatalf("Fuzz mismatch! args=%q got=%q want=%q", args, got, want)
		}
	})
}
