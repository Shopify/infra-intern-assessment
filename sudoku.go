package main

// SolveSudoku solves the Sudoku puzzle using backtracking algorithm
func SolveSudoku(grid [][]int) [][]int {
    if solve(grid) {
        return grid
    }
    return nil
}

// solve is the recursive function that implements the backtracking
func solve(grid [][]int) bool {
    emptyRow, emptyCol, found := findEmptySpot(grid)
    if !found {
        return true
    }
    for num := 1; num <= 9; num++ {
        if isValid(grid, emptyRow, emptyCol, num) {
            grid[emptyRow][emptyCol] = num
            if solve(grid) {
                return true
            }
            grid[emptyRow][emptyCol] = 0
        }
    }
    return false
}

// isValid checks if num can be placed at grid[row][col]
func isValid(grid [][]int, row int, col int, num int) bool {
    for i := 0; i < 9; i++ {
        if grid[row][i] == num || grid[i][col] == num {
            return false
        }
    }
    startRow, startCol := row-row%3, col-col%3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if grid[i+startRow][j+startCol] == num {
                return false
            }
        }
    }
    return true
}

// findEmptySpot finds an empty spot in the grid and returns its coordinates
func findEmptySpot(grid [][]int) (int, int, bool) {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if grid[i][j] == 0 {
                return i, j, true
            }
        }
    }
    return 0, 0, false
}