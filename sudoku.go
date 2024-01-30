package main

import "fmt"

// main is the entry point of the application.
// It demonstrates an example of how to use the SolveSudoku function.
func main() {
    // Example Sudoku puzzle
    input := [][]int{
        {5, 3, 0, 0, 7, 0, 0, 0, 0},
        {0, 9, 8, 0, 0, 0, 0, 6, 0},
        {8, 0, 0, 0, 6, 0, 0, 0, 3},
        {4, 0, 0, 8, 0, 3, 0, 0, 1},
        {7, 0, 0, 0, 2, 0, 0, 0, 6},
        {0, 6, 0, 0, 0, 0, 2, 8, 0},
        {0, 0, 0, 4, 1, 9, 0, 0, 5},
        {0, 0, 0, 0, 8, 0, 0, 7, 9},
    }

    // Solving the Sudoku puzzle
    SolveSudoku(input)

    // Printing the solved puzzle
    for _, row := range input {
        fmt.Println(row)
    }
}

// SolveSudoku solves the given Sudoku puzzle using a backtracking algorithm.
// It modifies the puzzle in place and returns true if the puzzle is solvable, else false.
func SolveSudoku(board [][]int) bool {
    if len(board) == 0 {
        return false
    }
    return solve(board)
}

// solve is a recursive helper function that applies the backtracking algorithm to solve the Sudoku.
func solve(board [][]int) bool {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            // Check if the current cell is empty
            if board[i][j] == 0 {
                // Try filling the cell with numbers 1 to 9
                for c := 1; c <= 9; c++ {
                    // Check if the number is valid in the current cell
                    if isValid(board, i, j, c) {
                        board[i][j] = c

                        // Continue with this setup. If it doesn't lead to a solution, backtrack.
                        if solve(board) {
                            return true
                        } else {
                            // Backtrack
                            board[i][j] = 0
                        }
                    }
                }
                // No valid number found, need to backtrack
                return false
            }
        }
    }
    // Solution found
    return true
}

// isValid checks if a given number can be placed in a specified cell without violating Sudoku rules.
func isValid(board [][]int, row, col, c int) bool {
    for i := 0; i < 9; i++ {
        // Check the row, column, and the 3x3 subgrid for violation
        if board[i][col] == c || board[row][i] == c || board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
            return false
        }
    }
    return true
}
