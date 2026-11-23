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

func hiddenCheckDeal(nodes []*ListNode, k int, got []*ListNode) string {
	if k <= 0 {
		if got != nil {
			return fmt.Sprintf("k = %d must return nil; got a slice of length %d", k, len(got))
		}
		return ""
	}
	if len(got) != k {
		return fmt.Sprintf("want %d hands; got %d", k, len(got))
	}
	for h := 0; h < k; h++ {
		curr := got[h]
		for pos, i := 0, h; i < len(nodes); pos, i = pos+1, i+k {
			if curr != nodes[i] {
				return fmt.Sprintf("hand %d, card %d is not the expected original node %d (reuse the nodes, keep the order)", h, pos, i)
			}
			curr = curr.Next
		}
		if curr != nil {
			return fmt.Sprintf("hand %d has extra cards (its last Next must be nil)", h)
		}
	}
	return ""
}

func hiddenCallLevel3WithTimeout(head *ListNode, k int, timeout time.Duration) ([]*ListNode, error) {
	type result struct {
		value    []*ListNode
		panicked interface{}
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{panicked: r}
			}
		}()
		ch <- result{value: DealCards(head, k)}
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
		k    int
	}{
		{"example_1", []int{1, 2, 3, 4, 5, 6, 7}, 3},
		{"example_2_more_hands_than_nodes", []int{10, 20}, 4},
		{"example_3_empty_list", []int{}, 2},
		{"k_is_one", []int{5, 6, 7}, 1},
		{"k_equals_length", []int{1, 2, 3, 4}, 4},
		{"k_is_length_minus_one", []int{1, 2, 3, 4}, 3},
		{"k_is_two", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2},
		{"single_node_single_hand", []int{42}, 1},
		{"single_node_many_hands", []int{42}, 5},
		{"empty_list_k_one", []int{}, 1},
		{"k_zero", []int{1, 2}, 0},
		{"k_zero_empty_list", []int{}, 0},
		{"k_negative", []int{1, 2, 3}, -5},
		{"k_very_negative", []int{1, 2, 3}, -1000},
		{"equal_values_reuse_by_node", []int{7, 7, 7, 7, 7, 7}, 3},
		{"huge_k_few_nodes", []int{1, 2, 3}, 10000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, nodes := hiddenBuild(tc.vals)
			got, err := hiddenCallLevel3WithTimeout(head, tc.k, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckDeal(nodes, tc.k, got); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel3_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(3003))

	for trial := 0; trial < 500; trial++ {
		vals := hiddenRandomVals(rng, rng.Intn(41), -3, 3)
		k := rng.Intn(18) - 2
		head, nodes := hiddenBuild(vals)

		got, err := hiddenCallLevel3WithTimeout(head, k, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckDeal(nodes, k, got); msg != "" {
			t.Fatalf("Property trial %d mismatch! vals=%v k=%d: %s", trial, vals, k, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel3_PerformanceAudit(t *testing.T) {
	vals := make([]int, 100000)
	for i := range vals {
		vals[i] = i
	}
	ks := []int{3, 1000, 50000}
	heads := make([]*ListNode, len(ks))
	allNodes := make([][]*ListNode, len(ks))
	for i := range ks {
		heads[i], allNodes[i] = hiddenBuild(vals)
	}

	start := time.Now()
	results := make([][]*ListNode, len(ks))
	for i, k := range ks {
		var err error
		results[i], err = hiddenCallLevel3WithTimeout(heads[i], k, 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed (k=%d): %v", k, err)
		}
	}
	elapsed := time.Since(start)

	for i, k := range ks {
		if msg := hiddenCheckDeal(allNodes[i], k, results[i]); msg != "" {
			t.Fatalf("Performance audit failed (k=%d): %s", k, msg)
		}
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenDealCards(f *testing.F) {
	f.Add(int64(1), uint8(7), int16(3))
	f.Add(int64(2), uint8(0), int16(0))
	f.Add(int64(3), uint8(50), int16(-4))

	f.Fuzz(func(t *testing.T, seed int64, n uint8, k int16) {
		rng := rand.New(rand.NewSource(seed))
		vals := hiddenRandomVals(rng, int(n)%60, -5, 5)
		kk := int(k) % 25
		head, nodes := hiddenBuild(vals)

		got := DealCards(head, kk)
		if msg := hiddenCheckDeal(nodes, kk, got); msg != "" {
			t.Fatalf("Fuzz mismatch! vals=%v k=%d: %s", vals, kk, msg)
		}
	})
}
