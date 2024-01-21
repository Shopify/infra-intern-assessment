# Shopify Intern Assessment Production Engineering

## Description

This is a Go program that solves a given Sudoku puzzle. The program takes a 9x9 grid as input, where empty cells may be represented by zeros (0), and output the solved Sudoku grid (see [Constraints](#constraints) section for more info).

A Sudoku puzzle is a 9x9 grid divided into nine 3x3 sub-grids. The goal is to fill in the empty cells with numbers from 1 to 9, such that each row, each column, and each sub-grid contains all the numbers from 1 to 9 without repetition.

## My Implementation
My implementation (see `sudoku.go`) uses a typical serial backtracking algorithm combined with memoization to efficiently solve a given 9 by 9 sudoku.

The motivation behind using a backtracking algorithm for this problem is because it is easy to implement and highly efficient in solving sudokus, perfectly matching what the Go language is known for: its simplicity and efficiency. Though Go is also known for its concurrency, from my experience, implementing a parallel algorithm to solve a 9 by 9 sudoku is not only complicated due to the inherent data dependencies, it also may not result in any runtime improvement at all since the synchronization and communication overheads introduced in a parallel program may outweigh the performance gain when the input size is small.

I also took adavantage of the given constraints in the problem statement that the input sudoku is guaranteed to be 9 by 9 and have exactly 1 solution, by adding memoization to improve the runtime of the constraint-checking function, which determines whether placing a number at an unfilled cell would still result in a valid sudoku. By doing so, the constraint-checking function only needs 1 boolean statement, as compared to a naive for loop. Despite both implementation being theoretically O(1) given a fixed-size input, practically, the memoization approach did result in some minor speedup (< 1 ms difference when tested with a difficult sudoku), therefore it is kept in the final implementation.

For more low-level details on how the backtracking & memoization worked (e.g. specific functions definitions, variables, data structures, etc.), see `sudoku.go`.

### Test Cases
More test cases are added in `sudoku_test.go`, which tests some edge cases such as the given sudoku is already completed, everything is filled in the sudoku except 1 cell, etc.

## Running the Program
To run the tests, open project root directory (e.g. `infra-intern-assessment`) in terminal, run `go test`.

Or, you can do the following to solve a custome sudoku:
1. Add a main function in `sudoku.go`
2. Make a sudoku of your choice in the form of 2d array
3. Call `SolveSudoku` function with the sudoku you just made
```Go
// Example main function with a user-defined sudoku
func main() {
  mySudoku := [][]int{
    {5, 3, 0, 0, 7, 0, 0, 0, 0},
    {6, 0, 0, 1, 9, 5, 0, 0, 0},
    {0, 9, 8, 0, 0, 0, 0, 6, 0},
    {8, 0, 0, 0, 6, 0, 0, 0, 3},
    {4, 0, 0, 8, 0, 3, 0, 0, 1},
    {7, 0, 0, 0, 2, 0, 0, 0, 6},
    {0, 6, 0, 0, 0, 0, 2, 8, 0},
    {0, 0, 0, 4, 1, 9, 0, 0, 5},
    {0, 0, 0, 0, 8, 0, 0, 7, 9},
  }
  // call function that solves the sudoku
  SolveSudoku(mySudoku)
  // print result
  PrintSudoku(mySudoku)
}
```

## Constraints:
- The input grid will be a 9x9 two-dimensional array of integers.
- The input grid will have exactly one solution.
- The input grid may contain zeros (0) to represent empty cells.

