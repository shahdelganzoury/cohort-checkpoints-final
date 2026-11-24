package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

// ---------- 1. LeetCode 1290: Binary number in linked list -> int ----------
func getDecimalValue(head *ListNode) int {
	result := 0
	for curr := head; curr != nil; curr = curr.Next {
		result = result*2 + curr.Val
	}
	return result
}

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

// ---------- 5. LeetCode 237: Delete node given only that node ----------
func deleteNodeNoHead(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

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

// ---------- 3. Build list from slice (dummy head + tail pointer, O(n)) ----------
func BuildLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, x := range nums {
		tail.Next = &ListNode{Val: x}
		tail = tail.Next
	}
	return dummy.Next
}

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

// ---------- 11. LeetCode 1472: Browser history (doubly linked list) ----------
type histNode struct {
	url        string
	prev, next *histNode
}

type BrowserHistory struct {
	curr *histNode
}

func NewBrowserHistory(homepage string) *BrowserHistory {
	return &BrowserHistory{curr: &histNode{url: homepage}}
}

func (b *BrowserHistory) Visit(url string) {
	n := &histNode{url: url, prev: b.curr}
	b.curr.next = n // drops all forward history
	b.curr = n
}

func (b *BrowserHistory) Back(steps int) string {
	for steps > 0 && b.curr.prev != nil {
		b.curr = b.curr.prev
		steps--
	}
	return b.curr.url
}

func (b *BrowserHistory) Forward(steps int) string {
	for steps > 0 && b.curr.next != nil {
		b.curr = b.curr.next
		steps--
	}
	return b.curr.url
}

// ---------- 12. LeetCode 142: Cycle entry ----------
func DetectCycleEntry(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// ---------- 13. LeetCode 138: Copy list with random pointer (O(1) space) ----------
type RNode struct {
	Val    int
	Next   *RNode
	Random *RNode
}

func copyRandomList(head *RNode) *RNode {
	if head == nil {
		return nil
	}
	// 1) interleave copies: A -> A' -> B -> B' ...
	for c := head; c != nil; c = c.Next.Next {
		c.Next = &RNode{Val: c.Val, Next: c.Next}
	}
	// 2) set random pointers on copies
	for c := head; c != nil; c = c.Next.Next {
		if c.Random != nil {
			c.Next.Random = c.Random.Next
		}
	}
	// 3) split the two lists
	newHead := head.Next
	for c := head; c != nil; {
		cp := c.Next
		c.Next = cp.Next
		if cp.Next != nil {
			cp.Next = cp.Next.Next
		}
		c = c.Next
	}
	return newHead
}

// ---------- 14. LeetCode 430: Flatten multilevel doubly linked list ----------
type MNode struct {
	Val   int
	Prev  *MNode
	Next  *MNode
	Child *MNode
}

func flatten(head *MNode) *MNode {
	for c := head; c != nil; c = c.Next {
		if c.Child == nil {
			continue
		}
		next := c.Next
		child := c.Child
		c.Next = child
		child.Prev = c
		c.Child = nil

		tail := child
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = next
		if next != nil {
			next.Prev = tail
		}
	}
	return head
}

// ---------- 15. LeetCode 146: LRU Cache (hash map + doubly linked list) ----------
type lruNode struct {
	key, val   int
	prev, next *lruNode
}

type LRUCache struct {
	cap        int
	m          map[int]*lruNode
	head, tail *lruNode // sentinels: head.next = most recent, tail.prev = least recent
}

func NewLRUCache(capacity int) *LRUCache {
	h, t := &lruNode{}, &lruNode{}
	h.next, t.prev = t, h
	return &LRUCache{cap: capacity, m: map[int]*lruNode{}, head: h, tail: t}
}

func (c *LRUCache) remove(n *lruNode) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) addFront(n *lruNode) {
	n.next = c.head.next
	n.prev = c.head
	c.head.next.prev = n
	c.head.next = n
}

func (c *LRUCache) Get(key int) int {
	n, ok := c.m[key]
	if !ok {
		return -1
	}
	c.remove(n)
	c.addFront(n)
	return n.val
}

