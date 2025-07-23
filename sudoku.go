// Karthik Nambiar Solution to Sudoku Problem

package main

// solves the sudoku problem using Depth-First Search.
func ActualSolve(board [][]int, row int, col int) bool {
    for i := row; i < 9; i++ {
        // If we start from a new row, the column should start from the beginning
        if i != row {
            col = 0 
        }
        for j := col; j < 9; j++ {
            // if the current position is already filled, skip to the next column
            if board[i][j] != 0 {
                continue 
            }
            for num := 1; num <= 9; num++ {
                // check if the number could be filled in the current position
                if isValid(board, i, j, num) {
                    board[i][j] = num
                    // recursively fill the remaining cells
                    if ActualSolve(board, i, j+1) {
                        return true 
                    }
                    // if the filling is invalid or leads to no solution, undo the current filling
                    board[i][j] = 0
                }
            }
            // if no number can be filled in current position, return false to backtrack
            return false 
        }
    }
    // if every position is filled return true
    return true 
}

// checks if a number could be filled in a certain position.
func isValid(board [][]int, row int, col int, num int) bool {
    // calculate the starting point of the block
    brow := (row/3)*3
	bcol := (col/3)*3 
    for i := 0; i < 9; i++ {
        // check if the number exist in the current row/col/block, if so return false
        if board[i][col] == num || board[row][i] == num || board[brow+i/3][bcol+i%3] == num {
            return false 
        }
    }
    // otherwise, the number can be filled in the position
    return true 
}

// solves Sudoku problem by using backtracking method.
func SolveSudoku(board [][]int) [][]int {
    ActualSolve(board, 0, 0)
	return board
}