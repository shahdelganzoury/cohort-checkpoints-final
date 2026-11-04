1	1290. Convert Binary Number in a Linked List to Integer	Test 1

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

2	Transmission Queue Tail Append	Test 2 (Level 2)

```go
// ---------- 2. Tail append ----------
func AppendNode(head *ListNode, val int) *ListNode {
	n := &ListNode{Val: val}
	if head == nil {
		return n
	}
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
	return head
}
```

3	Process Queue Node Purge (delete by value)	Test 2 (Level 3)

```go
// ---------- 5. LeetCode 237: Delete node given only that node ----------
func deleteNodeNoHead(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
```

4	237. Delete Node in a Linked List	Test 1

```go
// ---------- 4. Delete first node with value val ----------
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

5	Memory Block Chain Builder	Test 2 (Level 1)

```go
// ---------- 3. Build list from slice (dummy head + tail pointer, O(n)) ----------
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

6	206. Reverse Linked List / Stream Packet Sequence Inversion	Both

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

7	160. Intersection of Two Linked Lists	Test 1

```go
// ---------- 8. LeetCode 160: Intersection of two lists ----------
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

8	141. Linked List Cycle	Test 1

```go
// ---------- 7. LeetCode 141: Cycle detection ----------
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

9	19. Remove Nth Node From End of List	Test 1

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

10	Bidirectional Cache Node Insertion	Test 2 (Level 5)

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