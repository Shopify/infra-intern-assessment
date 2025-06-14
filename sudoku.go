//Arthur: Cailing Zhang
//Date: 2024/01/28
//Environment: MacOS Ventura 13.5 Apple M3 Chip
//go version go1.21.6 darwin/arm64

package main

import "fmt"

//Define the size of the Sudoku grid
const N = 9

// SolveSudoku tries to solve the Sudoku puzzle and returns the solved grid
func SolveSudoku(grid [][]int) [][]int {
    if solve(grid) {
        return grid // Return the solved grid
    }
    return nil // Return nil if no solution exists
}

// solve is the helper function that uses backtracking to solve the puzzle
func solve(grid [][]int) bool {
    row, col, isEmpty := findEmptyCell(grid)
    if isEmpty {
        return true // Puzzle solved
    }

    for num := 1; num <= N; num++ {
        if isSafe(grid, row, col, num) {
            grid[row][col] = num

            if solve(grid) {
                return true
            }
            grid[row][col] = 0 // Backtrack
        }
    }
    return false
}

// findEmptyCell finds an empty cell in the grid
func findEmptyCell(grid [][]int) (int, int, bool) {
    for row := 0; row < N; row++ {
        for col := 0; col < N; col++ {
            if grid[row][col] == 0 {
                return row, col, false
            }
        }
    }
    return 0, 0, true
}

// isSafe checks if it's safe to place a number in the cell
func isSafe(grid [][]int, row, col, num int) bool {
    return !usedInRow(grid, row, num) && !usedInCol(grid, col, num) && !usedInBox(grid, row-row%3, col-col%3, num)
}

// usedInRow checks if a number is used in the current row
func usedInRow(grid [][]int, row, num int) bool {
    for col := 0; col < N; col++ {
        if grid[row][col] == num {
            return true
        }
    }
    return false
}

// usedInCol checks if a number is used in the current column
func usedInCol(grid [][]int, col, num int) bool {
    for row := 0; row < N; row++ {
        if grid[row][col] == num {
            return true
        }
    }
    return false
}

// usedInBox checks if a number is used in the 3x3 subgrid
func usedInBox(grid [][]int, boxStartRow, boxStartCol, num int) bool {
    for row := 0; row < 3; row++ {
        for col := 0; col < 3; col++ {
            if grid[row+boxStartRow][col+boxStartCol] == num {
                return true
            }
        }
    }
    return false
}

// printGrid prints the Sudoku grid
func printGrid(grid [][]int) {
    for _, row := range grid {
        for _, val := range row {
            fmt.Printf("%d ", val)
        }
        fmt.Println()
    }
}

func main() {
   // Example input
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

    solvedGrid := SolveSudoku(grid)
    if solvedGrid != nil {
        printGrid(solvedGrid)
    } else {
        fmt.Println("No solution exists")
    }
}