func (c *LRUCache) Put(key, value int) {
	if n, ok := c.m[key]; ok {
		n.val = value
		c.remove(n)
		c.addFront(n)
		return
	}
	if len(c.m) == c.cap {
		lru := c.tail.prev
		c.remove(lru)
		delete(c.m, lru.key)
	}
	n := &lruNode{key: key, val: value}
	c.m[key] = n
	c.addFront(n)
}

// ---------- helpers + quick tests ----------
func toSlice(h *ListNode) []int {
	out := []int{}
	for ; h != nil; h = h.Next {
		out = append(out, h.Val)
	}
	return out
}

func main() {
	fmt.Println("1290:", getDecimalValue(BuildLinkedList([]int{1, 0, 1}))) // 5
	fmt.Println("append:", toSlice(AppendNode(BuildLinkedList([]int{1, 2}), 3)))
	fmt.Println("build empty is nil:", BuildLinkedList(nil) == nil)
	fmt.Println("delete head:", toSlice(DeleteNode(BuildLinkedList([]int{1, 2, 3}), 1)))
	fmt.Println("delete mid:", toSlice(DeleteNode(BuildLinkedList([]int{1, 2, 3}), 2)))

	l := BuildLinkedList([]int{4, 5, 1, 9})
	deleteNodeNoHead(l.Next) // delete 5
	fmt.Println("237:", toSlice(l))

	fmt.Println("reverse:", toSlice(ReverseLinkedList(BuildLinkedList([]int{1, 2, 3, 4}))))

	// cycle: 1->2->3->4->2
	c := BuildLinkedList([]int{1, 2, 3, 4})
	c.Next.Next.Next.Next = c.Next
	fmt.Println("hasCycle:", hasCycle(c), "entry val:", DetectCycleEntry(c).Val)
	fmt.Println("no cycle:", hasCycle(BuildLinkedList([]int{1, 2})))

	// intersection
	shared := BuildLinkedList([]int{8, 4, 5})
	a := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: shared}}
	b := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: shared}}}
	fmt.Println("intersection val:", getIntersectionNode(a, b).Val) // 8

	fmt.Println("remove 2nd from end:", toSlice(removeNthFromEnd(BuildLinkedList([]int{1, 2, 3, 4, 5}), 2)))
	fmt.Println("remove head via n=len:", toSlice(removeNthFromEnd(BuildLinkedList([]int{1, 2}), 2)))

	var d *DoublyListNode
	for _, v := range []int{5, 1, 3, 9, 7} {
		d = InsertSortedDoubly(d, v)
	}
	fwd, last := []int{}, d
	for n := d; n != nil; n = n.Next {
		fwd = append(fwd, n.Val)
		last = n
	}
	bwd := []int{}
	for n := last; n != nil; n = n.Prev {
		bwd = append(bwd, n.Val)
	}
	fmt.Println("doubly fwd:", fwd, "bwd:", bwd)

	bh := NewBrowserHistory("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")
	fmt.Println(bh.Back(1), bh.Back(1), bh.Forward(1))
	bh.Visit("linkedin.com")
	fmt.Println(bh.Forward(2), bh.Back(2), bh.Back(7))

	// random pointer: 1 -> 2, 1.random = 2, 2.random = 2
	r2 := &RNode{Val: 2}
	r1 := &RNode{Val: 1, Next: r2, Random: r2}
	r2.Random = r2
	cp := copyRandomList(r1)
	fmt.Println("copy ok:", cp != r1, cp.Random == cp.Next, cp.Next.Random == cp.Next, r1.Next == r2)

	// flatten: 1-2-3 with 2.child = 4-5
	n1, n2, n3 := &MNode{Val: 1}, &MNode{Val: 2}, &MNode{Val: 3}
	n4, n5 := &MNode{Val: 4}, &MNode{Val: 5}
	n1.Next, n2.Prev, n2.Next, n3.Prev = n2, n1, n3, n2
	n2.Child = n4
	n4.Next, n5.Prev = n5, n4
	out := []int{}
	for n := flatten(n1); n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	fmt.Println("flatten:", out)

	lru := NewLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))                         // 1
	lru.Put(3, 3)                                   // evicts 2
	fmt.Println(lru.Get(2))                         // -1
	lru.Put(4, 4)                                   // evicts 1
	fmt.Println(lru.Get(1), lru.Get(3), lru.Get(4)) // -1 3 4
}
