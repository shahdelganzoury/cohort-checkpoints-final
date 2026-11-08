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

func hiddenCheckDeleteNoHead(vals []int, idx int, head *ListNode) string {
	want := append(append([]int{}, vals[:idx]...), vals[idx+1:]...)
	got, ok := hiddenToSlice(head)
	if !ok {
		return "the list never ends (cycle?)"
	}
	if !hiddenEqual(got, want) {
		return fmt.Sprintf("got %v want %v", got, want)
	}
	return ""
}

func hiddenCallLevel4WithTimeout(node *ListNode, timeout time.Duration) error {
	done := make(chan struct{}, 1)
	failed := make(chan interface{}, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				failed <- r
			}
		}()
		DeleteNodeNoHead(node)
		done <- struct{}{}
	}()

	select {
	case <-done:
		return nil
	case r := <-failed:
		return fmt.Errorf("panic: %v", r)
	case <-time.After(timeout):
		return fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel4_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		idx  int
	}{
		{"two_nodes_delete_first", []int{4, 5}, 0},
		{"three_nodes_delete_first", []int{1, 2, 3}, 0},
		{"three_nodes_delete_middle", []int{1, 2, 3}, 1},
		{"four_nodes_second_to_last", []int{1, 2, 3, 4}, 2},
		{"negative_values", []int{-5, -4, -3, -2}, 1},
		{"zero_value", []int{0, 1, 2}, 0},
		{"long_list_delete_head", []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 0},
		{"long_list_second_to_last", []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}, 8},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenBuild(tc.vals)
			if err := hiddenCallLevel4WithTimeout(nodes[tc.idx], 1*time.Second); err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckDeleteNoHead(tc.vals, tc.idx, head); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel4_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(4004))

	for trial := 0; trial < 500; trial++ {
		n := 2 + rng.Intn(29)
		vals := rng.Perm(n) // unique values, as the problem guarantees
		idx := rng.Intn(n - 1)
		head, nodes := hiddenBuild(vals)

		if err := hiddenCallLevel4WithTimeout(nodes[idx], 500*time.Millisecond); err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckDeleteNoHead(vals, idx, head); msg != "" {
			t.Fatalf("Property trial %d mismatch! vals=%v idx=%d: %s", trial, vals, idx, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel4_PerformanceAudit(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		_, nodes := hiddenBuild([]int{1, 2, 3, 4})
		DeleteNodeNoHead(nodes[1])
	}
	elapsed := time.Since(start)

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenDeleteNodeNoHead(f *testing.F) {
	f.Add(int64(1), uint8(5), uint8(2))
	f.Add(int64(2), uint8(0), uint8(0))
	f.Add(int64(3), uint8(30), uint8(200))

	f.Fuzz(func(t *testing.T, seed int64, n uint8, idx uint8) {
		ln := int(n)%29 + 2
		rng := rand.New(rand.NewSource(seed))
		vals := rng.Perm(ln)
		i := int(idx) % (ln - 1)
		head, nodes := hiddenBuild(vals)

		DeleteNodeNoHead(nodes[i])
		if msg := hiddenCheckDeleteNoHead(vals, i, head); msg != "" {
			t.Fatalf("Fuzz mismatch! vals=%v idx=%d: %s", vals, i, msg)
		}
	})
}
