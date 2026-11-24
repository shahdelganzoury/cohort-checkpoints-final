## dealcards

### Instructions

Given the `head` of a singly linked list and an integer `k`, deal its nodes into `k` hands the way you deal cards: the 1st node goes to hand 0, the 2nd to hand 1, and so on, wrapping back to hand 0 after hand `k-1`. Return the heads of the hands as a slice.

Each hand keeps its nodes in their **original relative order**. You must **reuse the existing nodes**: do not create new ones. A hand that receives no node is `nil`, and if `k <= 0` the function returns `nil`.

Example 1:

```
Input: head = [1,2,3,4,5,6,7], k = 3
Output: [[1,4,7], [2,5], [3,6]]
Explanation: Nodes are dealt 1→hand 0, 2→hand 1, 3→hand 2, 4→hand 0, and so on.
```

Example 2:

```
Input: head = [10,20], k = 4
Output: [[10], [20], nil, nil]
Explanation: There are more hands than nodes, so the last two hands stay empty.
```

Example 3:

```
Input: head = [], k = 2
Output: [nil, nil]
Explanation: An empty list gives empty hands.
```

Constraints:

* The number of nodes is in the range `[0, 10000]`.
* `-1000 <= k <= 10000`
* Follow up: solve it in `O(n)` time without allocating any node.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Val  int
	Next *ListNode
}
```

### Expected function

```go
func DealCards(head *ListNode, k int) []*ListNode {

}
```

### Usage

Here is a possible program to test your function (it assumes the structure above is declared in your solution file):

```go
package main

import (
	"fmt"
)

func main() {
	show(DealCards(buildList(1, 2, 3, 4, 5, 6, 7), 3))
	show(DealCards(buildList(10, 20), 4))
	show(DealCards(buildList(5, 6, 7), 1))
	show(DealCards(nil, 2))
	show(DealCards(buildList(1, 2), 0))
}

func buildList(vals ...int) *ListNode {
	var head *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		head = &ListNode{Val: vals[i], Next: head}
	}
	return head
}

func show(hands []*ListNode) {
	if hands == nil {
		fmt.Println("no hands")
		return
	}
	for i, h := range hands {
		fmt.Printf("hand %d:", i)
		for ; h != nil; h = h.Next {
			fmt.Print(" ", h.Val)
		}
		fmt.Println()
	}
}
```

And its output:

```console
$ go run . | cat -e
hand 0: 1 4 7$
hand 1: 2 5$
hand 2: 3 6$
hand 0: 10$
hand 1: 20$
hand 2:$
hand 3:$
hand 0: 5 6 7$
hand 0:$
hand 1:$
no hands$
```

### Reference solution

**Level:** 1

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func DealCards(head *ListNode, k int) []*ListNode {
	if k <= 0 {
		return nil
	}
	heads := make([]*ListNode, k)
	tails := make([]*ListNode, k)
	i := 0
	for head != nil {
		next := head.Next
		head.Next = nil
		if tails[i] == nil {
			heads[i] = head
		} else {
			tails[i].Next = head
		}
		tails[i] = head
		head = next
		i = (i + 1) % k
	}
	return heads
}
```

**Approach:** Keep a head and a tail pointer per hand. Detach each node, append it to the tail of hand `i`, then move `i` to `(i+1) % k`.

**Complexity:** O(n) time, O(k) extra space for the hand pointers.

---

## crossover

### Instructions

Two chromosomes are stored as doubly linked lists of genes (`Val` is a letter). Given their heads `a` and `b`, cut `a` after its first `i` nodes and `b` after its first `j` nodes, then **swap the two tails**. Return the new heads of `a` and `b`, in that order.

If `i` is `0`, the whole of `a` is its tail. If `i` is at least the length of `a`, its tail is empty. The same goes for `j` and `b`.

After the swap **every `Prev` pointer must be correct**, and each head's `Prev` must be `nil`. Do not create new nodes.

