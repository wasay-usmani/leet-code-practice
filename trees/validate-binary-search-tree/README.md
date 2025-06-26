# Validate Binary Search Tree

LeetCode: [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)

## Problem
Given the `root` of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.

**Example:**

Input: root = [2,1,3]

Output: true

## Function Signature
```go
func isValidBST(root *TreeNode) bool
```
