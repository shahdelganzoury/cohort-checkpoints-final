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

// builds two lists that share a tail of ls nodes; the lists have la and lb private nodes in front
func hiddenMakeLists(la, lb, ls int, valFn func() int) (*ListNode, *ListNode, *ListNode) {
	if la+ls == 0 {
		la = 1
	}
	if lb+ls == 0 {
		lb = 1
	}
	build := func(n int, tail *ListNode) *ListNode {
		head := tail
		for i := 0; i < n; i++ {
			head = &ListNode{Val: valFn(), Next: head}
		}
		return head
	}
	shared := build(ls, nil) // nil when ls == 0
	a := build(la, shared)
	b := build(lb, shared)
	return a, b, shared
}

func hiddenCheckIntersect(want, got *ListNode, a, b *ListNode, aVals, bVals []int) string {
	if got != want {
		return fmt.Sprintf("got node %p want node %p", got, want)
	}
	aAfter, okA := hiddenToSlice(a)
	bAfter, okB := hiddenToSlice(b)
	if !okA || !okB || !hiddenEqual(aAfter, aVals) || !hiddenEqual(bAfter, bVals) {
		return "the lists must keep their original structure"
	}
	return ""
}

func hiddenCallLevel7WithTimeout(a, b *ListNode, timeout time.Duration) (*ListNode, error) {
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
		ch <- result{value: GetIntersectionNode(a, b)}
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
func TestHiddenLevel7_CornerCases(t *testing.T) {
	same := func() int { return 7 }
	rng := rand.New(rand.NewSource(7))
	small := func() int { return rng.Intn(3) }

	tests := []struct {
		name       string
		la, lb, ls int
		valFn      func() int
	}{
		{"classic_intersection", 3, 2, 3, small},
		{"no_intersection_single_nodes", 1, 1, 0, small},
		{"same_single_node", 0, 0, 1, small},
		{"same_whole_list", 0, 0, 5, small},
		{"a_is_part_of_b", 0, 4, 3, small},
		{"b_is_part_of_a", 5, 0, 2, small},
		{"intersect_at_last_node", 10, 1, 1, small},
		{"very_different_lengths_no_match", 1, 10, 0, small},
		{"equal_values_no_intersection", 3, 3, 0, same},
		{"equal_values_with_intersection", 3, 2, 4, same},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, b, want := hiddenMakeLists(tc.la, tc.lb, tc.ls, tc.valFn)
			aVals, _ := hiddenToSlice(a)
			bVals, _ := hiddenToSlice(b)

			got, err := hiddenCallLevel7WithTimeout(a, b, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckIntersect(want, got, a, b, aVals, bVals); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel7_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(7007))
	small := func() int { return rng.Intn(3) }

	for trial := 0; trial < 500; trial++ {
		la, lb, ls := rng.Intn(21), rng.Intn(21), rng.Intn(21)
		a, b, want := hiddenMakeLists(la, lb, ls, small)
		aVals, _ := hiddenToSlice(a)
		bVals, _ := hiddenToSlice(b)

		got, err := hiddenCallLevel7WithTimeout(a, b, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckIntersect(want, got, a, b, aVals, bVals); msg != "" {
			t.Fatalf("Property trial %d mismatch! la=%d lb=%d ls=%d: %s", trial, la, lb, ls, msg)
		}
	}
}

// Layer 3: Performance Audit (a nested-loop solution fails here)
func TestHiddenLevel7_PerformanceAudit(t *testing.T) {
	one := func() int { return 1 }
	a, b, want := hiddenMakeLists(100000, 50000, 100000, one)
	a2, b2, _ := hiddenMakeLists(100000, 100000, 0, one)

	start := time.Now()
	for i := 0; i < 20; i++ {
		got, err := hiddenCallLevel7WithTimeout(a, b, 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed: %v", err)
		}
		if got != want {
			t.Fatalf("Performance audit failed: wrong intersection node")
		}
		got, err = hiddenCallLevel7WithTimeout(a2, b2, 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed: %v", err)
		}
		if got != nil {
			t.Fatalf("Performance audit failed: expected nil for lists that do not intersect")
		}
	}
	elapsed := time.Since(start)

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenGetIntersectionNode(f *testing.F) {
	f.Add(int64(1), uint8(3), uint8(2), uint8(3))
	f.Add(int64(2), uint8(0), uint8(0), uint8(0))
	f.Add(int64(3), uint8(5), uint8(0), uint8(1))

	f.Fuzz(func(t *testing.T, seed int64, pa, pb, ps uint8) {
		rng := rand.New(rand.NewSource(seed))
		small := func() int { return rng.Intn(3) }
		a, b, want := hiddenMakeLists(int(pa)%20, int(pb)%20, int(ps)%20, small)
		aVals, _ := hiddenToSlice(a)
		bVals, _ := hiddenToSlice(b)

		got := GetIntersectionNode(a, b)
		if msg := hiddenCheckIntersect(want, got, a, b, aVals, bVals); msg != "" {
			t.Fatalf("Fuzz mismatch! pa=%d pb=%d ps=%d: %s", pa, pb, ps, msg)
		}
	})
}
