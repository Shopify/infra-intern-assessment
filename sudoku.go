package main

import "fmt"

// isSafe function checks if it is safe to place a number in a given cell:
// It does so by validating if the same number is not present in the current row, column, and the 3x3 subgrid.
func isSafe(board [][]int, row, col, num int) bool {
    // Check if 'num' is not present in the given row.
    for x := 0; x < 9; x++ {
        if board[row][x] == num {
            return false
        }
    }

    // Check if 'num' is not present in the given column.
    for x := 0; x < 9; x++ {
        if board[x][col] == num {
            return false
        }
    }

    // Check if 'num' is not present in the 3x3 subgrid.
    startRow := row - row%3
    startCol := col - col%3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i+startRow][j+startCol] == num {
                return false
            }
        }
    }

    // If the number is not present in row, column and 3x3 subgrid, return true.
    return true
}

// solveSudoku is the recursive function that uses backtracking to solve the Sudoku puzzle.
func solveSudoku(board [][]int) bool {
    row, col := -1, -1
    isEmpty := true

    // Find an unassigned location in the Sudoku board (represented by 0)
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == 0 {
                row, col = i, j // Assign row number and column number
                isEmpty = false // Mark that there is a space left to fill
                break
            }
        }
        if !isEmpty {
            break
        }
    }

    // No unassigned location is left, puzzle solved.
    if isEmpty {
        return true
    }

    // Try digits 1 to 9 in the current cell
    for num := 1; num <= 9; num++ {
        // Check safety before assigning the number
        if isSafe(board, row, col, num) {
            board[row][col] = num

            // Recursively try filling in rest of the board
            if solveSudoku(board) {
                return true
            }

            // If filling the current cell with 'num' leads to a solution,
            // return true, else undo the number (backtrack)
            board[row][col] = 0
        }
    }

    // Trigger backtracking
    return false
}

// SolveSudoku solves the given Sudoku puzzle using the solveSudoku function.
// It returns the solved puzzle if solvable, else returns an empty grid.
func SolveSudoku(board [][]int) [][]int {
    if solveSudoku(board) {
        return board
    }
    return [][]int{} // Return an empty grid if no solution exists
}

func main() {
    input := [][]int{
        // Input Sudoku grid with 0s representing empty cells
        {5, 3, 0, 0, 7, 0, 0, 0, 0},
        {6, 0, 0, 1, 9, 5, 0, 0, 0},
        // ... rest of the board
    }

    // Call SolveSudoku and print the result
    solution := SolveSudoku(input)
    fmt.Println("Solved Sudoku:")
    for _, row := range solution {
        fmt.Println(row)
    }
}