Example 1:

```
Input: a = "AAAA", b = "bbbbb", i = 2, j = 1
Output: a = "AAbbbb", b = "bAA"
Explanation: a keeps "AA" and takes b's tail "bbbb"; b keeps "b" and takes a's tail "AA".
```

Example 2:

```
Input: a = "ABC", b = "xyz", i = 0, j = 3
Output: a = "" (nil), b = "xyzABC"
Explanation: All of a is its tail and b's tail is empty, so a becomes empty.
```

Example 3:

```
Input: a = "ABC", b = "xy", i = 10, j = 1
Output: a = "ABCy", b = "x"
Explanation: i is past the end of a, so a gives away an empty tail.
```

Constraints:

* The number of nodes in each list is in the range `[0, 10000]`.
* `0 <= i, j <= 20000`
* Follow up: solve it in `O(i + j)` time and `O(1)` memory.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Val  byte
	Prev *ListNode
	Next *ListNode
}
```

### Expected function

```go
func Crossover(a, b *ListNode, i, j int) (*ListNode, *ListNode) {

}
```

### Usage

Here is a possible program to test your function (it assumes the structure above is declared in your solution file):

```go
package main

import (
	"fmt"
)

func main() {
	a, b := Crossover(buildList("AAAA"), buildList("bbbbb"), 2, 1)
	fmt.Println(show(a), show(b))
	a, b = Crossover(buildList("ABC"), buildList("xyz"), 0, 3)
	fmt.Println(show(a), show(b))
	a, b = Crossover(buildList("ABC"), buildList("xy"), 10, 1)
	fmt.Println(show(a), show(b))
	a, b = Crossover(nil, buildList("xy"), 0, 1)
	fmt.Println(show(a), show(b))
	a, b = Crossover(buildList("Q"), nil, 1, 0)
	fmt.Println(show(a), show(b))
}

func buildList(s string) *ListNode {
	var head, tail *ListNode
	for i := 0; i < len(s); i++ {
		n := &ListNode{Val: s[i], Prev: tail}
		if tail == nil {
			head = n
		} else {
			tail.Next = n
		}
		tail = n
	}
	return head
}

func show(h *ListNode) string {
	fwd, back := "", ""
	var last *ListNode
	for ; h != nil; h = h.Next {
		fwd += string(h.Val)
		last = h
	}
	for ; last != nil; last = last.Prev {
		back += string(last.Val)
	}
	return fmt.Sprintf("%q (back %q)", fwd, back)
}
```

And its output:

```console
$ go run . | cat -e
"AAbbbb" (back "bbbbAA") "bAA" (back "AAb")$
"" (back "") "xyzABC" (back "CBAzyx")$
"ABCy" (back "yCBA") "x" (back "x")$
"y" (back "y") "x" (back "x")$
"Q" (back "Q") "" (back "")$
```

### Reference solution

**Level:** 2

```go
package main

type ListNode struct {
	Val  byte
	Prev *ListNode
	Next *ListNode
}

// cut returns the head of the tail that starts after the first n nodes,
// and detaches it from the part before it.
func cut(head *ListNode, n int) (*ListNode, *ListNode) {
	if n <= 0 {
		return nil, head
	}
	cur := head
	for i := 1; i < n && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head, nil
	}
	tail := cur.Next
	cur.Next = nil
	if tail != nil {
		tail.Prev = nil
	}
	return head, tail
}

func join(front, back *ListNode) *ListNode {
	if front == nil {
		return back
	}
	last := front
	for last.Next != nil {
		last = last.Next
	}
	last.Next = back
	if back != nil {
		back.Prev = last
	}
	return front
}

