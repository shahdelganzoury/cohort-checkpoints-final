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

func hiddenMakeCycle(n, pos int, valFn func() int) (*ListNode, []*ListNode) {
	vals := make([]int, n)
	for i := range vals {
		vals[i] = valFn()
	}
	head, nodes := hiddenBuild(vals)
	if pos >= 0 && pos < n {
		nodes[n-1].Next = nodes[pos]
	}
	return head, nodes
}

func hiddenCheckStructure(nodes []*ListNode, pos int) string {
	for i, n := range nodes {
		var want *ListNode
		if i+1 < len(nodes) {
			want = nodes[i+1]
		} else if pos >= 0 && pos < len(nodes) {
			want = nodes[pos]
		}
		if n.Next != want {
			return fmt.Sprintf("node %d was modified; the list must not be changed", i)
		}
	}
	return ""
}

func hiddenCallLevel8WithTimeout(head *ListNode, timeout time.Duration) (bool, error) {
	type result struct {
		value    bool
		panicked interface{}
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{panicked: r}
			}
		}()
		ch <- result{value: HasCycle(head)}
	}()

	select {
	case res := <-ch:
		if res.panicked != nil {
			return false, fmt.Errorf("panic: %v", res.panicked)
		}
		return res.value, nil
	case <-time.After(timeout):
		return false, fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel8_CornerCases(t *testing.T) {
	same := func() int { return 7 }
	rng := rand.New(rand.NewSource(8))
	small := func() int { return rng.Intn(3) }

	tests := []struct {
		name  string
		n     int
		pos   int
		valFn func() int
	}{
		{"empty_list", 0, -1, small},
		{"single_no_cycle", 1, -1, small},
		{"single_self_loop", 1, 0, small},
		{"two_nodes_no_cycle", 2, -1, small},
		{"two_nodes_cycle_to_head", 2, 0, small},
		{"two_nodes_tail_self_loop", 2, 1, small},
		{"cycle_in_middle", 5, 2, small},
		{"cycle_to_second_to_last", 5, 3, small},
		{"five_nodes_no_cycle", 5, -1, small},
		{"long_cycle_to_head", 30, 0, small},
		{"long_cycle_in_middle", 30, 15, small},
		{"long_no_cycle", 30, -1, small},
		{"equal_values_no_cycle", 6, -1, same},
		{"equal_values_with_cycle", 6, 3, same},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenMakeCycle(tc.n, tc.pos, tc.valFn)
			want := tc.pos >= 0 && tc.pos < tc.n

			got, err := hiddenCallLevel8WithTimeout(head, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if got != want {
				t.Fatalf("Hidden test %q mismatch: n=%d cycleAt=%d got=%v want=%v", tc.name, tc.n, tc.pos, got, want)
			}
			if msg := hiddenCheckStructure(nodes, tc.pos); msg != "" {
				t.Fatalf("Hidden test %q: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel8_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(8008))
	small := func() int { return rng.Intn(3) }

	for trial := 0; trial < 500; trial++ {
		n := rng.Intn(31)
		pos := rng.Intn(n+1) - 1 // -1 means no cycle; n==0 always gives -1 or 0 (ignored)
		head, nodes := hiddenMakeCycle(n, pos, small)
		want := pos >= 0 && pos < n

		got, err := hiddenCallLevel8WithTimeout(head, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if got != want {
			t.Fatalf("Property trial %d mismatch! n=%d cycleAt=%d got=%v want=%v", trial, n, pos, got, want)
		}
		if msg := hiddenCheckStructure(nodes, pos); msg != "" {
			t.Fatalf("Property trial %d: %s", trial, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel8_PerformanceAudit(t *testing.T) {
	one := func() int { return 1 }
	withCycle, _ := hiddenMakeCycle(100000, 0, one)
	noCycle, _ := hiddenMakeCycle(100000, -1, one)

	start := time.Now()
	for i := 0; i < 20; i++ {
		got, err := hiddenCallLevel8WithTimeout(withCycle, 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed: %v", err)
		}
		if !got {
			t.Fatalf("Performance audit failed: cycle not detected")
		}
		got, err = hiddenCallLevel8WithTimeout(noCycle, 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed: %v", err)
		}
		if got {
			t.Fatalf("Performance audit failed: cycle reported on a plain list")
		}
	}
	elapsed := time.Since(start)

	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenHasCycle(f *testing.F) {
	f.Add(int64(1), uint8(5), int8(2))
	f.Add(int64(2), uint8(0), int8(-1))
	f.Add(int64(3), uint8(30), int8(0))

	f.Fuzz(func(t *testing.T, seed int64, n uint8, pos int8) {
		rng := rand.New(rand.NewSource(seed))
		small := func() int { return rng.Intn(3) }
		ln := int(n) % 40
		p := int(pos)
		if p < 0 || p >= ln {
			p = -1
		}
		head, _ := hiddenMakeCycle(ln, p, small)

		got := HasCycle(head)
		if got != (p >= 0) {
			t.Fatalf("Fuzz mismatch! n=%d cycleAt=%d got=%v want=%v", ln, p, got, p >= 0)
		}
	})
}
