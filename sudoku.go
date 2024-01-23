package main

import (
	"fmt"
)

// constant for the length of the Sudoku, which is a 9x9 grid
const GRID_LEN = 9

// printGrid takes a 9x9 integer grid and prints it to standard output.
func printGrid(grid [][]int) {
	fmt.Println("[")
    for r := 0; r < 8; r++ {
        fmt.Print("  [")
        for c := 0; c < 8; c++ {
            fmt.Printf("%d, ", grid[r][c])
        }
        fmt.Printf("%d],\n", grid[r][8])
    }

    fmt.Print("  [")
    for c := 0; c < 8; c++ {
        fmt.Printf("%d, ", grid[8][c])
    }
    fmt.Printf("%d]\n", grid[8][8])
    fmt.Println("]")
}

// SolveSudoku takes a 9x9 integer grid as a Sudoku, solves it, then prints it to standard output and returns the solved Sudoku.
func SolveSudoku(grid [][]int) [][]int {
	// fill the Sudoku starting from (r, c) = (0, 0)
	solved := fillPosition(grid, 0, 0)
	if solved != true {
		fmt.Println("Sudoku could not be solved!")
	}
	
	// print the solved Sudoku grid
    printGrid(grid)
	return grid
}

// fillPosition fills the position (row, col) in grid with a valid value from 1-9, according to the rules of Sudoku.
func fillPosition(grid [][]int, row, col int) bool {
	// look for the next empty position, going left from row and down from col
	solved, row, col := getNextEmptyPosition(grid, row, col)

	if solved {
		return solved
	}

	// iterate over values 1-9 and recurse if we find a value that is valid
	// backtrack if we did not solve the entire Sudoku grid
	for value := 1; value <= 9; value++ {
		if doesGridContainValue(grid, value, row, col) {
			// update grid with valid value
			grid[row][col] = value
			solved := fillPosition(grid, row, col + 1)
			if solved {
				return solved
			}

			// reset value as grid is invalid and a different value should be in this position when we backtrack
			grid[row][col] = 0
		}
	}

	return false
}

func getNextEmptyPosition(grid [][]int, row, col int) (bool, int, int) {
	// iterate through the grid to find the next empty spot, which is represented by a zero
	for row < GRID_LEN{
		for col < GRID_LEN {
			if grid[row][col] == 0 {
				return false, row, col
			}
			col++
		}
		// start from zero when we iterate over a new row
		col = 0
		row++
	}

	// if we have not found an empty value, we have solved the Sudoku grid
	return true, row, col
}

func doesGridContainValue(grid [][]int, value, row, col int) bool {
	// check if value is already in the row
	for c := 0; c < GRID_LEN; c++ {
		if grid[row][c] == value {
			return false
		}
	}

	// check if value is already in the column
	for r := 0; r < GRID_LEN; r++ {
		if grid[r][col] == value {
			return false
		}
	}

	// check if value is already in the 3x3 box
	rowBegin, rowEnd := (row / 3) * 3, (row / 3 + 1) * 3
	colBegin, colEnd := (col / 3) * 3, (col / 3 + 1) * 3
	for r := rowBegin; r < rowEnd; r++ {
		for c := colBegin; c < colEnd; c++ {
			if grid[r][c] == value {
				return false
			}
		}
	}

	// value is unique in the row, column, and box, so we can add it to the grid
	return true
}
