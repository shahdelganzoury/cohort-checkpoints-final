## insertsorteddoubly

### Instructions

Given the `head` of a doubly-linked list sorted in ascending order and an integer `val`, insert a new node holding `val` so that the list stays sorted. Return the head of the list.

Every node's `Prev` and `Next` pointers must be correct after the insertion. The `Prev` of the head and the `Next` of the tail are `nil`. If the new node becomes the first node, it is the new head. If `val` equals an existing value, the new node goes before the existing node.

Example 1:

```
Input: head = [3,7], val = 1
Output: [1,3,7]
Explanation: The new node becomes the head.
```

Example 2:

```
Input: head = [3,7], val = 5
Output: [3,5,7]
Explanation: The new node is inserted in the middle.
```

Example 3:

```
Input: head = [3,7], val = 9
Output: [3,7,9]
Explanation: The new node becomes the tail.
```

Example 4:

```
Input: head = [], val = 5
Output: [5]
```

Constraints:

* The number of nodes is in the range `[0, 1000]`.
* The list is sorted in ascending order.
* `-1000 <= Node.Val <= 1000`
* `-1000 <= val <= 1000`

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}
```

### Expected function

```go
func InsertSortedDoubly(head *DoublyListNode, val int) *DoublyListNode {

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
	var head *DoublyListNode

	head = InsertSortedDoubly(head, 5)
	printDoubly(head)

	head = InsertSortedDoubly(head, 3)
	printDoubly(head)

	head = InsertSortedDoubly(head, 9)
	printDoubly(head)

	head = InsertSortedDoubly(head, 7)
	printDoubly(head)

	head = InsertSortedDoubly(head, 1)
	printDoubly(head)
}

func printDoubly(head *DoublyListNode) {
	forward, backward := []int{}, []int{}
	var tail *DoublyListNode
	for curr := head; curr != nil; curr = curr.Next {
		forward = append(forward, curr.Val)
		tail = curr
	}
	for curr := tail; curr != nil; curr = curr.Prev {
		backward = append(backward, curr.Val)
	}
	fmt.Println(forward, backward)
}
```

And its output:

```console
$ go run . | cat -e
[5] [5]$
[3 5] [5 3]$
[3 5 9] [9 5 3]$
[3 5 7 9] [9 7 5 3]$
[1 3 5 7 9] [9 7 5 3 1]$
```
