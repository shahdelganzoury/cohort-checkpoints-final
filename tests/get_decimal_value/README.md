## getdecimalvalue

### Instructions

Given `head` which is a reference node to a singly-linked list. The value of each node in the linked list is either `0` or `1`. The linked list holds the binary representation of a number.

Return the decimal value of the number in the linked list.

The most significant bit is at the head of the linked list.

Example 1:

```
Input: head = [1,0,1]
Output: 5
Explanation: (101) in base 2 = (5) in base 10
```

Example 2:

```
Input: head = [0]
Output: 0
```

Constraints:

* The Linked List is not empty.
* Number of nodes will not exceed `30`.
* Each node's value is either `0` or `1`.

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
func GetDecimalValue(head *ListNode) int {

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
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}
	head2 := &ListNode{Val: 0}
	head3 := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}

	fmt.Println(GetDecimalValue(head1))
	fmt.Println(GetDecimalValue(head2))
	fmt.Println(GetDecimalValue(head3))
}
```

And its output:

```console
$ go run . | cat -e
5$
0$
7$
```
