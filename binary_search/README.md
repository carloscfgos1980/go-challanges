# Challenge 21: Binary Search Implementation

Problem Statement
Implement the binary search algorithm to efficiently find items in a sorted collection. Binary search is a divide-and-conquer algorithm that repeatedly divides the search space in half, making it much faster than linear search for sorted data.

You'll implement three versions of binary search:

- BinarySearch - Standard binary search that returns the index of a target value.
- BinarySearchRecursive - A recursive implementation of binary search.
- FindInsertPosition - Find the position where a value should be inserted to maintain sorted order.

**Function Signatures**

func BinarySearch(arr []int, target int) int
func BinarySearchRecursive(arr []int, target int, left int, right int) int
func FindInsertPosition(arr []int, target int) int

**Input Format**

For all functions, a sorted slice of integers arr and a target integer value.
For the recursive function, additional left and right parameters indicating the search range.

**Output Format**

BinarySearch and BinarySearchRecursive should return the index of the target if found, or -1 if not found.

FindInsertPosition should return the index where the target should be inserted to maintain sorted order.

**Requirements**

All functions must implement the binary search algorithm, which has O(log n) time complexity.

The arrays can be assumed to be sorted in ascending order.
BinarySearchRecursive must use recursion to solve the problem.
If multiple occurrences of the target exist, return the index of any occurrence.