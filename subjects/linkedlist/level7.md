# Level 7: Intersection of Two Linked Lists

**Source:** Test 1 (LeetCode 160)  
**Difficulty:** Easy  
**Focus:** Compare nodes by reference, not by value

## Problem

Given the heads of two singly linked lists, return the node where they merge into a shared tail, or `nil` if they never meet. The lists have no cycles and must keep their structure after the call.

## Function

```go
func getIntersectionNode(headA, headB *ListNode) *ListNode
```

## Examples

- `A = 4 → 1 → 8 → 4 → 5`, `B = 5 → 6 → 1 → 8 → 4 → 5`, sharing from the node `8`: returns that node (B's `1` is a different node with an equal value)
- `A = 2 → 6 → 4`, `B = 1 → 5`: returns `nil`

## Constraints

1 to 30,000 nodes per list. Follow-up: O(m + n) time, O(1) space.

## Solution

```go
// ---------- 7. LeetCode 160: Intersection of two lists ----------
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a // intersection node, or nil if none
}
```

**Key idea:** When a pointer reaches the end it jumps to the other list's head. Both then travel the same total distance and meet at the intersection, or both hit `nil`. O(m+n) time, O(1) space.