func Crossover(a, b *ListNode, i, j int) (*ListNode, *ListNode) {
	aFront, aTail := cut(a, i)
	bFront, bTail := cut(b, j)
	return join(aFront, bTail), join(bFront, aTail)
}
```

**Approach:** Walk `i` (and `j`) nodes to find each cut point and detach the tail, clearing the tail's `Prev` and the cut node's `Next`. Then attach `b`'s tail to `a`'s front and `a`'s tail to `b`'s front, fixing the joining `Prev`.

**Complexity:** O(i + j) time (only the kept fronts are walked), O(1) extra space.

---
## hopelimination

### Instructions

The nodes form a **circular** singly linked list, and `start` is one of them. Starting at `start`, repeat until the circle is empty: remove the current node and record its `Val`, then move `Val` steps forward around the **remaining** circle. Step 1 is the node right after the removed one, and the node you land on becomes the current node.

Return the recorded values in removal order. An empty circle (`nil`) returns an empty slice.

Values can be very large, so walking `Val` steps one by one is too slow: **take it modulo the current circle size**.

Example 1:

```
Input: circle = [3,1,4,1,5], start at 3
Output: [3,1,5,1,4]
Explanation: Remove 3 and hop 3 steps (1, 4, 1) to the second 1. Remove it and hop 1 step to 5. Remove 5 and hop 5 steps around [1,4] to land on 1. Then remove 1, then 4.
```

Example 2:

```
Input: circle = [1000000000,7,8]
Output: [1000000000,8,7]
Explanation: After removing the first node, 2 nodes remain, so the hop is 1000000000 mod 2 → 2 steps, landing on 8.
```

Example 3:

```
Input: circle = [9]
Output: [9]
Explanation: A single node is removed and the circle is empty.
```

Constraints:

* The number of nodes is in the range `[0, 5000]`.
* `1 <= Node.Val <= 1000000000`
* Follow up: never walk more than `size - 1` steps per hop.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Val  int
	Next *ListNode
}
```

### Expected function

```go
func HopEliminate(start *ListNode) []int {

}
```

### Usage

Here is a possible program to test your function (it assumes the structure above is declared in your solution file):

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(HopEliminate(ring(3, 1, 4, 1, 5)))
	fmt.Println(HopEliminate(ring(2, 2, 2, 2)))
	fmt.Println(HopEliminate(ring(1000000000, 7, 8)))
	fmt.Println(HopEliminate(ring(9)))
	fmt.Println(HopEliminate(nil))
}

func ring(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	tail := head
	for _, v := range vals[1:] {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	tail.Next = head
	return head
}
```

And its output:

```console
$ go run . | cat -e
[3 1 5 1 4]$
[2 2 2 2]$
[1000000000 8 7]$
[9]$
[]$
```

### Reference solution

**Level:** 3

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func HopEliminate(start *ListNode) []int {
	result := []int{}
	if start == nil {
		return result
	}
	size := 1
	prev := start
	for prev.Next != start {
		prev = prev.Next
		size++
	}
	cur := start
	for size > 0 {
		result = append(result, cur.Val)
		prev.Next = cur.Next // unlink cur
		size--
		if size == 0 {
			break
		}
		steps := (cur.Val-1)%size + 1 // step 1 = node right after cur
		cur.Next = nil
		cur = prev.Next
		for i := 1; i < steps; i++ {
			prev = cur
			cur = cur.Next
		}
	}
	return result
}
```

**Approach:** Count the circle once and find the node before `start`. Each round, unlink `cur` through `prev`, shrink the size, then walk `(Val-1) % size` extra steps from `prev.Next`, keeping `prev` one behind.

**Complexity:** O(n^2) worst-case time (each hop is under n steps), O(1) extra space besides the result.

---
## intervallist

### Instructions

Each node holds a closed interval `[Start, End]`. The list is kept **sorted by `Start`**, and no two of its intervals overlap. Given the `head` and a new interval `[start, end]`, add it and **merge every interval it overlaps on the fly**. Return the head.

Intervals that touch also merge: `[1,3]` and `[3,5]` become `[1,5]`, but `[1,2]` and `[3,4]` stay apart. One new interval can swallow many existing ones.

Create **at most one new node per call**, and unlink the nodes you merge away. If `start > end`, return the list unchanged.

Example 1:

```
Input: head = [[1,3],[5,6],[10,12]], add [3,4]
Output: [[1,4],[5,6],[10,12]]
Explanation: [3,4] touches [1,3], so they merge into [1,4].
```

Example 2:

```
Input: head = [[1,4],[5,6],[7,7],[8,9],[10,12]], add [2,11]
Output: [[1,12]]
Explanation: [2,11] overlaps every interval, so they all collapse into one.
```

Example 3:

```
Input: head = [], add [4,4]
Output: [[4,4]]
Explanation: The new interval becomes the only node.
```

Constraints:

* The number of nodes is in the range `[0, 10000]`.
* `-1000000000 <= start, end <= 1000000000`
* Follow up: one pass, `O(n)` time per call and `O(1)` extra memory.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Start int
	End   int
	Next  *ListNode
}
```

### Expected function

```go
func AddInterval(head *ListNode, start, end int) *ListNode {

}
```

### Usage

Here is a possible program to test your function (it assumes the structure above is declared in your solution file):

```go
package main

