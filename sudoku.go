package main



//Initial brute force solution




// SolveSudoku is the main function that is called to solve the Sudoku puzzle.
// It calls the helper function 'solve' and returns the solved board.
func SolveSudoku(board [][]int) [][]int {
    solve(board)
    return board
}

// solve is a helper function that uses backtracking to solve the Sudoku puzzle.
// It iterates over each cell in the board. If it finds an empty cell (0), it tries to fill it with a number from 1 to 9.
// If the number is valid (doesn't violate Sudoku rules), it recursively calls solve to fill in the rest of the board.
// If it can't find a valid number, it backtracks by setting the cell back to 0 and returning false.
func solve(board [][]int) bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == 0 {
                for num := 1; num <= 9; num++ {
                    if isValid(board, i, j, num) {
                        board[i][j] = num
                        if solve(board) {
                            return true
                        }
                        board[i][j] = 0
                    }
                }
                return false
            }
        }
    }
    return true
}

// isValid is a helper function that checks if a number is valid in a given cell by checking the row, column, and 3x3 box that the cell belongs to.
// It returns false if the number already exists in the row, column, or 3x3 box, and true otherwise.
func isValid(board [][]int, row, col, num int) bool {
    for i := 0; i < 9; i++ {
        if board[i][col] == num {
            return false
        }
        if board[row][i] == num {
            return false
        }
        if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
            return false
        }
    }
    return true
}




