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

func hiddenCallLevel2WithTimeout(head *ListNode, val int, timeout time.Duration) (*ListNode, error) {
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
		ch <- result{value: AppendNode(head, val)}
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

func hiddenCheckAppend(vals []int, val int, got *ListNode, nodes []*ListNode) string {
	want := append(append([]int{}, vals...), val)
	gotVals, ok := hiddenToSlice(got)
	if !ok {
		return "the result list never ends (cycle?)"
	}
	if !hiddenEqual(gotVals, want) {
		return fmt.Sprintf("got %v want %v", gotVals, want)
	}
	curr := got
	for i, n := range nodes {
		if curr != n {
			return fmt.Sprintf("node %d was replaced; the original nodes must be reused", i)
		}
		curr = curr.Next
	}
	return ""
}

// Layer 1: Corner Cases
func TestHiddenLevel2_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		val  int
	}{
		{"empty_list", []int{}, 5},
		{"empty_list_zero", []int{}, 0},
		{"single_node", []int{1}, 2},
		{"two_nodes", []int{1, 2}, 3},
		{"negative_value", []int{4, 5}, -7},
		{"duplicate_of_tail", []int{3, 3}, 3},
		{"all_equal_values", []int{7, 7, 7, 7}, 7},
		{"long_list", hiddenRepeat(1, 1000), 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenBuild(tc.vals)
			got, err := hiddenCallLevel2WithTimeout(head, tc.val, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckAppend(tc.vals, tc.val, got, nodes); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel2_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(2002))

	for trial := 0; trial < 500; trial++ {
		vals := hiddenRandomVals(rng, rng.Intn(51), -100, 100)
		val := rng.Intn(201) - 100
		head, nodes := hiddenBuild(vals)

		got, err := hiddenCallLevel2WithTimeout(head, val, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckAppend(vals, val, got, nodes); msg != "" {
			t.Fatalf("Property trial %d mismatch! vals=%v val=%d: %s", trial, vals, val, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel2_PerformanceAudit(t *testing.T) {
	start := time.Now()
	var head *ListNode
	for i := 0; i < 3000; i++ {
		head = AppendNode(head, i)
	}
	elapsed := time.Since(start)

	vals, ok := hiddenToSlice(head)
	if !ok || len(vals) != 3000 {
		t.Fatalf("Performance audit failed: expected 3000 nodes, got %d (ok=%v)", len(vals), ok)
	}
	for i, v := range vals {
		if v != i {
			t.Fatalf("Performance audit failed: node %d has value %d", i, v)
		}
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenAppendNode(f *testing.F) {
	f.Add(int64(1), uint8(3), int16(7))
	f.Add(int64(2), uint8(0), int16(-1))
	f.Add(int64(3), uint8(50), int16(0))

	f.Fuzz(func(t *testing.T, seed int64, n uint8, val int16) {
		rng := rand.New(rand.NewSource(seed))
		vals := hiddenRandomVals(rng, int(n)%60, -1000, 1000)
		head, nodes := hiddenBuild(vals)

		got := AppendNode(head, int(val))
		if msg := hiddenCheckAppend(vals, int(val), got, nodes); msg != "" {
			t.Fatalf("Fuzz mismatch! vals=%v val=%d: %s", vals, val, msg)
		}
	})
}
