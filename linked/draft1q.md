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

## transposeaccess

### Instructions

A self-organizing list follows the **transpose rule**: each time a value is accessed, its node swaps places with the node just before it, so popular values drift slowly toward the front.

Given the `head` of a singly linked list and a value `val`, find the first node holding `val`. Return the new head and that node's **1-based position before the swap**, then move the node one place toward the front.

You must swap by **relinking nodes**: changing `Val` fields is not allowed. If the node is already the head, nothing moves. If `val` is not in the list, return the list unchanged and `-1`.

Example 1:

```
Input: head = [4,7,9,2], val = 2
Output: position 4, list [4,7,2,9]
Explanation: 2 was 4th, and it swaps with 9, the node before it.
```

Example 2:

```
Input: head = [2,4,7,9], val = 2
Output: position 1, list [2,4,7,9]
Explanation: 2 is already the head, so nothing moves.
```

Example 3:

```
Input: head = [2,4,7,9], val = 5
Output: position -1, list [2,4,7,9]
Explanation: 5 is not in the list.
```

Constraints:

* The number of nodes is in the range `[0, 10000]`.
* `-100000 <= Node.Val, val <= 100000`
* Follow up: solve it in `O(n)` time and `O(1)` memory, without special-casing the head (use a sentinel).

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
func Access(head *ListNode, val int) (*ListNode, int) {

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
	head := buildList(4, 7, 9, 2)
	var pos int
	for _, q := range []int{2, 2, 2, 2, 5, 9} {
		head, pos = Access(head, q)
		fmt.Printf("Access(%d) = %d : %s\n", q, pos, str(head))
	}
	single := buildList(8)
	single, pos = Access(single, 8)
	fmt.Println(pos, str(single))
	var empty *ListNode
	empty, pos = Access(empty, 1)
	fmt.Println(pos, str(empty))
}

func buildList(vals ...int) *ListNode {
	var head *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		head = &ListNode{Val: vals[i], Next: head}
	}
	return head
}

func str(h *ListNode) string {
	s := ""
	for ; h != nil; h = h.Next {
		s += fmt.Sprint(h.Val) + " -> "
	}
	return s + "nil"
}
```

And its output:

```console
$ go run . | cat -e
Access(2) = 4 : 4 -> 7 -> 2 -> 9 -> nil$
Access(2) = 3 : 4 -> 2 -> 7 -> 9 -> nil$
Access(2) = 2 : 2 -> 4 -> 7 -> 9 -> nil$
Access(2) = 1 : 2 -> 4 -> 7 -> 9 -> nil$
Access(5) = -1 : 2 -> 4 -> 7 -> 9 -> nil$
Access(9) = 4 : 2 -> 4 -> 9 -> 7 -> nil$
1 8 -> nil$
-1 nil$
```

### Reference solution

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Access(head *ListNode, val int) (*ListNode, int) {
	dummy := &ListNode{Next: head}
	pp := dummy // node before prev
	pos := 1
	if head == nil {
		return head, -1
	}
	if head.Val == val {
		return head, 1
	}
	for pp.Next.Next != nil {
		prev := pp.Next
		cur := prev.Next
		pos++
		if cur.Val == val {
			prev.Next = cur.Next
			cur.Next = prev
			pp.Next = cur
			return dummy.Next, pos
		}
		pp = prev
	}
	return head, -1
}
```

**Approach:** Use a sentinel before the head and walk with three pointers (`pp -> prev -> cur`). When `cur` matches, relink `pp -> cur -> prev -> cur.Next`. The sentinel removes the head special case.

**Complexity:** O(n) time, O(1) extra space.

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

## paperfold

### Instructions

Think of a singly linked list as a strip of paper. Given its `head` and an integer `k`, fold the strip right after its `k`-th node, so that the part after the crease flips back over the first `k` nodes.

Nodes that land on top of each other become **one node whose `Val` is their sum**. The result is read **from the crease outward**: first `a[k] + a[k+1]`, then `a[k-1] + a[k+2]`, and so on. Whatever remains of the longer side follows, still in crease-outward order.

