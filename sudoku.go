// Sudoku.go takes a 9x9 grid as input, where empty cells are represented by zeros (0), and output the solved Sudoku grid
package main

import "fmt"

// isSafe function checks if it is safe to place a number in a given cell:
// It does so by validating if the same number is not in the current row, column, and the 3x3 subgrid
func isSafe(board [][]int, row, column, num int) bool {
   
    // Check if 'num' is not in the given column, return false if it is
    for i := 0; i < 9; i++ {
		if board[i][column] == num {return false}
	}

	// Check if 'num' is not in the given row, return false if it is
    for i := 0; i < 9; i++ {
		if board[row][i] == num {return false}
    }

    // Check if 'num' is not in the 3x3 subgrid, return false if it is
    startofRow := row - row%3
    startofColumn := column - column%3
    
	for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i+startofRow][j+startofColumn] == num {return false}
        }
    }

    // If the number is not present in row, column and 3x3 subgrid, return true, as it is "safe"
    return true
}

// solveSudoku is used as a recursive function which utilizes backtracking to solve the actual Sudoku puzzle

func solveSudoku(board [][]int) bool {
    row, column := -1, -1
    sudokuIsEmpty := true

    // Find an unassigned location in the Sudoku board (represented by a 0)
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == 0 {
                row, column = i, j // Assign row number and column number in the board
                sudokuIsEmpty = false // Mark that there is a space left to fill so the break does not occur (and exits loop)
                break
            }
        }
        if !sudokuIsEmpty {
            break
        }
    }

    // No unassigned location is left, puzzle solved.
    if sudokuIsEmpty {
        return true
    }

    // Try digits 1 to 9 in the current cell
    for num := 1; num <= 9; num++ {
        // Check safety before assigning the number
        if isSafe(board, row, column, num) {
            board[row][column] = num

            // Recursively try filling in rest of the board
            if solveSudoku(board) {
                return true
            }

            // If filling the current cell with 'num' leads to a solution,
            // return true, else undo the number (backtrack)
            board[row][column] = 0
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
    }

    // Call SolveSudoku and print the result
    solution := SolveSudoku(input)
    fmt.Println("Solved Sudoku:")
    for _, row := range solution {
        fmt.Println(row)
    }
}
