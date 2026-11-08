## reverselinkedlist

### Instructions

Given the `head` of a singly-linked list, reverse the list and return the new head.

Reverse the list in place by rewiring the `Next` pointers, using `O(1)` extra space.

Example 1:

```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```

Example 2:

```
Input: head = [1,2]
Output: [2,1]
```

Example 3:

```
Input: head = []
Output: []
```

Constraints:

* The number of nodes is in the range `[0, 5000]`.
* `-5000 <= Node.Val <= 5000`

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
func ReverseLinkedList(head *ListNode) *ListNode {

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
	printList(ReverseLinkedList(buildList(1, 2, 3, 4, 5)))
	printList(ReverseLinkedList(buildList(1, 2)))
	printList(ReverseLinkedList(buildList(7)))
	printList(ReverseLinkedList(nil))
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
5 -> 4 -> 3 -> 2 -> 1$
2 -> 1$
7$
nil$
```
