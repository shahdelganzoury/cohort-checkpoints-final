# Level 3: Process Queue Node Purge

**Source:** Test 2 (Level 3)  
**Difficulty:** Easy  
**Focus:** Head deletion vs internal node unlinking

## Problem

Remove the first node whose value equals `val` and return the head. If no node matches, return the list unchanged.

## Function

```go
func DeleteNode(head *ListNode, val int) *ListNode
```

## Examples

- `1 → 2 → 3`, val `1` gives `2 → 3` (head)
- `1 → 2 → 3`, val `2` gives `1 → 3` (middle)
- `1 → 2 → 3`, val `9` gives `1 → 2 → 3` (not found)
- `nil`, val `1` gives `nil`

## Constraints

Any list length, including 0.

## Solution

```go
// ---------- 3. Delete first node with value val ----------
func DeleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return dummy.Next
}
```

**Key idea:** A dummy head removes the special case for deleting the head. Stop one node before the target. O(n) time, O(1) space.
