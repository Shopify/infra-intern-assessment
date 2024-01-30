package main

// SolveSudoku(grid) : solves the sudoku puzzle using a backtracking algorithm
// returns the solved puzzle if it exists, nil otherwise
// asserts: len(grid) == 9, len(grid[i]) == 9 for 0 <= i < 9
func SolveSudoku(grid [][]int) [][]int {
    if solve(grid) {
        return grid
    }
    return nil
}

// solve(grid) recursively backtraces to solve puzzle
// returns true if puzzle is solved, false otherwise
// asserts: len(grid) == 9, len(grid[i]) == 9 for 0 <= i < 9
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

// isValid(grid, row, col, num) checks if num can be placed at the square grid[row][col]
// returns true if num can be placed, false otherwise
// asserts: len(grid) == 9, len(grid[i]) == 9 for 0 <= i < 9
func isValid(grid [][]int, row int, col int, num int) bool {
    blockRow, blockCol := row-row%3, col-col%3
    for i := 0; i < 9; i++ {
        if grid[row][i] == num || grid[i][col] == num {
            return false
        }
        if grid[blockRow+i/3][blockCol+i%3] == num {
            return false
        }
    }
    return true
}

// findEmptySpot(grid) finds an empty spot in the grid and returns its coordinates
// returns true if an empty spot is found, false otherwise
// asserts: len(grid) == 9, len(grid[i]) == 9 for 0 <= i < 9
func findEmptySpot(grid [][]int) (int, int, bool) {
    minOptions := 10
    minRow := -1
    minCol := -1
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if grid[i][j] == 0 {
                numOptions := countOptions(grid, i, j)
                if numOptions < minOptions {
                    minOptions = numOptions
                    minRow = i
                    minCol = j
                }
            }
        }
    }
    return minRow, minCol, minRow != -1
}

// countOptions(grid, row, col) counts the number of options for a square in the grid
// returns the number of options
// asserts: len(grid) == 9, len(grid[i]) == 9 for 0 <= i < 9
func countOptions(grid [][]int, row int, col int) int {
    options := 0
    for num := 1; num <= 9; num++ {
        if isValid(grid, row, col, num) {
            options++
        }
    }
    return options
}