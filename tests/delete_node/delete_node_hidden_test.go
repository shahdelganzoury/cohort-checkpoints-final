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
func hiddenRefDelete(vals []int, val int) ([]int, int) {
	idx := -1
	for i, v := range vals {
		if v == val {
			idx = i
			break
		}
	}
	if idx == -1 {
		return append([]int{}, vals...), -1
	}
	out := append([]int{}, vals[:idx]...)
	out = append(out, vals[idx+1:]...)
	return out, idx
}

func hiddenCheckDelete(vals []int, val int, got *ListNode, nodes []*ListNode) string {
	want, idx := hiddenRefDelete(vals, val)
	gotVals, ok := hiddenToSlice(got)
	if !ok {
		return "the result list never ends (cycle?)"
	}
	if !hiddenEqual(gotVals, want) {
		return fmt.Sprintf("got %v want %v", gotVals, want)
	}
	curr := got
	for i, n := range nodes {
		if i == idx {
			continue
		}
		if curr != n {
			return fmt.Sprintf("node %d was replaced; the remaining nodes must be reused", i)
		}
		curr = curr.Next
	}
	return ""
}

func hiddenCallLevel3WithTimeout(head *ListNode, val int, timeout time.Duration) (*ListNode, error) {
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
		ch <- result{value: DeleteNode(head, val)}
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
func TestHiddenLevel3_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		val  int
	}{
		{"empty_list", []int{}, 1},
		{"single_node_deleted", []int{1}, 1},
		{"single_node_not_found", []int{1}, 2},
		{"delete_head", []int{1, 2, 3}, 1},
		{"delete_middle", []int{1, 2, 3}, 2},
		{"delete_tail", []int{1, 2, 3}, 3},
		{"not_found", []int{1, 2, 3}, 9},
		{"duplicates_first_only", []int{2, 1, 2, 1, 2}, 2},
		{"duplicate_head_and_tail", []int{5, 6, 5}, 5},
		{"all_equal_values", []int{4, 4, 4}, 4},
		{"two_nodes_delete_second", []int{8, 9}, 9},
		{"negative_value", []int{-1, -2, -3}, -2},
		{"long_list_delete_tail", append(hiddenRepeat(1, 999), 2), 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenBuild(tc.vals)
			got, err := hiddenCallLevel3WithTimeout(head, tc.val, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckDelete(tc.vals, tc.val, got, nodes); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel3_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(3003))

	for trial := 0; trial < 500; trial++ {
		vals := hiddenRandomVals(rng, rng.Intn(31), -5, 5)
		val := rng.Intn(13) - 6
		head, nodes := hiddenBuild(vals)

		got, err := hiddenCallLevel3WithTimeout(head, val, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckDelete(vals, val, got, nodes); msg != "" {
			t.Fatalf("Property trial %d mismatch! vals=%v val=%d: %s", trial, vals, val, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel3_PerformanceAudit(t *testing.T) {
	head, _ := hiddenBuild(hiddenRepeat(1, 1000))

	start := time.Now()
	for i := 0; i < 5000; i++ {
		head = DeleteNode(head, 5000) // never found, walks the whole list
	}
	elapsed := time.Since(start)

	vals, ok := hiddenToSlice(head)
	if !ok || len(vals) != 1000 {
		t.Fatalf("Performance audit failed: list changed, got %d nodes (ok=%v)", len(vals), ok)
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenDeleteNode(f *testing.F) {
	f.Add(int64(1), uint8(5), int8(2))
	f.Add(int64(2), uint8(0), int8(0))
	f.Add(int64(3), uint8(30), int8(-3))

	f.Fuzz(func(t *testing.T, seed int64, n uint8, val int8) {
		rng := rand.New(rand.NewSource(seed))
		vals := hiddenRandomVals(rng, int(n)%40, -5, 5)
		head, nodes := hiddenBuild(vals)

		got := DeleteNode(head, int(val)%7)
		if msg := hiddenCheckDelete(vals, int(val)%7, got, nodes); msg != "" {
			t.Fatalf("Fuzz mismatch! vals=%v val=%d: %s", vals, int(val)%7, msg)
		}
	})
}
