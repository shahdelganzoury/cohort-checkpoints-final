package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

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

func hiddenBuildDoubly(vals []int) (*DoublyListNode, []*DoublyListNode) {
	var head, tail *DoublyListNode
	nodes := make([]*DoublyListNode, 0, len(vals))
	for _, v := range vals {
		n := &DoublyListNode{Val: v}
		nodes = append(nodes, n)
		if head == nil {
			head = n
		} else {
			tail.Next = n
			n.Prev = tail
		}
		tail = n
	}
	return head, nodes
}

// hidden reference - do not expose
func hiddenRefInsert(vals []int, val int) []int {
	idx := len(vals)
	for i, v := range vals {
		if v >= val {
			idx = i
			break
		}
	}
	out := append([]int{}, vals[:idx]...)
	out = append(out, val)
	return append(out, vals[idx:]...)
}

func hiddenCheckDoubly(vals []int, val int, got *DoublyListNode, nodes []*DoublyListNode) string {
	want := hiddenRefInsert(vals, val)
	if got == nil {
		return "the head must not be nil after an insertion"
	}
	if got.Prev != nil {
		return "the Prev of the head must be nil"
	}

	forward := []int{}
	var tail *DoublyListNode
	var prev *DoublyListNode
	j := 0
	newNodes := 0
	for curr := got; curr != nil; curr = curr.Next {
		if len(forward) > 100000 {
			return "the list never ends (cycle?)"
		}
		if curr.Prev != prev {
			return fmt.Sprintf("the Prev pointer of node %d is wrong", len(forward))
		}
		forward = append(forward, curr.Val)
		if j < len(nodes) && curr == nodes[j] {
			j++
		} else {
			newNodes++
		}
		prev = curr
		tail = curr
	}
	if !hiddenEqual(forward, want) {
		return fmt.Sprintf("forward order: got %v want %v", forward, want)
	}
	if j != len(nodes) || newNodes != 1 {
		return "the original nodes must be reused in order and exactly one new node added"
	}

	backward := []int{}
	for curr := tail; curr != nil; curr = curr.Prev {
		backward = append(backward, curr.Val)
		if len(backward) > 100000 {
			return "the Prev chain never ends"
		}
	}
	for i := range want {
		if backward[len(backward)-1-i] != want[i] {
			return fmt.Sprintf("backward order is wrong: got %v", backward)
		}
	}
	return ""
}

func hiddenCallLevel10WithTimeout(head *DoublyListNode, val int, timeout time.Duration) (*DoublyListNode, error) {
	type result struct {
		value    *DoublyListNode
		panicked interface{}
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{panicked: r}
			}
		}()
		ch <- result{value: InsertSortedDoubly(head, val)}
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

func hiddenSortedVals(rng *rand.Rand, n, lo, hi int) []int {
	vals := hiddenRandomVals(rng, n, lo, hi)
	for i := 1; i < len(vals); i++ {
		for j := i; j > 0 && vals[j-1] > vals[j]; j-- {
			vals[j-1], vals[j] = vals[j], vals[j-1]
		}
	}
	return vals
}

// Layer 1: Corner Cases
func TestHiddenLevel10_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		val  int
	}{
		{"empty_list", []int{}, 5},
		{"single_before", []int{5}, 1},
		{"single_after", []int{5}, 9},
		{"single_equal", []int{5}, 5},
		{"new_head", []int{3, 7}, 1},
		{"equal_to_head", []int{3, 7}, 3},
		{"middle", []int{3, 7}, 5},
		{"equal_to_tail", []int{3, 7}, 7},
		{"new_tail", []int{3, 7}, 9},
		{"between_duplicates", []int{2, 4, 4, 6}, 4},
		{"all_equal_values", []int{4, 4, 4}, 4},
		{"negative_values", []int{-9, -5, -1}, -6},
		{"zero_in_negatives_and_positives", []int{-2, -1, 1, 2}, 0},
		{"long_list_middle", []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 10},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenBuildDoubly(tc.vals)
			got, err := hiddenCallLevel10WithTimeout(head, tc.val, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckDoubly(tc.vals, tc.val, got, nodes); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel10_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1010))

	for trial := 0; trial < 500; trial++ {
		vals := hiddenSortedVals(rng, rng.Intn(31), -20, 20)
		val := rng.Intn(51) - 25
		head, nodes := hiddenBuildDoubly(vals)

		got, err := hiddenCallLevel10WithTimeout(head, val, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckDoubly(vals, val, got, nodes); msg != "" {
			t.Fatalf("Property trial %d mismatch! vals=%v val=%d: %s", trial, vals, val, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel10_PerformanceAudit(t *testing.T) {
	start := time.Now()
	var head *DoublyListNode
	for i := 0; i < 3000; i++ {
		head = InsertSortedDoubly(head, (i*7919)%1000)
	}
	elapsed := time.Since(start)

	count := 0
	prevVal := -1
	var prev *DoublyListNode
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val < prevVal || curr.Prev != prev {
			t.Fatalf("Performance audit failed: list is not sorted or Prev is wrong at node %d", count)
		}
		prevVal = curr.Val
		prev = curr
		count++
	}
	if count != 3000 {
		t.Fatalf("Performance audit failed: expected 3000 nodes, got %d", count)
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenInsertSortedDoubly(f *testing.F) {
	f.Add(int64(1), uint8(5), int8(3))
	f.Add(int64(2), uint8(0), int8(0))
	f.Add(int64(3), uint8(30), int8(-25))

	f.Fuzz(func(t *testing.T, seed int64, n uint8, val int8) {
		rng := rand.New(rand.NewSource(seed))
		vals := hiddenSortedVals(rng, int(n)%40, -20, 20)
		head, nodes := hiddenBuildDoubly(vals)

		got := InsertSortedDoubly(head, int(val))
		if msg := hiddenCheckDoubly(vals, int(val), got, nodes); msg != "" {
			t.Fatalf("Fuzz mismatch! vals=%v val=%d: %s", vals, val, msg)
		}
	})
}
