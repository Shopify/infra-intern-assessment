package main

func SolveSudoku(board [][]int) [][]int {
    if !isValidInput(board) {
        return [][]int{} // return an empty board if input is invalid
    }
    if backTracking(board, 0, 0) {
        return board
    }
    return [][]int{} // return an empty board if no solution is found
}

func backTracking(board [][]int, row, col int) bool {
    m, n := 9, 9
    if col == n {
        return backTracking(board, row+1, 0)
    }
    if row == m {
        return true
    }
    if board[row][col] != 0 {
        return backTracking(board, row, col+1)
    }

    for ch := 1; ch <= 9; ch++ {
        if !isValid(board, row, col, ch) {
            continue
        }
        board[row][col] = ch
        if backTracking(board, row, col+1) {
            return true
        }
        board[row][col] = 0
    }
    return false
}

func isValid(board [][]int, row, col, ch int) bool {
    for i := 0; i < 9; i++ {
        if board[row][i] == ch || board[i][col] == ch || board[(row/3)*3+i/3][(col/3)*3+i%3] == ch {
            return false
        }
    }
    return true
}

func isValidInput(board [][]int) bool {
    if len(board) != 9 {
        return false
    }
    for _, row := range board {
        if len(row) != 9 {
            return false
        }
    }
    return true
}
