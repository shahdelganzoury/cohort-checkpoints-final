## hopelimination

### Instructions

The nodes form a **circular** singly linked list, and `start` is one of them. Starting at `start`, repeat until the circle is empty: remove the current node and record its `Val`, then move `Val` steps forward around the **remaining** circle. Step 1 is the node right after the removed one, and the node you land on becomes the current node.

Return the recorded values in removal order. An empty circle (`nil`) returns an empty slice.

Values can be very large, so walking `Val` steps one by one is too slow: **take it modulo the current circle size**.

Example 1:

```
Input: circle = [3,1,4,1,5], start at 3
Output: [3,1,5,1,4]
Explanation: Remove 3 and hop 3 steps (1, 4, 1) to the second 1. Remove it and hop 1 step to 5. Remove 5 and hop 5 steps around [1,4] to land on 1. Then remove 1, then 4.
```

Example 2:

```
Input: circle = [1000000000,7,8]
Output: [1000000000,8,7]
Explanation: After removing the first node, 2 nodes remain, so the hop is 1000000000 mod 2 â†’ 2 steps, landing on 8.
```

Example 3:

```
Input: circle = [9]
Output: [9]
Explanation: A single node is removed and the circle is empty.
```

Constraints:

* The number of nodes is in the range `[0, 5000]`.
* `1 <= Node.Val <= 1000000000`
* Follow up: never walk more than `size - 1` steps per hop.

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
func HopEliminate(start *ListNode) []int {

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
	fmt.Println(HopEliminate(ring(3, 1, 4, 1, 5)))
	fmt.Println(HopEliminate(ring(2, 2, 2, 2)))
	fmt.Println(HopEliminate(ring(1000000000, 7, 8)))
	fmt.Println(HopEliminate(ring(9)))
	fmt.Println(HopEliminate(nil))
}

func ring(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	tail := head
	for _, v := range vals[1:] {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	tail.Next = head
	return head
}
```

And its output:

```console
$ go run . | cat -e
[3 1 5 1 4]$
[2 2 2 2]$
[1000000000 8 7]$
[9]$
[]$
```
