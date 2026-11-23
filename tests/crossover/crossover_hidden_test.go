package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func hiddenBuildDLL(s string) (*ListNode, []*ListNode) {
	var head, tail *ListNode
	nodes := make([]*ListNode, 0, len(s))
	for i := 0; i < len(s); i++ {
		n := &ListNode{Val: s[i], Prev: tail}
		nodes = append(nodes, n)
		if tail == nil {
			head = n
		} else {
			tail.Next = n
		}
		tail = n
	}
	return head, nodes
}

func hiddenRandomString(rng *rand.Rand, n int, base byte) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = base + byte(rng.Intn(4))
	}
	return string(b)
}

func hiddenMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// checks that head is exactly the chain "want" (same original nodes, same order, correct Prev pointers)
func hiddenCheckChain(name string, head *ListNode, want []*ListNode) string {
	if len(want) == 0 {
		if head != nil {
			return name + " should be empty (nil)"
		}
		return ""
	}
	if head == nil {
		return fmt.Sprintf("%s is nil; want %d nodes", name, len(want))
	}
	if head.Prev != nil {
		return name + ": the Prev of the head must be nil"
	}
	var prev *ListNode
	count := 0
	for c := head; c != nil; c = c.Next {
		if count >= len(want) {
			return name + ": the list is longer than expected (or never ends)"
		}
		if c != want[count] {
			return fmt.Sprintf("%s: node %d is %q but should be the original node %q (nodes must be reused, in order)", name, count, c.Val, want[count].Val)
		}
		if c.Prev != prev {
			return fmt.Sprintf("%s: the Prev pointer of node %d is wrong", name, count)
		}
		prev = c
		count++
	}
	if count != len(want) {
		return fmt.Sprintf("%s has %d nodes; want %d", name, count, len(want))
	}
	return ""
}

func hiddenCheckCross(as, bs string, i, j int, nodesA, nodesB []*ListNode, gotA, gotB *ListNode) string {
	ia, jb := hiddenMin(i, len(as)), hiddenMin(j, len(bs))
	wantA := append(append([]*ListNode{}, nodesA[:ia]...), nodesB[jb:]...)
	wantB := append(append([]*ListNode{}, nodesB[:jb]...), nodesA[ia:]...)
	if msg := hiddenCheckChain("a", gotA, wantA); msg != "" {
		return msg
	}
	return hiddenCheckChain("b", gotB, wantB)
}

func hiddenCallLevel2WithTimeout(a, b *ListNode, i, j int, timeout time.Duration) (*ListNode, *ListNode, error) {
	type result struct {
		a, b     *ListNode
		panicked interface{}
	}
	ch := make(chan result, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{panicked: r}
			}
		}()
		x, y := Crossover(a, b, i, j)
		ch <- result{a: x, b: y}
	}()

	select {
	case res := <-ch:
		if res.panicked != nil {
			return nil, nil, fmt.Errorf("panic: %v", res.panicked)
		}
		return res.a, res.b, nil
	case <-time.After(timeout):
		return nil, nil, fmt.Errorf("execution timed out after %v", timeout)
	}
}

