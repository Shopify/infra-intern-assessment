package main

import "fmt"

// Requires: The input grid will be a 9x9 two-dimensional array of integers.
//			 The input grid will have exactly one solution.
// Modifies: N/A
// Effects: Solves the sudoku, returns a 9 by 9 array of the solved sudoku
func SolveSudoku(sudoku [][]int) [][]int {
    for i := 0; i < len(sudoku); i++ {
        for j := 0; j < len(sudoku[i]); j++ {
            fmt.Printf("%d ", sudoku[i][j])
        }
        fmt.Println()
    }
	return sudoku
}

func main() {
	input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	SolveSudoku(input)
}
