package sudoku

const BOARDSIZE = 9

func SolveSudoku(board [][]int) [][]int {
    solve(board) // done in helper function to ease recursion
    return board
}

func solve(board [][]int) bool {
    for i := 0; i < BOARDSIZE; i++ {
        for j := 0; j < BOARDSIZE; j++ {

            if board[i][j] == 0 { // iterate through empty cells
                for num := 1; num <= BOARDSIZE; num++ { // check if any possible number is valid

                    if validCheck(board, i, j, num) {
                        board[i][j] = num
                        if solve(board) { // recursive call for checking board assuming valid number was part of final solution
                            return true
                        }
                        board[i][j] = 0 // empty if the valid cell was not part of final solution and keep trying
                    }
                }
                return false // return false if no valid solution found
            }
        }
    }
    return true
}

func validCheck(board [][]int, row, col, num int) bool {
    // check if number not in current row and column
    for i := 0; i < BOARDSIZE; i++ {

        if board[row][i] == num || board[i][col] == num {
            return false
        }
    }

    // check if number not in the current box
    startR, startC := 3 * (row/3), 3 * (col/3) // calculates top-left cell of current box
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {

            if board[startR + i][startC + j] == num {
                return false
            }
        }
    }

    return true
}