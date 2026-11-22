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
