// sudoku.go
package main

import "fmt"

const gridSize = 9

// isValid checks whether placing num at grid[row][col] is a valid move
func isValid(grid [][]int, row, col, num int) bool {
    // Check if num exists in the same row or column
    for i := 0; i < gridSize; i++ {
        if grid[row][i] == num || grid[i][col] == num {
            return false
        }
    }

    // Check if num exists in the 3x3 subgrid
    startRow, startCol := 3*(row/3), 3*(col/3)
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if grid[startRow+i][startCol+j] == num {
                return false
            }
        }
    }

    return true
}

// SolveSudoku uses backtracking to solve the Sudoku puzzle
func SolveSudoku(grid [][]int) [][]int {
    if solveSudoku(grid) {
        return grid
    }
    return nil
}


// solveSudoku recursively tries to solve the Sudoku puzzle
func solveSudoku(grid [][]int) bool {
    // Loop through each cell in the Sudoku grid
    for row := 0; row < gridSize; row++ {
        for col := 0; col < gridSize; col++ {
            // Check if the current cell is empty (contains 0)
            if grid[row][col] == 0 {
                // Try placing numbers 1 to 9 in the current cell
                for num := 1; num <= 9; num++ {
                    // Check if placing num in the current cell is valid
                    if isValid(grid, row, col, num) {
                        // If valid, set the cell to num and recursively try to solve the puzzle
                        grid[row][col] = num
                        if solveSudoku(grid) {
                            return true // Return true if a solution is found
                        }
                        // If the recursive call did not find a solution, backtrack by setting the cell back to 0
                        grid[row][col] = 0
                    }
                }
                // If no valid number is found for the current cell, backtrack
                return false
            }
        }
    }
    // If all cells are filled, print the solution
    printGrid(grid)
    return true
}

// printGrid prints the Sudoku grid from the given input grid
func printGrid(grid [][]int) {
    for _, row := range grid {
        fmt.Println(row)
    }
}


func main() {
	// Example input
	input := [][]int{
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

	// Solve the Sudoku puzzle
	solved := SolveSudoku(input)

	if solved != nil {
		fmt.Println("Solved Sudoku:")
		printGrid(solved)
	} else {
		fmt.Println("No solution exists.")
	}
}
