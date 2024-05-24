# Data Structures in Golang

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## Project Description

This repository contains implementations of various data structures in Golang. The primary goal of this project is to facilitate learning and understanding of these data structures. The data structures implemented include:

- [Red-Black Tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree)
- [B Tree](https://en.wikipedia.org/wiki/B-tree)
- [Heap](https://en.wikipedia.org/wiki/Heap_(data_structure))
- [Fibonacci Heap](https://en.wikipedia.org/wiki/Fibonacci_heap)
- [AVL Tree](https://en.wikipedia.org/wiki/AVL_tree)

## Data Structures

### Red-Black Tree
![Red-Black Tree](https://upload.wikimedia.org/wikipedia/commons/6/66/Red-black_tree_example.svg)
A Red-Black Tree is a balanced binary search tree where each node contains an extra bit for denoting the color of the node, either red or black. By constraining the node colors and ensuring specific properties are maintained, Red-Black Trees offer efficient insertion, deletion, and lookup operations, all in O(log n) time.

### B Tree
![B Tree](https://cdn.programiz.com/sites/tutorial2program/files/b-tree.png)

A B Tree is a self-balancing tree data structure that maintains sorted data and allows for efficient insertion, deletion, and search operations. Each node in a B Tree can contain more than one key and can have multiple children, making it highly efficient for systems that read and write large blocks of data.

### Heap
![Heap](https://media.geeksforgeeks.org/wp-content/uploads/20230323095300/maxh.png)

A Heap is a specialized tree-based data structure that satisfies the heap property: in a max heap, for any given node I, the value of I is greater than or equal to the values of its children; in a min heap, the value of I is less than or equal to the values of its children. Heaps are commonly used in priority queues and for implementing efficient sorting algorithms like heapsort.

### Fibonacci Heap
![Fibonacci Heap](https://www.programiz.com/sites/tutorial2program/files/fibonacci-heap.png)

A Fibonacci Heap is a collection of trees that are structured in a way that supports a set of heap operations that have better amortized time complexity than other heap data structures. It is particularly useful for algorithms that require repeated decrease-key operations, such as Dijkstra's algorithm for shortest paths.

### AVL Tree
![AVL Tree](https://miro.medium.com/v2/resize:fit:1024/0*Vi3aQ9sY9Yu4VNpa.png)

An AVL Tree is a self-balancing binary search tree where the difference between the heights of the left and right subtrees of any node is no more than one. This balance ensures that the AVL Tree maintains O(log n) time complexity for search, insert, and delete operations, making it highly efficient for dynamic data sets.

## How to Run the Code

To run the code for any of the data structures, use the following command:

```bash
go run main.go
```

Make sure you have [Golang](https://go.dev/) installed on your machine.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## References

- [Golang Documentation](https://go.dev/doc/)
- [Red-Black Tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree)
- [B Tree](https://en.wikipedia.org/wiki/B-tree)
- [Heap](https://en.wikipedia.org/wiki/Heap_(data_structure))
- [Fibonacci Heap](https://en.wikipedia.org/wiki/Fibonacci_heap)
- [AVL Tree](https://en.wikipedia.org/wiki/AVL_tree)

