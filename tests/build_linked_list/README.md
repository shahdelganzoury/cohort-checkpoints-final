## buildlinkedlist

### Instructions

Given a slice of integers `nums`, build a singly-linked list that contains the same values in the same order, and return its head.

If `nums` is empty or `nil`, return `nil`.

Your function should run in `O(n)` time.

Example 1:

```
Input: nums = [3,7,9]
Output: [3,7,9]
```

Example 2:

```
Input: nums = []
Output: []
Explanation: An empty slice produces an empty list (nil).
```

Constraints:

* `0 <= len(nums) <= 1000`
* `-1000 <= nums[i] <= 1000`

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
func BuildLinkedList(nums []int) *ListNode {

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
	printList(BuildLinkedList([]int{3, 7, 9}))
	printList(BuildLinkedList([]int{42}))
	printList(BuildLinkedList([]int{}))
	printList(BuildLinkedList(nil))
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
3 -> 7 -> 9$
42$
nil$
nil$
```
