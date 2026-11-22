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
