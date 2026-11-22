## dealcards

### Instructions

Given the `head` of a singly linked list and an integer `k`, deal its nodes into `k` hands the way you deal cards: the 1st node goes to hand 0, the 2nd to hand 1, and so on, wrapping back to hand 0 after hand `k-1`. Return the heads of the hands as a slice.

Each hand keeps its nodes in their **original relative order**. You must **reuse the existing nodes**: do not create new ones. A hand that receives no node is `nil`, and if `k <= 0` the function returns `nil`.

Example 1:

```
Input: head = [1,2,3,4,5,6,7], k = 3
Output: [[1,4,7], [2,5], [3,6]]
Explanation: Nodes are dealt 1â†’hand 0, 2â†’hand 1, 3â†’hand 2, 4â†’hand 0, and so on.
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
