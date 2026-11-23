package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// terms are {Coef, Exp} pairs in descending exponent order
func hiddenBuildPoly(terms [][2]int) (*ListNode, []*ListNode) {
	var head, tail *ListNode
	nodes := make([]*ListNode, 0, len(terms))
	for _, p := range terms {
		n := &ListNode{Coef: p[0], Exp: p[1]}
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

// random polynomial: each exponent in 0..maxExp is used with probability 1/2, coefficients are never 0
func hiddenRandomPoly(rng *rand.Rand, maxExp, maxCoef int) [][2]int {
	out := [][2]int{}
	for e := maxExp; e >= 0; e-- {
		if rng.Intn(2) == 0 {
			continue
		}
		c := rng.Intn(maxCoef) + 1
		if rng.Intn(2) == 0 {
			c = -c
		}
		out = append(out, [2]int{c, e})
	}
	return out
}

func hiddenEqualTerms(a, b [][2]int) bool {
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
func hiddenRefMul(a, b [][2]int) [][2]int {
	out := [][2]int{}
	if len(a) == 0 || len(b) == 0 {
		return out
	}
	acc := map[int]int{}
	for _, x := range a {
		for _, y := range b {
			acc[x[1]+y[1]] += x[0] * y[0]
		}
	}
	exps := []int{}
	for e, c := range acc {
		if c != 0 {
			exps = append(exps, e)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(exps)))
	for _, e := range exps {
		out = append(out, [2]int{acc[e], e})
	}
	return out
}

func hiddenCheckInputUnchanged(name string, terms [][2]int, nodes []*ListNode) string {
	for i, n := range nodes {
		var wantNext *ListNode
		if i+1 < len(nodes) {
			wantNext = nodes[i+1]
		}
		if n.Coef != terms[i][0] || n.Exp != terms[i][1] || n.Next != wantNext {
			return fmt.Sprintf("input %s was modified at term %d (the inputs must stay unchanged)", name, i)
		}
	}
	return ""
}

func hiddenCheckPoly(a, b [][2]int, nodesA, nodesB []*ListNode, got *ListNode) string {
	want := hiddenRefMul(a, b)
	inputs := map[*ListNode]bool{}
	for _, n := range nodesA {
		inputs[n] = true
	}
	for _, n := range nodesB {
		inputs[n] = true
	}
	gotTerms := [][2]int{}
	for c := got; c != nil; c = c.Next {
		if len(gotTerms) > len(want) {
			return "the result is longer than expected (or never ends)"
		}
		if inputs[c] {
			return "the result must be a NEW list; it reuses a node of an input"
		}
		gotTerms = append(gotTerms, [2]int{c.Coef, c.Exp})
	}
	if !hiddenEqualTerms(gotTerms, want) {
		return fmt.Sprintf("got %v want %v", gotTerms, want)
	}
	if msg := hiddenCheckInputUnchanged("a", a, nodesA); msg != "" {
		return msg
	}
	return hiddenCheckInputUnchanged("b", b, nodesB)
}

func hiddenCallLevel6WithTimeout(a, b *ListNode, timeout time.Duration) (*ListNode, error) {
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
		ch <- result{value: MultiplyPoly(a, b)}
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
func TestHiddenLevel6_CornerCases(t *testing.T) {
	dense := [][2]int{}
	for e := 9; e >= 0; e-- {
		dense = append(dense, [2]int{e + 1, e})
	}

	tests := []struct {
		name string
		a, b [][2]int
	}{
		{"example_1", [][2]int{{2, 2}, {3, 0}}, [][2]int{{1, 1}, {1, 0}}},
		{"example_2_terms_cancel", [][2]int{{1, 1}, {1, 0}}, [][2]int{{1, 1}, {-1, 0}}},
		{"example_3_nil_times_poly", nil, [][2]int{{5, 3}}},
		{"poly_times_nil", [][2]int{{5, 3}}, nil},
		{"nil_times_nil", nil, nil},
		{"constants", [][2]int{{7, 0}}, [][2]int{{-1, 0}}},
		{"times_the_constant_one", [][2]int{{3, 2}, {1, 0}}, [][2]int{{1, 0}}},
		{"one_times_poly", [][2]int{{1, 0}}, [][2]int{{3, 2}, {1, 0}}},
		{"constant_times_poly", [][2]int{{3, 0}}, [][2]int{{2, 4}, {-1, 2}, {5, 0}}},
		{"single_terms", [][2]int{{4, 3}}, [][2]int{{5, 2}}},
		{"big_exponents", [][2]int{{3, 5000}}, [][2]int{{-2, 5000}, {4, 0}}},
		{"middle_terms_cancel", [][2]int{{1, 2}, {1, 1}, {1, 0}}, [][2]int{{1, 1}, {-1, 0}}},
		{"difference_of_squares", [][2]int{{1, 1}, {1, 0}}, [][2]int{{1, 1}, {-1, 0}}},
		{"square_of_a_binomial", [][2]int{{1, 1}, {1, 0}}, [][2]int{{1, 1}, {1, 0}}},
		{"all_negative_coefficients", [][2]int{{-2, 3}, {-3, 1}}, [][2]int{{-4, 2}, {-1, 0}}},
		{"maximum_coefficients", [][2]int{{100, 2}, {-100, 0}}, [][2]int{{100, 2}, {100, 0}}},
		{"dense_times_dense", dense, dense},
		{"sparse_times_dense", [][2]int{{1, 5000}, {-1, 0}}, dense},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headA, nodesA := hiddenBuildPoly(tc.a)
			headB, nodesB := hiddenBuildPoly(tc.b)
			got, err := hiddenCallLevel6WithTimeout(headA, headB, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckPoly(tc.a, tc.b, nodesA, nodesB, got); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel6_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(6006))

	for trial := 0; trial < 500; trial++ {
		a := hiddenRandomPoly(rng, rng.Intn(13), 3)
		b := hiddenRandomPoly(rng, rng.Intn(13), 3)
		headA, nodesA := hiddenBuildPoly(a)
		headB, nodesB := hiddenBuildPoly(b)

		got, err := hiddenCallLevel6WithTimeout(headA, headB, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckPoly(a, b, nodesA, nodesB, got); msg != "" {
			t.Fatalf("Property trial %d mismatch! a=%v b=%v: %s", trial, a, b, msg)
		}
	}
}

// Layer 3: Performance Audit (inserting every product term into a sorted list fails here)
func TestHiddenLevel6_PerformanceAudit(t *testing.T) {
	a := [][2]int{}
	b := [][2]int{}
	for i := 199; i >= 0; i-- {
		ca, cb := i%100+1, (i*7)%100+1
		if i%2 == 1 {
			ca = -ca
		}
		if i%3 == 0 {
			cb = -cb
		}
		a = append(a, [2]int{ca, 25 * i})
		b = append(b, [2]int{cb, 24 * i})
	}
	headA, nodesA := hiddenBuildPoly(a)
	headB, nodesB := hiddenBuildPoly(b)

	var got *ListNode
	start := time.Now()
	for k := 0; k < 30; k++ {
		var err error
		got, err = hiddenCallLevel6WithTimeout(headA, headB, 1*time.Second)
		if err != nil {
			t.Fatalf("Performance audit failed: %v", err)
		}
	}
	elapsed := time.Since(start)

	if msg := hiddenCheckPoly(a, b, nodesA, nodesB, got); msg != "" {
		t.Fatalf("Performance audit failed: %s", msg)
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenMultiplyPoly(f *testing.F) {
	f.Add(int64(1), uint8(5), uint8(4))
	f.Add(int64(2), uint8(0), uint8(0))
	f.Add(int64(3), uint8(20), uint8(20))

	f.Fuzz(func(t *testing.T, seed int64, ea, eb uint8) {
		rng := rand.New(rand.NewSource(seed))
		a := hiddenRandomPoly(rng, int(ea)%25, 100)
		b := hiddenRandomPoly(rng, int(eb)%25, 100)
		headA, nodesA := hiddenBuildPoly(a)
		headB, nodesB := hiddenBuildPoly(b)

		got := MultiplyPoly(headA, headB)
		if msg := hiddenCheckPoly(a, b, nodesA, nodesB, got); msg != "" {
			t.Fatalf("Fuzz mismatch! a=%v b=%v: %s", a, b, msg)
		}
	})
}
