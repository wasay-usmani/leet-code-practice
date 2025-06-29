# Max Area of Island

LeetCode: [Max Area of Island](https://leetcode.com/problems/max-area-of-island/)

## Problem
Given a non-empty 2D binary grid grid of 0's (representing water) and 1's (representing land), an island is a maximal 4-directionally connected group of 1's. Return the area of the largest island in the grid. If there is no island, return 0.

**Example:**

Input: grid = [
  [0,0,1,0,0,0,0,1,0,0,0,0,0],
  [0,0,0,0,0,0,0,1,1,1,0,0,0],
  [0,1,1,0,1,0,0,0,0,0,0,0,0],
  [0,1,0,0,1,1,0,0,1,0,1,0,0],
  [0,1,0,0,1,1,0,0,1,1,1,0,0],
  [0,0,0,0,0,0,0,0,0,0,1,0,0],
  [0,0,0,0,0,0,0,1,1,1,0,0,0],
  [0,0,0,0,0,0,0,1,1,0,0,0,0]
]
Output: 6

## Function Signature
```go
func maxAreaOfIsland(grid [][]int) int
```
