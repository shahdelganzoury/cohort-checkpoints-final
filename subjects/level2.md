# Level 2: Transmission Queue Tail Append

**Source:** Test 2 (Level 2)  
**Difficulty:** Easy  
**Focus:** Stop at `curr.Next == nil`; empty head creates the list

## Problem

Add a new node holding `val` at the end of a singly linked list and return the head. If the list is empty, the new node becomes the head.

## Function

```go
func AppendNode(head *ListNode, val int) *ListNode
```

## Examples

- `1 → 2 → 3`, val `4` gives `1 → 2 → 3 → 4`
- `nil`, val `5` gives `5`

## Constraints

Any list length, including 0.

## Solution

```go
// ---------- 2. Tail append ----------
func AppendNode(head *ListNode, val int) *ListNode {
	n := &ListNode{Val: val}
	if head == nil {
		return n
	}
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
	return head
}
```

**Key idea:** Walk to the last node (`curr.Next == nil`, not `curr == nil`) and attach. O(n) time, O(1) space.
