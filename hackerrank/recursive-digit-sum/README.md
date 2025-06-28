# Recursive Digit Sum (Super Digit)

HackerRank: [Recursive Digit Sum](https://www.hackerrank.com/challenges/one-week-preparation-kit-recursive-digit-sum/problem)

## Problem
Given an integer as a string `n` and an integer `k`, create a new number `p` by concatenating the string `n` to itself `k` times. The super digit of `p` is defined recursively as follows:
- If `p` has only one digit, its super digit is `p`.
- Otherwise, the super digit of `p` is the super digit of the sum of the digits of `p`.

Return the super digit of `p`.

**Function Signature:**
```go
func superDigit(n string, k int32) int32
```
