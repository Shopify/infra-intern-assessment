# Sudoko Solver 🧩

Given the sudoku puzzle's presentation as a 9x9 grid with empty cells denoted by zeros, this program fills the grid in line with the following constraints: 

- The input grid will be a 9x9 two-dimensional array of integers.
- The input grid will have exactly one solution.
- The input grid may contain zeros (0) to represent empty cells.

# Implementation

This Sudoku Solver uses backtracking to solve the puzzle and print the solved grid to the console. 

This is accomplished through the following functions: 

## `isInCol(grid [][]int, col, num int) bool`

- Checks if a given number `num` is present in the current column `col` of the Sudoku grid.
- Returns `true` if `num` is present in the column, `false` otherwise.

## `isInRow(grid [][]int, row, num int) bool`

- Checks if a given number `num` is present in the current row `row` of the Sudoku grid.
- Returns `true` if `num` is present in the row, `false` otherwise.

## `isInSubgrid(grid [][]int, subgridStartRow, subgridStartCol, num int) bool`

- Checks if a given number `num` is present in the 3x3 subgrid starting at position `(subgridStartRow, subgridStartCol)` of the Sudoku grid.
- Returns `true` if `num` is present in the subgrid, `false` otherwise.

## `isValid(grid [][]int, row int, col int, num int) bool`

- Checks if placing a given number `num` in the specified cell `(row, col)` of the Sudoku grid is valid according to Sudoku rules.
- Returns `true` if placing `num` in the cell is valid, `false` otherwise.

## `solveSudokuHelper(grid [][]int, row int, col int) bool`

- A recursive helper function that uses backtracking to solve the Sudoku puzzle.
- Returns `true` if the puzzle is solved, `false` otherwise.

## `SolveSudoku(grid [][]int) [][]int`

- The main function to solve the Sudoku puzzle. It initializes the solving process using the `solveSudokuHelper` function.
- Returns the solved Sudoku grid, or `nil` if no solution is found.

