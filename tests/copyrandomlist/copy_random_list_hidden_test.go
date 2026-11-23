package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// spec: [value, index of the Random target or -1]
func hiddenBuildRandom(spec [][2]int) (*ListNode, []*ListNode) {
	nodes := make([]*ListNode, len(spec))
	for i := range spec {
		nodes[i] = &ListNode{Val: spec[i][0]}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	for i, s := range spec {
		if s[1] >= 0 {
			nodes[i].Random = nodes[s[1]]
		}
	}
	if len(nodes) == 0 {
		return nil, nodes
	}
	return nodes[0], nodes
}

func hiddenRandomSpec(rng *rand.Rand, n, lo, hi int) [][2]int {
	spec := make([][2]int, n)
	for i := range spec {
		spec[i] = [2]int{lo + rng.Intn(hi-lo+1), rng.Intn(n+1) - 1}
	}
	return spec
}

func hiddenCheckCopy(spec [][2]int, nodes []*ListNode, got *ListNode) string {
	// the original list must be exactly as it was
	for i, n := range nodes {
		var wantNext, wantRandom *ListNode
		if i+1 < len(nodes) {
			wantNext = nodes[i+1]
		}
		if spec[i][1] >= 0 {
			wantRandom = nodes[spec[i][1]]
		}
		if n.Val != spec[i][0] {
			return fmt.Sprintf("original node %d: Val was changed", i)
		}
		if n.Next != wantNext {
			return fmt.Sprintf("original node %d: Next was changed (restore the list)", i)
		}
		if n.Random != wantRandom {
			return fmt.Sprintf("original node %d: Random was changed", i)
		}
	}

	if len(spec) == 0 {
		if got != nil {
			return "an empty list must copy to nil"
		}
		return ""
	}

	originals := map[*ListNode]bool{}
	for _, n := range nodes {
		originals[n] = true
	}
	copies := []*ListNode{}
	index := map[*ListNode]int{}
	for c := got; c != nil; c = c.Next {
		if len(copies) > len(spec) {
			return "the copy is longer than the original (or never ends)"
		}
		if originals[c] {
			return fmt.Sprintf("copy node %d is an original node; the copy must use brand-new nodes", len(copies))
		}
		if _, dup := index[c]; dup {
			return "the same node appears twice in the copy"
		}
		index[c] = len(copies)
		copies = append(copies, c)
	}
	if len(copies) != len(spec) {
		return fmt.Sprintf("the copy has %d nodes; want %d", len(copies), len(spec))
	}
	for i, c := range copies {
		if c.Val != spec[i][0] {
			return fmt.Sprintf("copy node %d: Val = %d; want %d", i, c.Val, spec[i][0])
		}
		if spec[i][1] < 0 {
			if c.Random != nil {
				return fmt.Sprintf("copy node %d: Random should be nil", i)
			}
			continue
		}
		if c.Random == nil {
			return fmt.Sprintf("copy node %d: Random is nil; want node %d", i, spec[i][1])
		}
		if originals[c.Random] {
			return fmt.Sprintf("copy node %d: Random points into the ORIGINAL list", i)
		}
		idx, ok := index[c.Random]
		if !ok {
			return fmt.Sprintf("copy node %d: Random points outside the copy", i)
		}
		if idx != spec[i][1] {
			return fmt.Sprintf("copy node %d: Random points to node %d; want %d", i, idx, spec[i][1])
		}
	}
	return ""
}

func hiddenCallLevel1WithTimeout(head *ListNode, timeout time.Duration) (*ListNode, error) {
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
		ch <- result{value: CopyRandomList(head)}
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
func TestHiddenLevel1_CornerCases(t *testing.T) {
	long := make([][2]int, 1000)
	for i := range long {
		long[i] = [2]int{i%5 - 2, (i * 7) % 1000}
	}

	tests := []struct {
		name string
		spec [][2]int
	}{
		{"empty_list", [][2]int{}},
		{"single_no_random", [][2]int{{5, -1}}},
		{"single_self_random", [][2]int{{5, 0}}},
		{"two_nodes_point_at_each_other", [][2]int{{1, 1}, {2, 0}}},
		{"example_1", [][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}},
		{"example_2_equal_values", [][2]int{{3, -1}, {3, 0}, {3, -1}}},
		{"all_random_nil", [][2]int{{1, -1}, {2, -1}, {3, -1}}},
		{"all_random_to_head", [][2]int{{1, 0}, {2, 0}, {3, 0}, {4, 0}}},
		{"all_random_to_tail", [][2]int{{1, 3}, {2, 3}, {3, 3}, {4, 3}}},
		{"all_self_random", [][2]int{{1, 0}, {2, 1}, {3, 2}}},
		{"random_to_previous", [][2]int{{1, -1}, {2, 0}, {3, 1}, {4, 2}}},
		{"extreme_values", [][2]int{{-10000, 2}, {10000, 0}, {0, 1}}},
		{"all_equal_values_shuffled_randoms", [][2]int{{9, 2}, {9, 0}, {9, 3}, {9, 1}}},
		{"long_list_1000", long},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenBuildRandom(tc.spec)
			got, err := hiddenCallLevel1WithTimeout(head, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckCopy(tc.spec, nodes, got); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel1_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(1001))

	for trial := 0; trial < 500; trial++ {
		spec := hiddenRandomSpec(rng, rng.Intn(61), -3, 3) // few distinct values: many duplicates
		head, nodes := hiddenBuildRandom(spec)

		got, err := hiddenCallLevel1WithTimeout(head, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckCopy(spec, nodes, got); msg != "" {
			t.Fatalf("Property trial %d mismatch! spec=%v: %s", trial, spec, msg)
		}
	}
}

// Layer 3: Performance Audit (looking up each Random target by scanning fails here)
func TestHiddenLevel1_PerformanceAudit(t *testing.T) {
	n := 50000
	spec := make([][2]int, n)
	for i := range spec {
		spec[i] = [2]int{i % 100, (i * 31) % n}
	}
	head, nodes := hiddenBuildRandom(spec)

	var got *ListNode
	start := time.Now()
	for i := 0; i < 5; i++ {
		var err error
		got, err = hiddenCallLevel1WithTimeout(head, 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed: %v", err)
		}
	}
	elapsed := time.Since(start)

	if msg := hiddenCheckCopy(spec, nodes, got); msg != "" {
		t.Fatalf("Performance audit failed: %s", msg)
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenCopyRandomList(f *testing.F) {
	f.Add(int64(1), uint8(5))
	f.Add(int64(2), uint8(0))
	f.Add(int64(3), uint8(60))

	f.Fuzz(func(t *testing.T, seed int64, n uint8) {
		rng := rand.New(rand.NewSource(seed))
		spec := hiddenRandomSpec(rng, int(n)%80, -10000, 10000)
		head, nodes := hiddenBuildRandom(spec)

		got := CopyRandomList(head)
		if msg := hiddenCheckCopy(spec, nodes, got); msg != "" {
			t.Fatalf("Fuzz mismatch! spec=%v: %s", spec, msg)
		}
	})
}