You must **reuse the nodes**: do not create new nodes or copy values into a slice.

Example 1:

```
Input: head = [1,2,3,4,5,6,7], k = 3
Output: [7,7,7,7]
Explanation: 3+4, 2+5, 1+6, then 7 is left alone.
```

Example 2:

```
Input: head = [1,2,3,4,5,6,7], k = 5
Output: [11,11,3,2,1]
Explanation: 5+6, 4+7, then the leftover left side 3, 2, 1 read from the crease.
```

Example 3:

```
Input: head = [1,2,3,4], k = 9
Output: [4,3,2,1]
Explanation: The crease is past the end, so the whole strip is the left side, read back from the crease.
```

Constraints:

* The number of nodes is in the range `[0, 10000]`.
* `0 <= k <= 20000`
* Follow up: solve it in one pass, `O(n)` time and `O(1)` memory.

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
func PaperFold(head *ListNode, k int) *ListNode {

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
	show(PaperFold(buildList(1, 2, 3, 4, 5, 6, 7), 3))
	show(PaperFold(buildList(1, 2, 3, 4, 5, 6, 7), 5))
	show(PaperFold(buildList(1, 2, 3, 4), 0))
	show(PaperFold(buildList(1, 2, 3, 4), 9))
	show(PaperFold(buildList(42), 1))
	show(PaperFold(nil, 2))
}

func buildList(vals ...int) *ListNode {
	var head *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		head = &ListNode{Val: vals[i], Next: head}
	}
	return head
}

func show(h *ListNode) {
	for ; h != nil; h = h.Next {
		fmt.Print(h.Val, " -> ")
	}
	fmt.Println("nil")
}
```

And its output:

```console
$ go run . | cat -e
7 -> 7 -> 7 -> 7 -> nil$
11 -> 11 -> 3 -> 2 -> 1 -> nil$
1 -> 2 -> 3 -> 4 -> nil$
4 -> 3 -> 2 -> 1 -> nil$
42 -> nil$
nil$
```

### Reference solution

```go
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func PaperFold(head *ListNode, k int) *ListNode {
	// split: left = first k nodes, right = the rest
	var left *ListNode
	right := head
	for i := 0; i < k && right != nil; i++ {
		next := right.Next
		right.Next = left // reverse the left part while walking
		left = right
		right = next
	}
	// left now starts at the crease; merge the two strips node by node
	dummy := &ListNode{}
	tail := dummy
	for left != nil && right != nil {
		left.Val += right.Val
		nextL, nextR := left.Next, right.Next
		right.Next = nil
		tail.Next = left
		tail = left
		left, right = nextL, nextR
	}
	if left != nil {
		tail.Next = left
	} else {
		tail.Next = right
	}
	return dummy.Next
}
```

**Approach:** Reverse the first `k` nodes while walking to the crease, which leaves both strips starting at the crease. Zip them: add the right node's value into the left node, drop the right node, and append the longer strip's leftover.

**Complexity:** O(n) time, O(1) extra space.

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

## blockmove

### Instructions

A text document is a doubly linked list of lines that sits between two **sentinel** nodes. Write `NewDoc`, which builds a document from the given lines, and the method `MoveBlock`, which cuts lines `from..to` (1-based, inclusive) and pastes them right after line `after`. `after` is numbered in the **original** document, and `after = 0` means the top.

`MoveBlock` returns `false` and changes nothing if `from < 1`, `from > to`, `to` is past the last line, `after` is outside `0..n`, or `after` falls **inside the block**. Moving a block to where it already is returns `true`.

The move itself must be a constant number of relinks: **do not create or copy nodes**.

Example 1:

```
Input: doc = "ABCDEF", MoveBlock(2, 3, 5)
Output: true, doc = "ADEBCF"
Explanation: Lines B and C are cut and pasted after line 5 (E).
```

Example 2:

```
Input: doc = "ADEBCF", MoveBlock(4, 6, 0)
Output: true, doc = "BCFADE"
Explanation: B, C and F move to the top.
```

Example 3:

```
Input: doc = "CFADEB", MoveBlock(2, 4, 3)
Output: false, doc unchanged
Explanation: Line 3 is inside the block 2..4.
```

Constraints:

* The number of lines is in the range `[0, 10000]`.
* Finding the lines may take `O(n)`; the move itself must be `O(1)`.
* Follow up: write it so that no case needs a nil check, thanks to the sentinels.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Val  string
	Prev *ListNode
	Next *ListNode
}

type Doc struct {
	Head *ListNode // sentinel before the first line, holds no text
	Tail *ListNode // sentinel after the last line, holds no text
}
```

