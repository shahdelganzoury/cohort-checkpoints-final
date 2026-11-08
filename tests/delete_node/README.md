## deletenode

### Instructions

Given the `head` of a singly-linked list and an integer `val`, remove the first node whose value equals `val` and return the head of the list.

If no node has that value, return the list unchanged. If several nodes match, remove only the first one.

Example 1:

```
Input: head = [1,2,3], val = 1
Output: [2,3]
Explanation: The head node is removed.
```

Example 2:

```
Input: head = [1,2,3], val = 2
Output: [1,3]
```

Example 3:

```
Input: head = [1,2,3], val = 9
Output: [1,2,3]
Explanation: No node has the value 9.
```

Constraints:

* The number of nodes is in the range `[0, 1000]`.
* `-1000 <= Node.Val <= 1000`
* `-1000 <= val <= 1000`

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
func DeleteNode(head *ListNode, val int) *ListNode {

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
	// 1 -> 2 -> 3, delete the head
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	printList(DeleteNode(a, 1))

	// 1 -> 2 -> 3, delete a middle node
	b := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	printList(DeleteNode(b, 2))

	// 1 -> 2 -> 3, delete the tail
	c := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	printList(DeleteNode(c, 3))

	// 1 -> 2 -> 3, value not found
	d := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	printList(DeleteNode(d, 9))

	// 1 -> 2 -> 2, only the first match is removed
	e := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}
	printList(DeleteNode(e, 2))

	// empty list
	printList(DeleteNode(nil, 1))

	// single node
	f := &ListNode{Val: 1}
	printList(DeleteNode(f, 1))
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("nil")
		return
	}
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
```

And its output:

```console
$ go run . | cat -e
2 -> 3$
1 -> 3$
1 -> 2$
1 -> 2 -> 3$
1 -> 2$
nil$
nil$
```