import (
	"fmt"
)

func main() {
	var list *ListNode
	steps := [][2]int{{10, 12}, {1, 3}, {5, 6}, {3, 4}, {8, 9}, {7, 7}, {2, 11}, {20, 25}, {13, 13}, {9, 1}}
	for _, s := range steps {
		list = AddInterval(list, s[0], s[1])
		fmt.Printf("add [%d,%d]: ", s[0], s[1])
		show(list)
	}
	show(AddInterval(nil, 4, 4))
}

func show(h *ListNode) {
	for ; h != nil; h = h.Next {
		fmt.Printf("[%d,%d] -> ", h.Start, h.End)
	}
	fmt.Println("nil")
}
```

And its output:

```console
$ go run . | cat -e
add [10,12]: [10,12] -> nil$
add [1,3]: [1,3] -> [10,12] -> nil$
add [5,6]: [1,3] -> [5,6] -> [10,12] -> nil$
add [3,4]: [1,4] -> [5,6] -> [10,12] -> nil$
add [8,9]: [1,4] -> [5,6] -> [8,9] -> [10,12] -> nil$
add [7,7]: [1,4] -> [5,6] -> [7,7] -> [8,9] -> [10,12] -> nil$
add [2,11]: [1,12] -> nil$
add [20,25]: [1,12] -> [20,25] -> nil$
add [13,13]: [1,12] -> [13,13] -> [20,25] -> nil$
add [9,1]: [1,12] -> [13,13] -> [20,25] -> nil$
[4,4] -> nil$
```

### Reference solution

**Level:** 4

```go
package main

type ListNode struct {
	Start int
	End   int
	Next  *ListNode
}

