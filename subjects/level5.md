# Level 5: Memory Block Chain Builder

**Source:** Test 2 (Level 1)  
**Difficulty:** Easy  
**Focus:** Empty slice returns `nil`; dummy head + tail pointer for O(n)

## Problem

Given a slice of integers, build a singly linked list in the same order and return its head. An empty or nil slice returns `nil`.

## Function

```go
func BuildLinkedList(nums []int) *ListNode
```

## Examples

- `[3, 7, 9]` gives `3 → 7 → 9`
- `[]` gives `nil`

## Constraints

Aim for O(n). Calling `AppendNode` in a loop is O(n²).

## Solution

```go
// ---------- 5. Build list from slice (dummy head + tail pointer, O(n)) ----------
func BuildLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	tail := dummy
	for _, x := range nums {
		tail.Next = &ListNode{Val: x}
		tail = tail.Next
	}
	return dummy.Next
}
```

**Key idea:** Keep a `tail` pointer starting at a dummy node, attach each new node, and return `dummy.Next`. O(n) time, O(n) space for the new nodes.
