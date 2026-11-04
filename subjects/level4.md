# Level 4: Delete Node in a Linked List

**Source:** Test 1 (LeetCode 237)  
**Difficulty:** Easy (labelled Medium in Test 1)  
**Focus:** No access to `prev`, so copy the next node's value

## Problem

You get only a reference to a node in a singly linked list, not the head. Delete that node. It is guaranteed not to be the tail. After the call the value must be gone, the list one node shorter, and the rest in the same order.

## Function

```go
func deleteNodeNoHead(node *ListNode)
```

## Examples

- `4 → 5 → 1 → 9`, node = `5` gives `4 → 1 → 9`
- `4 → 5 → 1 → 9`, node = `1` gives `4 → 5 → 9`

## Constraints

2 to 1000 nodes, unique values, the node is not the tail.

## Solution

```go
// ---------- 4. LeetCode 237: Delete node given only that node ----------
func deleteNodeNoHead(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
```

**Key idea:** You cannot reach `prev`, so copy the next node's value into this node and unlink the next node instead. O(1) time and space.
