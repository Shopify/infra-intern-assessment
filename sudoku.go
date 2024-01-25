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
func solve(sudoku [][]int) bool {
    
    for i := 0; i < 9; i++ {
        
        for j := 0; j < 9; j++ {
            
            if sudoku[i][j] != 0 {
                continue;
            }
            for val := 1; val <= 9; val++ {
                if !isPossible(sudoku, i, j, val) {
                    continue;
                }
                sudoku[i][j] = val;
                if solve(sudoku) {
                    return true
                }
                sudoku[i][j] = 0
            }
            return false
        }
    }
    return true
}

func isPossible(sudoku [][]int, row int, col int, val int) bool {
    return uniqueInRow(sudoku, row, val) &&
        uniqueInCol(sudoku, col, val) &&
        uniqueInBox(sudoku, row, col, val) 
}

func uniqueInRow(sudoku [][]int, row int, val int) bool {
    for j := 0; j < 9; j++ {
        if sudoku[row][j] == val {
            return false
        }
    }
    return true
}

func uniqueInCol(sudoku [][]int, col int, val int) bool {
    for i := 0; i < 9; i++ {
        if sudoku[i][col] == val {
            return false
        }
    }
    return true
}

func uniqueInBox(sudoku [][]int, row int, col int, val int) bool {
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
