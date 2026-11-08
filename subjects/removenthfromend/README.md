## removenthfromend

### Instructions

Given the `head` of a singly-linked list, remove the `n`th node from the end of the list and return the head of the list.

Example 1:

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

Example 2:

```
Input: head = [1], n = 1
Output: []
```

Example 3:

```
Input: head = [1,2], n = 1
Output: [1]
```

Constraints:

* The number of nodes in the list is `sz`.
* `1 <= sz <= 30`
* `0 <= Node.Val <= 100`
* `1 <= n <= sz`
* Follow up: could you do this in one pass?

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
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

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
	printList(RemoveNthFromEnd(buildList(1, 2, 3, 4, 5), 2))
	printList(RemoveNthFromEnd(buildList(1), 1))
	printList(RemoveNthFromEnd(buildList(1, 2), 1))
	printList(RemoveNthFromEnd(buildList(1, 2), 2))
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
1 -> 2 -> 3 -> 5$
nil$
1$
2$
```