### Expected function

```go
func NewDoc(lines ...string) *Doc {

}

func (d *Doc) MoveBlock(from, to, after int) bool {

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
	d := NewDoc("A", "B", "C", "D", "E", "F")
	fmt.Println(d.MoveBlock(2, 3, 5))
	show(d)
	fmt.Println(d.MoveBlock(4, 6, 0))
	show(d)
	fmt.Println(d.MoveBlock(1, 1, 6))
	show(d)
	fmt.Println(d.MoveBlock(3, 4, 2))
	show(d)
	fmt.Println(d.MoveBlock(2, 4, 3))
	fmt.Println(d.MoveBlock(5, 7, 1))
	fmt.Println(d.MoveBlock(2, 2, 9))
	show(d)
	e := NewDoc()
	fmt.Println(e.MoveBlock(1, 1, 0))
	show(e)
}

func show(d *Doc) {
	fwd, back := "", ""
	for c := d.Head.Next; c != d.Tail; c = c.Next {
		fwd += c.Val
	}
	for c := d.Tail.Prev; c != d.Head; c = c.Prev {
		back += c.Val
	}
	fmt.Printf("[%s] [%s]\n", fwd, back)
}
```

And its output:

```console
$ go run . | cat -e
true$
[ADEBCF] [FCBEDA]$
true$
[BCFADE] [EDAFCB]$
true$
[CFADEB] [BEDAFC]$
true$
[CFADEB] [BEDAFC]$
false$
false$
false$
[CFADEB] [BEDAFC]$
false$
[] []$
```

### Reference solution

```go
package main

type ListNode struct {
	Val  string
	Prev *ListNode
	Next *ListNode
}

type Doc struct {
	Head *ListNode // sentinel before the first line
	Tail *ListNode // sentinel after the last line
}

func NewDoc(lines ...string) *Doc {
	d := &Doc{Head: &ListNode{}, Tail: &ListNode{}}
	d.Head.Next = d.Tail
	d.Tail.Prev = d.Head
	for _, t := range lines {
		n := &ListNode{Val: t, Prev: d.Tail.Prev, Next: d.Tail}
		d.Tail.Prev.Next = n
		d.Tail.Prev = n
	}
	return d
}

func (d *Doc) MoveBlock(from, to, after int) bool {
	if from < 1 || from > to || (after >= from && after <= to) || after < 0 {
		return false
	}
	var first, last, anchor *ListNode
	if after == 0 {
		anchor = d.Head
	}
	i := 0
	for cur := d.Head.Next; cur != d.Tail; cur = cur.Next {
		i++
		if i == from {
			first = cur
		}
		if i == to {
			last = cur
		}
		if i == after {
			anchor = cur
		}
	}
	if last == nil || anchor == nil {
		return false
	}
	if anchor.Next == first { // already in place
		return true
	}
	// cut [first..last]
	first.Prev.Next = last.Next
	last.Next.Prev = first.Prev
	// paste after anchor
	first.Prev = anchor
	last.Next = anchor.Next
	anchor.Next.Prev = last
	anchor.Next = first
	return true
}
```

**Approach:** One walk finds the first, last and anchor nodes. Cut the block by linking its outer neighbours together, then splice it after the anchor. The sentinels remove every head and tail special case.

**Complexity:** O(n) time, O(1) extra space.

---

## sparsegrid

### Instructions

A sparse matrix stores only its non-zero cells. Each cell is **one node that belongs to two lists at once**: its row list (linked by `Right`, sorted by `Col`) and its column list (linked by `Down`, sorted by `Row`).

