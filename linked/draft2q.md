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

---

## browserhistory

### Instructions

Build the history of one browser tab as a **doubly linked list of pages**. `NewBrowserHistory(homepage)` starts on `homepage`.

`Visit(url)` opens `url` right after the current page and returns its `*Page`. It **cuts off all forward history**: the new page's `Next` is `nil`, and the page that used to come after the current one is detached (its `Prev` becomes `nil`).

`Back(steps)` moves back at most `steps` pages, stopping at the oldest, and `Forward(steps)` moves forward at most `steps` pages, stopping at the newest. Both return the current `*Page`. Moving never creates pages: `Back` must return **the same node** that `Visit` returned earlier.

Example 1:

```
Input: NewBrowserHistory("leetcode.com"), Visit google.com, facebook.com, youtube.com, then Back(1), Back(1), Forward(1)
Output: facebook.com, google.com, facebook.com
Explanation: The second Back returns the very *Page that Visit("google.com") returned.
```

Example 2:

```
Input: then Visit("linkedin.com")
Output: linkedin.com, with Prev = facebook.com and Next = nil
Explanation: youtube.com is cut off, and its Prev becomes nil.
```

Example 3:

```
Input: then Forward(2), Back(2), Back(7)
Output: linkedin.com, google.com, leetcode.com
Explanation: Forward stops at the newest page, and Back stops at the homepage.
```

Constraints:

* `1 <= len(url) <= 20`
* `1 <= steps <= 100`, and at most `5000` calls in total.
* Do not store the history in a slice: walk the `Prev` and `Next` pointers.
* Follow up: `Visit` in `O(1)` and `Back`/`Forward` in `O(steps)`.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type Page struct {
	URL  string
	Prev *Page
	Next *Page
}

type BrowserHistory struct {
	Current *Page // the page on screen
}
```

### Expected function

```go
func NewBrowserHistory(homepage string) *BrowserHistory {

}

func (b *BrowserHistory) Visit(url string) *Page {

}

func (b *BrowserHistory) Back(steps int) *Page {

}

func (b *BrowserHistory) Forward(steps int) *Page {

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
	b := NewBrowserHistory("leetcode.com")
	home := b.Current
	google := b.Visit("google.com")
	b.Visit("facebook.com")
	youtube := b.Visit("youtube.com")
	fmt.Println(b.Back(1).URL)
	fmt.Println(b.Back(1) == google)
	fmt.Println(b.Forward(1).URL)
	linkedin := b.Visit("linkedin.com")
	fmt.Println(linkedin.Prev.URL, linkedin.Next == nil, youtube.Prev == nil)
	fmt.Println(b.Forward(2) == linkedin)
	fmt.Println(b.Back(2).URL)
	fmt.Println(b.Back(7) == home, home.Prev == nil)
	fmt.Println(walk(home))
	fmt.Println(walkBack(linkedin))

	solo := NewBrowserHistory("a.com")
	fmt.Println(solo.Back(3).URL, solo.Forward(3).URL)
}

func walk(p *Page) []string {
	var out []string
	for ; p != nil; p = p.Next {
		out = append(out, p.URL)
	}
	return out
}

func walkBack(p *Page) []string {
	var out []string
	for ; p != nil; p = p.Prev {
		out = append(out, p.URL)
	}
	return out
}
```

And its output:

```console
$ go run . | cat -e
facebook.com$
true$
facebook.com$
facebook.com true true$
true$
google.com$
true true$
[leetcode.com google.com facebook.com linkedin.com]$
[linkedin.com facebook.com google.com leetcode.com]$
a.com a.com$
```

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

---

## flattenmultilevel

### Instructions

Each node of a doubly linked list may also have a `Child` pointing to the head of another doubly linked list, which can have children of its own. Flatten everything into **one level**, depth first: a node's child list goes **right after that node and before its old `Next`**.

Return the head. Afterwards **every `Child` must be `nil`**, and walking the `Prev` pointers must give the exact reverse order.

Example 1:

```
Input: level 1 = 1-2-3-4-5-6, node 3 has child 7-8-9-10, node 8 has child 11-12
Output: 1 2 3 7 8 11 12 9 10 4 5 6
Explanation: Node 8's child goes right after 8, and the whole of node 3's child list comes before 4.
```

Example 2:

```
Input: level 1 = 1-2, node 1 has child 3
Output: 1 3 2
Explanation: The child goes between its parent and the parent's old Next.
```

Example 3:

```
Input: nil
Output: nil
Explanation: Nothing to flatten.
```

Constraints:

* The total number of nodes is in the range `[0, 1000]`.
* `1 <= Node.Val <= 100000`
* Follow up: `O(n)` time and `O(1)` extra memory, with no recursion or stack.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Val   int
	Prev  *ListNode
	Next  *ListNode
	Child *ListNode
}
```

