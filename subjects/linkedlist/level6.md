# Level 6: Reverse Linked List / Stream Packet Sequence Inversion

**Source:** Both tests (LeetCode 206)  
**Difficulty:** Easy  
**Focus:** `prev`, `curr`, `nextTemp` in O(1) space

## Problem

Reverse a singly linked list and return the new head.

## Function

```go
func ReverseLinkedList(head *ListNode) *ListNode
```

## Examples

- `1 → 2 → 3 → 4 → 5` gives `5 → 4 → 3 → 2 → 1`
- `1 → 2` gives `2 → 1`
- `nil` gives `nil`

## Constraints

0 to 5000 nodes. Do it in place. A recursive version is a common follow-up.

## Solution

```go
// ---------- 6. LeetCode 206: Reverse linked list (3 pointers) ----------
func ReverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}
```

**Key idea:** Save `Next` before overwriting it, point `curr` backwards, then advance both pointers. O(n) time, O(1) space.