// Layer 1: Corner Cases
func TestHiddenLevel2_CornerCases(t *testing.T) {
	long := hiddenRandomString(rand.New(rand.NewSource(2)), 10000, 'a')

	tests := []struct {
		name string
		a, b string
		i, j int
	}{
		{"example_1", "AAAA", "bbbbb", 2, 1},
		{"example_2_a_becomes_empty", "ABC", "xyz", 0, 3},
		{"example_3_i_past_the_end", "ABC", "xy", 10, 1},
		{"both_nil", "", "", 0, 0},
		{"a_nil", "", "xy", 0, 1},
		{"b_nil", "Q", "", 1, 0},
		{"a_nil_with_big_cuts", "", "xyz", 5, 5},
		{"full_swap_zero_zero", "ABC", "xyz", 0, 0},
		{"cut_at_exact_end_of_both", "ABC", "xyz", 3, 3},
		{"cut_past_the_end_of_both", "ABC", "xyz", 9, 9},
		{"a_cut_at_end_b_cut_at_start", "ABC", "xyz", 3, 0},
		{"a_cut_at_start_b_cut_at_end", "ABC", "xyz", 0, 3},
		{"single_nodes_swap", "A", "b", 0, 0},
		{"single_nodes_keep", "A", "b", 1, 1},
		{"one_each_middle", "ABCD", "wxyz", 1, 3},
		{"huge_cut_values", "ABC", "xyz", 20000, 20000},
		{"unequal_lengths", "A", "bcdefghij", 1, 4},
		{"long_lists", long, long, 5000, 7000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, nodesA := hiddenBuildDLL(tc.a)
			b, nodesB := hiddenBuildDLL(tc.b)
			gotA, gotB, err := hiddenCallLevel2WithTimeout(a, b, tc.i, tc.j, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckCross(tc.a, tc.b, tc.i, tc.j, nodesA, nodesB, gotA, gotB); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel2_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(2002))

	for trial := 0; trial < 500; trial++ {
		as := hiddenRandomString(rng, rng.Intn(21), 'A')
		bs := hiddenRandomString(rng, rng.Intn(21), 'a')
		i, j := rng.Intn(26), rng.Intn(26)
		a, nodesA := hiddenBuildDLL(as)
		b, nodesB := hiddenBuildDLL(bs)

		gotA, gotB, err := hiddenCallLevel2WithTimeout(a, b, i, j, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckCross(as, bs, i, j, nodesA, nodesB, gotA, gotB); msg != "" {
			t.Fatalf("Property trial %d mismatch! a=%q b=%q i=%d j=%d: %s", trial, as, bs, i, j, msg)
		}
	}
}

// Layer 3: Performance Audit
func TestHiddenLevel2_PerformanceAudit(t *testing.T) {
	rng := rand.New(rand.NewSource(3))
	as := hiddenRandomString(rng, 100000, 'A')
	bs := hiddenRandomString(rng, 100000, 'a')
	type pair struct {
		a, b           *ListNode
		nodesA, nodesB []*ListNode
	}
	pairs := make([]pair, 3)
	for k := range pairs {
		a, na := hiddenBuildDLL(as)
		b, nb := hiddenBuildDLL(bs)
		pairs[k] = pair{a, b, na, nb}
	}
	cuts := [][2]int{{50000, 30000}, {0, 99999}, {100000, 0}}

	start := time.Now()
	gotA := make([]*ListNode, 3)
	gotB := make([]*ListNode, 3)
	for k := range pairs {
		var err error
		gotA[k], gotB[k], err = hiddenCallLevel2WithTimeout(pairs[k].a, pairs[k].b, cuts[k][0], cuts[k][1], 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed: %v", err)
		}
	}
	elapsed := time.Since(start)

	for k := range pairs {
		if msg := hiddenCheckCross(as, bs, cuts[k][0], cuts[k][1], pairs[k].nodesA, pairs[k].nodesB, gotA[k], gotB[k]); msg != "" {
			t.Fatalf("Performance audit failed (cut %v): %s", cuts[k], msg)
		}
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenCrossover(f *testing.F) {
	f.Add(int64(1), uint8(4), uint8(5), uint8(2), uint8(1))
	f.Add(int64(2), uint8(0), uint8(0), uint8(0), uint8(0))
	f.Add(int64(3), uint8(30), uint8(3), uint8(200), uint8(1))

	f.Fuzz(func(t *testing.T, seed int64, la, lb, ci, cj uint8) {
		rng := rand.New(rand.NewSource(seed))
		as := hiddenRandomString(rng, int(la)%30, 'A')
		bs := hiddenRandomString(rng, int(lb)%30, 'a')
		i, j := int(ci)%40, int(cj)%40
		a, nodesA := hiddenBuildDLL(as)
		b, nodesB := hiddenBuildDLL(bs)

		gotA, gotB := Crossover(a, b, i, j)
		if msg := hiddenCheckCross(as, bs, i, j, nodesA, nodesB, gotA, gotB); msg != "" {
			t.Fatalf("Fuzz mismatch! a=%q b=%q i=%d j=%d: %s", as, bs, i, j, msg)
		}
	})
}
