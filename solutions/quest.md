Here are all 15 problems in the final order. The LeetCode statements are paraphrased from memory, so check exact limits on LeetCode. Test 2's problems (#2, #3, #4, #10, #12) are my reconstruction from the signatures and descriptions you gave me, since I don't have the original level files.

---

## 1. Convert Binary Number in a Linked List to Integer (LeetCode 1290, Test 1)

**Task:** Each node of a singly linked list holds a 0 or a 1. Together they form a binary number, with the most significant bit at the head. Return its decimal value.

**Function:** `getDecimalValue(head *ListNode) int`

**Examples**
- `1 → 0 → 1` returns `5`
- `0` returns `0`

**Constraints:** 1 to 30 nodes, each value 0 or 1.

---

## 2. Transmission Queue Tail Append (Test 2, Level 2)

**Task:** Add a new node holding `val` at the end of a singly linked list and return the head. If the list is empty, the new node becomes the head.

**Function:** `AppendNode(head *ListNode, val int) *ListNode`

**Examples**
- `1 → 2 → 3`, val `4` gives `1 → 2 → 3 → 4`
- `nil`, val `5` gives `5`

**Focus:** Stop at `curr.Next == nil`, not at `curr == nil`, and handle the empty list.

---

## 5. Delete Node in a Linked List (LeetCode 237, Test 1)

**Task:** You get only a reference to a node in a singly linked list, not the head. Delete that node. It is guaranteed not to be the tail, and all values are unique. After the call, the value must be gone, the list must be one node shorter, and the order of everything else must be unchanged.

**Function:** `deleteNodeNoHead(node *ListNode)` (returns nothing)

**Examples**
- `4 → 5 → 1 → 9`, node = `5` gives `4 → 1 → 9`
- `4 → 5 → 1 → 9`, node = `1` gives `4 → 5 → 9`

**Constraints:** 2 to 1000 nodes, unique values, the node is not the tail.

---

## 4. Process Queue Node Purge (Test 2, Level 3)

**Task:** Remove the first node whose value equals `val` and return the head. If no node matches, return the list unchanged.

**Function:** `DeleteNode(head *ListNode, val int) *ListNode`

**Examples**
- `1 → 2 → 3`, val `1` gives `2 → 3` (head deleted)
- `1 → 2 → 3`, val `2` gives `1 → 3` (middle node)
- `1 → 2 → 3`, val `9` gives `1 → 2 → 3` (not found)
- `nil`, val `1` gives `nil`

**Focus:** Deleting the head is the special case. A dummy head removes it. The original may expect deleting all matches, so check that.

---

## 3. Memory Block Chain Builder (Test 2, Level 1)

**Task:** Given a slice of integers, build a singly linked list with the same order and return its head. An empty or nil slice returns `nil`.

**Function:** `BuildLinkedList(nums []int) *ListNode`

**Examples**
- `[3, 7, 9]` gives `3 → 7 → 9`
- `[]` gives `nil`

**Focus:** Aim for O(n) time. Calling `AppendNode` in a loop is O(n²), so keep a tail pointer or dummy head.

---

## 6. Reverse Linked List (LeetCode 206, Both tests)

**Task:** Reverse a singly linked list and return the new head.

**Function:** `ReverseLinkedList(head *ListNode) *ListNode`

**Examples**
- `1 → 2 → 3 → 4 → 5` gives `5 → 4 → 3 → 2 → 1`
- `1 → 2` gives `2 → 1`
- `nil` gives `nil`

**Constraints:** 0 to 5000 nodes. Do it in place with O(1) extra space. A recursive version is a common follow-up.

---

## 8. Intersection of Two Linked Lists (LeetCode 160, Test 1)

**Task:** Given the heads of two singly linked lists, return the node where they merge into a shared tail, or `nil` if they never meet. The lists have no cycles and must keep their structure after the call.

**Function:** `getIntersectionNode(headA, headB *ListNode) *ListNode`

**Examples**
- `A = 4 → 1 → 8 → 4 → 5`, `B = 5 → 6 → 1 → 8 → 4 → 5`, sharing from the node `8`: returns that node. B's `1` is *not* the intersection, since it is a different node with an equal value.
- `A = 1 → 9 → 1 → 2 → 4`, `B = 3 → 2 → 4`, sharing from `2`: returns that node.
- `A = 2 → 6 → 4`, `B = 1 → 5`: returns `nil`.

**Constraints:** Each list has 1 to 30,000 nodes. Follow-up: O(m + n) time and O(1) space.

---

## 7. Linked List Cycle (LeetCode 141, Test 1)

**Task:** Return `true` if following `Next` pointers from the head ever loops back to a node already visited, otherwise `false`.

**Function:** `hasCycle(head *ListNode) bool`

**Examples**
- `3 → 2 → 0 → -4`, with the tail pointing back to the node `2`: `true`
- `1 → 2`, with the tail pointing back to the head: `true`
- `1` with no cycle: `false`

**Constraints:** 0 to 10,000 nodes. Follow-up: use O(1) memory.

---

## 9. Remove Nth Node From End of List (LeetCode 19, Test 1)

**Task:** Remove the nth node counting from the tail and return the head.

**Function:** `removeNthFromEnd(head *ListNode, n int) *ListNode`

**Examples**
- `1 → 2 → 3 → 4 → 5`, n = 2 gives `1 → 2 → 3 → 5`
- `1`, n = 1 gives `nil`
- `1 → 2`, n = 1 gives `1`

