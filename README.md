# Technical Instructions
1. Fork this repo to your local Github account.
2. Create a new branch to complete all your work in.
3. Test your work using the provided tests
4. Create a Pull Request against the Shopify Main branch when you're done and all tests are passing

# Shopify Intern Assessment Production Engineering

## Description

Write a Go program that solves a given Sudoku puzzle. The program should take a 9x9 grid as input, where empty cells are represented by zeros (0), and output the solved Sudoku grid.

A Sudoku puzzle is a 9x9 grid divided into nine 3x3 sub-grids. The goal is to fill in the empty cells with numbers from 1 to 9, such that each row, each column, and each sub-grid contains all the numbers from 1 to 9 without repetition.

Your program should implement an efficient algorithm to solve the Sudoku puzzle and print the solved grid to the console.

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

## Instructions:
1. Write a function called SolveSudoku that takes a 9x9 grid as input and returns the solved Sudoku grid.
2. Implement an efficient algorithm to solve the Sudoku puzzle. You can use any approach or technique you prefer.
3. Confirm the validity of your code against the tests found in this repo.
4. Ensure that your code is well-documented and easy to understand.

## Constraints:
- The input grid will be a 9x9 two-dimensional array of integers.
- The input grid will have exactly one solution.
- The input grid may contain zeros (0) to represent empty cells.

## Validation: 
To validate the correctness of the solution, you can compare the output of the program with the expected output for a set of test cases containing unsolved Sudoku puzzles.



---------------------------------- Instructions ---------------------------------------------------

To run the program enter "go run sudoku.go" into the terminal

To run the test cases enter "go test" into the terminal