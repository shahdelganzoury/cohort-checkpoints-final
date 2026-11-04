# Level 1: Convert Binary Number in a Linked List to Integer

**Source:** Test 1 (LeetCode 1290)  
**Difficulty:** Easy  
**Focus:** Read-only traversal, no pointer changes

## Problem

Each node holds 0 or 1. Together they form a binary number with the most significant bit at the head. Return its decimal value.

## Function

```go
func getDecimalValue(head *ListNode) int
```

## Examples

- `1 → 0 → 1` returns `5`
- `0` returns `0`

## Constraints

1 to 30 nodes, each value 0 or 1.

## Solution

```go
// ---------- 1. LeetCode 1290: Binary number in linked list -> int ----------
func getDecimalValue(head *ListNode) int {
	result := 0
	for curr := head; curr != nil; curr = curr.Next {
		result = result*2 + curr.Val
	}
	return result
}
```

**Key idea:** Shift the accumulator left (`*2`) and add the current bit. O(n) time, O(1) space.
