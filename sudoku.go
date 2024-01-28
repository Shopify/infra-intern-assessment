package main

const gridSize = 9 // the grid size is 9x9 and does not change

func SolveSudoku(grid [][]int) [][]int {
	solveSudokuHelper(grid)
	return grid
}

// solveSudokuHelper is a helper function that updates the input grid recursively
func solveSudokuHelper(grid [][]int) bool {
    emptyCell := findEmptyCell(grid)
    if emptyCell == nil {
        return true
    }

	// row and column of the empty cell
    row, col := emptyCell[0], emptyCell[1]

	// check all the number possibilities from 1..9
    for num := 1; num <= gridSize; num++ {
        if isValidMove(grid, num, []int{row, col}) {
            grid[row][col] = num // put the new number in the grid

            if solveSudokuHelper(grid) {
                return true
            }

			// backtrack if the number is not valid, and empty the cell
            grid[row][col] = 0
        }
    }

    return false
}

// isValidMove is a helper function that checks if the new added number is valid in the row, column, and sub-grid
func isValidMove(grid [][]int, num int, pos []int) bool {
    return isValidRow(grid, num, pos) && isValidColumn(grid, num, pos) && isValidSubGrid(grid, num, pos)
}

// isValidRow is a helper function that checks if the add number is a valid entry in the row
func isValidRow(grid [][]int, num int, pos []int) bool {
	// row and column of the current number
	numRow, numCol := pos[0], pos[1]

    for col := 0; col < gridSize; col++ {
        if grid[numRow][col] == num && numCol != col {
            return false
        }
    }
    return true
}

// isValidColumn is a helper function that checks if the add number is a valid entry in the column
func isValidColumn(grid [][]int, num int, pos []int) bool {
	// row and column of the current number
	numRow, numCol := pos[0], pos[1]

    for row := 0; row < gridSize; row++ {
        if grid[row][numCol] == num && numRow != row {
            return false
        }
    }
    return true
}

// isValidSubGrid is a helper function that checks if the add number is valid in the sub-grid
func isValidSubGrid(grid [][]int, num int, pos []int) bool {
	// row and column of the current number
	numRow, numCol := pos[0], pos[1]

	// find the top left (first) number of the sub-grid
    subGridStartRow, subGridStartCol := (numRow/3)*3, (numCol/3)*3

    for row := subGridStartRow; row < subGridStartRow+3; row++ {
        for col := subGridStartCol; col < subGridStartCol+3; col++ {
            if grid[row][col] == num && (row != numRow || col != numCol) {
                return false
            }
        }
    }
    return true
}

// findEmptyCell is a helper function that returns the position of the cell if it contains 0
// meaning that the cell is empty
func findEmptyCell(grid [][]int) []int {
    for row := 0; row < gridSize; row++ {
        for col := 0; col < gridSize; col++ {
            if grid[row][col] == 0 {
                return []int{row, col}
            }
        }
    }
    return nil
}
