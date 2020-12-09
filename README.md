# ADS #
## Introduction ##
ADS(Advancing data structre) is learning workspace/lab to understand data structure and algorithms. Programming languge used is GO(Golang).

## Complaxity study summary: ##

Dynamic Array:
* Double the size of new array and copy old array once the original array is full.
Get: O(1)ST
Set: O(1)ST
Insert: -end- O(1) after amortise analysis,
	     -Start/Middle) - O(n)
Traverse: O(n)T, O(1)S

Linklist:
* Does not require memory allocation one after other, point next node to current node
	Get/Set: O(i)ST
	Insert: begening: O(1)ST; ending: O(1)ST; middle: O(i)T O(1)S
	initialise/Copy: O(n)ST
	Traverse: O(n)T O(1)S

HashTable:
* can store any key datatype unlike array, later key has to be converted in int by using hash function.
* Collision: if has keys are same it stores in a linklist with key value pair on same hash index.
* Resizing: hash table is smart enough to identify the resize, based on array size and linklist size.
	Insert/Delete/Search: O(1)ST, O(n)-worst case
	Initialising: O(n)ST

Stack:
* LIFO (last in first out) -
* implement using dynamic array, so amortise analysis is constant for both add/delete(push/pop) operation(since both operation is on last index of the array).
	Example: book store;
	insert/delete: O(1)ST
	Searching: O(n)

Queue:
* FIFO (first in first out),
* can’t implement using dynamic array to achieve the const complexity on both insert/delete (enqueue/dequeue) operation. So has to implement using linked list.
	linked list implementation : has to have the tail pointer to delete(dequeue) at the last of the queue.
	Example: line for booking tickets,
	insert/delete: O(1)ST
	Searching: O(n)T, O(1)S
	initialising: O(1); because initially it would be length of zero and after every operation it takes constant time

String:
* Strings are immutable in most of the language.
* Instead concating string on many times in an operation , convert it to first byte array and later convert back to string to minimise the complexity cost. (since concat string is an O(n) complexity).
	get: O(1)ST
	set: can’t override an index
	Traverse: O(n)T, O(1)S
	Copy: O(n) ST

Graph:
	Vertex, Edges, Connected/Disconnected Graph, Directed/Undirected Graph(flights)(friendship), cyclic graph - avoid cycle by marking node as visited,
	Store: O(V+E)S
	Traverse: DFS, BFS, O(V+E)T

Trees:
* Sort of Graph; Never has a cycle; Only one parent for each node; Not allowed to have disconnected tree.
* Example: Company hierarchy, family hierarchy.

Binary Tree, BST, Min/Max Heap, Ternary Tree(most 3 child node),K-Ary Tree, Balance Tree(RedBlack Tree, AVL Tree).
Full Tree: all node has to fill from left
Complete Tree: all node has to have 2 nodes except child
Perfect Tree: all leaf node has same depth
Store: O(n)ST
Traverse: O(n)T, O(log(n)) from root to one path

### Resources ###

# DATA STRUCTURE #
* https://www.hackerearth.com/practice/data-structures/

# TREE #
* BinaryTree: https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/992/
* https://www.geeksforgeeks.org/binary-tree-data-structure/
* https://www.hackerearth.com/practice/data-structures/trees/binary-and-nary-trees/tutorial/

# HEAP #
* https://www.geeksforgeeks.org/heap-data-structure/
* https://www.hackerearth.com/practice/data-structures/trees/heapspriority-queues/tutorial/
