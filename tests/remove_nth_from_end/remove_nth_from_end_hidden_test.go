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
func hiddenCheckRemoveNth(vals []int, n int, got *ListNode, nodes []*ListNode) string {
	idx := len(vals) - n
	want := append(append([]int{}, vals[:idx]...), vals[idx+1:]...)
	gotVals, ok := hiddenToSlice(got)
	if !ok {
		return "the result list never ends (cycle?)"
	}
	if !hiddenEqual(gotVals, want) {
		return fmt.Sprintf("got %v want %v", gotVals, want)
	}
	curr := got
	for i, node := range nodes {
		if i == idx {
			continue
		}
		if curr != node {
			return fmt.Sprintf("node %d was replaced; the remaining nodes must be reused", i)
		}
		curr = curr.Next
	}
	return ""
}

func hiddenCallLevel9WithTimeout(head *ListNode, n int, timeout time.Duration) (*ListNode, error) {
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
		ch <- result{value: RemoveNthFromEnd(head, n)}
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
func TestHiddenLevel9_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		n    int
	}{
		{"single_node", []int{1}, 1},
		{"two_nodes_remove_tail", []int{1, 2}, 1},
		{"two_nodes_remove_head", []int{1, 2}, 2},
		{"three_nodes_remove_middle", []int{1, 2, 3}, 2},
		{"remove_head_of_five", []int{1, 2, 3, 4, 5}, 5},
		{"remove_second_node", []int{1, 2, 3, 4, 5}, 4},
		{"remove_second_to_last", []int{1, 2, 3, 4, 5}, 2},
		{"remove_tail_of_five", []int{1, 2, 3, 4, 5}, 1},
		{"duplicates", []int{7, 7, 7, 7}, 2},
		{"long_list_remove_middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 10},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenBuild(tc.vals)
			got, err := hiddenCallLevel9WithTimeout(head, tc.n, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckRemoveNth(tc.vals, tc.n, got, nodes); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel9_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(9009))

	for trial := 0; trial < 500; trial++ {
		size := 1 + rng.Intn(30)
		vals := hiddenRandomVals(rng, size, 0, 100)
		n := 1 + rng.Intn(size)
		head, nodes := hiddenBuild(vals)

		got, err := hiddenCallLevel9WithTimeout(head, n, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckRemoveNth(vals, n, got, nodes); msg != "" {
			t.Fatalf("Property trial %d mismatch! vals=%v n=%d: %s", trial, vals, n, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel9_PerformanceAudit(t *testing.T) {
	size := 100000
	vals := make([]int, size)
	for i := range vals {
		vals[i] = i
	}

	start := time.Now()
	for _, n := range []int{1, size / 2, size} {
		head, nodes := hiddenBuild(vals)
		got, err := hiddenCallLevel9WithTimeout(head, n, 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed (n=%d): %v", n, err)
		}
		if msg := hiddenCheckRemoveNth(vals, n, got, nodes); msg != "" {
			t.Fatalf("Performance audit failed (n=%d): %s", n, msg)
		}
	}
	elapsed := time.Since(start)

	if elapsed > 2*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 2s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenRemoveNthFromEnd(f *testing.F) {
	f.Add(int64(1), uint8(5), uint8(2))
	f.Add(int64(2), uint8(1), uint8(1))
	f.Add(int64(3), uint8(30), uint8(30))

	f.Fuzz(func(t *testing.T, seed int64, ln uint8, nn uint8) {
		rng := rand.New(rand.NewSource(seed))
		size := int(ln)%30 + 1
		n := int(nn)%size + 1
		vals := hiddenRandomVals(rng, size, 0, 100)
		head, nodes := hiddenBuild(vals)

		got := RemoveNthFromEnd(head, n)
		if msg := hiddenCheckRemoveNth(vals, n, got, nodes); msg != "" {
			t.Fatalf("Fuzz mismatch! vals=%v n=%d: %s", vals, n, msg)
		}
	})
}
