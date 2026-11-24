# Checkpoint: 22 Problem-Solving Questions (Go)

Prepared for the next Nextera Education coding checkpoint. All questions are written for Go. Each question has a title, a statement, the input and output, constraints, a Go signature, worked examples, and instructions. Each is followed by an **Expected answer** section for mentors with the approach, complexity, a verified Go reference solution, and the expected output of every example.

> **Mentor note:** The *Expected answer* sections are for mentors only. Remove them before you hand the paper to students.

## Table of contents

- [Are these linked-list questions?](#are-these-linked-list-questions)
- [Difficulty order at a glance](#difficulty-order-at-a-glance)
- [Shared conventions](#shared-conventions)
- [Tier 1: Warm-up](#tier-1-warm-up)
  - [Q1. Design Browser History](#q1-design-browser-history)
  - [Q2. Minimum Loss](#q2-minimum-loss)
  - [Q3. Journey to the Moon](#q3-journey-to-the-moon)
- [Tier 2: Standard medium](#tier-2-standard-medium)
  - [Q4. Copy List with Random Pointer](#q4-copy-list-with-random-pointer)
  - [Q5. Queries with Fixed Length](#q5-queries-with-fixed-length)
  - [Q6. Flatten a Multilevel Doubly Linked List](#q6-flatten-a-multilevel-doubly-linked-list)
  - [Q7. LRU Cache](#q7-lru-cache)
  - [Q8. Min-Max Queue](#q8-min-max-queue)
  - [Q9. Debug and Extend: Recent Files Menu](#q9-debug-and-extend-recent-files-menu)
- [Tier 3: Upper medium](#tier-3-upper-medium)
  - [Q10. Multiply Sparse Polynomials](#q10-multiply-sparse-polynomials)
  - [Q11. Flatten by Levels](#q11-flatten-by-levels)
  - [Q12. Reversing Circle](#q12-reversing-circle)
  - [Q13. Queue (Codeforces 91B)](#q13-queue-codeforces-91b)
  - [Q14. Commit History](#q14-commit-history)
  - [Q15. Abbreviation](#q15-abbreviation)
  - [Q16. Mysterious Present (Codeforces 4D)](#q16-mysterious-present-codeforces-4d)
  - [Q17. Poisonous Plants](#q17-poisonous-plants)
- [Tier 4: Hard](#tier-4-hard)
  - [Q18. Longest Regular Bracket Sequence (Codeforces 5C)](#q18-longest-regular-bracket-sequence-codeforces-5c)
  - [Q19. Self-Organizing List Cost](#q19-self-organizing-list-cost)
  - [Q20. Expiring LRU Cache](#q20-expiring-lru-cache)
  - [Q21. Sereja and Brackets (Codeforces 380C)](#q21-sereja-and-brackets-codeforces-380c)
  - [Q22. Circular Range Min with Range Add (Codeforces 52C)](#q22-circular-range-min-with-range-add-codeforces-52c)
- [Verification notes](#verification-notes)

---

## Are these linked-list questions?

**Not all of them.** Only 7 of the 22 are true linked-list questions. In 3 more, a linked list is optional or only used by the naive solution. The other 12 are array, string, graph, tree, or design problems, and they need no linked list. Commit History stores parent pointers, but it is a tree problem (lowest common ancestor), not a list problem.

| Group | Transcript numbers (22-question list) | Questions |
|---|---|---|
| Truly linked list | 5, 6, 7, 9, 10, 11, 20 | Copy List with Random Pointer, Flatten a Multilevel Doubly Linked List, Debug and Extend: Recent Files Menu, LRU Cache, Multiply Sparse Polynomials, Flatten by Levels, Expiring LRU Cache |
| Optional or naive only | 3, 16, 19 | Design Browser History, Reversing Circle, Self-Organizing List Cost |
| Not linked list | 1, 2, 4, 8, 12, 13, 14, 15, 17, 18, 21, 22 | Minimum Loss, Journey to the Moon, Queries with Fixed Length, Min-Max Queue, Abbreviation, Mysterious Present, Queue (CF 91B), Commit History (parent-pointer tree), Poisonous Plants, Longest Regular Bracket Sequence, Sereja and Brackets, Circular Range Min with Range Add |

**Compact table** (listed in the transcript's order; *New #* is the position in this paper)

| Transcript # | New # | Question | Main technique | Linked-list question? |
|---|---|---|---|---|
| 1 | Q2 | Minimum Loss | Sort by price keeping indices, compare neighbours | No |
| 2 | Q3 | Journey to the Moon | Union-find, pair counting in int64 | No |
| 3 | Q1 | Design Browser History | Slice + current index (doubly linked list also works) | Optional |
| 4 | Q5 | Queries with Fixed Length | Monotonic deque (sliding-window max) | No |
| 5 | Q4 | Copy List with Random Pointer | Singly list with a second pointer; hash map or interleaving | **Yes** (singly) |
| 6 | Q6 | Flatten a Multilevel Doubly Linked List | Doubly list with child pointers; DFS splicing | **Yes** (doubly) |
| 7 | Q9 | Debug and Extend: Recent Files Menu | Doubly list with sentinels + hash map; second list for pins | **Yes** (doubly) |
| 8 | Q8 | Min-Max Queue | Queue + two monotonic deques | No |
| 9 | Q7 | LRU Cache | Doubly list + hash map | **Yes** (doubly) |
| 10 | Q10 | Multiply Sparse Polynomials | Singly list of terms; accumulate by exponent | **Yes** (singly) |
| 11 | Q11 | Flatten by Levels | Doubly list with child pointers; BFS using the list as its own queue | **Yes** (doubly) |
| 12 | Q15 | Abbreviation | 2D DP reduced to 1D | No |
| 13 | Q16 | Mysterious Present | Sort + LIS with binary search and parent array | No |
| 14 | Q13 | Queue (CF 91B) | Suffix minimum + binary search | No |
| 15 | Q14 | Commit History | Parent-pointer tree, hash map of branches, LCA / binary lifting | No (tree) |
| 16 | Q12 | Reversing Circle | Josephus-style simulation with k mod size | Optional |
| 17 | Q17 | Poisonous Plants | Monotonic stack tracking death days | No |
| 18 | Q18 | Longest Regular Bracket Sequence | Stack of indices (or DP) | No |
| 19 | Q19 | Self-Organizing List Cost | Naive: move-to-front list. Optimal: Fenwick tree over timestamps | Naive only |
| 20 | Q20 | Expiring LRU Cache | Doubly list + hash map + min-heap with lazy deletion | **Yes** (doubly) |
| 21 | Q21 | Sereja and Brackets | Segment tree merging (pairs, open, close) | No |
| 22 | Q22 | Circular Range Min with Range Add | Lazy segment tree; wrapped ranges split in two | No |

---

## Difficulty order at a glance

This order starts from the transcript's four tiers. The ranking is based on the main task. A few questions move up when their follow-up is required, and the notes below say which.

| # | Question | Tier | Linked list? |
|---|---|---|---|
| 1 | Design Browser History | 1 Warm-up | Optional |
| 2 | Minimum Loss | 1 Warm-up | No |
| 3 | Journey to the Moon | 1 Warm-up | No |
| 4 | Copy List with Random Pointer | 2 Standard medium | Yes |
| 5 | Queries with Fixed Length | 2 Standard medium | No |
| 6 | Flatten a Multilevel Doubly Linked List | 2 Standard medium | Yes |
| 7 | LRU Cache | 2 Standard medium | Yes |
| 8 | Min-Max Queue | 2 Standard medium | No |
| 9 | Debug and Extend: Recent Files Menu | 2 Standard medium | Yes |
| 10 | Multiply Sparse Polynomials | 3 Upper medium | Yes |
| 11 | Flatten by Levels | 3 Upper medium | Yes |
| 12 | Reversing Circle | 3 Upper medium | Optional |
| 13 | Queue (CF 91B) | 3 Upper medium | No |
| 14 | Commit History | 3 Upper medium | No (tree) |
| 15 | Abbreviation | 3 Upper medium | No |
| 16 | Mysterious Present (CF 4D) | 3 Upper medium | No |
| 17 | Poisonous Plants | 3 Upper medium | No |
| 18 | Longest Regular Bracket Sequence (CF 5C) | 4 Hard | No |
| 19 | Self-Organizing List Cost | 4 Hard | Naive only |
| 20 | Expiring LRU Cache | 4 Hard | Yes |
| 21 | Sereja and Brackets (CF 380C) | 4 Hard | No |
| 22 | Circular Range Min with Range Add (CF 52C) | 4 Hard | No |

**Changes from the transcript's order, and why**

- Design Browser History now opens the paper. It is the easiest problem in the set because it needs only a slice and an index.
- Copy List with Random Pointer comes before Queries with Fixed Length. The hash-map version is a short, standard exercise. A monotonic deque is a new idea for most students.
- LRU Cache now comes before Debug and Extend. The Recent Files Menu is an LRU with bugs plus a pin feature, so students meet the clean version first.
- Reversing Circle, Queue (CF 91B) and Commit History come before Abbreviation, Mysterious Present and Poisonous Plants. The first three are simulation, binary search and tree walks. The last three need DP, LIS with reconstruction, or a non-obvious monotonic-stack invariant.
- Self-Organizing List Cost stays in Tier 4 only because its O((n + q) log n) target is required here. Without that requirement, the naive linked-list version is a Tier 1 problem.

---

## Shared conventions

- Unless a question says otherwise, students implement the given function or type in `package main`. Mentors test it by calling the function directly, so stdin and stdout are not used.
- "Input" describes the arguments, and "Output" describes the return value.
- Go's `int` is 64 bits on all checkpoint platforms (amd64 and arm64). Where a value can exceed about 2·10^9, the question says so, and the signature uses `int64` where that makes the intent clearer.
- Each question lists any **Skills needed beyond linked lists**, which shows what else students must know.

---
# Tier 1: Warm-up

## Q1. Design Browser History

*Source: LeetCode 1472 (Q11 in the original set). Linked list: optional.*

**Skills needed beyond linked lists:** Dynamic arrays (slices) with a cursor index, and clamping with `min` and `max`. A doubly linked list is an acceptable alternative, but students don't need one.

**Statement**
Build a type that simulates the history of one browser tab.

- `NewBrowserHistory(homepage)` starts on `homepage`.
- `Visit(url)` goes to `url` from the current page and clears all forward history.
- `Back(steps)` moves back at most `steps` pages, stopping at the oldest page, and returns the current URL.
- `Forward(steps)` moves forward at most `steps` pages, stopping at the newest page, and returns the current URL.

**Input / Output**
- Input: a sequence of method calls on one `*BrowserHistory`.
- Output: `Back` and `Forward` return the URL of the current page after the move.

**Constraints**
- `1 ≤ len(url), len(homepage) ≤ 20`, lowercase letters and `.`
- `1 ≤ steps ≤ 100`
- At most 5000 calls in total.

**Required signature (Go)**
```go
type BrowserHistory struct{ /* your fields */ }

func NewBrowserHistory(homepage string) *BrowserHistory
func (b *BrowserHistory) Visit(url string)
func (b *BrowserHistory) Back(steps int) string
func (b *BrowserHistory) Forward(steps int) string
```

**Example**
```
b := NewBrowserHistory("leetcode.com")
b.Visit("google.com"); b.Visit("facebook.com"); b.Visit("youtube.com")
b.Back(1)          → "facebook.com"
b.Back(1)          → "google.com"
b.Forward(1)       → "facebook.com"
b.Visit("linkedin.com")            // forward history ("youtube.com") is cleared
b.Forward(2)       → "linkedin.com"
b.Back(2)          → "google.com"
b.Back(7)          → "leetcode.com"
```

**Instructions**
- Each operation should run in O(1) time, amortized for `Visit`, or in O(steps) time.
- You may use a slice or a doubly linked list. Explain your choice in one sentence.

### Expected answer — Q1

**Approach.** Store the pages in a slice. `cur` is the index of the page on screen, and `last` is the index of the newest page that can still be reached. `Visit` moves `cur` forward, then overwrites or appends the URL and sets `last = cur`. It doesn't delete anything: entries after `last` are stale and get overwritten later. `Back` and `Forward` clamp `cur` to `[0, last]`.

**Complexity.** Each call is O(1) time, amortized for `Visit`. Space is O(number of visits).

**Reference solution (Go)**

```go
package main

// BrowserHistory keeps every page in a slice. cur is the page on screen and
// last is the newest page that is still reachable with Forward.
type BrowserHistory struct {
	pages []string
	cur   int
	last  int
}

func NewBrowserHistory(homepage string) *BrowserHistory {
	return &BrowserHistory{pages: []string{homepage}}
}

// Visit opens url and discards all forward history in O(1) amortized time.
func (b *BrowserHistory) Visit(url string) {
	b.cur++
	if b.cur < len(b.pages) {
		b.pages[b.cur] = url // overwrite a stale forward entry
	} else {
		b.pages = append(b.pages, url)
	}
	b.last = b.cur // anything after cur is no longer reachable
}

func (b *BrowserHistory) Back(steps int) string {
	b.cur = max(0, b.cur-steps)
	return b.pages[b.cur]
}

func (b *BrowserHistory) Forward(steps int) string {
	b.cur = min(b.last, b.cur+steps)
	return b.pages[b.cur]
}
```

**Expected output of the example:** `facebook.com`, `google.com`, `facebook.com`, `linkedin.com`, `google.com`, `leetcode.com`.

---

## Q2. Minimum Loss

*Source: HackerRank, "Minimum Loss". Linked list: no.*

**Skills needed beyond linked lists:** Sorting while keeping original indices, an argument about why only adjacent values matter, and `int64` (prices go up to 10^16).

**Statement**
You know the price of a house in each of `n` years, and all prices are different. You must buy the house in one year and sell it in a **later** year for a **strictly lower** price. Return the smallest possible loss (buy price − sell price).

**Input / Output**
- Input: `price`, where `price[i]` is the price in year `i`.
- Output: the minimum loss as `int64`.

**Constraints**
- `2 ≤ n ≤ 2·10^5`
- `1 ≤ price[i] ≤ 10^16`, all distinct
- At least one valid loss always exists.

**Required signature (Go)**
```go
func MinLoss(price []int64) int64
```

**Examples**
```
MinLoss([]int64{20, 7, 8, 2, 5})  → 2   // buy at 7 (year 1), sell at 5 (year 4)
MinLoss([]int64{5, 10, 3})        → 2   // buy at 5, sell at 3
MinLoss([]int64{3, 2, 1})         → 1
```

**Instructions**
- Target O(n log n). An O(n²) double loop is too slow at the maximum size.
- Hint: sort by price and compare neighbours, or use an ordered structure.

### Expected answer — Q2

**Approach.** Sort the year indices by price. The best pair is always next to each other in sorted order. Suppose a better pair (H, then L later) had some price M between them. If M's year is after H's year, selling at M gives a smaller loss. If M's year is before H's year, it is also before L's year, so buying at M gives a smaller loss. Either way, the pair (H, L) wasn't optimal. So for each adjacent pair (`low`, `high`) in sorted order, it is a valid sale when `high`'s year comes before `low`'s year. Take the minimum of those differences.

**Complexity.** O(n log n) time and O(n) space.

**Reference solution (Go)**

```go
package main

import "sort"

// MinLoss returns the smallest price[i]-price[j] with i < j and price[i] > price[j].
// The best pair is always adjacent in sorted order, so sort the years by price
// and check only neighbours whose years are in the right order.
func MinLoss(price []int64) int64 {
	idx := make([]int, len(price))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(x, y int) bool { return price[idx[x]] < price[idx[y]] })

	best := int64(-1)
	for k := 0; k+1 < len(idx); k++ {
		low, high := idx[k], idx[k+1] // price[low] < price[high]
		if high < low {               // buy the higher price first, sell later
			loss := price[high] - price[low]
			if best == -1 || loss < best {
				best = loss
			}
		}
	}
	return best
}
```

**Expected output:** `2`, `2`, `1`.

---

## Q3. Journey to the Moon

*Source: HackerRank, "Journey to the Moon". Linked list: no.*

**Skills needed beyond linked lists:** Union-find (disjoint set union) with path compression and union by size, connected components, counting pairs in a way that avoids O(k²), and `int64` overflow (n = 10^5 gives about 5·10^9 pairs).

**Statement**
`n` astronauts are labeled `0 … n-1`. Each given pair `(a, b)` means that astronauts `a` and `b` are from the same country. This relation is transitive. Count the ways to choose two astronauts from **different** countries.

**Input / Output**
- Input: `n` and the list of same-country pairs.
- Output: the number of cross-country pairs as `int64`.

**Constraints**
- `1 ≤ n ≤ 10^5`
- `0 ≤ len(same) ≤ 10^6`, `0 ≤ a, b < n`

**Required signature (Go)**
```go
func CrossCountryPairs(n int, same [][2]int) int64
```

**Examples**
```
CrossCountryPairs(5, [][2]int{{0,1},{2,3},{0,4}})  → 6   // countries {0,1,4} and {2,3}: 3·2
CrossCountryPairs(4, nil)                          → 6   // four countries of size 1: C(4,2)
CrossCountryPairs(4, [][2]int{{0,1},{1,2},{2,3}})  → 0   // one country
```

**Instructions**
- Target near-linear time, O((n + p) · α(n)).
- The answer must not overflow. Explain why `int32` would fail.

### Expected answer — Q3

**Approach.** Union all pairs with union-find, then read off the size of each component. Go through the components in any order. Each new component of size `s` forms `s · seen` pairs with the astronauts already counted, and then `seen += s`. Alternatively, compute `C(n,2) − Σ C(s,2)`.

**Complexity.** O((n + p) · α(n)) time and O(n) space.

**Reference solution (Go)**

```go
package main

// CrossCountryPairs counts pairs of astronauts from different countries.
func CrossCountryPairs(n int, same [][2]int) int64 {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]] // path halving
			x = parent[x]
		}
		return x
	}
	for _, p := range same {
		a, b := find(p[0]), find(p[1])
		if a == b {
			continue
		}
		if size[a] < size[b] {
			a, b = b, a
		}
		parent[b] = a // union by size
		size[a] += size[b]
	}

	// Each new country pairs with everyone counted so far.
	var total, seen int64
	for i := 0; i < n; i++ {
		if find(i) == i {
			s := int64(size[i])
			total += s * seen
			seen += s
		}
	}
	return total
}
```

**Expected output:** `6`, `6`, `0`. For n = 100000 with no pairs, the answer is 4,999,950,000, which needs 64 bits.

---

# Tier 2: Standard medium

## Q4. Copy List with Random Pointer

*Source: LeetCode 138 (Q13 in the original set). Linked list: **yes** (singly, with a second pointer).*

**Skills needed beyond linked lists:** A hash map from original node to copy (pointer identity as the key), or the O(1)-space interleaving trick. It also helps to understand that equal values do not mean equal nodes.

**Statement**
Each node has a `Next` pointer and a `Random` pointer. `Random` may point to any node in the list, or be `nil`. Build a **deep copy**: `n` brand-new nodes with the same values, whose `Next` and `Random` pointers form the same structure but point only to nodes in the copy, never to the original. The original list must be unchanged when the function returns.

**Input / Output**
- Input: the head of the original list. In the examples, each node is written as `[value, index of Random target or null]`.
- Output: the head of the copy.

**Constraints**
- `0 ≤ n ≤ 1000`
- `-10^4 ≤ Val ≤ 10^4`
- `Random` is `nil` or points to a node in the same list.

**Required signature (Go)**
```go
type RNode struct {
	Val    int
	Next   *RNode
	Random *RNode
}

func copyRandomList(head *RNode) *RNode
```

**Examples**
```
[[7,null],[13,0],[11,4],[10,2],[1,0]]  → a copy with the same shape: [[7,null],[13,0],[11,4],[10,2],[1,0]]
[[1,1],[2,1]]                          → [[1,1],[2,1]]
[[3,null],[3,0],[3,null]]              → [[3,null],[3,0],[3,null]]   (equal values, so nodes can't be matched by value)
[] (nil)                               → nil
```

**Instructions**
- Main task: O(n) time. A hash map is fine.
- Follow-up: O(1) extra space, not counting the copy itself. Restore the original list.
- No node of the copy may be a node of the original.

### Expected answer — Q4

**Approach.** *Hash-map version:* in one pass, map each original node to a new node. In a second pass, set `copy.Next = m[orig.Next]` and `copy.Random = m[orig.Random]`.
*Follow-up version (reference):* (1) Insert each copy right after its original: `A → A' → B → B' …`. (2) Set `X'.Random = X.Random.Next`. (3) Split the two lists apart, which also restores the original `Next` pointers.

**Complexity.** O(n) time. The map version uses O(n) extra space, and the interleaving version uses O(1).

**Reference solution (Go)**

```go
package main

type RNode struct {
	Val    int
	Next   *RNode
	Random *RNode
}

// copyRandomList deep-copies the list in O(n) time and O(1) extra space
// by interleaving each copy right after its original.
func copyRandomList(head *RNode) *RNode {
	if head == nil {
		return nil
	}
	// 1. A -> A' -> B -> B' -> ...
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &RNode{Val: cur.Val, Next: cur.Next}
	}
	// 2. The copy of X.Random is X.Random.Next.
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	// 3. Separate the two lists and restore the original.
	newHead := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		cp := cur.Next
		cur.Next = cp.Next
		if cp.Next != nil {
			cp.Next = cp.Next.Next
		}
	}
	return newHead
}
```

**Expected output:** Each copy serializes to exactly the input shape. The original serializes the same before and after the call, and the copy shares no nodes with the original. For example 4, the result is `nil`.

---

## Q5. Queries with Fixed Length

*Source: HackerRank, "Queries with Fixed Length". Linked list: no.*

**Skills needed beyond linked lists:** Sliding windows and a monotonic deque of indices (sliding-window maximum).

**Statement**
You are given an array `a` and several queries. For a query `d`, look at every window of `d` consecutive elements and take the maximum of each window. The answer is the **minimum** of those maxima.

**Input / Output**
- Input: the array `a` and the list of queries.
- Output: one answer per query, in order.

**Constraints**
- `1 ≤ n ≤ 10^5`, `0 ≤ a[i] < 10^6`
- `1 ≤ len(queries) ≤ 100`, `1 ≤ d ≤ n`

**Required signature (Go)**
```go
func FixedLength(a []int, queries []int) []int
```

**Examples**
```
FixedLength([]int{2,3,4,5,6}, []int{2,3})              → [3 4]
FixedLength([]int{5,5,5}, []int{1,3})                  → [5 5]
FixedLength([]int{33,11,44,11,55}, []int{1,2,3,4,5})   → [11 33 44 44 55]
```

**Instructions**
- Each query must run in O(n), for O(n · q) in total. The naive O(n · d) per query is too slow.
- Use a monotonic deque. Say what invariant it maintains.

### Expected answer — Q5

**Approach.** For each query, slide a window over the array and keep a deque of indices whose values are strictly decreasing from front to back. Before pushing `i`, pop from the back every index with value ≤ `a[i]`. Drop the front once it leaves the window. Once the window is full, `a[dq[0]]` is its maximum. Track the smallest of these maxima.

**Complexity.** O(n) per query, because each index is pushed and popped at most once. That gives O(n · q) time in total and O(n) space.

**Reference solution (Go)**

```go
package main

// FixedLength answers each query d with min over all windows of size d of the
// window maximum. Each query is O(n) with a monotonic deque of indices whose
// values decrease from front to back.
func FixedLength(a []int, queries []int) []int {
	res := make([]int, 0, len(queries))
	for _, d := range queries {
		dq := make([]int, 0, len(a)) // indices; dq[0] is the current max
		best := -1
		for i, x := range a {
			for len(dq) > 0 && a[dq[len(dq)-1]] <= x {
				dq = dq[:len(dq)-1]
			}
			dq = append(dq, i)
			if dq[0] <= i-d { // front fell out of the window
				dq = dq[1:]
			}
			if i >= d-1 {
				if m := a[dq[0]]; best == -1 || m < best {
					best = m
				}
			}
		}
		res = append(res, best)
	}
	return res
}
```

**Expected output:** `[3 4]`, `[5 5]`, `[11 33 44 44 55]`.

---

## Q6. Flatten a Multilevel Doubly Linked List

*Source: LeetCode 430 (Q14 in the original set). Linked list: **yes** (doubly, with child pointers).*

**Skills needed beyond linked lists:** Depth-first traversal, with recursion or an explicit stack. No other data structure is required.

**Statement**
Each node has `Prev`, `Next` and `Child`. `Child` points to the head of another doubly linked list, which can have children of its own. Flatten everything into a single-level doubly linked list. A node's child list goes **right after that node and before its old `Next`** (depth-first order). Set every `Child` to `nil` and keep every `Prev` pointer correct.

**Input / Output**
- Input: the head of the top level. The head of each child list has `Prev == nil`.
- Output: the head of the flattened list.

**Constraints**
- `0 ≤` total nodes `≤ 1000`
- `1 ≤ Val ≤ 10^5`

**Required signature (Go)**
```go
type MNode struct {
	Val   int
	Prev  *MNode
	Next  *MNode
	Child *MNode
}

func flatten(head *MNode) *MNode
```

**Examples**
```
Level 1: 1-2-3-4-5-6; node 3 has child list 7-8-9-10; node 8 has child list 11-12
  → 1 2 3 7 8 11 12 9 10 4 5 6
Level 1: 1-2; node 1 has child 3
  → 1 3 2
nil → nil
```

**Instructions**
- O(n) time.
- After flattening, walking the `Prev` pointers must give the reverse order, and every `Child` must be `nil`.

### Expected answer — Q6

**Approach.** Walk forward through the list. When `cur` has a child, find the tail of the child list (one level only), then splice: `cur ⇄ child … tail ⇄ oldNext`, and set `cur.Child = nil`. The walker then enters the child list, so any deeper children are spliced when it reaches them. A recursive DFS that returns the tail of each flattened sub-list is also fine.

**Complexity.** O(n) time, because each node is walked at most twice: once as part of a tail search and once by the main walker. Space is O(1) extra, or O(depth) for recursion.

**Reference solution (Go)**

```go
package main

type MNode struct {
	Val   int
	Prev  *MNode
	Next  *MNode
	Child *MNode
}

// flatten splices every child list in right after its parent (depth-first
// order). Walking forward means nested children are handled when reached.
// O(n) time, O(1) extra space.
func flatten(head *MNode) *MNode {
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Child == nil {
			continue
		}
		child := cur.Child
		tail := child
		for tail.Next != nil {
			tail = tail.Next
		}
		// cur <-> child ... tail <-> oldNext
		tail.Next = cur.Next
		if cur.Next != nil {
			cur.Next.Prev = tail
		}
		cur.Next = child
		child.Prev = cur
		cur.Child = nil
	}
	return head
}
```

**Expected output:** `1 2 3 7 8 11 12 9 10 4 5 6`, `1 3 2`, `nil`. The reference was also checked for correct `Prev` links and cleared `Child` pointers.

---

## Q7. LRU Cache

*Source: LeetCode 146 (Q15 in the original set). Linked list: **yes** (doubly, plus a hash map).*

**Skills needed beyond linked lists:** Hash maps that store node pointers, and sentinel (dummy) nodes.

**Statement**
Design a cache with a fixed capacity. When it is full, it evicts the **least recently used** entry.

- `NewLRUCache(capacity)` creates the cache.
- `Get(key)` returns the value, or `-1` if the key is absent. A hit counts as a use.
- `Put(key, value)` updates the key if it exists and inserts it otherwise. If an insert would exceed the capacity, it first evicts the least recently used key. `Put` counts as a use.

**Input / Output**
- Input: a sequence of `Get` and `Put` calls.
- Output: `Get` returns the value or `-1`.

**Constraints**
- `1 ≤ capacity ≤ 3000`, `0 ≤ key ≤ 10^4`, `0 ≤ value ≤ 10^5`
- Up to 2·10^5 calls.

**Required signature (Go)**
```go
type LRUCache struct{ /* your fields */ }

func NewLRUCache(capacity int) *LRUCache
func (c *LRUCache) Get(key int) int
func (c *LRUCache) Put(key, value int)
```

**Example**
```
c := NewLRUCache(2)
c.Put(1, 1); c.Put(2, 2)
c.Get(1)    → 1
c.Put(3, 3)            // evicts key 2
c.Get(2)    → -1
c.Put(4, 4)            // evicts key 1
c.Get(1)    → -1
c.Get(3)    → 3
c.Get(4)    → 4
```

**Instructions**
- `Get` and `Put` must both be O(1) on average.
- Implement the list yourself. Do not use `container/list`.

### Expected answer — Q7

**Approach.** Use a map from key to node and a doubly linked list with head and tail sentinels, where `head.next` is the most recent entry. A use unlinks the node and pushes it to the front. Eviction removes `tail.prev` from both the list and the map.

**Complexity.** O(1) average per operation and O(capacity) space.

**Reference solution (Go)**

```go
package main

type lruNode struct {
	key, val   int
	prev, next *lruNode
}

// LRUCache: a hash map finds a node in O(1); a doubly linked list with
// sentinels keeps recency order (head.next = most recent).
type LRUCache struct {
	cap        int
	items      map[int]*lruNode
	head, tail *lruNode
}

func NewLRUCache(capacity int) *LRUCache {
	h, t := &lruNode{}, &lruNode{}
	h.next, t.prev = t, h
	return &LRUCache{cap: capacity, items: make(map[int]*lruNode), head: h, tail: t}
}

func (c *LRUCache) unlink(n *lruNode) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) pushFront(n *lruNode) {
	n.prev, n.next = c.head, c.head.next
	c.head.next.prev = n
	c.head.next = n
}

func (c *LRUCache) Get(key int) int {
	n, ok := c.items[key]
	if !ok {
		return -1
	}
	c.unlink(n)
	c.pushFront(n)
	return n.val
}

func (c *LRUCache) Put(key, value int) {
	if n, ok := c.items[key]; ok {
		n.val = value
		c.unlink(n)
		c.pushFront(n)
		return
	}
	if len(c.items) == c.cap {
		lru := c.tail.prev
		c.unlink(lru)
		delete(c.items, lru.key)
	}
	n := &lruNode{key: key, val: value}
	c.items[key] = n
	c.pushFront(n)
}
```

**Expected output:** `1`, `-1`, `-1`, `3`, `4`.

---

## Q8. Min-Max Queue

*Source: original question from the transcript's set. Linked list: no.*

**Skills needed beyond linked lists:** A FIFO queue, monotonic deques (one for the minimum, one for the maximum), handling duplicate values, and amortized analysis. Optionally, a queue built from two stacks.

**Statement**
Design a FIFO queue that can also report its smallest and largest current elements.

- `Push(x)` adds `x` to the back.
- `Pop()` removes and returns the front element.
- `Min()` and `Max()` return the smallest and largest elements currently in the queue.

`Pop`, `Min` and `Max` are never called on an empty queue.

**Input / Output**
- Input: a sequence of calls.
- Output: the return values of `Pop`, `Min` and `Max`.

**Constraints**
- Up to 10^6 calls, `-10^9 ≤ x ≤ 10^9`.

**Required signature (Go)**
```go
type MinMaxQueue struct{ /* your fields */ }

func NewMinMaxQueue() *MinMaxQueue
func (m *MinMaxQueue) Push(x int)
func (m *MinMaxQueue) Pop() int
func (m *MinMaxQueue) Min() int
func (m *MinMaxQueue) Max() int
```

**Example**
```
Push(3), Push(1), Push(5)
Min() → 1    Max() → 5
Pop() → 3    Min() → 1    Max() → 5
Pop() → 1    Min() → 5    Max() → 5
Push(2)      Min() → 2    Max() → 5
Pop() → 5    Min() → 2    Max() → 2
```

**Instructions**
- Every operation must be amortized O(1). Scanning the queue in `Min` or `Max` doesn't count.
- Duplicates must work, for example `Push(1), Push(1), Pop()` followed by `Min()`.
- Follow-up: build it using only two stacks, where each stack tracks its own min and max.

### Expected answer — Q8

**Approach.** Keep the queue itself plus two monotonic deques. `mins` is non-decreasing, so its front is the minimum. `maxs` is non-increasing, so its front is the maximum. On `Push(x)`, pop from the back of `mins` every value **strictly** greater than `x`, and from `maxs` every value strictly smaller, so that equal values are kept. On `Pop`, if the removed value equals the front of a deque, pop that front too.
*Two-stack follow-up:* each stack entry stores `(value, min so far, max so far)`. Push onto the in-stack. When the out-stack is empty, move everything over. `Min` combines the two stack tops.

**Complexity.** Amortized O(1) per operation and O(n) space.

**Reference solution (Go)**

```go
package main

// MinMaxQueue is a FIFO queue plus two monotonic deques:
// mins is non-decreasing (front = minimum), maxs is non-increasing (front = maximum).
// Every value enters and leaves each deque at most once, so all operations
// are amortized O(1).
type MinMaxQueue struct {
	q, mins, maxs []int
}

func NewMinMaxQueue() *MinMaxQueue { return &MinMaxQueue{} }

func (m *MinMaxQueue) Push(x int) {
	m.q = append(m.q, x)
	for len(m.mins) > 0 && m.mins[len(m.mins)-1] > x { // keep equal values
		m.mins = m.mins[:len(m.mins)-1]
	}
	m.mins = append(m.mins, x)
	for len(m.maxs) > 0 && m.maxs[len(m.maxs)-1] < x {
		m.maxs = m.maxs[:len(m.maxs)-1]
	}
	m.maxs = append(m.maxs, x)
}

func (m *MinMaxQueue) Pop() int {
	x := m.q[0]
	m.q = m.q[1:]
	if m.mins[0] == x {
		m.mins = m.mins[1:]
	}
	if m.maxs[0] == x {
		m.maxs = m.maxs[1:]
	}
	return x
}

func (m *MinMaxQueue) Min() int { return m.mins[0] }
func (m *MinMaxQueue) Max() int { return m.maxs[0] }
```

**Expected output:** `1 5`, `3 1 5`, `1 5 5`, `2 5`, `5 2 2`. The reference also matched a brute-force model on 200 random operation sequences with many duplicates.

---
## Q9. Debug and Extend: Recent Files Menu

*Source: original debug-and-extend question from the transcript. Linked list: **yes** (doubly, with sentinels, plus a hash map).*

**Skills needed beyond linked lists:** Hash maps, sentinel nodes, and reading and debugging existing code (pointer-update order, off-by-one errors, keeping two structures in sync). For Part 2, an entry that belongs to two linked lists at once.

This question has two parts. Students receive the buggy code below, fix it, and then add a feature on top.

**Context.** A text editor keeps a "Recent Files" menu with a fixed capacity. Opening a file moves it to the top. When the menu is full, opening a new file removes the least recently opened one. The code below is supposed to do this, but it has bugs.

### Part 1: Fix the bugs

**Intended behaviour**
- `NewRecentFiles(capacity)` creates an empty menu (`1 ≤ capacity ≤ 1000`).
- `Open(name)` moves `name` to the front, or inserts it at the front. If the menu is full and `name` is new, it first evicts the least recently opened file.
- `List()` returns names from most to least recent.
- `Contains(name)` and `Size()` do what their names say.

**Buggy code (give this to students exactly as written)**
```go
package main

type entry struct {
	name       string
	prev, next *entry
}

type RecentFiles struct {
	cap        int
	index      map[string]*entry
	head, tail *entry // sentinels: head.next is the most recent
}

func NewRecentFiles(capacity int) *RecentFiles {
	h, t := &entry{}, &entry{}
	h.next = t
	t.prev = h
	return &RecentFiles{cap: capacity, index: map[string]*entry{}, head: h, tail: t}
}

func (r *RecentFiles) unlink(e *entry) {
	e.prev.next = e.next
	e.next.prev = e.prev
}

func (r *RecentFiles) pushFront(e *entry) {
	e.next = r.head.next
	e.prev = r.head
	r.head.next = e
	r.head.next.prev = e
}


func (r *RecentFiles) Open(name string) {
	if e, ok := r.index[name]; ok {
		r.pushFront(e)
		return
	}
	if len(r.index) > r.cap {
		last := r.tail.prev
		r.unlink(last)
	}
	e := &entry{name: name}
	r.index[name] = e
	r.pushFront(e)
}

func (r *RecentFiles) List() []string {
	var out []string
	for e := r.head.next; e != r.tail; e = e.next {
		out = append(out, e.name)
	}
	return out
}

func (r *RecentFiles) Contains(name string) bool {
	_, ok := r.index[name]
	return ok
}

func (r *RecentFiles) Size() int { return len(r.index) }
```

**Failing test**
```
r := NewRecentFiles(3)
Open("a"), Open("b"), Open("c")   // List → [c, b, a]
Open("a")                         // List → [a, c, b]
Open("d")                         // evicts b, List → [d, a, c]
Open("c")                         // List → [c, d, a]
Open("e")                         // evicts a, List → [e, c, d]
Size() → 3, Contains("a") → false, Contains("b") → false
```

**Part 1 tasks**
1. Find and fix **every** bug without rewriting the program from scratch.
2. Explain each bug in one sentence.
3. Show that the failing test above passes.

### Part 2: Add the pin feature

After Part 1 passes, extend the type so users can pin files.

- `Open(name string) bool` now returns whether the file is in the menu afterwards.
- `Pin(name string) bool` pins a file that is already in the menu. It returns `false` if the file is absent, and `true` if the file is already pinned. **Pinning does not count as a use.**
- `Unpin(name string) bool` unpins a file. It returns `false` if the file is absent or not pinned. **A successful unpin counts as a use: the file moves to the front of the menu.** (This rule was added so that `Unpin` can be O(1). See the mentor note in the answer.)
- Pinned files still count toward the capacity and keep their recency order in `List()`, but they are never evicted.
- When a new file is opened into a full menu, evict the least recent **unpinned** file. If every file is pinned, `Open` returns `false`, adds nothing and changes nothing.
- Opening a file that is already in the menu always returns `true` and moves it to the front, whether it is pinned or not.

**Required signatures after Part 2 (Go)**
```go
func NewRecentFiles(capacity int) *RecentFiles
func (r *RecentFiles) Open(name string) bool
func (r *RecentFiles) Pin(name string) bool
func (r *RecentFiles) Unpin(name string) bool
func (r *RecentFiles) List() []string
func (r *RecentFiles) Contains(name string) bool
func (r *RecentFiles) Size() int
```

**Example**
```
r := NewRecentFiles(3)
Open("a"), Open("b"), Open("c")   // List → [c, b, a]
Pin("a")          → true
Open("d")         → true          // a is the oldest but pinned, so b is evicted
List()            → [d, c, a]
Pin("c")          → true
Pin("d")          → true          // all three are pinned
Open("e")         → false         // nothing can be evicted, list unchanged
List()            → [d, c, a]
Unpin("a")        → true          // counts as a use
List()            → [a, d, c]
Open("e")         → true          // evicts a (the only unpinned file)
List()            → [e, d, c]
Pin("zzz")        → false
Unpin("e")        → false         // e is not pinned
```

**Constraints.** Up to 200,000 calls. `Open`, `Pin` and `Unpin` must be O(1) amortized.
**Follow-up.** Finding the least recent unpinned file must not scan past pinned entries on every call. Keep pinned files out of the eviction list, for example by putting unpinned files in a second list, so that eviction stays O(1).

### Expected answer — Q9

**Bugs in the given code**

| # | Bug (one sentence) | Fix |
|---|---|---|
| 1 | In `pushFront`, `r.head.next = e` runs before `r.head.next.prev = e`, so the second line sets `e.prev = e` and never updates the old first node's `prev`. | Run `r.head.next.prev = e` first, then `r.head.next = e`. |
| 2 | Opening a file that already exists calls `pushFront` without `unlink`, so the node is linked twice and the list becomes a cycle. | Call `r.unlink(e)` before `r.pushFront(e)`. |
| 3 | The eviction check `len(r.index) > r.cap` is off by one, so the menu grows to `cap + 1` before it evicts anything. | Use `>=`. |
| 4 | Eviction unlinks the node but never removes it from `r.index`, so `Contains` and `Size` are wrong and the map leaks. | Add `delete(r.index, last.name)`. |

What a mentor will see with the unmodified code: `List()` after the first three opens is correct, because forward links survive bug 1. The `Open("a")` re-open (bug 2) turns the list into a cycle, so the next `List()` call **never returns** (verified: the test hangs until a timeout kills it).

**Part 1 reference (fixed code, with each change marked)**

**Reference solution (Go)**

```go
package main

type entry struct {
	name       string
	prev, next *entry
}

type RecentFiles struct {
	cap        int
	index      map[string]*entry
	head, tail *entry // sentinels: head.next is the most recent
}

func NewRecentFiles(capacity int) *RecentFiles {
	h, t := &entry{}, &entry{}
	h.next = t
	t.prev = h
	return &RecentFiles{cap: capacity, index: map[string]*entry{}, head: h, tail: t}
}

func (r *RecentFiles) unlink(e *entry) {
	e.prev.next = e.next
	e.next.prev = e.prev
}

func (r *RecentFiles) pushFront(e *entry) {
	e.next = r.head.next
	e.prev = r.head
	r.head.next.prev = e // FIX 1: update the old first node before moving head.next
	r.head.next = e
}

func (r *RecentFiles) Open(name string) {
	if e, ok := r.index[name]; ok {
		r.unlink(e) // FIX 2: detach before re-inserting at the front
		r.pushFront(e)
		return
	}
	if len(r.index) >= r.cap { // FIX 3: evict when full, not when over-full
		last := r.tail.prev
		r.unlink(last)
		delete(r.index, last.name) // FIX 4: keep the map in sync with the list
	}
	e := &entry{name: name}
	r.index[name] = e
	r.pushFront(e)
}

func (r *RecentFiles) List() []string {
	var out []string
	for e := r.head.next; e != r.tail; e = e.next {
		out = append(out, e.name)
	}
	return out
}

func (r *RecentFiles) Contains(name string) bool {
	_, ok := r.index[name]
	return ok
}

func (r *RecentFiles) Size() int { return len(r.index) }
```

**Part 1 expected output:** `[c b a]`, `[a c b]`, `[d a c]`, `[c d a]`, `[e c d]`, `Size() = 3`, `Contains("a") = false`, `Contains("b") = false`. Walking the `prev` links from the tail gives `[d c e]`, which confirms the backward links are consistent.

**Part 2 approach.** Every entry stays in the recency list that `List()` walks. Unpinned entries are also threaded through a second list (`uprev`/`unext`) in the same recency order. Eviction takes the tail of that second list. If the second list is empty, every file is pinned and `Open` returns `false`. `Pin` removes the entry from the eviction list. `Unpin` counts as a use, so the entry rejoins **both** lists at the front, which is the position its recency requires. Every operation is O(1) worst case, with O(capacity) space.

> **Mentor note on `Unpin`:** The transcript's design note points out that if `Unpin` does *not* count as a use, an unpinned file has to be re-inserted into the eviction list at its old recency position. That needs either a walk past the pinned entries (O(number of pinned files)) or an ordered structure (O(log n)). This paper adopts the transcript's suggested rule ("unpinned files rejoin at the front") so the O(1) requirement can be met. The transcript's original example gives the same results under this rule. If you prefer the other rule, relax the `Unpin` requirement to O(p) or O(log n).

**Part 2 reference (complete extended type)**

**Reference solution (Go)**

```go
package main

// Part 2: pinning. Every entry lives in the recency list (prev/next), which
// List() walks. Unpinned entries are ALSO in a second list (uprev/unext) in
// the same recency order; eviction takes its tail, so it never scans pinned
// files. All operations are O(1).
type entry struct {
	name         string
	pinned       bool
	prev, next   *entry // recency list: every file
	uprev, unext *entry // eviction list: unpinned files only
}

type RecentFiles struct {
	cap          int
	index        map[string]*entry
	head, tail   *entry // sentinels of the recency list
	uhead, utail *entry // sentinels of the eviction list
}

func NewRecentFiles(capacity int) *RecentFiles {
	r := &RecentFiles{cap: capacity, index: map[string]*entry{},
		head: &entry{}, tail: &entry{}, uhead: &entry{}, utail: &entry{}}
	r.head.next, r.tail.prev = r.tail, r.head
	r.uhead.unext, r.utail.uprev = r.utail, r.uhead
	return r
}

func (r *RecentFiles) unlink(e *entry) {
	e.prev.next = e.next
	e.next.prev = e.prev
}

func (r *RecentFiles) pushFront(e *entry) {
	e.next = r.head.next
	e.prev = r.head
	r.head.next.prev = e
	r.head.next = e
}

func (r *RecentFiles) uUnlink(e *entry) {
	e.uprev.unext = e.unext
	e.unext.uprev = e.uprev
}

func (r *RecentFiles) uPushFront(e *entry) {
	e.unext = r.uhead.unext
	e.uprev = r.uhead
	r.uhead.unext.uprev = e
	r.uhead.unext = e
}

// touch makes e the most recent file in both lists.
func (r *RecentFiles) touch(e *entry) {
	r.unlink(e)
	r.pushFront(e)
	if !e.pinned {
		r.uUnlink(e)
		r.uPushFront(e)
	}
}

func (r *RecentFiles) Open(name string) bool {
	if e, ok := r.index[name]; ok {
		r.touch(e)
		return true
	}
	if len(r.index) >= r.cap {
		victim := r.utail.uprev
		if victim == r.uhead { // every file is pinned
			return false
		}
		r.uUnlink(victim)
		r.unlink(victim)
		delete(r.index, victim.name)
	}
	e := &entry{name: name}
	r.index[name] = e
	r.pushFront(e)
	r.uPushFront(e)
	return true
}

// Pin does not count as a use: the file keeps its place in List().
func (r *RecentFiles) Pin(name string) bool {
	e, ok := r.index[name]
	if !ok {
		return false
	}
	if !e.pinned {
		e.pinned = true
		r.uUnlink(e)
	}
	return true
}

// Unpin counts as a use (stated rule): the file moves to the front, which is
// exactly where it must rejoin the eviction list, so this stays O(1).
func (r *RecentFiles) Unpin(name string) bool {
	e, ok := r.index[name]
	if !ok || !e.pinned {
		return false
	}
	e.pinned = false
	r.uPushFront(e)
	r.unlink(e)
	r.pushFront(e)
	return true
}

func (r *RecentFiles) List() []string {
	var out []string
	for e := r.head.next; e != r.tail; e = e.next {
		out = append(out, e.name)
	}
	return out
}

func (r *RecentFiles) Contains(name string) bool {
	_, ok := r.index[name]
	return ok
}

func (r *RecentFiles) Size() int { return len(r.index) }
```

**Part 2 expected output:** `true`, `true`, `[d c a]`, `true`, `true`, `false`, `[d c a]`, `true`, `[a d c]`, `true`, `[e d c]`, `false`, `false`. The reference also passed two extra checks: the Part 1 sequence still ends at `[e c d]`, and re-opening a pinned file returns `true` and moves it to the front.

---

# Tier 3: Upper medium

## Q10. Multiply Sparse Polynomials

*Source: original question from the transcript's set. Linked list: **yes** (singly list of terms).*

**Skills needed beyond linked lists:** Accumulating by exponent (a slice indexed by exponent, or a hash map plus sorting), building an output list with a dummy head, and dropping coefficients that cancel to zero.

**Statement**
A polynomial is stored as a linked list of terms `Term{Coef, Exp, Next}`. The terms are sorted by `Exp` in **descending** order, with no duplicate exponents and no zero coefficients. Return the product of two polynomials as a **new** list in the same form. Terms whose coefficients cancel to 0 must not appear. If the product is zero (either input is `nil`), return `nil`.

**Input / Output**
- Input: two polynomial lists, `a` and `b`. The examples write each term as `(coef, exp)`.
- Output: the head of the product list.

**Constraints**
- Up to 200 terms per polynomial
- `0 ≤ Exp ≤ 5000`, `1 ≤ |Coef| ≤ 100`

**Required signature (Go)**
```go
type Term struct {
	Coef int
	Exp  int
	Next *Term
}

func MultiplyPoly(a, b *Term) *Term
```

**Examples**
```
[(2,2),(3,0)] × [(1,1),(1,0)]   → [(2,3),(2,2),(3,1),(3,0)]     // (2x²+3)(x+1)
[(1,1),(1,0)] × [(1,1),(-1,0)]  → [(1,2),(-1,0)]                // (x+1)(x−1): the x terms cancel
nil × [(5,3)]                   → nil
[(3,5000)] × [(-2,5000),(4,0)]  → [(-6,10000),(12,5000)]
```

**Instructions**
- Do not modify the input lists.
- Target O(|a|·|b| + E) time, where E is the largest result exponent, or O(|a|·|b| · log) if you use a map and sort.
- Explain how you avoid emitting zero terms.

### Expected answer — Q10

**Approach.** The heads hold the largest exponents, so the largest product exponent is `a.Exp + b.Exp`. Allocate an accumulator of that size + 1 and add `x.Coef · y.Coef` into slot `x.Exp + y.Exp` for every pair of terms. Then walk the exponents from high to low, appending a new term for each non-zero slot (using a dummy head). A k-way merge of `a × y` for each term `y` of `b` is a pure-list alternative.

**Complexity.** O(|a|·|b| + E) time and O(E) extra space, where E ≤ 10,000.

**Reference solution (Go)**

```go
package main

type Term struct {
	Coef int
	Exp  int
	Next *Term
}

// MultiplyPoly multiplies two sparse polynomials stored as linked lists in
// descending exponent order. Products are accumulated in a slice indexed by
// exponent (the heads hold the largest exponents, so the size is known),
// then the result list is built from the highest exponent down.
func MultiplyPoly(a, b *Term) *Term {
	if a == nil || b == nil {
		return nil
	}
	acc := make([]int, a.Exp+b.Exp+1)
	for x := a; x != nil; x = x.Next {
		for y := b; y != nil; y = y.Next {
			acc[x.Exp+y.Exp] += x.Coef * y.Coef
		}
	}
	dummy := &Term{}
	tail := dummy
	for e := len(acc) - 1; e >= 0; e-- {
		if acc[e] != 0 { // cancelled terms are skipped
			tail.Next = &Term{Coef: acc[e], Exp: e}
			tail = tail.Next
		}
	}
	return dummy.Next
}
```

**Expected output:** `[(2,3),(2,2),(3,1),(3,0)]`, `[(1,2),(-1,0)]`, `nil`, `[(-6,10000),(12,5000)]`.

---

## Q11. Flatten by Levels

*Source: original question from the transcript's set (variant of LeetCode 430). Linked list: **yes** (doubly, with child pointers).*

**Skills needed beyond linked lists:** Breadth-first (level-order) traversal. The basic version uses a queue. The follow-up uses the list itself as the queue, keeping a tail pointer.

**Statement**
Each node has `Prev`, `Next` and `Child`, as in Q6. Flatten the structure into a single-level doubly linked list in **level order**: all level-1 nodes, then all level-2 nodes, then level 3, and so on. Within a level, child lists appear in the order their parents appear in the flattened result. Set every `Child` to `nil` and keep every `Prev` pointer correct.

**Input / Output**
- Input: the head of level 1.
- Output: the head of the flattened list.

**Constraints**
- `0 ≤` total nodes `≤ 5000`

**Required signature (Go)**
```go
// MNode is the same type as in Q6.
func FlattenByLevel(head *MNode) *MNode
```

**Examples**
```
Level 1: 1-2-3. Node 1 has child list 9-10, node 2 has child list 4-5,
node 3 has child 6, and node 4 has child list 7-8.
  → 1 2 3 9 10 4 5 6 7 8
Level 1: 1-2; node 1 has child 3
  → 1 2 3        (compare Q6's depth-first result: 1 3 2)
nil → nil
```

**Instructions**
- O(n) time.
- Follow-up: O(1) extra space, with no queue or recursion.

### Expected answer — Q11

**Approach (meets the follow-up).** Find the tail of level 1. Walk `cur` from the head. Whenever `cur` has a child, append the whole child list after `tail`, clear `cur.Child`, and move `tail` to the end of what was appended. Parents are visited in level order, so child lists are appended in level order: the list itself acts as the BFS queue. A simpler version collects the child lists in a queue level by level, which uses O(width) extra space.

**Complexity.** O(n) time, because each node is passed once by `cur` and once by the tail walk. O(1) extra space.

**Reference solution (Go)**

```go
package main

type MNode struct {
	Val   int
	Prev  *MNode
	Next  *MNode
	Child *MNode
}

// FlattenByLevel walks the list from the head while keeping a pointer to the
// current end. Whenever the walker meets a child list, the whole child list
// is appended at the end. Because parents are visited in level order, the
// children are appended in level order too: this is BFS with the list itself
// acting as the queue. O(n) time, O(1) extra space.
func FlattenByLevel(head *MNode) *MNode {
	if head == nil {
		return nil
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Child == nil {
			continue
		}
		child := cur.Child
		cur.Child = nil
		tail.Next = child
		child.Prev = tail
		for tail.Next != nil {
			tail = tail.Next
		}
	}
	return head
}
```

**Expected output:** `1 2 3 9 10 4 5 6 7 8`, `1 2 3`, `nil`. `Prev` links and cleared `Child` pointers were also verified.

---

## Q12. Reversing Circle

*Source: original question from the transcript's set (Josephus variant). Linked list: optional (a circular doubly linked list or a slice both work).*

**Skills needed beyond linked lists:** Modular arithmetic (reducing `k` modulo the current size, because `k` can be 10^9), index bookkeeping after removing an element, and simulation.

**Statement**
`n` people stand in a circle, numbered `1 … n` in clockwise order. Counting starts at person 1 and goes clockwise. The starting person counts as 1, and the `k`-th person counted is eliminated. After each elimination, the direction **reverses**. Counting restarts from the first remaining person next to the eliminated one in the new direction, and that person counts as 1. Return the elimination order, with the last survivor at the end.

**Input / Output**
- Input: `n` and `k`.
- Output: a slice of length `n` with the elimination order. The last element is the survivor.

**Constraints**
- `1 ≤ n ≤ 5000`
- `1 ≤ k ≤ 10^9`

**Required signature (Go)**
```go
func EliminationOrder(n, k int) []int
```

**Examples**
```
EliminationOrder(5, 3) → [3 5 4 2 1]
EliminationOrder(4, 2) → [2 4 3 1]
EliminationOrder(3, 1) → [1 3 2]
EliminationOrder(1, 7) → [1]
```
Trace of the first example: count 1,2,**3** clockwise, so 3 goes. Reverse and start at 2: count 2,1,**5**, so 5 goes. Reverse and start at 1: count 1,2,**4**, so 4 goes. Reverse and start at 2: count 2,1,**2**, so 2 goes. Person 1 survives.

**Instructions**
- O(n²) is acceptable. Stepping `k` times per elimination is not, because k can be 10^9.
- You may use a circular doubly linked list (step `(k−1) mod size` times) or a slice with index arithmetic.

### Expected answer — Q12

**Approach.** Keep the remaining people in a slice in clockwise order, along with `pos` (the index of the person who counts as 1) and `dir` (±1). The eliminated index is `(pos + dir·((k−1) mod m)) mod m`. After removing it, flip `dir`. If the new direction is clockwise, the next start is the person who slid into the removed index (`idx mod (m−1)`). If it is counter-clockwise, the next start is the person just before it (`(idx−1) mod (m−1)`).

**Complexity.** O(n²) time because of slice deletions (25·10^6 element moves at most), and O(n) space. A circular doubly linked list gives the same bound, with O(min(k, m)) steps per elimination.

**Reference solution (Go)**

```go
package main

// EliminationOrder simulates the reversing circle on a slice that lists the
// remaining people in clockwise order. Reducing (k-1) modulo the current size
// handles huge k. O(n^2) time because of slice deletions, which is fine for
// n <= 5000.
func EliminationOrder(n, k int) []int {
	people := make([]int, n)
	for i := range people {
		people[i] = i + 1
	}
	order := make([]int, 0, n)
	pos, dir := 0, 1 // pos = index of the person who is counted as 1
	for len(people) > 0 {
		m := len(people)
		step := (k - 1) % m
		idx := ((pos+dir*step)%m + m) % m
		order = append(order, people[idx])
		people = append(people[:idx], people[idx+1:]...)
		dir = -dir
		if m-1 == 0 {
			break
		}
		if dir == 1 {
			pos = idx % (m - 1) // next person clockwise slid into idx
		} else {
			pos = (idx - 1 + (m - 1)) % (m - 1) // previous person clockwise
		}
	}
	return order
}
```

**Expected output:** `[3 5 4 2 1]`, `[2 4 3 1]`, `[1 3 2]`, `[1]`. The reference was also checked against a step-by-step circular doubly linked list simulation on 2000 random cases, and once with k = 10^9.

---

## Q13. Queue (Codeforces 91B)

*Source: Codeforces 91B "Queue" (about 1500–1800 rating). Linked list: no.*

**Skills needed beyond linked lists:** Suffix minimums, binary search on a monotone array (`sort.Search`). A monotonic stack is an alternative.

**Statement**
`n` walruses stand in a line. Walrus `0` is at the end of the queue, and walrus `n−1` is at the front. Walrus `i` has age `a[i]`. Walrus `i` is displeased if some walrus **ahead** of it (index `j > i`) is strictly younger. Its displeasure is the number of walruses standing between it and the **farthest** such younger walrus, which is `j − i − 1` for the largest such `j`. If nobody ahead is younger, the answer is `−1`.

**Input / Output**
- Input: the ages `a`.
- Output: the displeasure of each walrus.

**Constraints**
- `2 ≤ n ≤ 10^5`
- `1 ≤ a[i] ≤ 10^9`

**Required signature (Go)**
```go
func Displeasure(a []int) []int
```

**Examples**
```
Displeasure([]int{10, 8, 5, 3, 50, 45})     → [2 1 0 -1 0 -1]
Displeasure([]int{10, 9, 8})                → [1 0 -1]
Displeasure([]int{1, 2, 3})                 → [-1 -1 -1]
Displeasure([]int{10, 4, 6, 3, 2, 8, 15})   → [4 2 1 0 -1 -1 -1]
```

**Instructions**
- Target O(n log n) or O(n). The O(n²) scan is too slow.
- Hint: the suffix minimum `suf[j] = min(a[j..n−1])` is non-decreasing from left to right.

### Expected answer — Q13

**Approach.** Build `suf`. For walrus `i`, the farthest younger walrus is the **last** index `j` with `suf[j] < a[i]`. At that index, `suf[j+1] ≥ a[i]`, so `a[j]` itself is the minimum and is younger. Because `suf` is non-decreasing, `j` is one less than the first index where `suf ≥ a[i]`, which binary search finds. If `j > i`, the answer is `j − i − 1`. Otherwise it is `−1`.

**Complexity.** O(n log n) time and O(n) space.

**Reference solution (Go)**

```go
package main

import "sort"

// Displeasure: suf[j] = min(a[j..n-1]) is non-decreasing, so the farthest
// younger walrus ahead of i is the last j with suf[j] < a[i], found by
// binary search. O(n log n) time, O(n) space.
func Displeasure(a []int) []int {
	n := len(a)
	suf := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		suf[i] = a[i]
		if i+1 < n && suf[i+1] < suf[i] {
			suf[i] = suf[i+1]
		}
	}
	res := make([]int, n)
	for i, x := range a {
		// first index whose suffix minimum is >= x; the one before it is the answer
		j := sort.Search(n, func(j int) bool { return suf[j] >= x }) - 1
		if j > i {
			res[i] = j - i - 1
		} else {
			res[i] = -1
		}
	}
	return res
}
```

**Expected output:** `[2 1 0 -1 0 -1]`, `[1 0 -1]`, `[-1 -1 -1]`, `[4 2 1 0 -1 -1 -1]`. The reference also matched a brute-force O(n²) solution on 2000 random arrays.

---

## Q14. Commit History

*Source: original question from the transcript's set. Linked list: **no**. This is a parent-pointer tree, and each branch head walks up a chain of parents.*

**Skills needed beyond linked lists:** Trees stored as parent arrays, hash maps (branch name → head commit), lowest common ancestor (LCA) by equalising depths, and binary lifting for the O(log n) follow-up.

**Statement**
Design a tiny version-control history.

- `NewRepo()` creates a repo with one branch, `"main"`, pointing at root commit `0`. `"main"` is checked out.
- `Commit() int` creates a new commit whose parent is the head of the current branch. It moves the current branch to the new commit and returns the new id. Ids count up from 1.
- `Branch(name)` creates a new branch at the current head. It does not switch to it.
- `Checkout(name)` switches to an existing branch.
- `MergeBase(a, b string) int` returns the most recent common ancestor of the heads of branches `a` and `b`. A commit counts as its own ancestor.
- `Log() []int` returns commit ids from the current head back to the root.

**Input / Output**
- Input: a sequence of calls. Branch names given to `Checkout` and `MergeBase` always exist, and names given to `Branch` are always new.
- Output: the return values of `Commit`, `MergeBase` and `Log`.

**Constraints**
- Up to 10^5 operations in total. The total length of all `Log` outputs is at most 10^6.

**Required signature (Go)**
```go
type Repo struct{ /* your fields */ }

func NewRepo() *Repo
func (r *Repo) Commit() int
func (r *Repo) Branch(name string)
func (r *Repo) Checkout(name string)
func (r *Repo) MergeBase(a, b string) int
func (r *Repo) Log() []int
```

**Example**
```
r := NewRepo()
r.Commit() → 1
r.Commit() → 2
r.Branch("dev"); r.Checkout("dev")
r.Commit() → 3
r.Commit() → 4
r.Checkout("main")
r.Commit() → 5
r.MergeBase("main", "dev") → 2
r.Log()                    → [5 2 1 0]
r.Checkout("dev"); r.Branch("fix"); r.Checkout("fix")
r.Commit() → 6
r.MergeBase("fix", "dev")  → 4
r.MergeBase("fix", "main") → 2
r.Log()                    → [6 4 3 2 1 0]
```

**Instructions**
- Main task: `MergeBase` in O(depth). Store a parent and a depth for each commit, and walk the deeper head up first.
- Follow-up: make `MergeBase` O(log n) with binary lifting.

### Expected answer — Q14

**Approach.** Each commit stores its depth and a binary-lifting table `up[v][k]`, the 2^k-th ancestor of v. The root's ancestors are itself. `Commit` fills its table in O(log n) from its parent's table. `MergeBase` lifts the deeper head to the other's depth. If the two heads are now equal, that commit is the answer. Otherwise, lift both together for each `k` from high to low while their ancestors differ. The answer is the parent of either. `Log` follows `up[v][0]` back to 0.

**Complexity.** `Commit` is O(log n). `MergeBase` is O(log n). `Log` is O(length of the output). Space is O(n log n).

**Reference solution (Go)**

```go
package main

const logN = 17 // 2^17 > 100,000 commits

// Repo stores commits as a parent-pointer tree. up[v][k] is the 2^k-th
// ancestor of commit v (binary lifting), so MergeBase is O(log n).
type Repo struct {
	depth    []int
	up       [][logN]int
	branches map[string]int // branch name -> head commit
	current  string
}

func NewRepo() *Repo {
	r := &Repo{branches: map[string]int{"main": 0}, current: "main"}
	r.depth = append(r.depth, 0)
	r.up = append(r.up, [logN]int{}) // root's ancestors are itself (0)
	return r
}

func (r *Repo) Commit() int {
	parent := r.branches[r.current]
	id := len(r.depth)
	var anc [logN]int
	anc[0] = parent
	for k := 1; k < logN; k++ {
		anc[k] = r.up[anc[k-1]][k-1]
	}
	r.depth = append(r.depth, r.depth[parent]+1)
	r.up = append(r.up, anc)
	r.branches[r.current] = id
	return id
}

func (r *Repo) Branch(name string)   { r.branches[name] = r.branches[r.current] }
func (r *Repo) Checkout(name string) { r.current = name }

func (r *Repo) MergeBase(a, b string) int {
	u, v := r.branches[a], r.branches[b]
	if r.depth[u] < r.depth[v] {
		u, v = v, u
	}
	for k := logN - 1; k >= 0; k-- { // lift u to v's depth
		if r.depth[u]-(1<<k) >= r.depth[v] {
			u = r.up[u][k]
		}
	}
	if u == v {
		return u
	}
	for k := logN - 1; k >= 0; k-- { // lift both to just below the LCA
		if r.up[u][k] != r.up[v][k] {
			u, v = r.up[u][k], r.up[v][k]
		}
	}
	return r.up[u][0]
}

func (r *Repo) Log() []int {
	var out []int
	for c := r.branches[r.current]; ; c = r.up[c][0] {
		out = append(out, c)
		if c == 0 {
			return out
		}
	}
}
```

**Expected output:** `1, 2, 3, 4, 5`, `2`, `[5 2 1 0]`, `6`, `4`, `2`, `[6 4 3 2 1 0]`. The reference also matched a naive ancestor-set check on 20,000 random queries.

---

## Q15. Abbreviation

*Source: HackerRank, "Abbreviation". Linked list: no.*

**Skills needed beyond linked lists:** 2D dynamic programming over two strings, then reducing it to a 1D rolling array (the loop order matters), and character-case handling.

**Statement**
You are given strings `a` and `b`. You may capitalize zero or more of the lowercase letters in `a`, then delete all the lowercase letters that remain. In effect, each lowercase letter is either capitalized or deleted, and uppercase letters already in `a` can never be deleted. Can you turn `a` into exactly `b`?

**Input / Output**
- Input: `a`, made of English letters, and `b`, made of uppercase letters.
- Output: `true` if `a` can become `b`, otherwise `false`.

**Constraints**
- `1 ≤ |a| ≤ 1000`, `0 ≤ |b| ≤ 1000`

**Required signature (Go)**
```go
func CanAbbreviate(a, b string) bool
```

**Examples**
```
CanAbbreviate("daBcd", "ABC")  → true    // capitalize a and c, delete both d's
CanAbbreviate("AbcDE", "AFDE") → false   // no F is available
CanAbbreviate("aaa", "")       → true    // delete everything
CanAbbreviate("KXzQ", "K")     → false   // X and Q are uppercase and can't be deleted
```

**Instructions**
- Main task: O(|a|·|b|) dynamic programming.
- Follow-up: reduce the memory to O(|b|).

### Expected answer — Q15

**Approach.** Let `dp[i][j]` mean "`a[:i]` can become `b[:j]`", with `dp[0][0] = true`. For character `c = a[i]`:
- `dp[i+1][j+1] |= dp[i][j]` if `upper(c) == b[j]` (capitalize or keep `c` and match it);
- `dp[i+1][j] |= dp[i][j]` if `c` is lowercase (delete it).

To use one row, iterate `j` **downward** so `dp[j]` still holds the previous row when `dp[j+1]` is computed. Then set `dp[0] = dp[0] && lowercase(c)`.

**Complexity.** O(|a|·|b|) time. O(|b|) space in the follow-up version.

**Reference solution (Go)**

```go
package main

// CanAbbreviate: dp[j] says whether the processed prefix of a can become
// b[:j]. For each character c of a (j runs downward so dp[j] is still the
// previous row):
//   - capitalize c and match b[j]:   needs old dp[j]
//   - delete c (only if lowercase):   needs old dp[j+1]
// O(|a|*|b|) time, O(|b|) space.
func CanAbbreviate(a, b string) bool {
	dp := make([]bool, len(b)+1)
	dp[0] = true
	for i := 0; i < len(a); i++ {
		c := a[i]
		lower := c >= 'a' && c <= 'z'
		up := c
		if lower {
			up = c - 'a' + 'A'
		}
		for j := len(b) - 1; j >= 0; j-- {
			dp[j+1] = (up == b[j] && dp[j]) || (lower && dp[j+1])
		}
		dp[0] = dp[0] && lower // an uppercase letter can never be deleted
	}
	return dp[len(b)]
}
```

**Expected output:** `true`, `false`, `true`, `false`. The reference also matched an exhaustive recursive search on 5000 random small cases.

---

## Q16. Mysterious Present (Codeforces 4D)

*Source: Codeforces 4D "Mysterious Present" (about 1700 rating). Linked list: no.*

**Skills needed beyond linked lists:** Sorting with a custom comparator (width ascending, and **height descending** for equal widths), the longest strictly increasing subsequence with binary search in O(n log n), and rebuilding the actual chain with a parent array.

**Statement**
You have a card of size `(cw, ch)` and `n` envelopes, where envelope `i` (1-based) has size `(w_i, h_i)`. An item fits into an envelope only if its width **and** height are both **strictly** smaller. Build the longest chain of envelopes in which each envelope fits inside the next one and the card fits into the smallest. Envelopes can't be rotated. Return the chain length and the envelope indices from smallest to largest. If no envelope can hold the card, return `0` and an empty slice.

**Input / Output**
- Input: card size `cw, ch`, and `env[i] = {w, h}` for envelope `i+1`.
- Output: `(length, indices)`. If several longest chains exist, **any one is accepted**.

**Constraints**
- `1 ≤ n ≤ 5000`
- `1 ≤ w, h, cw, ch ≤ 10^6`

**Required signature (Go)**
```go
func Chain(cw, ch int, env [][2]int) (int, []int)
```

**Examples**
```
Chain(1, 1, [][2]int{{2,2},{2,2}})          → (1, [1])   // (1, [2]) is equally correct
Chain(3, 3, [][2]int{{5,4},{12,11},{9,8}})  → (3, [1 3 2])
Chain(10, 10, [][2]int{{5,20},{20,5},{10,11}}) → (0, [])  // no envelope is larger than the card in both dimensions
```

**Instructions**
- Main task: an O(n²) DP after sorting is accepted.
- Follow-up: O(n log n). Be careful with equal widths, because two envelopes with the same width can never both be in a chain.

### Expected answer — Q16

**Approach.** Discard envelopes that can't hold the card. Sort the rest by width ascending, and by height **descending** when widths are equal, so that a strictly increasing run of heights can never use two envelopes of the same width. Then find the longest **strictly** increasing subsequence of heights with patience sorting. `tails[l]` is the envelope ending the best chain of length `l+1` that has the smallest height. Binary-search for the first tail with height `≥ h`. Store `parent[id] = tails[l−1]`, then walk the parents back from the last tail to rebuild the chain.

**Complexity.** O(n log n) time and O(n) space.

**Reference solution (Go)**

```go
package main

import "sort"

// Chain: keep only envelopes larger than the card, sort by width ascending
// and, for equal widths, height DESCENDING (so two envelopes of the same
// width can never both be picked), then take a strictly increasing LIS on
// heights with binary search. tails[l] holds the index of the envelope that
// ends the best chain of length l+1 with the smallest height; parent[]
// rebuilds the chain. O(n log n) time, O(n) space.
func Chain(cw, ch int, env [][2]int) (int, []int) {
	var ids []int // 0-based indices of usable envelopes
	for i, e := range env {
		if e[0] > cw && e[1] > ch {
			ids = append(ids, i)
		}
	}
	sort.Slice(ids, func(x, y int) bool {
		a, b := env[ids[x]], env[ids[y]]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1] > b[1]
	})

	parent := make([]int, len(env))
	var tails []int
	for _, id := range ids {
		h := env[id][1]
		// first length whose tail height is >= h (strictly increasing LIS)
		l := sort.Search(len(tails), func(k int) bool { return env[tails[k]][1] >= h })
		parent[id] = -1
		if l > 0 {
			parent[id] = tails[l-1]
		}
		if l == len(tails) {
			tails = append(tails, id)
		} else {
			tails[l] = id
		}
	}
	if len(tails) == 0 {
		return 0, []int{}
	}
	chain := make([]int, len(tails))
	for k, id := len(tails)-1, tails[len(tails)-1]; k >= 0; k, id = k-1, parent[id] {
		chain[k] = id + 1 // 1-based for the answer
	}
	return len(chain), chain
}
```

**Expected output (reference):** `(1, [2])` for the first example. That is valid, and `[1]` is equally correct, so check that the chain is valid rather than comparing it with a fixed answer. Then `(3, [1 3 2])` and `(0, [])`. The reference was also checked against an O(n²) DP on 3000 random cases, and every returned chain was validated.

---

## Q17. Poisonous Plants

*Source: HackerRank, "Poisonous Plants". Linked list: no. A day-by-day simulation that relinks a list is the naive approach.*

**Skills needed beyond linked lists:** A monotonic stack that carries extra state (the day each plant dies), and reasoning about simultaneous events.

**Statement**
Plants stand in a row, and each has a pesticide level. Every day, each plant with **strictly more** pesticide than the plant immediately to its left dies. All of a day's deaths happen at the same time, after which the survivors close ranks. Return the number of days after which no plant dies any more.

**Input / Output**
- Input: the pesticide levels `p`, left to right.
- Output: the number of days until the row is stable.

**Constraints**
- `1 ≤ n ≤ 10^5`
- `0 ≤ p[i] ≤ 10^9`

**Required signature (Go)**
```go
func DaysUntilStable(p []int) int
```

**Examples**
```
DaysUntilStable([]int{6, 5, 8, 4, 7, 10, 9})  → 2
   day 1: 8, 7, 10 die → [6 5 4 9];  day 2: 9 dies → [6 5 4];  day 3: nothing
DaysUntilStable([]int{1, 2, 3, 4})            → 1   // 2, 3 and 4 all die on day 1
DaysUntilStable([]int{5, 4, 3})               → 0
DaysUntilStable([]int{3, 6, 2, 7, 5})         → 2
DaysUntilStable([]int{4, 3, 7, 5, 6, 4, 2})   → 3
```

**Instructions**
- A day-by-day simulation is O(n²) in the worst case and is too slow.
- Target O(n) with a monotonic stack that tracks the day on which each plant dies.

### Expected answer — Q17

**Approach.** Scan left to right with a stack of `(pesticide, deathDay)` whose pesticide values strictly increase from bottom to top. For plant `x`, pop every entry with pesticide ≥ `x` and remember the largest death day among them. Those plants protect nobody to their right once they are gone, so `x` must outlive them. If the stack is now empty, nothing to the left is smaller and `x` never dies (day 0). Otherwise, `x` dies one day after the last of the popped plants: `maxDay + 1`. The answer is the largest death day.

**Complexity.** O(n) time, because each plant is pushed and popped once, and O(n) space.

**Reference solution (Go)**

```go
package main

// DaysUntilStable: scan left to right with a stack of (pesticide, deathDay)
// whose pesticide values strictly increase from bottom to top. A plant dies
// one day after the latest-dying plant between it and the nearest plant to
// its left with strictly less pesticide; if no such plant exists, it never
// dies (day 0). O(n) time, O(n) space.
func DaysUntilStable(p []int) int {
	type item struct{ val, day int }
	var st []item
	ans := 0
	for _, x := range p {
		maxDay := 0
		for len(st) > 0 && st[len(st)-1].val >= x {
			maxDay = max(maxDay, st[len(st)-1].day)
			st = st[:len(st)-1]
		}
		day := 0
		if len(st) > 0 {
			day = maxDay + 1
		}
		ans = max(ans, day)
		st = append(st, item{x, day})
	}
	return ans
}
```

**Expected output:** `2`, `1`, `0`, `2`, `3`. The reference also matched a literal day-by-day simulation on 5000 random cases.

---

# Tier 4: Hard

## Q18. Longest Regular Bracket Sequence (Codeforces 5C)

*Source: Codeforces 5C (about 1900 rating). Linked list: no.*

**Skills needed beyond linked lists:** A stack of indices with a "base" sentinel (or a DP over the string), and counting ties.

**Statement**
You are given a string of `(` and `)`. Find the length of the longest substring that is a regular (valid) bracket sequence, and how many substrings of that length exist. If there is no such substring, return `(0, 1)`.

**Input / Output**
- Input: the string `s`.
- Output: `(length, count)`.

**Constraints**
- `1 ≤ |s| ≤ 10^6`

**Required signature (Go)**
```go
func LongestRegular(s string) (length int, count int)
```

**Examples**
```
LongestRegular(")((())))(()())")  → (6, 2)   // "((()))" and "(()())"
LongestRegular("))(")             → (0, 1)
LongestRegular("()(())")          → (6, 1)
```

**Instructions**
- O(n) time is required.
- Use a stack of indices or a DP. Explain why the count is correct, so that a long run is not also counted as its shorter pieces.

### Expected answer — Q18

**Approach.** Keep a stack of indices that starts with `−1`, the base just before the current valid run. For `(`, push its index. For `)`, pop. If the stack is now empty, this `)` is unmatched, so push its index as the new base. Otherwise, `i − top` is the length of the **longest** valid substring ending at `i`. Update `(best, count)`: a strictly longer length resets the count to 1, and an equal length adds 1. Each end position contributes only its longest valid substring, so the maximal runs are counted exactly once.

**Complexity.** O(n) time and O(n) space.

**Reference solution (Go)**

```go
package main

// LongestRegular uses a stack of indices. The bottom of the stack is always
// the index just before the current valid run (initially -1). When a ')'
// closes a '(' and the stack is not empty afterwards, i - top is the length
// of the longest valid substring ending at i. O(n) time, O(n) space.
func LongestRegular(s string) (length int, count int) {
	stack := []int{-1}
	best, cnt := 0, 1 // "(0, 1)" when nothing is valid
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i) // unmatched ')' becomes the new base
			continue
		}
		l := i - stack[len(stack)-1]
		if l > best {
			best, cnt = l, 1
		} else if l == best {
			cnt++
		}
	}
	return best, cnt
}
```

**Expected output:** `(6, 2)`, `(0, 1)`, `(6, 1)`. The reference also matched a brute-force O(n³) solution on 5000 random strings.

---

## Q19. Self-Organizing List Cost

*Source: original question from the transcript's set. Linked list: naive solution only.*

**Skills needed beyond linked lists:** A Fenwick tree (binary indexed tree) over a timeline of slots, hash maps (value → slot), the "move-to-front means a new, smaller timestamp" modelling trick, and 64-bit totals (up to 10^10).

**Statement**
A singly linked list of **distinct** integers starts in the given order. For each query value, find its node, add its 1-based position to a running total (the cost), and then move that node to the front of the list. Return the total cost.

**Input / Output**
- Input: the head of the list and the queries. Every queried value exists in the list.
- Output: the total cost. Only the total is returned. The final list order shown below is for illustration.

**Constraints**
- `1 ≤ n ≤ 10^5` nodes, `0 ≤ q ≤ 10^5` queries, `|Val| ≤ 10^9`
- The total cost can reach 10^10.

**Required signature (Go)**
```go
type Node struct {
	Val  int
	Next *Node
}

func MoveToFrontCost(head *Node, queries []int) int
```

**Examples**
```
4 → 7 → 9 → 2, queries [9 9 4 2 7]
   costs 3 + 1 + 2 + 4 + 4 → 14     (the list ends as 7 → 2 → 4 → 9)
5, queries [5 5 5]          → 3
1 → 2 → 3, queries []       → 0
```

**Instructions**
- The naive linked-list simulation (O(n·q)) passes small tests only. Write it first as a reference.
- Required: O((n + q) log n).
- The answer must not overflow.

### Expected answer — Q19

**Approach.** *Naive:* walk the list to find the value, add its position, and unlink and relink it at the head.
*Optimised:* create `q + n` slots. Put the initial list in slots `q+1 … q+n` and mark them in a Fenwick tree. The k-th move-to-front uses slot `q − k + 1`, which is to the left of every occupied slot, so it is the new front. A node's position is the number of occupied slots up to and including its slot, which is a prefix sum. For each query, add `prefix(slot[v])`, clear the old slot, mark the new front slot, and record it in the map.

**Complexity.** Naive is O(n·q). Optimised is O((n + q) log(n + q)) time and O(n + q) space.

**Reference solution (Go)**

```go
package main

type Node struct {
	Val  int
	Next *Node
}

// MoveToFrontCostNaive is the direct simulation on the linked list:
// O(n) per query, O(n*q) total. It also leaves the list reordered.
func MoveToFrontCostNaive(head *Node, queries []int) (int, *Node) {
	total := 0
	for _, v := range queries {
		var prev *Node
		pos := 1
		cur := head
		for cur.Val != v {
			prev, cur = cur, cur.Next
			pos++
		}
		total += pos
		if prev != nil { // move cur to the front
			prev.Next = cur.Next
			cur.Next = head
			head = cur
		}
	}
	return total, head
}

// MoveToFrontCost is the O((n+q) log n) version. Give every node a slot on a
// timeline of q+n slots: the initial list occupies slots q+1..q+n, and the
// k-th move-to-front uses slot q-k+1, which is left of everything so far.
// A Fenwick tree marks occupied slots, so a node's position in the list is
// the number of occupied slots up to and including its own.
func MoveToFrontCost(head *Node, queries []int) int {
	q := len(queries)
	slot := map[int]int{}
	size := q
	for c := head; c != nil; c = c.Next {
		size++
		slot[c.Val] = size
	}
	bit := make([]int, size+1)
	add := func(i, d int) {
		for ; i <= size; i += i & -i {
			bit[i] += d
		}
	}
	prefix := func(i int) int {
		s := 0
		for ; i > 0; i -= i & -i {
			s += bit[i]
		}
		return s
	}
	for i := q + 1; i <= size; i++ {
		add(i, 1)
	}
	total := 0 // up to 1e5 * 1e5 = 1e10: needs 64-bit (Go int is 64-bit)
	front := q
	for _, v := range queries {
		s := slot[v]
		total += prefix(s)
		add(s, -1)
		add(front, 1)
		slot[v] = front
		front--
	}
	return total
}
```

**Expected output:** `14` (the naive version also leaves `7 → 2 → 4 → 9`), `3`, `0`. The Fenwick version matched the naive version on 2000 random cases, and a 10^5 × 10^5 worst case gives 10,000,000,000, which confirms 64-bit arithmetic.

---
## Q20. Expiring LRU Cache

*Source: original question from the transcript's set. Linked list: **yes** (doubly, plus a hash map and a min-heap).*

**Skills needed beyond linked lists:** Hash maps, a min-heap (`container/heap`) keyed on expiry time, **lazy deletion** of stale heap entries, careful time semantics (an entry is valid while `now < expire`), and amortized analysis.

**Statement**
Design a cache with a fixed capacity in which every entry also has a time-to-live.

- `NewTTLCache(capacity)` creates the cache.
- `Put(key, value, ttl, now)` stores the entry, which is valid while `time < now + ttl`. If the key already exists, its value and expiry are replaced. `Put` counts as a use.
- `Get(key, now)` returns the value, or `−1` if the key is absent or expired. A hit counts as a use.
- When a **new** key is inserted into a full cache, first remove **all** expired entries. If the cache is still full, evict the least recently used entry.

`now` never decreases from one call to the next.

**Input / Output**
- Input: a sequence of calls with non-decreasing `now`.
- Output: `Get` returns the value or `−1`.

**Constraints**
- `1 ≤ capacity ≤ 3000`, up to 2·10^5 calls
- `0 ≤ key ≤ 10^5`, `0 ≤ value ≤ 10^9`, `1 ≤ ttl ≤ 10^9`, `0 ≤ now ≤ 10^9`

**Required signature (Go)**
```go
type TTLCache struct{ /* your fields */ }

func NewTTLCache(capacity int) *TTLCache
func (c *TTLCache) Put(key, value, ttl, now int)
func (c *TTLCache) Get(key, now int) int
```

**Example**
```
c := NewTTLCache(2)
c.Put(1, 10, 5, 0)      // key 1 valid until time 5 (exclusive)
c.Put(2, 20, 10, 1)     // key 2 valid until time 11
c.Get(1, 4)   → 10
c.Put(3, 30, 3, 5)      // full; key 1 has expired and is purged, so nothing is evicted
c.Get(1, 5)   → -1
c.Get(2, 6)   → 20
c.Put(4, 40, 10, 7)     // full and nothing has expired; evicts the LRU key, 3
c.Get(3, 7)   → -1
c.Get(4, 7)   → 40
c.Get(2, 12)  → -1      // expired at 11
```

**Instructions**
- O(log n) per operation is acceptable, amortized over the heap pops.
- Follow-up: `Get` in amortized O(1).
- Implement the list yourself. You may use `container/heap`.

### Expected answer — Q20

**Approach.** Use an LRU cache (map plus a doubly linked list with sentinels, as in Q7) together with a min-heap of `(expireAt, key)`.
- **Get:** if the key exists but `now ≥ expire`, remove it from the map and list and return `−1`. Its heap entry becomes stale and is skipped later. Otherwise, move the node to the front. This is O(1), which meets the follow-up.
- **Put on an existing key:** update the value and expiry, push a new heap entry (the old one becomes stale), and move the node to the front.
- **Put on a new key into a full cache:** pop heap entries with `expireAt ≤ now`. Remove an entry only if the key is still present **and** its current expiry equals the popped one (lazy deletion). If the cache is still full, evict `tail.prev`. Then insert at the front and push a heap entry.

**Complexity.** `Get` is O(1). `Put` is O(log H) plus the pops, and each heap entry is popped at most once, so `Put` is amortized O(log H), where H ≤ number of `Put` calls. Space is O(capacity + H).

**Reference solution (Go)**

```go
package main

import "container/heap"

type ttlNode struct {
	key, val, expire int // valid while now < expire
	prev, next       *ttlNode
}

// expiry heap entry; stale entries (key removed or expiry changed) are
// skipped when popped: lazy deletion.
type exp struct{ at, key int }
type expHeap []exp

func (h expHeap) Len() int           { return len(h) }
func (h expHeap) Less(i, j int) bool { return h[i].at < h[j].at }
func (h expHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *expHeap) Push(x any)        { *h = append(*h, x.(exp)) }
func (h *expHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// TTLCache = LRU cache (map + doubly linked list with sentinels)
// + min-heap of expiry times with lazy deletion.
type TTLCache struct {
	cap        int
	items      map[int]*ttlNode
	head, tail *ttlNode
	exps       expHeap
}

func NewTTLCache(capacity int) *TTLCache {
	h, t := &ttlNode{}, &ttlNode{}
	h.next, t.prev = t, h
	return &TTLCache{cap: capacity, items: map[int]*ttlNode{}, head: h, tail: t}
}

func (c *TTLCache) unlink(n *ttlNode) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *TTLCache) pushFront(n *ttlNode) {
	n.prev, n.next = c.head, c.head.next
	c.head.next.prev = n
	c.head.next = n
}

func (c *TTLCache) remove(n *ttlNode) {
	c.unlink(n)
	delete(c.items, n.key)
}

// purge removes every entry with expire <= now. O(k log n) for k heap pops.
func (c *TTLCache) purge(now int) {
	for len(c.exps) > 0 && c.exps[0].at <= now {
		e := heap.Pop(&c.exps).(exp)
		if n, ok := c.items[e.key]; ok && n.expire == e.at { // skip stale entries
			c.remove(n)
		}
	}
}

// Get is O(1): an expired hit is removed on the spot; its heap entry becomes stale.
func (c *TTLCache) Get(key, now int) int {
	n, ok := c.items[key]
	if !ok {
		return -1
	}
	if now >= n.expire {
		c.remove(n)
		return -1
	}
	c.unlink(n)
	c.pushFront(n)
	return n.val
}

func (c *TTLCache) Put(key, value, ttl, now int) {
	if n, ok := c.items[key]; ok {
		n.val, n.expire = value, now+ttl
		heap.Push(&c.exps, exp{n.expire, key}) // old entry becomes stale
		c.unlink(n)
		c.pushFront(n)
		return
	}
	if len(c.items) == c.cap {
		c.purge(now)
		if len(c.items) == c.cap {
			c.remove(c.tail.prev) // least recently used
		}
	}
	n := &ttlNode{key: key, val: value, expire: now + ttl}
	c.items[key] = n
	c.pushFront(n)
	heap.Push(&c.exps, exp{n.expire, key})
}
```

**Expected output:** `10`, `-1`, `20`, `-1`, `40`, `-1`. The reference also matched a simple slice-based model on 300 random sequences of 300 operations.

---

## Q21. Sereja and Brackets (Codeforces 380C)

*Source: Codeforces 380C (about 2000 rating). Linked list: no.*

**Skills needed beyond linked lists:** Segment trees with a non-trivial **associative merge** (matched pairs, unmatched `(`, unmatched `)`). The merge order matters, so the left and right parts of a query must be combined in order. You also need memory awareness for |s| = 10^6.

**Statement**
You are given a bracket string `s` and `m` queries `(l, r)`, 1-based and inclusive. For each query, return the length of the longest regular bracket **subsequence** of `s[l..r]`. A subsequence may skip characters.

**Input / Output**
- Input: `s` and the queries.
- Output: one answer per query, in order.

**Constraints**
- `1 ≤ |s| ≤ 10^6`, `1 ≤ m ≤ 10^5`, `1 ≤ l ≤ r ≤ |s|`

**Required signature (Go)**
```go
func BracketQueries(s string, queries [][2]int) []int
```

**Example**
```
s = "())(())(())("
(1,1)  → 0
(2,3)  → 0
(1,2)  → 2
(1,12) → 10
(8,12) → 4
(5,11) → 6
(2,10) → 6
```

**Instructions**
- Each query must take O(log n). Scanning each query range is O(n·m), which is too slow.
- Hint: each segment stores (matched pairs, unmatched `)`, unmatched `(`), and two segments merge in O(1).

### Expected answer — Q21

**Approach.** For any substring, greedy matching leaves a summary `(pairs, open, close)`: the matched pairs, the unmatched `(` on the right, and the unmatched `)` on the left. When L is followed by R, `t = min(L.open, R.close)` new pairs form, so `pairs = L.pairs + R.pairs + t`, `open = L.open + R.open − t`, and `close = L.close + R.close − t`. Build a segment tree over these summaries. For a query, combine the covering nodes **in left-to-right order**. The iterative tree keeps separate left and right accumulators for this. The answer is `2 · pairs`.

**Complexity.** O(n) build and O(log n) per query, so O(n + m log n) in total. Space is O(n), with 2n nodes of three `int32` values (about 24 MB at n = 10^6).

**Reference solution (Go)**

```go
package main

// seg summarises a substring after greedily matching brackets inside it.
type seg struct {
	pairs int32 // matched pairs
	open  int32 // unmatched '('
	close int32 // unmatched ')'
}

// merge(L, R): L's unmatched '(' can match R's unmatched ')'.
func merge(l, r seg) seg {
	t := min(l.open, r.close)
	return seg{l.pairs + r.pairs + t, l.open + r.open - t, l.close + r.close - t}
}

// BracketQueries answers each 1-based query (l, r) with the length of the
// longest regular subsequence of s[l..r]. Iterative segment tree:
// O(n) build, O(log n) per query.
func BracketQueries(s string, queries [][2]int) []int {
	n := len(s)
	t := make([]seg, 2*n)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			t[n+i] = seg{open: 1}
		} else {
			t[n+i] = seg{close: 1}
		}
	}
	for i := n - 1; i > 0; i-- {
		t[i] = merge(t[2*i], t[2*i+1])
	}
	res := make([]int, len(queries))
	for qi, q := range queries {
		var left, right seg // order matters: left part, then right part
		lo, hi := q[0]-1+n, q[1]+n // half-open [lo, hi)
		for lo < hi {
			if lo&1 == 1 {
				left = merge(left, t[lo])
				lo++
			}
			if hi&1 == 1 {
				hi--
				right = merge(t[hi], right)
			}
			lo >>= 1
			hi >>= 1
		}
		res[qi] = 2 * int(merge(left, right).pairs)
	}
	return res
}
```

**Expected output:** `[0 0 2 10 4 6 6]`. The reference also matched a greedy per-query scan on 2000 random strings (10,000 queries).

---

## Q22. Circular Range Min with Range Add (Codeforces 52C)

*Source: Codeforces 52C "Circular RMQ" (about 2200 rating). Linked list: no.*

**Skills needed beyond linked lists:** A segment tree with **lazy propagation** (range add, range minimum), splitting a wrapped circular range into two linear ranges, and `int64` values (sums can reach about 2·10^11).

**Statement**
You are given a circular array `a` of length `n`, indexed `0 … n−1`. Support two operations:
- `Inc(l, r, v)`: add `v` to every element from `l` to `r`, inclusive.
- `Min(l, r)`: return the minimum over that range.

If `l > r`, the range wraps around the end of the array: `l, l+1, …, n−1, 0, …, r`. If `l == r`, it is the single element.

**Input / Output**
- Input: the initial array and a list of operations, where `Op{Inc: true, L, R, V}` is an increment and `Op{L, R}` is a minimum query.
- Output: the answers to the `Min` operations, in order.

**Constraints**
- `1 ≤ n ≤ 2·10^5`, `1 ≤ m ≤ 2·10^5` operations
- `−10^6 ≤ a[i], v ≤ 10^6`

**Required signature (Go)**
```go
type Op struct {
	Inc  bool  // true: Inc(L, R, V); false: Min(L, R)
	L, R int
	V    int64
}

func Process(a []int64, ops []Op) []int64
```

**Example**
```
a = [1 2 3 4]
Inc(3, 0, -1)   // wraps: indices 3, 0  → a = [0 2 3 3]
Min(0, 1)  → 0
Min(2, 1)  → 0  // wraps: indices 2, 3, 0, 1
Min(1, 3)  → 2
```
Original Codeforces sample: `a = [1 2 3 4]`, then `Min(3,0)`, `Inc(3,0,−1)`, `Min(0,1)`, `Min(2,1)` gives `1 0 0`.

**Instructions**
- O(log n) per operation is required.
- State why `int32` is not safe here.

### Expected answer — Q22

**Approach.** Build a segment tree in which `mn[x]` is the minimum of node x's range **including** the pending addition `add[x]` stored at that node. When an update covers a node completely, add `v` to both `mn[x]` and `add[x]` and stop. On the way back up, set `mn[x] = min(children) + add[x]`. A query adds the `add[x]` values of the nodes it passes through. This is the "never push down" form of lazy propagation. The usual push-down form is equally fine. Split a wrapped range into `[l, n−1]` and `[0, r]`.

**Complexity.** O(n) build, O(log n) per operation, and O(n) space.

**Reference solution (Go)**

```go
package main

import "math"

type Op struct {
	Inc  bool  // true: Inc(L, R, V); false: Min(L, R)
	L, R int   // 0-based; L > R means the range wraps around the end
	V    int64 // only used by Inc
}

// Segment tree for range add / range min. add[x] is a pending addition that
// applies to the whole node; mn[x] already includes add[x] (this "lazy tag
// that is never pushed down" variant needs no push function).
type segTree struct {
	n       int
	mn, add []int64
}

func (t *segTree) build(x, lo, hi int, a []int64) {
	if lo == hi {
		t.mn[x] = a[lo]
		return
	}
	mid := (lo + hi) / 2
	t.build(2*x, lo, mid, a)
	t.build(2*x+1, mid+1, hi, a)
	t.mn[x] = min(t.mn[2*x], t.mn[2*x+1])
}

func (t *segTree) update(x, lo, hi, l, r int, v int64) {
	if r < lo || hi < l {
		return
	}
	if l <= lo && hi <= r {
		t.mn[x] += v
		t.add[x] += v
		return
	}
	mid := (lo + hi) / 2
	t.update(2*x, lo, mid, l, r, v)
	t.update(2*x+1, mid+1, hi, l, r, v)
	t.mn[x] = min(t.mn[2*x], t.mn[2*x+1]) + t.add[x]
}

func (t *segTree) query(x, lo, hi, l, r int) int64 {
	if r < lo || hi < l {
		return math.MaxInt64
	}
	if l <= lo && hi <= r {
		return t.mn[x]
	}
	mid := (lo + hi) / 2
	return min(t.query(2*x, lo, mid, l, r), t.query(2*x+1, mid+1, hi, l, r)) + t.add[x]
}

// Process runs the operations and returns the answers to the Min operations.
// A wrapped range [L..n-1] + [0..R] is split into two ordinary ranges.
// O((n + m) log n) time, O(n) space. Values can reach about 2e11, so int64.
func Process(a []int64, ops []Op) []int64 {
	n := len(a)
	t := &segTree{n: n, mn: make([]int64, 4*n), add: make([]int64, 4*n)}
	t.build(1, 0, n-1, a)
	var out []int64
	for _, op := range ops {
		var ranges [][2]int
		if op.L <= op.R {
			ranges = [][2]int{{op.L, op.R}}
		} else {
			ranges = [][2]int{{op.L, n - 1}, {0, op.R}}
		}
		if op.Inc {
			for _, rg := range ranges {
				t.update(1, 0, n-1, rg[0], rg[1], op.V)
			}
		} else {
			best := int64(math.MaxInt64)
			for _, rg := range ranges {
				best = min(best, t.query(1, 0, n-1, rg[0], rg[1]))
			}
			out = append(out, best)
		}
	}
	return out
}
```

**Expected output:** `[0 0 2]` for the example and `[1 0 0]` for the Codeforces sample. The reference also matched a brute-force simulation on 1000 random arrays (20 operations each, with many wrapped ranges).

---

## Verification notes

Every reference solution in this file was saved as a Go program on the build machine and run with Go 1.27.3 (`go vet` clean). Each program checks the reference against every example in its question. Where a brute-force solution exists, it also compares the reference with that brute force on random inputs.

| Q | Example checks | Extra randomized / stress check | Result |
|---|---|---|---|
| 1 | 6 | — | Pass |
| 2 | 3 (+ 10^16 case) | 2000 vs O(n²) brute force | Pass |
| 3 | 3 (+ overflow case) | — | Pass |
| 4 | 4 lists × 3 properties | — | Pass |
| 5 | 3 | 2000 vs brute force | Pass |
| 6 | 3 | — | Pass |
| 7 | 5 | — | Pass |
| 8 | 13 | 200 sequences vs slice model | Pass |
| 9 | Part 1: 9; Part 2: 19 | Buggy original confirmed to hang on the failing test | Pass |
| 10 | 4 | — | Pass |
| 11 | 3 | — | Pass |
| 12 | 4 (+ k = 10^9) | 2000 vs circular-list simulation | Pass |
| 13 | 4 | 2000 vs brute force | Pass |
| 14 | 11 | 20,000 MergeBase queries vs naive | Pass |
| 15 | 4 | 5000 vs exhaustive search | Pass |
| 16 | 3 | 3000 vs O(n²) DP + chain validation | Pass |
| 17 | 5 | 5000 vs day-by-day simulation | Pass |
| 18 | 3 | 5000 vs brute force | Pass |
| 19 | 5 | 2000 Fenwick vs naive; 10^5 × 10^5 worst case | Pass |
| 20 | 6 | 300 sequences vs slice model | Pass |
| 21 | 7 | 2000 strings vs greedy scan | Pass |
| 22 | 3 + CF sample | 1000 vs brute force | Pass |

**Corrections to the transcript**
- **Poisonous Plants (Q17):** the transcript gives `[1,2,3,4] → 3`. The correct answer is **1**, because 2, 3 and 4 each have more pesticide than their left neighbour and all die on day 1. This matches HackerRank's rules and was confirmed by simulation.
- **Mysterious Present (Q16):** the transcript gives `[1]` for the first example. `[2]` is equally valid, and the original problem accepts any longest chain, so grade by checking that the chain is valid.
- **Debug and Extend (Q9):** the explicit rule "a successful Unpin counts as a use" was added, as the transcript's own design note recommends, so that the O(1) requirement can be met. The transcript's example is unchanged under this rule. One `List()` line was added after `Unpin("a")` to make the rule visible.
- **Circular RMQ (Q22):** the signature uses `int64` instead of `int` to make the overflow requirement explicit.
