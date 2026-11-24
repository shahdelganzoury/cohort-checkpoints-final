1290. Convert Binary Number in a Linked List to Integer (Easy)
Focus: Basic singly linked list traversal and accumulator logic.
Why it tests understanding: Verifies they can iterate through nodes sequentially without modifying the structure.
206. Reverse Linked List (Easy)
Focus: Fundamental single-pointer reversal.
Why it tests understanding: Tests if they understand how to rewire next pointers dynamically without losing references (prev, curr, next_temp).
141. Linked List Cycle (Easy)
Focus: Two-pointer technique (Floyd’s Cycle Finding Algorithm).
Why it tests understanding: Tests whether they grasp pointer movement at different speeds (fast/slow) rather than just linear iteration.
160. Intersection of Two Linked Lists (Easy)
Focus: Aligning list lengths and multi-pointer traversal.
Why it tests understanding: Verifies they understand reference equality vs. value equality and how memory addresses connect.
237. Delete Node in a Linked List (Medium)
Focus: Singly linked list structural limitations.
Why it tests understanding: Given only the target node (no head access), they must realize they cannot traverse backward to update prev.next, forcing a value-copying workaround.
1472. Design Browser History (Medium)
Focus: Direct Doubly Linked List implementation.
Why it tests understanding: Requires implementing forward and backward movement (prev and next), perfectly showing why doubly linked lists are useful for bidirectionality.
19. Remove Nth Node From End of List (Medium)
Focus: One-pass two-pointer deletion and dummy head usage.
Why it tests understanding: Tests edge cases (like removing the head node) and managing distance offsets between pointers.
138. Copy List with Random Pointer (Medium)
Focus: Complex multi-pointer node duplication.
Why it tests understanding: Deep-copying nodes with arbitrary random pointers requires either hash mapping or clever inline pointer interleave tricks.
430. Flatten a Multilevel Doubly Linked List (Medium)
Focus: Advanced Doubly Linked List pointer maintenance.
Why it tests understanding: Requires inserting nested sub-lists while keeping both prev and next pointers fully intact across all levels without breaking the chain.
146. LRU Cache (Medium / Hard)
Focus: Combining a Hash Map with a Doubly Linked List.
Why it tests understanding: The ultimate test of why DLLs exist. Demonstrates how $O(1)$ node removal and insertion combined with hash lookups build real-world data structures.