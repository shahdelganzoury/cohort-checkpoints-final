# Linked List Checkpoints

Ten linked list problems ordered from easiest to hardest. Each level has its own file with the problem, function signature, examples, constraints and a Go solution.

## Shared types

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}
```

## Levels

| Level | Problem | Function | Concept | Source | Focus |
|---|---|---|---|---|---|
| **1** | [Convert Binary Number in a Linked List to Integer](./level1.md) | `func getDecimalValue(head *ListNode) int` | Singly list traversal + accumulator | Test 1 (LeetCode 1290) | Read-only traversal, no pointer changes |
| **2** | [Transmission Queue Tail Append](./level2.md) | `func AppendNode(head *ListNode, val int) *ListNode` | Traversal & tail pointer insertion | Test 2 (Level 2) | Stop at `curr.Next == nil`; empty head creates the list |
| **3** | [Process Queue Node Purge](./level3.md) | `func DeleteNode(head *ListNode, val int) *ListNode` | Singly list deletion | Test 2 (Level 3) | Head deletion vs internal node unlinking |
| **4** | [Delete Node in a Linked List](./level4.md) | `func deleteNodeNoHead(node *ListNode)` | Singly list structural limits | Test 1 (LeetCode 237) | No access to `prev`, so copy the next node's value |
| **5** | [Memory Block Chain Builder](./level5.md) | `func BuildLinkedList(nums []int) *ListNode` | Singly list construction from a slice | Test 2 (Level 1) | Empty slice returns `nil`; dummy head + tail pointer for O(n) |
| **6** | [Reverse Linked List / Stream Packet Sequence Inversion](./level6.md) | `func ReverseLinkedList(head *ListNode) *ListNode` | In-place 3-pointer rewiring | Both tests (LeetCode 206) | `prev`, `curr`, `nextTemp` in O(1) space |
| **7** | [Intersection of Two Linked Lists](./level7.md) | `func getIntersectionNode(headA, headB *ListNode) *ListNode` | Multi-pointer traversal, reference equality | Test 1 (LeetCode 160) | Compare nodes by reference, not by value |
| **8** | [Linked List Cycle](./level8.md) | `func hasCycle(head *ListNode) bool` | Two pointers (Floyd's cycle detection) | Test 1 (LeetCode 141) | Fast and slow pointers moving at different speeds |
| **9** | [Remove Nth Node From End of List](./level9.md) | `func removeNthFromEnd(head *ListNode, n int) *ListNode` | One-pass two pointers + dummy head | Test 1 (LeetCode 19) | Pointer gap of n, removing the head as an edge case |
| **10** | [Bidirectional Cache Node Insertion](./level10.md) | `func InsertSortedDoubly(head *DoublyListNode, val int) *DoublyListNode` | Doubly linked list updates | Test 2 (Level 5) | Managing both `Next` and `Prev` at the head, middle and tail |
