# Tower Breakers

HackerRank: [Tower Breakers](https://www.hackerrank.com/challenges/one-week-preparation-kit-tower-breakers-1/problem)

## Problem

Two players are playing a game of Tower Breakers! The rules are as follows:

- There are n towers, each of height m.
- The players move in alternating turns.
- In each turn, a player can choose a tower of height x > 1 and reduce its height to y, where 1 <= y < x and y evenly divides x.
- If a player cannot make a move, they lose.
- Player 1 always moves first.

Given n and m, determine which player wins the game. If both players play optimally, return 1 if Player 1 wins, or 2 if Player 2 wins.

**Function Signature:**
```go
func towerBreakers(n int32, m int32) int32
```
