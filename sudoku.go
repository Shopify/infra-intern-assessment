package main

const GRIDSIZE = 9

/*
Helper function to ensure current number works within its box
*/
func CheckBoxValid(grid [][]int, xCoord int, yCoord int, number int) bool {
	boxXCoord := xCoord - (xCoord % 3)
	boxYCoord := yCoord - (yCoord % 3)
	for xInc := 0; xInc <= 2; xInc++ {
		for yInc := 0; yInc <= 2; yInc++ {
			if grid[boxXCoord+xInc][boxYCoord+yInc] == number {
				return false
			}
		}
	}
	return true
}

/*
Helper function to ensure current number works within its column
*/
func CheckColumnValid(grid [][]int, xCoord int, number int) bool {
	for yCoord := 0; yCoord < 9; yCoord++ {
		if grid[xCoord][yCoord] == number {
			return false
		}
	}
	return true
}

/*
Helper function to ensure current number works within its row
*/
func CheckRowValid(grid [][]int, yCoord int, number int) bool {
	for xCoord := 0; xCoord < 9; xCoord++ {
		if grid[xCoord][yCoord] == number {
			return false
		}
	}
	return true
}

/*
Helper function to determine next cell to visit
*/
func GetNextCell(xCoord int, yCoord int) (int, int) {
	if xCoord == GRIDSIZE-1 {
		return 0, yCoord + 1
	}
	return xCoord + 1, yCoord
}

/*
Helper recursive function to perform backtracking
the function returns boolean value to determine whether
a correct value has been found. Since slices are pass
by reference by default in Go, the grid does not need to be
modified
*/
func SolveCell(grid [][]int, xCoord int, yCoord int) bool {
	if yCoord == GRIDSIZE {
		return true
	}

	newX, newY := GetNextCell(xCoord, yCoord)

	if grid[xCoord][yCoord] != 0 {
		return SolveCell(grid, newX, newY)
	}

	for possibleValue := 1; possibleValue <= 9; possibleValue++ {
		if CheckColumnValid(grid, xCoord, possibleValue) &&
			CheckRowValid(grid, yCoord, possibleValue) &&
			CheckBoxValid(grid, xCoord, yCoord, possibleValue) {
			grid[xCoord][yCoord] = possibleValue
			result := SolveCell(grid, newX, newY)
			if result {
				return result
			} else {
				grid[xCoord][yCoord] = 0
			}
		}
	}
	return false
}

/*
Main entry point for solving sudoku, using
backtracking algorithm combined with recursion
*/
func SolveSudoku(grid [][]int) [][]int {
	SolveCell(grid, 0, 0)
	return grid
}