### Expected function

```go
func Flatten(head *ListNode) *ListNode {

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
	top := buildList(1, 2, 3, 4, 5, 6)
	mid := buildList(7, 8, 9, 10)
	nth(top, 2).Child = mid
	nth(mid, 1).Child = buildList(11, 12)
	show(Flatten(top))

	small := buildList(1, 2)
	small.Child = buildList(3)
	show(Flatten(small))

	deep := buildList(1)
	deep.Child = buildList(2)
	deep.Child.Child = buildList(3)
	show(Flatten(deep))

	show(Flatten(buildList(4)))
	show(Flatten(nil))
}

func buildList(vals ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v, Prev: tail}
		tail = tail.Next
	}
	if dummy.Next != nil {
		dummy.Next.Prev = nil
	}
	return dummy.Next
}

func nth(h *ListNode, i int) *ListNode {
	for ; i > 0; i-- {
		h = h.Next
	}
	return h
}

func show(h *ListNode) {
	fwd, back := []int{}, []int{}
	children := 0
	var last *ListNode
	for c := h; c != nil; c = c.Next {
		fwd = append(fwd, c.Val)
		if c.Child != nil {
			children++
		}
		last = c
	}
	for c := last; c != nil; c = c.Prev {
		back = append(back, c.Val)
	}
	fmt.Println(fwd, back, "children left:", children)
}
```

And its output:

```console
$ go run . | cat -e
[1 2 3 7 8 11 12 9 10 4 5 6] [6 5 4 10 9 12 11 8 7 3 2 1] children left: 0$
[1 3 2] [2 3 1] children left: 0$
[1 2 3] [3 2 1] children left: 0$
[4] [4] children left: 0$
[] [] children left: 0$
```

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

---

## lrucache

### Instructions

Design a cache with a fixed capacity that evicts the **least recently used** key when it is full. Recency order is kept in a doubly linked list between the sentinels `Head` and `Tail`, and `Items` maps each key to its node.

`NewLRUCache(capacity)` creates an empty cache. `Get(key)` returns the value, or `-1` if the key is absent, and a hit counts as a use. `Put(key, value)` updates or inserts the key and counts as a use. If inserting a **new** key would exceed the capacity, it first evicts the least recently used one, removing it from **both** the list and the map.

Implement the list yourself: `container/list` is not allowed.

Example 1:

```
Input: NewLRUCache(2), Put(1,1), Put(2,2), Get(1)
Output: 1, order [1 2]
Explanation: Get(1) makes key 1 the most recent.
```

Example 2:

```
Input: then Put(3,3), Get(2)
Output: -1, order [3 1]
Explanation: The cache is full, so key 2 (the least recent) is evicted.
```

Example 3:

```
Input: then Put(3,30)
Output: no eviction, order [3 ...]
Explanation: Updating an existing key changes its value and moves it to the front.
```

Constraints:

* `1 <= capacity <= 3000`
* `0 <= key <= 10000`, `0 <= value <= 100000`, and up to `200000` calls.
* Follow up: `Get` and `Put` in `O(1)` on average, with no special case for an empty list (thanks to the sentinels).

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Key  int
	Val  int
	Prev *ListNode
	Next *ListNode
}

