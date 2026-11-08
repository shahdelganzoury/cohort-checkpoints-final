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

func hiddenCheckReverse(vals []int, got *ListNode, nodes []*ListNode) string {
	want := make([]int, len(vals))
	for i, v := range vals {
		want[len(vals)-1-i] = v
	}
	gotVals, ok := hiddenToSlice(got)
	if !ok {
		return "the result list never ends (cycle?)"
	}
	if !hiddenEqual(gotVals, want) {
		return fmt.Sprintf("got %v want %v", gotVals, want)
	}
	curr := got
	for i := len(nodes) - 1; i >= 0; i-- {
		if curr != nodes[i] {
			return "the list must be reversed in place by reusing the original nodes"
		}
		curr = curr.Next
	}
	return ""
}

func hiddenCallLevel6WithTimeout(head *ListNode, timeout time.Duration) (*ListNode, error) {
	type result struct {
		value    *ListNode
		panicked interface{}
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{panicked: r}
			}
		}()
		ch <- result{value: ReverseLinkedList(head)}
	}()

	select {
	case res := <-ch:
		if res.panicked != nil {
			return nil, fmt.Errorf("panic: %v", res.panicked)
		}
		return res.value, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel6_CornerCases(t *testing.T) {
	increasing := make([]int, 1000)
	for i := range increasing {
		increasing[i] = i
	}

	tests := []struct {
		name string
		vals []int
	}{
		{"empty_list", []int{}},
		{"single_node", []int{1}},
		{"two_nodes", []int{1, 2}},
		{"three_nodes", []int{1, 2, 3}},
		{"duplicates", []int{4, 4, 5, 5}},
		{"all_equal_values", []int{9, 9, 9}},
		{"negative_values", []int{-1, -2, -3}},
		{"increasing_1000", increasing},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenBuild(tc.vals)
			got, err := hiddenCallLevel6WithTimeout(head, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckReverse(tc.vals, got, nodes); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel6_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(6006))

	for trial := 0; trial < 500; trial++ {
		vals := hiddenRandomVals(rng, rng.Intn(51), -100, 100)
		head, nodes := hiddenBuild(vals)

		got, err := hiddenCallLevel6WithTimeout(head, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckReverse(vals, got, nodes); msg != "" {
			t.Fatalf("Property trial %d mismatch! vals=%v: %s", trial, vals, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel6_PerformanceAudit(t *testing.T) {
	vals := make([]int, 100000)
	for i := range vals {
		vals[i] = i
	}
	head, nodes := hiddenBuild(vals)

	start := time.Now()
	reversed, err := hiddenCallLevel6WithTimeout(head, 1*time.Second)
	if err != nil {
		t.Fatalf("Performance audit failed: %v", err)
	}
	restored, err := hiddenCallLevel6WithTimeout(reversed, 1*time.Second)
	if err != nil {
		t.Fatalf("Performance audit failed: %v", err)
	}
	elapsed := time.Since(start)

	if reversed != nodes[len(nodes)-1] || restored != nodes[0] {
		t.Fatalf("Performance audit failed: the original nodes must be reused")
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenReverseLinkedList(f *testing.F) {
	f.Add(int64(1), uint8(5))
	f.Add(int64(2), uint8(0))
	f.Add(int64(3), uint8(100))

	f.Fuzz(func(t *testing.T, seed int64, n uint8) {
		rng := rand.New(rand.NewSource(seed))
		vals := hiddenRandomVals(rng, int(n)%80, -1000, 1000)
		head, nodes := hiddenBuild(vals)

		got := ReverseLinkedList(head)
		if msg := hiddenCheckReverse(vals, got, nodes); msg != "" {
			t.Fatalf("Fuzz mismatch! vals=%v: %s", vals, msg)
		}
	})
}
