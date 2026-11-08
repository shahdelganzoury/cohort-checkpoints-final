# Level 9: Remove Nth Node From End of List

**Source:** Test 1 (LeetCode 19)  
**Difficulty:** Medium  
**Focus:** Pointer gap of n, removing the head as an edge case

## Problem

Remove the nth node counting from the tail and return the head.

## Function

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode
```

## Examples

- `1 → 2 → 3 → 4 → 5`, n = 2 gives `1 → 2 → 3 → 5`
- `1`, n = 1 gives `nil`
- `1 → 2`, n = 1 gives `1`

## Constraints

1 to 30 nodes, `1 ≤ n ≤ length`. Follow-up: do it in one pass.

## Solution

```go
// ---------- 9. LeetCode 19: Remove nth node from end ----------
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
```

**Key idea:** Start `fast` n steps ahead, then move both until `fast` is on the last node. `slow` is then just before the target. The dummy head handles removing the real head. O(n) time, O(1) space.