type LRUCache struct {
	Capacity int
	Items    map[int]*ListNode
	Head     *ListNode // sentinel: Head.Next is the most recently used
	Tail     *ListNode // sentinel: Tail.Prev is the least recently used
}
```

### Expected function

```go
func NewLRUCache(capacity int) *LRUCache {

}

func (c *LRUCache) Get(key int) int {

}

func (c *LRUCache) Put(key, value int) {

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
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1), order(c))
	c.Put(3, 3)
	fmt.Println(c.Get(2), order(c))
	c.Put(4, 4)
	fmt.Println(c.Get(1), c.Get(3), c.Get(4), order(c))
	c.Put(3, 30)
	fmt.Println(c.Get(3), order(c), len(c.Items))

	one := NewLRUCache(1)
	fmt.Println(one.Get(5))
	one.Put(5, 50)
	one.Put(6, 60)
	fmt.Println(one.Get(5), one.Get(6), order(one))
}

// order lists keys from most to least recent, then checks the Prev links.
func order(c *LRUCache) string {
	fwd, back := []int{}, []int{}
	for n := c.Head.Next; n != c.Tail; n = n.Next {
		fwd = append(fwd, n.Key)
	}
	for n := c.Tail.Prev; n != c.Head; n = n.Prev {
		back = append(back, n.Key)
	}
	return fmt.Sprint(fwd, back)
}
```

And its output:

```console
$ go run . | cat -e
1 [1 2] [2 1]$
-1 [3 1] [1 3]$
-1 3 4 [4 3] [3 4]$
30 [3 4] [4 3] 2$
-1$
-1 60 [6] [6]$
```

---

## recentfiles

### Instructions

A text editor keeps a **Recent Files** menu with a fixed capacity. Opening a file moves it to the top. When the menu is full, opening a new file evicts the least recently opened one. `List()` returns names from most to least recent, and `Contains` and `Size` do what their names say.

**Part 1: fix the bugs.** The code below is supposed to do this, but it has bugs. Find and fix **every** bug without rewriting it, and explain each one in a sentence.

```go
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

**Part 2: add pinning.** `Open` now returns whether the file is in the menu afterwards. `Pin(name)` pins a file already in the menu: it returns `false` if the file is absent, and pinning **does not count as a use**. `Unpin(name)` returns `false` if the file is absent or not pinned. A successful unpin **counts as a use**, so the file moves to the front.

Pinned files still count toward the capacity and keep their place in `List()`, but they are **never evicted**. A new file evicts the least recent **unpinned** file. If every file is pinned, `Open` returns `false` and changes nothing. Re-opening a file that is already in the menu always returns `true` and moves it to the front.

Example 1:

```
Input: NewRecentFiles(3), Open a, b, c, a, d, c, e
Output: [c b a], [a c b], [d a c], [c d a], [e c d]; Size 3, Contains("a") false
Explanation: This is the Part 1 test. d evicts b and e evicts a. With the original code, re-opening a creates a cycle and List never returns.
```

Example 2:

```
Input: NewRecentFiles(3), Open a, b, c, Pin("a"), Open("d")
Output: true, true, [d c a]
Explanation: a is the oldest but pinned, so b is evicted instead.
```

Example 3:

```
Input: then Pin c and d, Open("e"), Unpin("a"), Open("e")
Output: false [d c a], true [a d c], true [e d c]
Explanation: With everything pinned, e cannot get in. Unpinning a moves it to the front, and a is then the only file that can be evicted.
```

Constraints:

* `1 <= capacity <= 1000`, and up to `200000` calls.
* `Open`, `Pin` and `Unpin` must be `O(1)`.
* Follow up: finding the least recent unpinned file must not scan past pinned files. Keep unpinned files in a second list (`uprev`/`unext`) as well.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type entry struct {
	name         string
	pinned       bool
	prev, next   *entry // recency list: every file
	uprev, unext *entry // eviction list: unpinned files only
}

type RecentFiles struct {
	cap          int
	index        map[string]*entry
	head, tail   *entry // sentinels of the recency list: head.next is the most recent
	uhead, utail *entry // sentinels of the eviction list
}
```

### Expected function

```go
func NewRecentFiles(capacity int) *RecentFiles {

}

