package main

/*
	Function smallGrid checks if the 3x3 grid around row i and col j is valid
	under traditional sudoku rules.
	Returns true if grid beginning and i and j is valid, false otherwise
*/
func smallGrid(i int, j int, grid [][]int) bool {
	var seen [10]int
	for col := i; col < i+3; col++ {
		for row := j; row < j+3; row++ {
			curr := grid[row][col]
			if curr != 0 && seen[curr] == 1 {
				return false
			} else if curr != 0 {
				seen[curr] = 1
			}
		}
	}
	return true
}

/*
	Function isValidSudoku checks if the given sudoku board is valid
	The function checks every row, column, and 3x3 grid to see if it is valid.
	Returns true if valid, false otherwise
*/
func isValidSudoku(grid [][]int) bool {

	// check if every row is valid
	for i := 0; i < len(grid); i++ {
		var seen [10]int
		for j := 0; j < len(grid[0]); j++ {
			curr := grid[i][j]
			if curr != 0 && seen[curr] == 1 {
				return false
			} else if curr != 0 {
				seen[curr] = 1
			}
		}
	}

	// check if every column is valid
	for j := 0; j < len(grid); j++ {
		var seen [10]byte
		for i := 0; i < len(grid[0]); i++ {
			curr := grid[i][j]
			if curr != 0 && seen[curr] == 1 {
				return false
			} else if curr != 0 {
				seen[curr] = 1
			}
		}
	}

	// check all three by three grids in the board
	for col := 0; col < 7; col += 3 {
		for row := 0; row < 7; row += 3 {
			if !(smallGrid(row, col, grid)) {
				return false
			}
		}
	}
	return true
}
/*
	Function doSolving solves the sudoku puzzle using backtracking
	Returns true if the puzzle is solved, calls itself recurisvely
	trying to replace 0s with the correct number to solve the puzzle
*/
func doSolving(grid [][]int) bool{
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					grid[i][j] = num
					if isValidSudoku(grid) && doSolving(grid) {
						return true
					}
					grid[i][j] = 0
				}
				return false
			}
		}
	}
	return true
}

/*
	Function SolveSudoku solves the sudoku puzzle by calling doSolving
	Returns the original grid that was passed in since its been modified
*/

func SolveSudoku(grid [][]int) [][]int {
	doSolving(grid)
	return grid
}