**Constraints:** 1 to 30 nodes, `1 ≤ n ≤ length`. Follow-up: do it in one pass.

---

## 10. Bidirectional Cache Node Insertion (Test 2, Level 5)

**Task:** You get the head of a sorted (ascending) doubly linked list. Insert a node with `val` so the list stays sorted and every `Prev` and `Next` pointer is correct. The head's `Prev` and the tail's `Next` are `nil`. Return the head, which changes if the new value is smallest.

**Function:** `InsertSortedDoubly(head *DoublyListNode, val int) *DoublyListNode`

**Examples**
- `nil`, val `5` gives `5`
- `3 ⇄ 7`, val `1` gives `1 ⇄ 3 ⇄ 7` (new head)
- `3 ⇄ 7`, val `5` gives `3 ⇄ 5 ⇄ 7` (middle)
- `3 ⇄ 7`, val `9` gives `3 ⇄ 7 ⇄ 9` (new tail)

**Focus:** The new node needs `Prev` and `Next`, and its neighbors need updating too. My solution puts duplicates before existing equal values, so state your own rule if the original differs.

---

## 11. Design Browser History (LeetCode 1472, Test 1)

**Task:** Build a class that simulates browser history.

- `BrowserHistory(homepage)` starts on the homepage.
- `visit(url)` goes to `url` from the current page and clears all forward history.
- `back(steps)` moves back at most `steps` pages, stopping at the oldest, and returns the current URL.
- `forward(steps)` moves forward at most `steps` pages, stopping at the newest, and returns the current URL.

**Example**
```
BrowserHistory("leetcode.com")
visit("google.com"), visit("facebook.com"), visit("youtube.com")
back(1)       → "facebook.com"
back(1)       → "google.com"
forward(1)    → "facebook.com"
visit("linkedin.com")        // forward history cleared
forward(2)    → "linkedin.com"
back(2)       → "google.com"
back(7)       → "leetcode.com"
```

**Constraints:** URLs are 1 to 20 characters, `1 ≤ steps ≤ 100`, at most 5000 calls.

---

## 12. Infinite Router Loop Interceptor (LeetCode 142, Test 2, Level 6)

**Task:** Given the head of a singly linked list that may contain a cycle, return the node where the cycle begins, or `nil` if there is no cycle. Do not modify the list.

**Function:** `DetectCycleEntry(head *ListNode) *ListNode`

**Examples**
- `3 → 2 → 0 → -4`, tail pointing back to `2`: returns the node `2`
- `1 → 2`, tail pointing back to `1`: returns the node `1`
- `1` with no cycle: returns `nil`

**Constraints:** 0 to 10,000 nodes. Follow-up: use O(1) space.

---

## 13. Copy List with Random Pointer (LeetCode 138, Test 1)

**Task:** Each node has a `Next` pointer and a `Random` pointer, which may point to any node in the list or be `nil`. Build a deep copy: `n` brand-new nodes with the same values, where the copy's `Next` and `Random` pointers form the same structure but point only to nodes in the copy, never to the original.

**Function:** `copyRandomList(head *RNode) *RNode`

**Examples** (each node written as `[value, index of Random target or null]`)
- `[[7,null],[13,0],[11,4],[10,2],[1,0]]` gives a copy with the same shape
- `[[1,1],[2,1]]` gives a copy with the same shape
- `[[3,null],[3,0],[3,null]]` gives a copy with the same shape (equal values, so you can't match nodes by value)

**Constraints:** 0 to 1000 nodes, values from -10,000 to 10,000. Follow-up: O(1) extra space.

---

## 14. Flatten a Multilevel Doubly Linked List (LeetCode 430, Test 1)

**Task:** Each node has `Prev`, `Next` and `Child`. A `Child` points to the head of another doubly linked list, which can have children of its own. Flatten everything into a single-level doubly linked list. A node's child list goes right after that node and before its old `Next`. Set every `Child` to `nil` and keep all `Prev` pointers correct.

**Function:** `flatten(head *MNode) *MNode`

**Examples**
- Level 1 is `1-2-3-4-5-6`, node 3 has child list `7-8-9-10`, and node 8 has child list `11-12`. The result is `1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6`.
- Level 1 is `1-2`, and node 1 has child `3`. The result is `1, 3, 2`.
- `nil` returns `nil`.

**Constraints:** 0 to 1000 nodes.

---

## 15. LRU Cache (LeetCode 146, Test 1)

**Task:** Design a cache with a fixed capacity that evicts the least recently used entry when full.

- `LRUCache(capacity)` creates the cache.
- `get(key)` returns the value, or `-1` if absent. It counts as a use.
- `put(key, value)` updates the key if it exists, otherwise inserts it. If this exceeds capacity, evict the least recently used key first. It counts as a use.

Both operations must run in O(1) average time.

**Example**
```
LRUCache(2)
put(1,1), put(2,2)
get(1)   → 1
put(3,3)          // evicts key 2
get(2)   → -1
put(4,4)          // evicts key 1
get(1)   → -1
get(3)   → 3
get(4)   → 4
```

**Constraints:** `1 ≤ capacity ≤ 3000`, keys 0 to 10,000, values 0 to 100,000, up to 200,000 calls.

---

I can put these into a Word doc or PDF if you want them as a printable set.