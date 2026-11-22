## intervallist

### Instructions

Each node holds a closed interval `[Start, End]`. The list is kept **sorted by `Start`**, and no two of its intervals overlap. Given the `head` and a new interval `[start, end]`, add it and **merge every interval it overlaps on the fly**. Return the head.

Intervals that touch also merge: `[1,3]` and `[3,5]` become `[1,5]`, but `[1,2]` and `[3,4]` stay apart. One new interval can swallow many existing ones.

Create **at most one new node per call**, and unlink the nodes you merge away. If `start > end`, return the list unchanged.

Example 1:

```
Input: head = [[1,3],[5,6],[10,12]], add [3,4]
Output: [[1,4],[5,6],[10,12]]
Explanation: [3,4] touches [1,3], so they merge into [1,4].
```

Example 2:

```
Input: head = [[1,4],[5,6],[7,7],[8,9],[10,12]], add [2,11]
Output: [[1,12]]
Explanation: [2,11] overlaps every interval, so they all collapse into one.
```

Example 3:

```
Input: head = [], add [4,4]
Output: [[4,4]]
Explanation: The new interval becomes the only node.
```

Constraints:

* The number of nodes is in the range `[0, 10000]`.
* `-1000000000 <= start, end <= 1000000000`
* Follow up: one pass, `O(n)` time per call and `O(1)` extra memory.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Start int
	End   int
	Next  *ListNode
}
```

### Expected function

```go
func AddInterval(head *ListNode, start, end int) *ListNode {

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
	var list *ListNode
	steps := [][2]int{{10, 12}, {1, 3}, {5, 6}, {3, 4}, {8, 9}, {7, 7}, {2, 11}, {20, 25}, {13, 13}, {9, 1}}
	for _, s := range steps {
		list = AddInterval(list, s[0], s[1])
		fmt.Printf("add [%d,%d]: ", s[0], s[1])
		show(list)
	}
	show(AddInterval(nil, 4, 4))
}

func show(h *ListNode) {
	for ; h != nil; h = h.Next {
		fmt.Printf("[%d,%d] -> ", h.Start, h.End)
	}
	fmt.Println("nil")
}
```

And its output:

```console
$ go run . | cat -e
add [10,12]: [10,12] -> nil$
add [1,3]: [1,3] -> [10,12] -> nil$
add [5,6]: [1,3] -> [5,6] -> [10,12] -> nil$
add [3,4]: [1,4] -> [5,6] -> [10,12] -> nil$
add [8,9]: [1,4] -> [5,6] -> [8,9] -> [10,12] -> nil$
add [7,7]: [1,4] -> [5,6] -> [7,7] -> [8,9] -> [10,12] -> nil$
add [2,11]: [1,12] -> nil$
add [20,25]: [1,12] -> [20,25] -> nil$
add [13,13]: [1,12] -> [13,13] -> [20,25] -> nil$
add [9,1]: [1,12] -> [13,13] -> [20,25] -> nil$
[4,4] -> nil$
```
