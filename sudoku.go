package main

import "fmt"

const dim = 9
const emptyCell = 0

//  solves the Sudoku puzzle using backtracking
func solve(grid [][]int) bool {
	row, col, found := findEmptyCell(grid)

	// check if successful
	if !found {
		return true
	}

	// Filling empty cells
	for num := 1; num <= dim; num++ {
		if checkIfSafe(grid, row, col, num) {
			grid[row][col] = num
			// recursive step to solve puzzle
			if solve(grid) {
				return true
			}

			// If solution is invalid, backtrack
			grid[row][col] = emptyCell
		}
	}
	return false
}

// finds the position of empty cells
func findEmptyCell(grid [][]int) (int, int, bool) {
	for row := 0; row < dim; row++ {
		for col := 0; col < dim; col++ {
			if grid[row][col] == emptyCell {
				return row, col, true
			}
		}
	}
	return -1, -1, false
}

// checks if specified value is present in the given row
func checkRow(grid [][]int, row int, val int) bool {
	for columnIndex := 0; columnIndex < dim; columnIndex++ {
		if grid[row][columnIndex] == val {
			return true
		}
	}
	return false
}

// checks if specified value is present in the given column
func checkCol(grid [][]int, col int, val int) bool {
	for rowIndex := 0; rowIndex < dim; rowIndex++ {
		if grid[rowIndex][col] == val {
			return true
		}
	}
	return false
}

// checks if specified value is present in the given section
func checkSection(grid [][]int, row int, col int, val int) bool {
	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		for colIndex := 0; colIndex < 3; colIndex++ {
			if grid[rowIndex+row][colIndex+col] == val {
				return true
			}
		}
	}
	return false
}

// checks whether the given value is able to be placed in the given position
func checkIfSafe(grid [][]int, row int, col int, val int) bool {
    return grid[row][col] == emptyCell && !checkRow(grid, row, val) && !checkCol(grid, col, val) && !checkSection(grid, row-row%3, col-col%3, val)
}

// prints the grid
func printGrid(grid [][]int) {
	for rowIndex := 0; rowIndex < dim; rowIndex++ {
		for colIndex := 0; colIndex < dim; colIndex++ {
			fmt.Printf("%d ", grid[rowIndex][colIndex])
		}
		fmt.Println()
	}
}

// solves the puzzle
func SolveSudoku(input [][]int) [][]int {
	// checks if input is 9x9
	if len(input) != 9 {
		return nil
	}
	for _, row := range input {
		if len(row) != 9 {
			return nil
		}
	}

	if !isValidInput(input){
		return nil
	}
	
	if !solve(input) {
		return nil
	}

	return input
}

func isValidInput(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		if !checkValidRow(grid, i) || !checkValidCol(grid, i) || !checkValidSection(grid, i) {
			return false
		}
	}
	return true
}

// checks if the specified row is valid
func checkValidRow(grid [][]int, row int) bool {
	visited := make(map[int]bool)
	for _, value := range grid[row] {
		if value != 0 {
			if visited[value] {
				return false
			}
			visited[value] = true
		}
	}
	return true
}

// checks if the specified column is valid
func checkValidCol(grid [][]int, col int) bool {
    visited := make(map[int]bool)
    for row := 0; row < 9; row++ {
        value := grid[row][col]
        if value != 0 {
            if visited[value] {
                return false
            }
            visited[value] = true
        }
    }
    return true
}

// checks if the specified 3x3 box is valid
func checkValidSection(grid [][]int, box int) bool {
    visited := make(map[int]bool)
    rowOffset := (box / 3) * 3
    colOffset := (box % 3) * 3
	// section row
    for i := 0; i < 3; i++ {
		// section column
        for j := 0; j < 3; j++ {
            value := grid[rowOffset+i][colOffset+j]
            if value != 0 {
                if visited[value] {
                    return false
                }
                visited[value] = true
            }
        }
    }
    return true
}

// main function, used for debugging
func main() {
	grid := [][]int {
		{0, 0, 1, 0, 5, 0, 0, 4, 0},
		{0, 0, 3, 1, 0, 0, 2, 6, 0},
		{0, 0, 0, 0, 7, 0, 1, 0, 0},
		{0, 0, 2, 7, 9, 0, 0, 1, 0},
		{0, 0, 0, 6, 0, 0, 5, 2, 0},
		{0, 6, 0, 0, 0, 0, 7, 0, 0},
		{0, 5, 0, 0, 0, 0, 0, 9, 0},
		{0, 0, 4, 0, 2, 0, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 6, 0, 5},
	}


	solved := SolveSudoku(grid)
	if solved != nil {
		printGrid(solved)
		return
	}
	fmt.Println("Cannot be solved")
}
