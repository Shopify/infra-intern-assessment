package main

// Return false if the number is already in that column, row, or box. True otherwise
func IsSafe(board [][]int, row, col, num int) bool {
    
    // check row and column
    for x := 0; x < 9; x++ {
        if board[row][x] == num || board[x][col] == num {
            return false
        }
    }

    // check box
    startRow := row - row%3
    startCol := col - col%3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i + startRow][j + startCol] == num {
                return false
            }
        }
    }

    return true
}

func SolveSudoku(board [][]int) [][]int {
    
    row := -1
    col := -1
    isEmpty := true

    // Find the first empty cell
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == 0 {
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

    // If there are no empty cells, the sudoku is solved
    if isEmpty {
        return board
    }

    for num := 1; num <= 9; num++ {
        if IsSafe(board, row, col, num) {
            board[row][col] = num
            // Recursively try to solve the rest of the board
            if solvedBoard := SolveSudoku(board); solvedBoard != nil {
                return solvedBoard
            }
            board[row][col] = 0 // Backtrack and remove the number
        }
    }
    return nil
}
