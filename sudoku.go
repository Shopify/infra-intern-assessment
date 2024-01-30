package main

import ("fmt")

const N = 9

// SolveSudoku solves the Sudoku puzzle using backtracking algorithm
// Case 1: no solution exists*
func SolveSudoku(grid [N][N]int) [N][N]int {
    if solver(&grid, 0, 0) {
        return grid
    }
    return [N][N]int{} 
}

// solver is a recursive helper function for SolveSudoku
func solver(grid *[N][N]int, row, col int) bool {
    // If we have reached the last cell, return true
    if row == N-1 && col == N {
        return true
    }

    // Move to the next row
    if col == N {
        row++
        col = 0
    }

    // Skip filled cells
    if grid[row][col] != 0 {
        return solver(grid, row, col+1)
    }

    // Check if it's safe to place the number
    for num := 1; num <= N; num++ {
        if isSafe(grid, row, col, num) {
            grid[row][col] = num

            if solver(grid, row, col+1) {
                return true
            }
            // Backtrack
            grid[row][col] = 0
        }
    }
    return false
}

// isSafe checks if it's safe to place a number in a cell by checking the row, column, and 3x3 subgrid
func isSafe(grid *[N][N]int, row, col, num int) bool {
    for x := 0; x < N; x++ {
        if grid[row][x] == num || grid[x][col] == num || grid[3*(row/3)+x/3][3*(col/3)+x%3] == num {
            return false
        }
    }
    return true
}

func main() {
    puzzle := [N][N]int{
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

    solution := SolveSudoku(puzzle)
    fmt.Printf("Solved Sudoku:\n%v\n", solution)
}