Write `NewGrid` and the methods `Set` and `Get`. `Set(r, c, val)` creates or updates the cell, or removes it when `val == 0`, keeping both lists sorted. Setting an empty cell to `0` does nothing. `Get` returns `0` for an empty cell. Both methods **ignore coordinates outside the grid**.

A cell is never duplicated: the node reached through its row must be **the same node** reached through its column.

Example 1:

```
Input: NewGrid(3, 4), Set(1,2,5), Set(0,2,7), Set(1,0,3)
Output: row 1 = (1,0)=3 -> (1,2)=5, col 2 = (0,2)=7 -> (1,2)=5
Explanation: (1,2) is a single node, linked from row 1 and from column 2.
```

Example 2:

```
Input: then Set(1,2,0)
Output: row 1 = (1,0)=3, col 2 = (0,2)=7
Explanation: Setting 0 unlinks the cell from both lists.
```

Example 3:

```
Input: Get(2,1), Get(-1,0)
Output: 0 0
Explanation: Empty and out-of-range cells read as 0.
```

Constraints:

* `0 <= rows, cols <= 1000`
* `-1000000 <= val <= 1000000`
* Each call must be `O(row length + column length)`.
* Follow up: insert and remove without special-casing the list heads (pointer to pointer).

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type Cell struct {
	Row   int
	Col   int
	Val   int
	Right *Cell
	Down  *Cell
}

type Grid struct {
	RowHead []*Cell // RowHead[r] is the first cell of row r
	ColHead []*Cell // ColHead[c] is the first cell of column c
}
```

### Expected function

```go
func NewGrid(rows, cols int) *Grid {

}

func (g *Grid) Set(r, c, val int) {

}

func (g *Grid) Get(r, c int) int {

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
	g := NewGrid(3, 4)
	g.Set(1, 2, 5)
	g.Set(0, 2, 7)
	g.Set(1, 0, 3)
	g.Set(2, 2, 9)
	g.Set(1, 3, 4)
	g.Set(5, 5, 1) // outside, ignored
	g.Set(1, 2, 6) // update
	dump(g)
	fmt.Println(g.Get(1, 2), g.Get(2, 1), g.Get(-1, 0))
	fmt.Println(g.RowHead[1].Right == g.ColHead[2].Down)
	g.Set(1, 2, 0) // remove
	g.Set(0, 0, 0) // removing an empty cell does nothing
	dump(g)
	empty := NewGrid(0, 0)
	empty.Set(0, 0, 1)
	fmt.Println(empty.Get(0, 0), len(empty.RowHead))
}

func dump(g *Grid) {
	for r, h := range g.RowHead {
		fmt.Printf("row %d:", r)
		for ; h != nil; h = h.Right {
			fmt.Printf(" (%d,%d)=%d", h.Row, h.Col, h.Val)
		}
		fmt.Println()
	}
	for c, h := range g.ColHead {
		fmt.Printf("col %d:", c)
		for ; h != nil; h = h.Down {
			fmt.Printf(" (%d,%d)=%d", h.Row, h.Col, h.Val)
		}
		fmt.Println()
	}
}
```

And its output:

```console
$ go run . | cat -e
row 0: (0,2)=7$
row 1: (1,0)=3 (1,2)=6 (1,3)=4$
row 2: (2,2)=9$
col 0: (1,0)=3$
col 1:$
col 2: (0,2)=7 (1,2)=6 (2,2)=9$
col 3: (1,3)=4$
6 0 0$
true$
row 0: (0,2)=7$
row 1: (1,0)=3 (1,3)=4$
row 2: (2,2)=9$
col 0: (1,0)=3$
col 1:$
col 2: (0,2)=7 (2,2)=9$
col 3: (1,3)=4$
0 0$
```

### Reference solution

```go
package main

type Cell struct {
	Row   int
	Col   int
	Val   int
	Right *Cell
	Down  *Cell
}

type Grid struct {
	RowHead []*Cell
	ColHead []*Cell
}

