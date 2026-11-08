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

func hiddenCallLevel5WithTimeout(nums []int, timeout time.Duration) (*ListNode, error) {
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
		ch <- result{value: BuildLinkedList(nums)}
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

func hiddenCheckBuild(nums []int, got *ListNode) string {
	if len(nums) == 0 {
		if got != nil {
			return "an empty input must return nil"
		}
		return ""
	}
	gotVals, ok := hiddenToSlice(got)
	if !ok {
		return "the result list never ends (cycle?)"
	}
	if !hiddenEqual(gotVals, nums) {
		return fmt.Sprintf("got %v want %v", gotVals, nums)
	}
	seen := map[*ListNode]bool{}
	for curr := got; curr != nil; curr = curr.Next {
		if seen[curr] {
			return "the same node appears twice in the list"
		}
		seen[curr] = true
	}
	return ""
}

// Layer 1: Corner Cases
func TestHiddenLevel5_CornerCases(t *testing.T) {
	increasing := make([]int, 1000)
	for i := range increasing {
		increasing[i] = i
	}

	tests := []struct {
		name string
		nums []int
	}{
		{"nil_slice", nil},
		{"empty_slice", []int{}},
		{"single_value", []int{1}},
		{"single_zero", []int{0}},
		{"two_values", []int{1, 2}},
		{"negative_values", []int{-3, -2, -1}},
		{"duplicates", []int{5, 5, 5, 5}},
		{"unsorted", []int{9, 1, 8, 2, 7}},
		{"large_values", []int{1<<31 - 1, -(1 << 31)}},
		{"increasing_1000", increasing},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			original := append([]int{}, tc.nums...)
			got, err := hiddenCallLevel5WithTimeout(tc.nums, 1*time.Second)
			if err != nil {
				t.Fatalf("Hidden test %q failed: %v", tc.name, err)
			}
			if msg := hiddenCheckBuild(tc.nums, got); msg != "" {
				t.Fatalf("Hidden test %q mismatch: %s", tc.name, msg)
			}
			if !hiddenEqual(tc.nums, original) {
				t.Fatalf("Hidden test %q: the input slice must not be modified", tc.name)
			}
		})
	}
}

// Layer 2: Seeded Property Tests (500 trials under plain `go test ./...`)
func TestHiddenLevel5_SeededProperty(t *testing.T) {
	rng := rand.New(rand.NewSource(5005))

	for trial := 0; trial < 500; trial++ {
		nums := hiddenRandomVals(rng, rng.Intn(101), -1000, 1000)

		got, err := hiddenCallLevel5WithTimeout(nums, 500*time.Millisecond)
		if err != nil {
			t.Fatalf("Property trial %d failed: %v", trial, err)
		}
		if msg := hiddenCheckBuild(nums, got); msg != "" {
			t.Fatalf("Property trial %d mismatch! nums=%v: %s", trial, nums, msg)
		}
	}
}

// Layer 3: Performance Audit (an O(n^2) build that re-walks the list fails here)
func TestHiddenLevel5_PerformanceAudit(t *testing.T) {
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i
	}

	start := time.Now()
	got, err := hiddenCallLevel5WithTimeout(nums, 1*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("Performance audit failed: %v (expected an O(n) build)", err)
	}
	if msg := hiddenCheckBuild(nums, got); msg != "" {
		t.Fatalf("Performance audit failed: %s", msg)
	}
	if elapsed > 1*time.Second {
		t.Fatalf("Performance audit failed: took %v (expected < 1s)", elapsed)
	}
}

// Layer 4: Native Go Fuzzing
func FuzzHiddenBuildLinkedList(f *testing.F) {
	f.Add(int64(1), uint16(5))
	f.Add(int64(2), uint16(0))
	f.Add(int64(3), uint16(1500))

	f.Fuzz(func(t *testing.T, seed int64, n uint16) {
		rng := rand.New(rand.NewSource(seed))
		nums := hiddenRandomVals(rng, int(n)%2000, -100000, 100000)

		got := BuildLinkedList(nums)
		if msg := hiddenCheckBuild(nums, got); msg != "" {
			t.Fatalf("Fuzz mismatch! %d values: %s", len(nums), msg)
		}
	})
}
