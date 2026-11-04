# Level 10: Bidirectional Cache Node Insertion

**Source:** Test 2 (Level 5)  
**Difficulty:** Medium  
**Focus:** Managing both `Next` and `Prev` at the head, middle and tail

## Problem

Given the head of a sorted (ascending) doubly linked list, insert a node with `val` so the list stays sorted and every `Prev` and `Next` pointer is correct. The head's `Prev` and the tail's `Next` are `nil`. Return the head, which changes if the new value is smallest.

## Function

```go
func InsertSortedDoubly(head *DoublyListNode, val int) *DoublyListNode
```

## Examples

- `nil`, val `5` gives `5`
- `3 ⇄ 7`, val `1` gives `1 ⇄ 3 ⇄ 7` (new head)
- `3 ⇄ 7`, val `5` gives `3 ⇄ 5 ⇄ 7` (middle)
- `3 ⇄ 7`, val `9` gives `3 ⇄ 7 ⇄ 9` (new tail)

## Constraints

Duplicates are inserted before existing equal values in this solution.

## Solution

```go
// ---------- 10. Insert into sorted doubly linked list ----------
func InsertSortedDoubly(head *DoublyListNode, val int) *DoublyListNode {
	n := &DoublyListNode{Val: val}
	if head == nil {
		return n
	}
	if val <= head.Val { // new head
		n.Next = head
		head.Prev = n
		return n
	}
	curr := head
	for curr.Next != nil && curr.Next.Val < val {
		curr = curr.Next
	}
	n.Prev = curr
	n.Next = curr.Next
	if curr.Next != nil { // middle (not tail)
		curr.Next.Prev = n
	}
	curr.Next = n
	return head
}
```

**Key idea:** Handle the new head first, then find the node just before the insertion point and update four links: `n.Prev`, `n.Next`, `next.Prev` (only if it exists) and `prev.Next`. O(n) time, O(1) space.
