## appendnode

### Instructions

Given the `head` of a singly-linked list and an integer `val`, add a new node holding `val` at the end of the list and return the head of the list.

If the list is empty (`head` is `nil`), the new node becomes the head.

Example 1:

```
Input: head = [1,2,3], val = 4
Output: [1,2,3,4]
```

Example 2:

```
Input: head = [], val = 5
Output: [5]
Explanation: The list was empty, so the new node is the head.
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
func AppendNode(head *ListNode, val int) *ListNode {

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
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}

	head = AppendNode(head, 3)
	printList(head)

	head = AppendNode(head, 4)
	printList(head)

	printList(AppendNode(nil, 5))
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
1 -> 2 -> 3$
1 -> 2 -> 3 -> 4$
5$
```
