package main

import (
    "fmt"
)
'''
In this implementation:

Cross-Hatching Elements: The isSafe function embodies a principle similar to cross-hatching. It methodically checks each row, column, 
and 3x3 subgrid to see if a number can be safely placed, similar to how cross-hatching works by systematically eliminating possibilities.

Backtracking Algorithm: The solveSudoku function is a classic demonstration of backtracking. It tries placing numbers in empty cells and 
backtracks when a placement leads to a dead end, thus exploring all possible configurations until a solution is found.

This combination of checking feasibility (akin to cross-hatching) and recursive exploration with backtracking makes the algorithm both efficient 
and robust for solving Sudoku puzzles.

'''
// isSafe checks if it's safe to place a number in a specific cell.
// This function embodies a principle similar to cross-hatching, where it checks
// rows, columns, and subgrids to determine the feasibility of placing a number.
func isSafe(board [][]int, row, col, num int) bool {
    // Check row and column: Ensures the number is not repeated in the same row or column.
    for x := 0; x < 9; x++ {
        if board[row][x] == num || board[x][col] == num {
            return false
        }
    }

    // Check 3x3 subgrid: Ensures the number is not repeated in the corresponding subgrid.
    startRow := row - row%3
    startCol := col - col%3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i+startRow][j+startCol] == num {
                return false
            }
        }
    }

    return true
}

// solveSudoku is the heart of the backtracking algorithm. It works recursively,
// trying numbers in empty cells and backtracking when a number leads to no solution.
func solveSudoku(board [][]int) bool {
    row, col := -1, -1
    isEmpty := true

    // Find the first empty cell (marked with 0). This is a linear scan, akin to
    // scanning rows and columns in cross-hatching but for a different purpose.
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == 0 {
                row, col = i, j
                isEmpty = false
                break
            }
        }
        if !isEmpty {
            break
        }
    }

    // If no empty cell is found, the puzzle is solved.
    if isEmpty {
        return true
    }

    // Try placing numbers 1-9 in the found empty cell
    // This is where backtracking plays a crucial role.
    for num := 1; num <= 9; num++ {
        if isSafe(board, row, col, num) {
            board[row][col] = num

            // Recursively try to solve the rest of the board.
            if solveSudoku(board) {
                return true
            }

            // If the choice of number doesn't lead to a solution,
            // backtrack by resetting the cell and trying the next number.
            board[row][col] = 0
        }
    }
    // Trigger backtracking
    return false
}

// SolveSudoku is the public interface of the solver.
// It takes the Sudoku board as input and returns the solved board.
func SolveSudoku(board [][]int) [][]int {
    if solveSudoku(board) {
        return board
    }
    return [][]int{} // Return an empty board if no solution is found.
}

func printBoard(board [][]int) {
    for _, row := range board {
        for _, num := range row {
            fmt.Printf("%d ", num)
        }
        fmt.Println()
    }
}

func main() {
	puzzle := [][]int{
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

	solved := SolveSudoku(puzzle)
	printBoard(solved)
}