func (r *RecentFiles) Open(name string) bool {

}

func (r *RecentFiles) Pin(name string) bool {

}

func (r *RecentFiles) Unpin(name string) bool {

}

func (r *RecentFiles) List() []string {

}

func (r *RecentFiles) Contains(name string) bool {

}

func (r *RecentFiles) Size() int {

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
	// Part 1: the failing test, which must pass after the fixes.
	r := NewRecentFiles(3)
	r.Open("a")
	r.Open("b")
	r.Open("c")
	fmt.Println(r.List())
	r.Open("a")
	fmt.Println(r.List())
	r.Open("d")
	fmt.Println(r.List())
	r.Open("c")
	fmt.Println(r.List())
	r.Open("e")
	fmt.Println(r.List(), backward(r))
	fmt.Println(r.Size(), r.Contains("a"), r.Contains("b"))

	// Part 2: pinning.
	p := NewRecentFiles(3)
	p.Open("a")
	p.Open("b")
	p.Open("c")
	fmt.Println(p.Pin("a"), p.Open("d"), p.List())
	fmt.Println(p.Pin("c"), p.Pin("d"), p.Open("e"), p.List())
	fmt.Println(p.Unpin("a"), p.List())
	fmt.Println(p.Open("e"), p.List(), backward(p))
	fmt.Println(p.Open("c"), p.List())
	fmt.Println(p.Pin("zzz"), p.Unpin("e"), p.Pin("d"), p.Size())

	one := NewRecentFiles(1)
	fmt.Println(one.List(), one.Open("x"), one.Open("y"), one.List(), one.Contains("x"))
}

// backward walks the prev links from the tail sentinel.
func backward(r *RecentFiles) []string {
	var out []string
	for e := r.tail.prev; e != r.head; e = e.prev {
		out = append(out, e.name)
	}
	return out
}
```

And its output:

```console
$ go run . | cat -e
[c b a]$
[a c b]$
[d a c]$
[c d a]$
[e c d] [d c e]$
3 false false$
true true [d c a]$
true true false [d c a]$
true [a d c]$
true [e d c] [c d e]$
true [c e d]$
false false true 3$
[] true true [y] false$
```

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

---

## flattenbylevels

### Instructions

As in `flattenmultilevel`, each node of a doubly linked list may have a `Child` list, and children can have children. This time, flatten it **level by level**: all level-1 nodes, then all level-2 nodes, then level 3, and so on. Within a level, child lists appear in the order their parents appear in the result.

Return the head. Every `Child` must be `nil` afterwards, and the `Prev` pointers must give the exact reverse order.

Example 1:

```
Input: level 1 = 1-2-3; 1 has child 9-10, 2 has child 4-5, 3 has child 6, 4 has child 7-8
Output: 1 2 3 9 10 4 5 6 7 8
Explanation: Level 2 is 9 10 (from 1), 4 5 (from 2), 6 (from 3). Level 3 is 7 8.
```

Example 2:

```
Input: level 1 = 1-2, node 1 has child 3
Output: 1 2 3
Explanation: Compare the depth-first result of flattenmultilevel: 1 3 2.
```

Example 3:

```
Input: nil
Output: nil
Explanation: Nothing to flatten.
```

Constraints:

* The total number of nodes is in the range `[0, 5000]`.
* Follow up: `O(n)` time and `O(1)` extra memory, with no queue or recursion: use the list itself as the queue.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Val   int
	Prev  *ListNode
	Next  *ListNode
	Child *ListNode
}
```

### Expected function

