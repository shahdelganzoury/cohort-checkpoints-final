## getintersectionnode

### Instructions

Given the heads of two singly-linked lists `headA` and `headB`, return the node at which the two lists intersect. If the two lists have no intersection at all, return `nil`.

Two lists intersect when they share the same node (the same reference in memory) and from there on are the same list. Nodes that only have equal values are **not** an intersection.

The lists contain no cycles, and they must keep their original structure after your function returns.

Example 1:

```
Input: listA = [4,1,8,4,5], listB = [5,6,1,8,4,5] (both end with the shared nodes 8 -> 4 -> 5)
Output: the node with value 8
Explanation: The lists meet at the node with value 8. The node with value 1 in A and the node with value 1 in B have equal values but are different nodes.
```

Example 2:

```
Input: listA = [2,6,4], listB = [1,5]
Output: nil
Explanation: The lists do not intersect.
```

Constraints:

* The number of nodes in each list is in the range `[1, 30000]`.
* `1 <= Node.Val <= 100000`
* Compare nodes by reference, not by value.
* Follow up: solve it in `O(m + n)` time and `O(1)` memory.

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
func GetIntersectionNode(headA, headB *ListNode) *ListNode {

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
	shared := buildList(8, 4, 5)
	a := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: shared}}
	b := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: shared}}}
	fmt.Println(GetIntersectionNode(a, b) == shared)
	fmt.Println(GetIntersectionNode(a, b).Val)

	c := buildList(2, 6, 4)
	d := buildList(1, 5)
	fmt.Println(GetIntersectionNode(c, d) == nil)

	e := buildList(1, 2, 3)
	f := buildList(1, 2, 3)
	fmt.Println(GetIntersectionNode(e, f) == nil)
}

func buildList(vals ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}
```

And its output:

```console
$ go run . | cat -e
true$
8$
true$
true$
```
