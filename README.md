# Shopify Online Assessment - Coding Challenge

## Description

This repository contains my submission for the Shopify infrastructure engineering online assessment coding challenge. The challenge involved is a Go program that solves a given Sudoku puzzle. The program takes a 9x9 grid as input, where empty cells are represented by zeros, and output the solved Sudoku puzzle.

### Example: Input:
```
[
  [5, 3, 0, 0, 7, 0, 0, 0, 0],
  [6, 0, 0, 1, 9, 5, 0, 0, 0],
  [0, 9, 8, 0, 0, 0, 0, 6, 0],
  [8, 0, 0, 0, 6, 0, 0, 0, 3],
  [4, 0, 0, 8, 0, 3, 0, 0, 1],
  [7, 0, 0, 0, 2, 0, 0, 0, 6],
  [0, 6, 0, 0, 0, 0, 2, 8, 0],
  [0, 0, 0, 4, 1, 9, 0, 0, 5],
  [0, 0, 0, 0, 8, 0, 0, 7, 9]
]
```

### Program Output:
```
[
  [5, 3, 4, 6, 7, 8, 9, 1, 2],
  [6, 7, 2, 1, 9, 5, 3, 4, 8],
  [1, 9, 8, 3, 4, 2, 5, 6, 7],
  [8, 5, 9, 7, 6, 1, 4, 2, 3],
  [4, 2, 6, 8, 5, 3, 7, 9, 1],
  [7, 1, 3, 9, 2, 4, 8, 5, 6],
  [9, 6, 1, 5, 3, 7, 2, 8, 4],
  [2, 8, 7, 4, 1, 9, 6, 3, 5],
  [3, 4, 5, 2, 8, 6, 1, 7, 9]
]
```

## Implementation

The implementation of the Sudoku solver program can be found in the `sudoku.go` file.

My initial approach to this problem involved using a brute-force algorithm to generate all possible ways to fill up the Sudoku grid. However, I quickly realized that this would be extremely inefficient with 9^81 possibilities.

Instead, a more efficient approach to solving this problem can be achieved using backtracking. Backtracking is an algorithmic technique that explores different possibilities and "backtracks" when it realizes that the current path of choices does not lead to a valid Sudoku solution. In our case, the backtracking algorithm continues until a Sudoku solution is found or until all possibilities are exhausted.

My Sudoku solver implementation consists of three methods: `SolveSudoku`, `computeSudoku`, and `validSudoku`.

The `SolveSudoku` method is the entry point for solving the Sudoku puzzle. It calls the `computeSudoku` method to start the solving process.

The `computeSudoku` method is the backtracking function using recursion that attempts to fill in the Sudoku board. It uses a nested loop to iterate over each cell in the board. If the current cell is empty with a zero, it tries numbers from 1 to 9 in a trial-and-error manner. For each trial, it checks if placing the current digit is valid using the `validSudoku` method. If the board is not valid, it backtracks and tries another possibility until a solution is found.

The `validSudoku` method checks if placing a particular digit at a given position (row, col) is valid according to Sudoku rules. It verifies that the digit does not appear in the same row, column, or 3x3 block.

## Challenges

Some challenges I faced during this coding challenge included verifying the correct indices for every 3x3 grid. However, I realized that using integer division and the modulo operator, there is a cleaver way to iterate through each subgrid. Furthermore, I struggled to validate my backtracking solution since there was only one given test case. To address this, I created multiple custom test cases to ensure the correctness of my Sudoku solver.

