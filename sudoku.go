package main

import (
    "fmt"
)

const N = 9 // 9x9 sudoku puzzle

// Function to print the grid to visualize it
func printGrid(grid [][]int) {
    fmt.Println("\nSolved Sudoko Puzzle:")
	fmt.Println("+-------+-------+-------+")
    for i := 0; i < N; i++ {
		fmt.Print("| ")
        for j := 0; j < N; j++ {
			if j == 3 || j == 6 {
				fmt.Print("| ")
			}
            fmt.Printf("%d ", grid[i][j])
			if j == 8 {
				fmt.Print("|")
			}
        }

		if i == 2 || i == 5 || i == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}

    }
    fmt.Println()
}

// Function to check if it's safe to place a number in the given row, column, and 3x3 box
func isSafe(grid [][]int, row, col, num int) bool {
    // Check row and column
    for x := 0; x < N; x++ {
        if grid[row][x] == num || grid[x][col] == num {
            return false
        }
    }

    // Check 3x3 box
    startRow := row - row%3
    startCol := col - col%3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if grid[i+startRow][j+startCol] == num {
                return false
            }
        }
    }
    return true
}


// Function that returns the solved grid, or the original grid if no solution is found
func SolveSudoku(grid [][]int) [][]int {
    if solve(&grid) {
        printGrid(grid)
        return grid
    }
    return grid // Return the original grid if no solution is found
}

// SolveSudoku Helper Function to solve the Sudoku puzzle using backtracking algorithm
func solve(grid *[][]int) bool {
    row := -1
    col := -1
    isEmpty := true

    for i := 0; i < N; i++ {
        for j := 0; j < N; j++ {
            if (*grid)[i][j] == 0 {
                row = i
                col = j
                isEmpty = false
                break
            }
        }
        if !isEmpty {
            break
        }
    }

    // No empty space left
    if isEmpty {
        return true
    }

    for num := 1; num <= N; num++ {
        if isSafe(*grid, row, col, num) {
            (*grid)[row][col] = num
            if solve(grid) {
                return true
            }
            (*grid)[row][col] = 0 // Undo the current cell for backtracking
        }
    }
    return false
}
