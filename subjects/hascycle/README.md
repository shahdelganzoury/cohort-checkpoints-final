## hascycle

### Instructions

Given `head`, the head of a singly-linked list, return `true` if the list contains a cycle, otherwise return `false`.

A list has a cycle if some node can be reached again by following the `Next` pointers continuously.

Example 1:

```
Input: head = [3,2,0,-4], the tail's Next points to the node at index 1
Output: true
Explanation: There is a cycle in the list, where the tail connects to the second node.
```

Example 2:

```
Input: head = [1,2], the tail's Next points to the node at index 0
Output: true
Explanation: There is a cycle in the list, where the tail connects to the first node.
```

Example 3:

```
Input: head = [1], no cycle
Output: false
```

Constraints:

* The number of nodes is in the range `[0, 10000]`.
* `-100000 <= Node.Val <= 100000`
* Follow up: can you solve it using `O(1)` memory?

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
func HasCycle(head *ListNode) bool {

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
	head := buildList(3, 2, 0, -4)
	head.Next.Next.Next.Next = head.Next // the tail points back to the node with value 2
	fmt.Println(HasCycle(head))

	head2 := buildList(1, 2)
	head2.Next.Next = head2 // the tail points back to the head
	fmt.Println(HasCycle(head2))

	fmt.Println(HasCycle(buildList(1)))
	fmt.Println(HasCycle(nil))
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
true$
false$
false$
```
