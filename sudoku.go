package main

import "fmt"

// CanPlace checks if a number can be placed at the given row and column in the Sudoku board.
func CanPlace(board [][]int, row, col, num int) bool {
    for x := 0; x < 9; x++ {
        // Check if the number is already in the row or column.
        if board[row][x] == num || board[x][col] == num {
            return false
        }
        // Check if the number is in the 3x3 subgrid.
        if board[3*(row/3)+x/3][3*(col/3)+x%3] == num {
            return false
        }
    }
    return true
}

// solve is a recursive helper function that attempts to solve the Sudoku puzzle.
func solve(board [][]int) bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            // Look for an empty cell (denoted by 0).
            if board[i][j] == 0 {
                // Try placing numbers 1 through 9 in the empty cell.
                for num := 1; num <= 9; num++ {
                    if CanPlace(board, i, j, num) {
                        board[i][j] = num
                        // Recursively attempt to solve the rest of the board.
                        if solve(board) {
                            return true
                        }
                        // Backtrack if placing num doesn't lead to a solution.
                        board[i][j] = 0
                    }
                }
                // Return false if no number can be placed in the empty cell.
                return false
            }
        }
    }
    // Return true if all cells are filled, meaning the puzzle is solved.
    return true
}

// SolveSudoku attempts to solve the given Sudoku puzzle and returns the solved board.
func SolveSudoku(board [][]int) [][]int {
    if solve(board) {
        // If the puzzle is solvable, return the solved board.
        return board
    }
    // Return nil if no solution exists for the puzzle.
    return nil
}

// printBoard prints the Sudoku board to the console.
func printBoard(board [][]int) {
    for _, row := range board {
        for _, num := range row {
            fmt.Printf("%d ", num)
        }
        fmt.Println()
    }
}

// 
//func main() {
    // Define the initial Sudoku board with some cells filled in.
//  board := [][]int{
//        {5, 3, 0, 0, 7, 0, 0, 0, 0},
//        {6, 0, 0, 1, 9, 5, 0, 0, 0},
//        {0, 9, 8, 0, 0, 0, 0, 6, 0},
//        {8, 0, 0, 0, 6, 0, 0, 0, 3},
//        {4, 0, 0, 8, 0, 3, 0, 0, 1},
//        {7, 0, 0, 0, 2, 0, 0, 0, 6},
//        {0, 6, 0, 0, 0, 0, 2, 8, 0},
//        {0, 0, 0, 4, 1, 9, 0, 0, 5},
//        {0, 0, 0, 0, 8, 0, 0, 7, 9},
//    }

    // Attempt to solve the Sudoku puzzle.
//    solved := SolveSudoku(board)
//    if solved != nil {
        // If a solution exists, print the solved board.
//        printBoard(solved)
//    } else {
        // If no solution exists, inform the user.
//        fmt.Println("No solution exists")
//    }
//}
