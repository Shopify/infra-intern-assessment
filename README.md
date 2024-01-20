# Sudoku Solver

This is a Sudoku solver implemented in Go that utilizes backtracking and recursion to find a solution for a given Sudoku puzzle. The solver is designed to handle a 9x9 Sudoku grid.

## Table of Contents

- [Sudoku Solver](#sudoku-solver)
  - [Table of Contents](#table-of-contents)
  - [File Structure](#file-structure)
  - [Algorithm](#algorithm)
  - [Optimization](#optimization)


## File Structure

The codebase is organized into multiple files for clarity and maintainability:

- **sudoku_solver.go:**
  - Contains the main Sudoku solving logic using backtracking and recursion.

- **sudoku_helper.go:**
  - Includes helper functions for stack operations, grid printing, and validity checks.
  - Defines the `EmptyCell` type and the `Stack` data structure.


## Algorithm

The solver uses a backtracking algorithm, a popular technique for solving constraint satisfaction problems like Sudoku. The main idea is to systematically explore possible solutions and backtrack when an invalid placement is encountered.

The algorithm works as follows:

1. **Finding Empty Cells:**
   - The Sudoku grid is scanned to identify empty cells (cells with the value 0).
   - The locations of these empty cells are stored in a stack.

2. **Backtracking with Recursion:**
   - The solver employs a recursive function to attempt different number placements in each empty cell.
   - If a placement is valid, the solver proceeds to the next empty cell.
   - If no valid placement is found, the algorithm backtracks to the previous empty cell and tries a different number.

3. **Base Case:**
   - The algorithm continues until all empty cells are filled, indicating the completion of a valid Sudoku solution.


## Optimization

The algorithm optimizes the process of finding empty cells by utilizing a stack. The stack keeps track of the locations of empty cells beforehand, allowing for efficient retrieval and removal. This optimization helps in reducing the time complexity of the Sudoku solving process.