func AddInterval(head *ListNode, start, end int) *ListNode {
	if start > end {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	// skip intervals that end before the new one starts
	for prev.Next != nil && prev.Next.End < start {
		prev = prev.Next
	}
	// swallow every interval that overlaps [start, end]
	cur := prev.Next
	for cur != nil && cur.Start <= end {
		if cur.Start < start {
			start = cur.Start
		}
		if cur.End > end {
			end = cur.End
		}
		next := cur.Next
		cur.Next = nil
		cur = next
	}
	prev.Next = &ListNode{Start: start, End: end, Next: cur}
	return dummy.Next
}
```

**Approach:** With a sentinel, skip intervals that end before `start`. Then swallow every interval whose `Start <= end`, widening `[start, end]`, and splice one node holding the merged range in their place.

**Complexity:** O(n) time per call, O(1) extra space.

---

## copyrandomlist

### Instructions

Each node has a `Next` pointer and a `Random` pointer, which may point to any node in the list or be `nil`. Return a **deep copy**: brand-new nodes with the same values, whose `Next` and `Random` pointers form the same shape but point **only to nodes in the copy**, never to the original.

Values may repeat, so you cannot match nodes by value. The original list must be **unchanged** when the function returns.

Example 1:

```
Input: [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
Explanation: Each node is written as [Val, index of its Random target]. The copy has the same shape but brand-new nodes.
```

Example 2:

```
Input: [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]
Explanation: All values are equal, so Random targets must be matched by node, not by value.
```

Example 3:

```
Input: []
Output: []
Explanation: An empty list copies to nil.
```

Constraints:

* The number of nodes is in the range `[0, 1000]`.
* `-10000 <= Node.Val <= 10000`
* `Random` is `nil` or points to a node in the same list.
* Follow up: `O(n)` time and `O(1)` extra memory (no map), restoring the original list.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}
```

### Expected function

```go
func CopyRandomList(head *ListNode) *ListNode {

}
```

### Usage

Here is a possible program to test your function (it assumes the structure above is declared in your solution file):

```go
package main

import (
	"fmt"
)

func main() {
	tests := [][][2]int{
		{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		{{1, 1}, {2, 1}},
		{{3, -1}, {3, 0}, {3, -1}},
		{{5, 0}},
		{},
	}
	for _, t := range tests {
		orig := buildList(t)
		before := encode(orig)
		cp := CopyRandomList(orig)
		fmt.Println(encode(cp), encode(orig) == before, shares(orig, cp))
	}
}

// buildList takes [value, index of Random or -1] pairs.
func buildList(spec [][2]int) *ListNode {
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
		return nil
	}
	return nodes[0]
}

func encode(head *ListNode) string {
	index := map[*ListNode]int{}
	i := 0
	for c := head; c != nil; c = c.Next {
		index[c] = i
		i++
	}
	s := "["
	for c := head; c != nil; c = c.Next {
		r := "null"
		if c.Random != nil {
			r = fmt.Sprint(index[c.Random])
		}
		s += fmt.Sprintf("[%d,%s]", c.Val, r)
	}
	return s + "]"
}

func shares(a, b *ListNode) bool {
	seen := map[*ListNode]bool{}
	for c := a; c != nil; c = c.Next {
		seen[c] = true
	}
	for c := b; c != nil; c = c.Next {
		if seen[c] || (c.Random != nil && seen[c.Random]) {
			return true
		}
	}
	return false
}
```

And its output:

```console
$ go run . | cat -e
[[7,null][13,0][11,4][10,2][1,0]] true false$
[[1,1][2,1]] true false$
[[3,null][3,0][3,null]] true false$
[[5,0]] true false$
[] true false$
```

### Reference solution

**Level:** 5

```go
package main

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

func CopyRandomList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 1. A -> A' -> B -> B' -> ...
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &ListNode{Val: cur.Val, Next: cur.Next}
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

**Approach:** Weave each copy right after its original (`A -> A' -> B -> B'`), set `X'.Random = X.Random.Next`, then unweave the two lists, which also restores the original `Next` pointers. A map from original to copy is the simpler O(n)-space alternative.

**Complexity:** O(n) time, O(1) extra space besides the copy.

---
## multiplypolynomials

### Instructions

A polynomial is a linked list of terms sorted by `Exp` in **descending** order, with no duplicate exponents and no zero coefficients. Given two polynomials `a` and `b`, return their product as a **new** list in the same form.

Terms whose coefficients cancel to `0` **must not appear**. If either input is `nil`, the product is zero: return `nil`. Do not modify the input lists.

Example 1:

```
Input: a = [(2,2),(3,0)], b = [(1,1),(1,0)]
Output: [(2,3),(2,2),(3,1),(3,0)]
Explanation: (2x^2 + 3)(x + 1) = 2x^3 + 2x^2 + 3x + 3. Each term is written as (Coef,Exp).
```

Example 2:

```
Input: a = [(1,1),(1,0)], b = [(1,1),(-1,0)]
Output: [(1,2),(-1,0)]
Explanation: (x + 1)(x - 1): the x terms cancel, so they are left out.
```

Example 3:

```
Input: a = nil, b = [(5,3)]
Output: nil
Explanation: Anything times the zero polynomial is zero.
```

Constraints:

* Each polynomial has at most `200` terms.
* `0 <= Exp <= 5000`, `1 <= |Coef| <= 100`
* Follow up: `O(|a| * |b| + E)` time, where `E` is the largest result exponent.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Coef int
	Exp  int
	Next *ListNode
}
```

### Expected function

```go
func MultiplyPoly(a, b *ListNode) *ListNode {

}
```

### Usage

Here is a possible program to test your function (it assumes the structure above is declared in your solution file):

```go
package main

import (
	"fmt"
)

func main() {
	a := buildPoly(2, 2, 3, 0)
	b := buildPoly(1, 1, 1, 0)
	fmt.Println(show(MultiplyPoly(a, b)))
	fmt.Println(show(a), show(b))
	fmt.Println(show(MultiplyPoly(buildPoly(1, 1, 1, 0), buildPoly(1, 1, -1, 0))))
	fmt.Println(show(MultiplyPoly(nil, buildPoly(5, 3))))
	fmt.Println(show(MultiplyPoly(buildPoly(3, 5000), buildPoly(-2, 5000, 4, 0))))
	fmt.Println(show(MultiplyPoly(buildPoly(7, 0), buildPoly(-1, 0))))
	fmt.Println(show(MultiplyPoly(buildPoly(1, 2, 1, 1), buildPoly(1, 1, -1, 0))))
}

// buildPoly takes coef, exp pairs in descending exponent order.
func buildPoly(pairs ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for i := 0; i+1 < len(pairs); i += 2 {
		tail.Next = &ListNode{Coef: pairs[i], Exp: pairs[i+1]}
		tail = tail.Next
	}
	return dummy.Next
}

func show(p *ListNode) string {
	if p == nil {
		return "nil"
	}
	s := "["
	for ; p != nil; p = p.Next {
		s += fmt.Sprintf("(%d,%d)", p.Coef, p.Exp)
	}
	return s + "]"
}
```

And its output:

```console
$ go run . | cat -e
[(2,3)(2,2)(3,1)(3,0)]$
[(2,2)(3,0)] [(1,1)(1,0)]$
[(1,2)(-1,0)]$
nil$
[(-6,10000)(12,5000)]$
[(-7,0)]$
[(1,3)(-1,1)]$
```

### Reference solution

**Level:** 6

```go
package main

type ListNode struct {
	Coef int
	Exp  int
	Next *ListNode
}

func MultiplyPoly(a, b *ListNode) *ListNode {
	if a == nil || b == nil {
		return nil
	}
	acc := make([]int, a.Exp+b.Exp+1) // heads hold the largest exponents
	for x := a; x != nil; x = x.Next {
		for y := b; y != nil; y = y.Next {
			acc[x.Exp+y.Exp] += x.Coef * y.Coef
		}
	}
	dummy := &ListNode{}
	tail := dummy
	for e := len(acc) - 1; e >= 0; e-- {
		if acc[e] != 0 { // cancelled terms are skipped
			tail.Next = &ListNode{Coef: acc[e], Exp: e}
			tail = tail.Next
		}
	}
	return dummy.Next
}
```

**Approach:** The heads hold the largest exponents, so the top product exponent is `a.Exp + b.Exp`. Accumulate every `x.Coef * y.Coef` into a slice slot `x.Exp + y.Exp`. Then build the result from high to low with a dummy head, skipping zero slots.

**Complexity:** O(|a|·|b| + E) time, O(E) extra space, where E <= 10000.

---