```go
func FlattenByLevel(head *ListNode) *ListNode {

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
	top := buildList(1, 2, 3)
	top.Child = buildList(9, 10)
	top.Next.Child = buildList(4, 5)
	top.Next.Child.Child = buildList(7, 8)
	top.Next.Next.Child = buildList(6)
	show(FlattenByLevel(top))

	small := buildList(1, 2)
	small.Child = buildList(3)
	show(FlattenByLevel(small))

	deep := buildList(1)
	deep.Child = buildList(2, 3)
	deep.Child.Next.Child = buildList(4)
	deep.Child.Child = buildList(5)
	show(FlattenByLevel(deep))

	show(FlattenByLevel(buildList(8)))
	show(FlattenByLevel(nil))
}

func buildList(vals ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v, Prev: tail}
		tail = tail.Next
	}
	if dummy.Next != nil {
		dummy.Next.Prev = nil
	}
	return dummy.Next
}

func show(h *ListNode) {
	fwd, back := []int{}, []int{}
	children := 0
	var last *ListNode
	for c := h; c != nil; c = c.Next {
		fwd = append(fwd, c.Val)
		if c.Child != nil {
			children++
		}
		last = c
	}
	for c := last; c != nil; c = c.Prev {
		back = append(back, c.Val)
	}
	fmt.Println(fwd, back, "children left:", children)
}
```

And its output:

```console
$ go run . | cat -e
[1 2 3 9 10 4 5 6 7 8] [8 7 6 5 4 10 9 3 2 1] children left: 0$
[1 2 3] [3 2 1] children left: 0$
[1 2 3 5 4] [4 5 3 2 1] children left: 0$
[8] [8] children left: 0$
[] [] children left: 0$
```

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

---

## reversingcircle

### Instructions

People stand in a **circular doubly linked list**, where `Next` goes clockwise. Counting starts at `start`, which counts as 1, and goes clockwise. The `k`-th person counted is **eliminated** (unlinked from the circle). After each elimination the **direction reverses**, and counting restarts from the nearest remaining person in the new direction, who counts as 1.

Return the values in elimination order and the **survivor node** itself. The survivor must be one of the original nodes, left as a circle of one (its `Next` and `Prev` point to itself). A `nil` start returns an empty slice and `nil`.

`k` can be huge: **step `(k-1)` modulo the current size**.

Example 1:

```
Input: circle 1..5, start at 1, k = 3
Output: order [3 5 4 2], survivor 1
Explanation: Count 1,2,3 clockwise and eliminate 3. Reverse, so 2,1,5 eliminates 5. Reverse, so 1,2,4 eliminates 4. Reverse, so 2,1,2 eliminates 2.
```

Example 2:

```
Input: circle 1..6, start at 1, k = 1000000000
Output: order [4 5 3 2 1], survivor 6
Explanation: Stepping one billion times is too slow, but 999999999 mod 6 is 3 steps.
```

Example 3:

```
Input: circle 1, k = 7
Output: order [], survivor 1
Explanation: A single person survives at once.
```

Constraints:

* The number of nodes is in the range `[0, 5000]`.
* `1 <= k <= 1000000000`
* Follow up: never step more than `size - 1` times per elimination, for `O(n^2)` in the worst case.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Val  int
	Prev *ListNode // counter-clockwise neighbour
	Next *ListNode // clockwise neighbour
}
```

### Expected function

```go
func ReversingCircle(start *ListNode, k int) ([]int, *ListNode) {

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
	for _, t := range [][2]int{{5, 3}, {4, 2}, {3, 1}, {6, 1000000000}, {1, 7}} {
		start := buildCircle(t[0])
		order, survivor := ReversingCircle(start, t[1])
		fmt.Println(t[0], t[1], order, survivor.Val, survivor.Next == survivor, survivor.Prev == survivor)
	}
	order, survivor := ReversingCircle(nil, 3)
	fmt.Println(order, survivor == nil)
}

// buildCircle links 1..n clockwise through Next and returns person 1.
func buildCircle(n int) *ListNode {
	head := &ListNode{Val: 1}
	tail := head
	for v := 2; v <= n; v++ {
		node := &ListNode{Val: v, Prev: tail}
		tail.Next = node
		tail = node
	}
	tail.Next = head
	head.Prev = tail
	return head
}
```

And its output:

```console
$ go run . | cat -e
5 3 [3 5 4 2] 1 true true$
4 2 [2 4 3] 1 true true$
3 1 [1 3] 2 true true$
6 1000000000 [4 5 3 2 1] 6 true true$
1 7 [] 1 true true$
[] true$
```

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
