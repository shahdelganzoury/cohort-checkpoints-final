package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func hiddenBuild(vals []int) (*ListNode, []*ListNode) {
	var head, tail *ListNode
	nodes := make([]*ListNode, 0, len(vals))
	for _, v := range vals {
		n := &ListNode{Val: v}
		nodes = append(nodes, n)
		if head == nil {
			head = n
		} else {
			tail.Next = n
		}
		tail = n
	}
	return head, nodes
}

func hiddenToSlice(head *ListNode) ([]int, bool) {
	out := []int{}
	steps := 0
	for curr := head; curr != nil; curr = curr.Next {
		out = append(out, curr.Val)
		steps++
		if steps > 1000000 {
			return out, false
		}
	}
	return out, true
}

func hiddenEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func hiddenRepeat(v, n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = v
	}
	return out
}

func hiddenRandomVals(rng *rand.Rand, n, lo, hi int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = lo + rng.Intn(hi-lo+1)
	}
	return out
}

// hidden reference - do not expose
func hiddenRefDecimal(bits []int) int {
	result := 0
	for _, b := range bits {
		result = result*2 + b
	}
	return result
}

func hiddenCallLevel1WithTimeout(head *ListNode, timeout time.Duration) (int, error) {
	type result struct {
		value    int
		panicked interface{}
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{panicked: r}
			}
		}()
		ch <- result{value: GetDecimalValue(head)}
	}()

	select {
	case res := <-ch:
		if res.panicked != nil {
			return 0, fmt.Errorf("panic: %v", res.panicked)
		}
		return res.value, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel1_CornerCases(t *testing.T) {
	alternating := []int{}
	for i := 0; i < 30; i++ {
		alternating = append(alternating, (i+1)%2)
	}

	tests := []struct {
		name string
		bits []int
	}{
		{"single_zero", []int{0}},
		{"single_one", []int{1}},
		{"two_bits_10", []int{1, 0}},
		{"two_bits_01", []int{0, 1}},
		{"leading_zeros", []int{0, 0, 0, 1}},
		{"all_ones_30", hiddenRepeat(1, 30)},
		{"all_zeros_30", hiddenRepeat(0, 30)},
		{"one_then_29_zeros", append([]int{1}, hiddenRepeat(0, 29)...)},
		{"alternating_30", alternating},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, _ := hiddenBuild(tc.bits)
			want := hiddenRefDecimal(tc.bits)
			got, err := hiddenCallLevel1WithTimeout(head, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if got != want {
				t.Fatalf("Hidden test %q mismatch: bits=%v got=%d want=%d", tc.name, tc.bits, got, want)
			}
			after, _ := hiddenToSlice(head)
			if !hiddenEqual(after, tc.bits) {
				t.Fatalf("Hidden test %q: the list must not be modified; got %v want %v", tc.name, after, tc.bits)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel1_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1001))

	for trial := 0; trial < 500; trial++ {
		n := 1 + rng.Intn(30)
		bits := hiddenRandomVals(rng, n, 0, 1)
		head, _ := hiddenBuild(bits)

		want := hiddenRefDecimal(bits)
		got, err := hiddenCallLevel1WithTimeout(head, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if got != want {
			t.Fatalf("Property trial %d mismatch! bits=%v got=%d want=%d", trial, bits, got, want)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel1_PerformanceAudit(t *testing.T) {
	head, _ := hiddenBuild(hiddenRepeat(1, 30))

	start := time.Now()
	for i := 0; i < 200000; i++ {
		GetDecimalValue(head)
	}
	elapsed := time.Since(start)

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenGetDecimalValue(f *testing.F) {
	f.Add(uint32(5), uint8(3))
	f.Add(uint32(0), uint8(1))
	f.Add(uint32(0xFFFFFFFF), uint8(30))

	f.Fuzz(func(t *testing.T, pattern uint32, length uint8) {
		n := int(length%30) + 1
		bits := make([]int, n)
		for i := range bits {
			bits[i] = int((pattern >> uint(i)) & 1)
		}
		head, _ := hiddenBuild(bits)

		want := hiddenRefDecimal(bits)
		got := GetDecimalValue(head)
		if got != want {
			t.Fatalf("Fuzz mismatch! bits=%v -> got %d want %d", bits, got, want)
		}
	})
}
