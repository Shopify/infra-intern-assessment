package main

// entry point for tests
func SolveSudoku(sudoku [][]int) [][]int {
    if solve(sudoku) {
        return sudoku
    }
    return nil
}

//  solve solves a 2D 9x9 sudoku grid with constraints:
//      there is exactly 1 solution,
//      empty cells are represented as 0,
//      non-empty cells have a value between 1-9.
//  The result will be a 9x9 grid where:
//      no value appears twice in any row or column,
//      no value appears twice in any of the 9 mutually exclusive 3x3 grids.
//  Returns true if the board is solved, false otherwise.
func solve(sudoku [][]int) bool {
    
    for i := 0; i < 9; i++ {
        
        for j := 0; j < 9; j++ {
            //  don't need to worry about cells with values; only replace
            //  empties
            if sudoku[i][j] != 0 {
                continue;
            }
            for val := 1; val <= 9; val++ {
                if !isPossible(sudoku, i, j, val) {
                    continue;
                }
                //  val is a potential candidate for this cell; try it and 
                //  solve the resulting board. if it didn't work, remove it
                //  and try a different val
                sudoku[i][j] = val;
                if solve(sudoku) {
                    return true
                }
                sudoku[i][j] = 0
            }
            // every val doesn't work for this cell; board is unsolvable
            return false
        }
    }
    // every cell is filled
    return true
}

//  isPossible checks if val can be placed in a cell (specified by its row
//  and col)  without rendering the current board invalid, if so returns true
//  and false otherwise.
func isPossible(sudoku [][]int, row int, col int, val int) bool {
    return uniqueInRow(sudoku, row, val) &&
        uniqueInCol(sudoku, col, val) &&
        uniqueInBox(sudoku, row, col, val) 
}

//  uniqueInRow checks if val can be placed in row without being duplicate,
//  if so returns true, and false otherwise.
func uniqueInRow(sudoku [][]int, row int, val int) bool {
    for j := 0; j < 9; j++ {
        if sudoku[row][j] == val {
            return false
        }
    }
    return true
}

//  uniqueInCol checks if val can be placed in col without being duplicate,
//  if so returns true, and false otherwise.
func uniqueInCol(sudoku [][]int, col int, val int) bool {
    for i := 0; i < 9; i++ {
        if sudoku[i][col] == val {
            return false
        }
    }
    return true
}

//  uniqueInBox checks if val can be placed in the box that cell (row, col) is
//  in without being duplicate, if so returns true, and false otherwise.
func uniqueInBox(sudoku [][]int, row int, col int, val int) bool {
    // coordinates of the top left cell of the box specified by (row, col)
    boxCol := (col / 3)*3;
    boxRow := (row / 3)*3;

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if sudoku[boxRow + i][boxCol + j] == val {
                return false
        
            }
        }
    }
    return true
}
