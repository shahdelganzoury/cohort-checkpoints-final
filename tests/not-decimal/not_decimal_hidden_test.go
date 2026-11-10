package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// // hidden reference - do not expose
func hiddenRefNotDecimal(s string) string {
	if s == "" {
		return "\n"
	}
	i := 0
	neg := false
	if s[0] == '-' {
		neg = true
		i = 1
	}
	if i >= len(s) {
		return s + "\n"
	}
	dot := -1
	digits := 0
	for j := i; j < len(s); j++ {
		c := s[j]
		switch {
		case c == '.':
			if dot != -1 || digits == 0 {
				return s + "\n"
			}
			dot = j
		case c >= '0' && c <= '9':
			digits++
		default:
			return s + "\n"
		}
	}
	if dot == -1 {
		return s + "\n"
	}
	frac := s[dot+1:]
	if strings.Trim(frac, "0") == "" {
		return s[:dot] + "\n"
	}
	all := s[i:dot] + frac
	all = strings.TrimLeft(all, "0")
	if all == "" {
		all = "0"
	}
	if neg {
		all = "-" + all
	}
	return all + "\n"
}

func hiddenCallNotDecimalWithTimeout(s string, timeout time.Duration) (string, error) {
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
		ch <- result{out: NotDecimal(s)}
	}()

	select {
	case res := <-ch:
		return res.out, res.err
	case <-time.After(timeout):
		return "", fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel1_NotDecimal_CornerCases(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"only_zero_after_dot", "10.0", "10\n"},
		{"many_zeros_after_dot", "5.00", "5\n"},
		{"zero_dot_zero", "0.0", "0\n"},
		{"leading_zeros_in_fraction", "0.05", "5\n"},
		{"negative_small", "-0.5", "-5\n"},
		{"trailing_zero_kept", "1.50", "150\n"},
		{"zero_integer", "0", "0\n"},
		{"big_integer", "100", "100\n"},
		{"letters_only", "abc", "abc\n"},
		{"two_dots", "1.2.3", "1.2.3\n"},
		{"letter_in_number", "12a.5", "12a.5\n"},
		{"only_minus", "-", "-\n"},
		{"negative_integer", "-42", "-42\n"},
		{"negative_with_zero_fraction", "-5.0", "-5\n"},
		{"long_fraction", "3.14159265", "314159265\n"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if ref := hiddenRefNotDecimal(tc.input); ref != tc.want {
				t.Fatalf("internal: reference disagrees with table for %q: %q vs %q", tc.input, ref, tc.want)
			}
			got, err := hiddenCallNotDecimalWithTimeout(tc.input, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if got != tc.want {
				t.Fatalf("Hidden test %q mismatch: input=%q got=%q want=%q", tc.name, tc.input, got, tc.want)
			}
		})
	}
}

func hiddenRandomNotDecimalInput(rng *rand.Rand) string {
	// integer part
	n := rng.Intn(6) + 1
	var ib strings.Builder
	for i := 0; i < n; i++ {
		if i == 0 && n > 1 {
			ib.WriteByte(byte('1' + rng.Intn(9)))
		} else {
			ib.WriteByte(byte('0' + rng.Intn(10)))
		}
	}
	intPart := ib.String()

	kind := rng.Intn(100)
	if kind < 15 { // integer only
		if rng.Intn(2) == 0 && intPart != "0" {
			return "-" + intPart
		}
		return intPart
	}

	m := rng.Intn(8) + 1
	var fb strings.Builder
	allZero := kind < 25
	for i := 0; i < m; i++ {
		if allZero {
			fb.WriteByte('0')
		} else {
			fb.WriteByte(byte('0' + rng.Intn(10)))
		}
	}
	frac := fb.String()
	s := intPart + "." + frac

	neg := rng.Intn(2) == 0
	if intPart == "0" && strings.Trim(frac, "0") == "" {
		neg = false
	}
	if neg {
		s = "-" + s
	}

	if kind >= 90 { // garbage: inject a letter
		pos := rng.Intn(len(s))
		s = s[:pos] + "x" + s[pos+1:]
	}
	return s
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel1_NotDecimal_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1001))

	for trial := 0; trial < 500; trial++ {
		in := hiddenRandomNotDecimalInput(rng)
		want := hiddenRefNotDecimal(in)
		got, err := hiddenCallNotDecimalWithTimeout(in, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v (input=%q)", trial, err, in)
		}
		if got != want {
			t.Fatalf("Property trial %d mismatch! input=%q got=%q want=%q", trial, in, got, want)
		}
	}
}

// Layer 3: Performance Audit (100,000 iterations in < 1s)
func TestHiddenLevel1_NotDecimal_PerformanceAudit(t *testing.T) {
	inputs := []string{"174.2", "0.1255", "1.20525856", "-19.525856", "1952", "abc", "10.0"}
	start := time.Now()
	for i := 0; i < 100000; i++ {
		NotDecimal(inputs[i%len(inputs)])
	}
	elapsed := time.Since(start)

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing (structured numeric inputs)
func FuzzHiddenNotDecimal(f *testing.F) {
	f.Add(uint32(0), uint32(1), uint8(0), false)
	f.Add(uint32(174), uint32(2), uint8(0), false)
	f.Add(uint32(1), uint32(20525856), uint8(7), false)
	f.Add(uint32(19), uint32(525856), uint8(5), true)

	f.Fuzz(func(t *testing.T, ip uint32, fp uint32, fd uint8, neg bool) {
		digits := int(fd%8) + 1
		pow := uint32(1)
		for i := 0; i < digits; i++ {
			pow *= 10
		}
		frac := fmt.Sprintf("%0*d", digits, fp%pow)
		if ip == 0 && strings.Trim(frac, "0") == "" {
			neg = false
		}
		s := fmt.Sprintf("%d.%s", ip, frac)
		if neg {
			s = "-" + s
		}

		want := hiddenRefNotDecimal(s)
		got := NotDecimal(s)
		if got != want {
			t.Fatalf("Fuzz mismatch! input=%q got=%q want=%q", s, got, want)
		}
	})
}
