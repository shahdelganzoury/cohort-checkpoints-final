package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// builds a circle of the given values; the returned slice holds its nodes in order
func hiddenRing(vals []int) []*ListNode {
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
	}
	for i := range nodes {
		nodes[i].Next = nodes[(i+1)%len(nodes)]
	}
	return nodes
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
// vals is the circle in node order, s is the index of the start node
func hiddenRefHop(vals []int, s int) []int {
	if len(vals) == 0 {
		return []int{}
	}
	circle := append(append([]int{}, vals[s:]...), vals[:s]...)
	out := []int{}
	cur := 0
	for len(circle) > 0 {
		v := circle[cur]
		out = append(out, v)
		circle = append(circle[:cur], circle[cur+1:]...)
		m := len(circle)
		if m == 0 {
			break
		}
		cur = (cur + v - 1) % m
	}
	return out
}

func hiddenCallLevel4WithTimeout(start *ListNode, timeout time.Duration) ([]int, error) {
	type result struct {
		value    []int
		panicked interface{}
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{panicked: r}
			}
		}()
		ch <- result{value: HopEliminate(start)}
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

func hiddenStart(nodes []*ListNode, s int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}
	return nodes[s]
}

// Layer 1: Corner Cases
func TestHiddenLevel4_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		s    int
	}{
		{"empty_circle", []int{}, 0},
		{"single_node", []int{9}, 0},
		{"single_node_huge_value", []int{1000000000}, 0},
		{"two_nodes", []int{1, 2}, 0},
		{"two_nodes_start_second", []int{1, 2}, 1},
		{"example_1", []int{3, 1, 4, 1, 5}, 0},
		{"example_2_huge_value", []int{1000000000, 7, 8}, 0},
		{"all_equal_values", []int{2, 2, 2, 2}, 0},
		{"value_one_everywhere", []int{1, 1, 1, 1, 1}, 0},
		{"value_equals_size", []int{3, 3, 3}, 0},
		{"value_multiple_of_size", []int{6, 6, 6}, 0},
		{"start_in_the_middle", []int{5, 4, 3, 2, 1}, 2},
		{"start_at_the_last_node", []int{5, 4, 3, 2, 1}, 4},
		{"descending_values", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 0},
		{"large_mixed_values", []int{1000000000, 999999999, 123456789, 1, 2}, 3},
		{"values_smaller_than_size", []int{1, 2, 1, 2, 1, 2, 1, 2}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nodes := hiddenRing(tc.vals)
			want := hiddenRefHop(tc.vals, tc.s)
			got, err := hiddenCallLevel4WithTimeout(hiddenStart(nodes, tc.s), 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if !hiddenEqual(got, want) {
				t.Fatalf("Hidden test %q mismatch: vals=%v start=%d got=%v want=%v", tc.name, tc.vals, tc.s, got, want)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel4_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(4004))

	for trial := 0; trial < 500; trial++ {
		n := rng.Intn(41)
		var vals []int
		if rng.Intn(4) == 0 {
			vals = hiddenRandomVals(rng, n, 999999000, 1000000000) // very large hops
		} else {
			vals = hiddenRandomVals(rng, n, 1, 30)
		}
		s := 0
		if n > 0 {
			s = rng.Intn(n)
		}
		nodes := hiddenRing(vals)
		want := hiddenRefHop(vals, s)

		got, err := hiddenCallLevel4WithTimeout(hiddenStart(nodes, s), 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if !hiddenEqual(got, want) {
			t.Fatalf("Property trial %d mismatch! vals=%v start=%d got=%v want=%v", trial, vals, s, got, want)
		}
	}
}

// Layer 3: Performance Audit (walking Val steps one by one, with no modulo, fails here)
func TestHiddenLevel4_PerformanceAudit(t *testing.T) {
	rng := rand.New(rand.NewSource(44))
	big := hiddenRepeat(1000000000, 5000)
	mixed := hiddenRandomVals(rng, 5000, 1, 1000000000)

	start := time.Now()
	for _, vals := range [][]int{big, mixed} {
		nodes := hiddenRing(vals)
		want := hiddenRefHop(vals, 0)

		got, err := hiddenCallLevel4WithTimeout(nodes[0], 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed: %v (take Val modulo the circle size)", err)
		}
		if !hiddenEqual(got, want) {
			t.Fatalf("Performance audit failed: wrong removal order")
		}
	}
	elapsed := time.Since(start)

	if elapsed > 2*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 2s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenHopEliminate(f *testing.F) {
	f.Add(int64(1), uint8(5), uint8(0), true)
	f.Add(int64(2), uint8(0), uint8(0), false)
	f.Add(int64(3), uint8(30), uint8(7), true)

	f.Fuzz(func(t *testing.T, seed int64, n uint8, s uint8, huge bool) {
		rng := rand.New(rand.NewSource(seed))
		size := int(n) % 40
		var vals []int
		if huge {
			vals = hiddenRandomVals(rng, size, 1, 1000000000)
		} else {
			vals = hiddenRandomVals(rng, size, 1, 20)
		}
		start := 0
		if size > 0 {
			start = int(s) % size
		}
		nodes := hiddenRing(vals)
		want := hiddenRefHop(vals, start)

		got := HopEliminate(hiddenStart(nodes, start))
		if !hiddenEqual(got, want) {
			t.Fatalf("Fuzz mismatch! vals=%v start=%d got=%v want=%v", vals, start, got, want)
		}
	})
}
