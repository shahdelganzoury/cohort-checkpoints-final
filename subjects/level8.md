# Level 8: Linked List Cycle

**Source:** Test 1 (LeetCode 141)  
**Difficulty:** Easy  
**Focus:** Fast and slow pointers moving at different speeds

## Problem

Return `true` if following `Next` pointers from the head ever loops back to a node already visited, otherwise `false`.

## Function

```go
func hasCycle(head *ListNode) bool
```

## Examples

- `3 → 2 → 0 → -4`, tail pointing back to `2`: `true`
- `1 → 2`, tail pointing back to the head: `true`
- `1` with no cycle: `false`

## Constraints

0 to 10,000 nodes. Follow-up: use O(1) memory.

## Solution

```go
// ---------- 8. LeetCode 141: Cycle detection ----------
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```

**Key idea:** Slow moves 1 step, fast moves 2. If there is a cycle they must meet; otherwise fast reaches `nil`. O(n) time, O(1) space.
