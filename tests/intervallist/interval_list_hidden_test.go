package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func hiddenBuildIntervals(iv [][2]int) (*ListNode, []*ListNode) {
	var head, tail *ListNode
	nodes := make([]*ListNode, 0, len(iv))
	for _, p := range iv {
		n := &ListNode{Start: p[0], End: p[1]}
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

// sorted, non-overlapping and non-touching intervals
func hiddenRandomIntervals(rng *rand.Rand, n int) [][2]int {
	out := [][2]int{}
	cur := rng.Intn(20) - 10
	for i := 0; i < n; i++ {
		length := rng.Intn(5)
		out = append(out, [2]int{cur, cur + length})
		cur = cur + length + 1 + rng.Intn(4)
	}
	return out
}

func hiddenEqualIv(a, b [][2]int) bool {
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

// hidden reference - do not expose
func hiddenRefAdd(iv [][2]int, s, e int) [][2]int {
	out := [][2]int{}
	if s > e {
		return append(out, iv...)
	}
	ns, ne := s, e
	i := 0
	for i < len(iv) && iv[i][1] < s {
		out = append(out, iv[i])
		i++
	}
	for i < len(iv) && iv[i][0] <= e {
		if iv[i][0] < ns {
			ns = iv[i][0]
		}
		if iv[i][1] > ne {
			ne = iv[i][1]
		}
		i++
	}
	out = append(out, [2]int{ns, ne})
	for ; i < len(iv); i++ {
		out = append(out, iv[i])
	}
	return out
}

func hiddenCheckInterval(iv [][2]int, s, e int, nodes []*ListNode, got *ListNode) string {
	want := hiddenRefAdd(iv, s, e)
	originals := map[*ListNode]bool{}
	for _, n := range nodes {
		originals[n] = true
	}
	gotIv := [][2]int{}
	newNodes := 0
	for c := got; c != nil; c = c.Next {
		if len(gotIv) > len(iv)+1 {
			return "the list is too long (or never ends)"
		}
		gotIv = append(gotIv, [2]int{c.Start, c.End})
		if !originals[c] {
			newNodes++
		}
	}
	if !hiddenEqualIv(gotIv, want) {
		return fmt.Sprintf("got %v want %v", gotIv, want)
	}
	if s > e {
		if newNodes != 0 || (len(nodes) > 0 && got != nodes[0]) {
			return "start > end: the list must be returned unchanged, without creating a node"
		}
		return ""
	}
	if newNodes > 1 {
		return fmt.Sprintf("%d new nodes were created; at most one is allowed per call", newNodes)
	}
	return ""
}

func hiddenCallLevel5WithTimeout(head *ListNode, start, end int, timeout time.Duration) (*ListNode, error) {
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
		ch <- result{value: AddInterval(head, start, end)}
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
func TestHiddenLevel5_CornerCases(t *testing.T) {
	tests := []struct {
		name string
		iv   [][2]int
		s, e int
	}{
		{"example_1_touching_merge", [][2]int{{1, 3}, {5, 6}, {10, 12}}, 3, 4},
		{"example_2_swallow_everything", [][2]int{{1, 4}, {5, 6}, {7, 7}, {8, 9}, {10, 12}}, 2, 11},
		{"example_3_empty_list", [][2]int{}, 4, 4},
		{"before_everything", [][2]int{{5, 6}, {8, 9}}, 1, 2},
		{"after_everything", [][2]int{{5, 6}, {8, 9}}, 20, 30},
		{"between_without_touching", [][2]int{{1, 2}, {8, 9}}, 4, 5},
		{"adjacent_integers_do_not_merge", [][2]int{{1, 2}}, 3, 4},
		{"adjacent_on_the_left_do_not_merge", [][2]int{{3, 4}}, 1, 2},
		{"touch_the_end_of_an_interval", [][2]int{{1, 3}}, 3, 5},
		{"touch_the_start_of_an_interval", [][2]int{{5, 7}}, 1, 5},
		{"bridge_two_intervals", [][2]int{{1, 2}, {4, 5}}, 2, 4},
		{"bridge_without_touching_either_end", [][2]int{{1, 2}, {5, 6}}, 3, 4},
		{"inside_an_existing_interval", [][2]int{{1, 10}}, 3, 4},
		{"equal_to_an_existing_interval", [][2]int{{1, 3}, {5, 6}}, 5, 6},
		{"swallow_the_middle", [][2]int{{1, 2}, {4, 5}, {7, 8}, {10, 11}}, 3, 9},
		{"swallow_the_first_ones", [][2]int{{1, 2}, {4, 5}, {7, 8}}, 0, 5},
		{"swallow_the_last_ones", [][2]int{{1, 2}, {4, 5}, {7, 8}}, 5, 20},
		{"point_intervals_stay_apart", [][2]int{{1, 1}, {3, 3}}, 2, 2},
		{"point_interval_inside", [][2]int{{1, 5}}, 3, 3},
		{"negative_values", [][2]int{{-10, -5}, {-3, -1}}, -6, -2},
		{"extreme_values_swallow_all", [][2]int{{-5, 5}, {10, 20}}, -1000000000, 1000000000},
		{"extreme_values_in_list", [][2]int{{-1000000000, -999999999}, {999999999, 1000000000}}, 0, 1},
		{"invalid_interval_empty_list", [][2]int{}, 9, 1},
		{"invalid_interval_non_empty_list", [][2]int{{1, 2}, {5, 6}}, 9, 1},
		{"invalid_interval_that_would_overlap", [][2]int{{1, 10}}, 5, 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenBuildIntervals(tc.iv)
			got, err := hiddenCallLevel5WithTimeout(head, tc.s, tc.e, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckInterval(tc.iv, tc.s, tc.e, nodes, got); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel5_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(5005))

	for trial := 0; trial < 500; trial++ {
		iv := hiddenRandomIntervals(rng, rng.Intn(16))
		s := rng.Intn(70) - 8
		e := s + rng.Intn(20)
		if rng.Intn(10) == 0 {
			e = s - 1 - rng.Intn(5) // sometimes an invalid interval
		}
		head, nodes := hiddenBuildIntervals(iv)

		got, err := hiddenCallLevel5WithTimeout(head, s, e, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckInterval(iv, s, e, nodes, got); msg != "" {
			t.Fatalf("Property trial %d mismatch! iv=%v add [%d,%d]: %s", trial, iv, s, e, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel5_PerformanceAudit(t *testing.T) {
	n := 100000
	iv := make([][2]int, n)
	for i := range iv {
		iv[i] = [2]int{3 * i, 3*i + 1}
	}
	head, nodes := hiddenBuildIntervals(iv)

	start := time.Now()
	for k := 0; k < 30; k++ {
		x := 3 * (k * 3000) // a point that is already inside an existing interval
		var err error
		head, err = hiddenCallLevel5WithTimeout(head, x, x, 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed: %v", err)
		}
	}
	got, err := hiddenCallLevel5WithTimeout(head, 0, 400000, 1*time.Second)
	if err != nil {
		t.Fatalf("Performance audit failed: %v", err)
	}
	elapsed := time.Since(start)

	_ = nodes
	if got == nil || got.Next != nil || got.Start != 0 || got.End != 400000 {
		t.Fatalf("Performance audit failed: expected the single interval [0,400000]")
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenAddInterval(f *testing.F) {
	f.Add(int64(1), uint8(5), int8(3), int8(8))
	f.Add(int64(2), uint8(0), int8(0), int8(0))
	f.Add(int64(3), uint8(15), int8(20), int8(-5))

	f.Fuzz(func(t *testing.T, seed int64, n uint8, s, e int8) {
		rng := rand.New(rand.NewSource(seed))
		iv := hiddenRandomIntervals(rng, int(n)%20)
		head, nodes := hiddenBuildIntervals(iv)

		got := AddInterval(head, int(s), int(e))
		if msg := hiddenCheckInterval(iv, int(s), int(e), nodes, got); msg != "" {
			t.Fatalf("Fuzz mismatch! iv=%v add [%d,%d]: %s", iv, s, e, msg)
		}
	})
}
