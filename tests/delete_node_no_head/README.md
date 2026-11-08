## deletenodenohead

### Instructions

Write a function to delete a node in a singly-linked list. You are **not** given the head of the list. Instead, you are given the node to be deleted directly.

It is guaranteed that the node to be deleted is not the tail of the list.

Deleting the node does not mean freeing it from memory. It means that after your function runs:

* The value of the given node no longer exists in the list.
* The number of nodes in the list is one less.
* All the values before the given node are in the same order.
* All the values after the given node are in the same order.

Example 1:

```
Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation: You are given the second node with value 5, the list becomes 4 -> 1 -> 9 after calling your function.
```

Example 2:

```
Input: head = [4,5,1,9], node = 1
Output: [4,5,9]
Explanation: You are given the third node with value 1, the list becomes 4 -> 5 -> 9 after calling your function.
```

Constraints:

* The number of nodes is in the range `[2, 1000]`.
* `-1000 <= Node.Val <= 1000`
* The value of each node is unique.
* The given node is in the list and is not the tail.

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
func DeleteNodeNoHead(node *ListNode) {

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
	node5 := &ListNode{Val: 5}
	head := &ListNode{Val: 4, Next: node5}
	node5.Next = &ListNode{Val: 1, Next: &ListNode{Val: 9}}
	DeleteNodeNoHead(node5) // the node with value 5
	printList(head)

	node1 := &ListNode{Val: 1, Next: &ListNode{Val: 9}}
	head2 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: node1}}
	DeleteNodeNoHead(node1) // the node with value 1
	printList(head2)
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
4 -> 1 -> 9$
4 -> 5 -> 9$
```
