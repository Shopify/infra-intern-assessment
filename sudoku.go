// Package main implements a Sudoku solver.
// This package provides functionality to solve a 9x9 Sudoku puzzle using backtracking.
// To solve for higher grid sizes, change the GridSize constant.

package main

const GridSize = 9 // Sudoku grid size

// isValid checks if placing 'num' in grid[row][col] is valid.
func isValid(grid [][]int, row, col, num int) bool {
    subgridRowStart, subgridColStart := row-row%3, col-col%3

    for i := 0; i < GridSize; i++ {
        // Check row, column, and 3x3 subgrid
        if grid[row][i] == num || grid[i][col] == num ||
           grid[subgridRowStart+i/3][subgridColStart+i%3] == num {
            return false
        }
    }
    return true
}

// Solve the Sudoku puzzle using backtracking.
func solveSudokuHelper(grid [][]int, row, col int) bool {
    if row == GridSize {
        return true // Puzzle solved
    }

    nextRow, nextCol := row, col+1
    if nextCol == GridSize {
        nextRow, nextCol = row+1, 0
    }

    if grid[row][col] != 0 {
        return solveSudokuHelper(grid, nextRow, nextCol)
    }

    for num := 1; num <= GridSize; num++ {
        if isValid(grid, row, col, num) {
            grid[row][col] = num
            if solveSudokuHelper(grid, nextRow, nextCol) {
                return true
            }
            grid[row][col] = 0 // Backtrack
        }
    }
    return false
}

// SolveSudoku solves the Sudoku puzzle.
// Entry point for backtracking.
func SolveSudoku(grid [][]int) [][]int {
    if solveSudokuHelper(grid, 0, 0) {
        return grid
    }
    return nil
}