func NewGrid(rows, cols int) *Grid {
	return &Grid{RowHead: make([]*Cell, rows), ColHead: make([]*Cell, cols)}
}

func (g *Grid) inside(r, c int) bool {
	return r >= 0 && c >= 0 && r < len(g.RowHead) && c < len(g.ColHead)
}

func (g *Grid) Get(r, c int) int {
	if !g.inside(r, c) {
		return 0
	}
	for cur := g.RowHead[r]; cur != nil && cur.Col <= c; cur = cur.Right {
		if cur.Col == c {
			return cur.Val
		}
	}
	return 0
}

func (g *Grid) Set(r, c, val int) {
	if !g.inside(r, c) {
		return
	}
	// find slot in the row list
	rp := &g.RowHead[r]
	for *rp != nil && (*rp).Col < c {
		rp = &(*rp).Right
	}
	// find slot in the column list
	cp := &g.ColHead[c]
	for *cp != nil && (*cp).Row < r {
		cp = &(*cp).Down
	}
	existing := *rp
	if existing != nil && existing.Col == c {
		if val != 0 {
			existing.Val = val
			return
		}
		*rp = existing.Right // unlink from row
		*cp = existing.Down  // unlink from column (same node)
		existing.Right, existing.Down = nil, nil
		return
	}
	if val == 0 {
		return
	}
	n := &Cell{Row: r, Col: c, Val: val, Right: *rp, Down: *cp}
	*rp = n
	*cp = n
}
```

**Approach:** Walk the row list to the slot before column `c` and the column list to the slot before row `r`; a pointer-to-pointer avoids head cases. Then update in place, unlink from both lists, or insert one node into both.

**Complexity:** O(row length + column length) per call, O(1) extra space.

---

## greaterpointers

### Instructions

Every node has a normal `Next` pointer and a `Greater` pointer to the **first later node with a strictly bigger value**, or `nil` if there is none.

Given the `head`, a position `pos` and a value `val`, insert a new node so that it ends up at 0-based position `pos`, and keep **every** `Greater` pointer in the list correct. Return the head. A `pos` below `0` inserts at the front, and a `pos` past the end appends.

The new node needs its own `Greater`, and some **earlier** nodes may now have to point at it. Rebuilding every pointer from scratch is too slow.

Example 1:

```
Input: list = 4(5) 5(-) 3(-) 1(-), InsertAt(2, 8)
Output: 4(5) 5(8) 8(-) 3(-) 1(-)
Explanation: 5 now points at 8. 4 keeps 5, which still comes first. Each node is written as value(Greater's value), with - meaning nil.
```

Example 2:

```
Input: list = ... 3(-) 1(-), InsertAt(99, 2)
Output: ... 3(-) 1(2) 2(-)
Explanation: 2 is appended, and 1 is the only earlier node smaller than 2 without a closer target.
```

Example 3:

```
Input: list = nil, InsertAt(0, 5)
Output: 5(-)
Explanation: The first node has no Greater.
```

Constraints:

* The number of nodes is in the range `[0, 10000]`.
* `-100000 <= val <= 100000`
* Follow up: `O(n)` per insert. Find the new node's `Greater` by hopping along `Greater` pointers.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Val     int
	Next    *ListNode
	Greater *ListNode // first later node with a strictly bigger Val, or nil
}
```

### Expected function

```go
func InsertAt(head *ListNode, pos, val int) *ListNode {

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
	show(list)
	steps := [][2]int{{0, 5}, {1, 3}, {2, 1}, {0, 4}, {2, 8}, {99, 2}, {3, 6}, {-1, 9}, {4, 6}}
	for _, s := range steps {
		list = InsertAt(list, s[0], s[1])
		fmt.Printf("InsertAt(%d, %d): ", s[0], s[1])
		show(list)
	}
}

func show(h *ListNode) {
	for ; h != nil; h = h.Next {
		if h.Greater == nil {
			fmt.Printf("%d(-) ", h.Val)
		} else {
			fmt.Printf("%d(%d) ", h.Val, h.Greater.Val)
		}
	}
	fmt.Println("|")
}
```

And its output:

```console
$ go run . | cat -e
|$
InsertAt(0, 5): 5(-) |$
InsertAt(1, 3): 5(-) 3(-) |$
InsertAt(2, 1): 5(-) 3(-) 1(-) |$
InsertAt(0, 4): 4(5) 5(-) 3(-) 1(-) |$
InsertAt(2, 8): 4(5) 5(8) 8(-) 3(-) 1(-) |$
InsertAt(99, 2): 4(5) 5(8) 8(-) 3(-) 1(2) 2(-) |$
InsertAt(3, 6): 4(5) 5(8) 8(-) 6(-) 3(-) 1(2) 2(-) |$
InsertAt(-1, 9): 9(-) 4(5) 5(8) 8(-) 6(-) 3(-) 1(2) 2(-) |$
InsertAt(4, 6): 9(-) 4(5) 5(8) 8(-) 6(-) 6(-) 3(-) 1(2) 2(-) |$
```

### Reference solution

```go
package main

type ListNode struct {
	Val     int
	Next    *ListNode
	Greater *ListNode
}

func InsertAt(head *ListNode, pos, val int) *ListNode {
	n := &ListNode{Val: val}
	// collect the nodes that will sit before n
	var before []*ListNode
	cur := head
	for i := 0; i < pos && cur != nil; i++ {
		before = append(before, cur)
		cur = cur.Next
	}
	// link n in
	n.Next = cur
	if len(before) == 0 {
		head = n
	} else {
		before[len(before)-1].Next = n
	}
	// n.Greater: jump along Greater pointers from the node after n
	g := n.Next
	for g != nil && g.Val <= val {
		g = g.Greater
	}
	n.Greater = g
	// earlier nodes: their target changes to n only if it used to lie after n
	inBefore := make(map[*ListNode]bool, len(before))
	for _, b := range before {
		inBefore[b] = true
	}
	for _, b := range before {
		if b.Val < val && !inBefore[b.Greater] {
			b.Greater = n
		}
	}
	return head
}
```

**Approach:** Link the node in. Its `Greater` comes from hopping `Greater` pointers from `n.Next` while values are `<= val`. An earlier node switches to `n` only if its value is `< val` and its old target lies after `n`, that is, outside the prefix (tracked in a set).

**Complexity:** O(n) time per insert, O(n) extra space for the prefix set.

---

## timerwheel

### Instructions

A timer wheel is a **circular** list of `Size` slots. Each slot holds a doubly linked list of entries behind a sentinel. Write `NewWheel` and the methods `Schedule`, `Cancel` and `Tick`.

`Schedule(id, delay)` adds an entry that fires on the `delay`-th `Tick` from now and returns it as a handle. A `delay` below `1` counts as `1`. Delays longer than one rotation must **wrap around** using `Rounds`.

`Cancel(e)` removes a pending entry in **O(1)** and returns `true`. It returns `false` for `nil` or for an entry that has already fired or been cancelled.

`Tick()` moves `Current` to the next slot. Entries there with `Rounds == 0` fire and are removed, and every other entry there loses one round. It returns the fired IDs in scheduling order, or an empty slice if none fired.

Example 1:

```
Input: NewWheel(4), Schedule("b", 4), Schedule("c", 5), Schedule("d", 9), then 9 Ticks
Output: b fires on tick 4, c on tick 5, d on tick 9
Explanation: c and d land in the same slot, which is 1 step ahead, with Rounds 1 and 2. They pass through that slot on tick 1 (and tick 5 for d), losing one round each time.
```

Example 2:

```
Input: e := Schedule("e", 2), Cancel(e), Cancel(e)
Output: true, false
Explanation: The second cancel finds e already detached.
```

Example 3:

```
Input: NewWheel(1), Schedule("x", 3), 4 Ticks
Output: [] [] [x] []
Explanation: With a single slot, every delay wraps around.
```

Constraints:

* `1 <= Size <= 1000`
* `1 <= delay <= 1000000000`, and up to `100000` operations.
* `Cancel` must be `O(1)`, and `Tick` must be `O(entries in that slot)`.
* Follow up: make `Schedule` independent of `Size` by keeping the slots in a slice as well.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type Entry struct {
	ID     string
	Rounds int // full rotations left before the entry fires
	Prev   *Entry
	Next   *Entry
}

type Slot struct {
	Entries *Entry // sentinel of a circular doubly linked list of entries
	Next    *Slot
}

type Wheel struct {
	Current *Slot // slots form a circular singly linked list
	Size    int
}
```

### Expected function

```go
func NewWheel(size int) *Wheel {

}

func (w *Wheel) Schedule(id string, delay int) *Entry {

}

func (w *Wheel) Cancel(e *Entry) bool {

}

func (w *Wheel) Tick() []string {

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
	w := NewWheel(4)
	a := w.Schedule("a", 1)
	w.Schedule("b", 4)
	c := w.Schedule("c", 5)
	w.Schedule("d", 9)
	e := w.Schedule("e", 2)
	w.Schedule("f", 2)
	fmt.Println(w.Cancel(e))
	fmt.Println(w.Cancel(e))
	for t := 1; t <= 10; t++ {
		fired := w.Tick()
		fmt.Printf("tick %d: %v\n", t, fired)
		if t == 3 {
			fmt.Println(w.Cancel(c))
			w.Schedule("g", 1)
		}
	}
	fmt.Println(w.Cancel(a))
	fmt.Println(w.Cancel(nil))
	one := NewWheel(1)
	one.Schedule("x", 3)
	fmt.Println(one.Tick(), one.Tick(), one.Tick(), one.Tick())
}
```

And its output:

```console
$ go run . | cat -e
true$
false$
tick 1: [a]$
tick 2: [f]$
tick 3: []$
true$
tick 4: [b g]$
tick 5: []$
tick 6: []$
tick 7: []$
tick 8: []$
tick 9: [d]$
tick 10: []$
false$
false$
[] [] [x] []$
```

### Reference solution

```go
package main

type Entry struct {
	ID     string
	Rounds int
	Prev   *Entry
	Next   *Entry
}

type Slot struct {
	Entries *Entry // sentinel of a circular doubly linked list
	Next    *Slot
}

type Wheel struct {
	Current *Slot
	Size    int
}

func NewWheel(size int) *Wheel {
	if size < 1 {
		size = 1
	}
	var first, last *Slot
	for i := 0; i < size; i++ {
		s := &Slot{Entries: &Entry{}}
		s.Entries.Prev, s.Entries.Next = s.Entries, s.Entries
		if first == nil {
			first = s
		} else {
			last.Next = s
		}
		last = s
	}
	last.Next = first
	return &Wheel{Current: first, Size: size}
}

func (w *Wheel) Schedule(id string, delay int) *Entry {
	if delay < 1 {
		delay = 1
	}
	steps := delay % w.Size
	s := w.Current
	for i := 0; i < steps; i++ {
		s = s.Next
	}
	e := &Entry{ID: id, Rounds: (delay - 1) / w.Size}
	sent := s.Entries
	e.Prev, e.Next = sent.Prev, sent
	sent.Prev.Next = e
	sent.Prev = e
	return e
}

func (w *Wheel) Cancel(e *Entry) bool {
	if e == nil || e.Prev == nil {
		return false
	}
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	e.Prev, e.Next = nil, nil
	return true
}

func (w *Wheel) Tick() []string {
	w.Current = w.Current.Next
	fired := []string{}
	sent := w.Current.Entries
	for e := sent.Next; e != sent; {
		next := e.Next
		if e.Rounds == 0 {
			fired = append(fired, e.ID)
			w.Cancel(e)
		} else {
			e.Rounds--
		}
		e = next
	}
	return fired
}
```

**Approach:** The target slot is `delay % Size` steps ahead, with `Rounds = (delay-1) / Size`. Each slot's sentinel makes unlinking O(1), and clearing `Prev`/`Next` marks an entry as detached so a second cancel returns false. Tick advances, then fires or decrements.

**Complexity:** Schedule O(Size), Cancel O(1), Tick O(k) for k entries in the slot; O(Size + entries) space.
