//Moez Bajwa
//Shopify Infrastructure Engineering Intern Assessment

package main

func SolveSudoku(grid [][]int) [][]int {

	//size of sudoku grid
	s := 9

	//Checks if a position on the board is valid for the time being.
	//Is provided the coordinates of the point on the sudoku grid (x,y), the proposed value for that point (n), and the grid.
	//Returns true or false based off checks.
	isValid := func(x int, y int, n int, grid [][]int) bool {
		//Checks row of point for redundancy of n.
		for i := 0; i < s; i++ {
			if grid[x][i] == n {
				return false
			}
		}

		//Checks column of point for redundancy of n.
		for i := 0; i < s; i++ {
			if grid[i][y] == n {
				return false
			}
		}

		//Checks parent 3x3 grid of point for redundancy of n.

		//Find and set start points to the top left corner of its parent 3x3 grid in order to check entire grid.
		x = (x / 3) * 3
		y = (y / 3) * 3

		for i := x; i < x+3; i++ {
			for j := y; j < y+3; j++ {
				if grid[i][j] == n {
					return false
				}
			}
		}

		//Not found in any row, column or 3x3 so it is considered valid for the time being.
		return true

	}

	//Checks if all Sudoku board positions have been filled.
	isFull := func(grid [][]int) bool {
		for i := 0; i < s; i++ {
			for j := 0; j < s; j++ {
				if grid[i][j] == 0 {
					return false
				}
			}
		}
		return true
	}

	//Gets the next empty position, going from top left to bottom right of Sudoku board.
	//Returns coordinates of the next empty position and -1, -1 if full.
	getNextEmpty := func(grid [][]int) (int, int) {
		for i := 0; i < s; i++ {
			for j := 0; j < s; j++ {
				if grid[i][j] == 0 {
					return i, j
				}
			}
		}
		return -1, -1
	}

	//Recursively solves Sudoku board, returning true if there is a solution.
	var solve func([][]int) bool

	solve = func(grid [][]int) bool {
		//Base case checks if board is filled, and then returns true.
		if isFull(grid) {
			return true
		}

		//Get the next empty cell on the board.
		px, py := getNextEmpty(grid)

		//Loop through possible values for that cell (1-9).
		for i := 1; i <= s; i++ {

			//If a value is valid (not in any other row, column or 3x3), assign value to that cell and recursively solve the next empty cell.
			if isValid(px, py, i, grid) {
				grid[px][py] = i
				if solve(grid) {
					return true
				} else {
					//Backtracks if that value did not ent up working, and returns false.
					grid[px][py] = 0
				}
			}
		}
		return false
	}

	//Call solve function and return solved grid.
	solve(grid)
	return grid
}
