# Shopify Intern Assessment Production Engineering

## Description:
Author: Hridyansh Dugar

This Go program solves the given Sudoku puzzle using a backtracking algorithm. The program takes a 9x9 grid as an input(anything else will fail the validation check). The program then creates an empty grid of size 9x9 and copies the grid to the newly created grid and uses the back track propagation algorithm on it to solve the sudoku, if no solution is found it returns nil else it returns the solved grid.

## Testing:
Additional test cases have been added to check for the algorithm's correctness, tests to check for validation error and cases where no solution is found are also added.

## Original Idea:
Originally I had an idea to implement a pipelined version of the backtracking algorithm where many subroutines would concurrently try to solve the sudoku but since Im fairly new to Go (I learnt it while coding this assessment out), i wasnt able to implement that well, and i found that concurrency is well abstracted in Go which is a good thing but it makes debugging harder for beginners. Overall i had fun implementing the sudoku solver in a language new to me :) !

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
