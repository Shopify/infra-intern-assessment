package main

import "fmt"

// SolveSudoku takes a 9x9 Sudoku grid with zeros representing empty cells and solves the puzzle.
func SolveSudoku(grid [][]int) bool {
    if isSolved(grid) {
        return true
    }

    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            if grid[row][col] == 0 {
                for num := 1; num <= 9; num++ {
                    if isValid(grid, row, col, num) {
                        grid[row][col] = num
                        if SolveSudoku(grid) {
                            return true
                        }
                        grid[row][col] = 0 // backtrack
                    }
                }
                return false // trigger backtracking
            }
        }
    }
    return true // puzzle is solved
}

// isSolved checks if the Sudoku puzzle is already solved.
func isSolved(grid [][]int) bool {
    for row := range grid {
        for col := range grid[row] {
            if grid[row][col] == 0 {
                return false
            }
        }
    }
    return true
}

// isValid checks if placing num in grid[row][col] is valid per Sudoku rules.
func isValid(grid [][]int, row, col, num int) bool {
    for i := 0; i < 9; i++ {
        // Check row and column
        if grid[row][i] == num || grid[i][col] == num {
            return false
        }

        // Check 3x3 subgrid
        subRow := 3 * (row / 3)
        subCol := 3 * (col / 3)
        if grid[subRow+i/3][subCol+i%3] == num {
            return false
        }
    }
    return true
}

func main() {
    grid := [][]int{
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

    if SolveSudoku(grid) {
        fmt.Println("Sudoku Solved:")
        for _, row := range grid {
            fmt.Println(row)
        }
    } else {
        fmt.Println("No solution exists")
    }
}
