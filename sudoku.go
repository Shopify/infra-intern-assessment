package main

// - Here is my final solution for the Shopifly Technical Challenge
// - I created a backtracking algorithm to create the sudoku solver.
// - My solution has three functions. The first function (SolveSudoku) simply calls the main helper function
// which is called solve_sudoku on the grid. solve_sudoku is the main recursive function that does the solving
// and backtracking.
// - Lastly solve_sudoku often calls is_legal to check if the number to
// put on an empty spot is legal by sudoku rules.

// - This is SolveSudoku
// - The input for this function is an 9x9 Sudoku grid as described in the handout
// - It then calls the solve_sudoku helper function adding 0 for the row and column arguments.
// - This ensures that solve_sudoku fills in every empty cell in the grid since it starts at grid[0][0] at recurses through the grid.
// - solve_sudoku mutates the grid and this function then returns the completed grid.
// Preconditions: The input grid must follow all instructions as mentioned in the challenge Readme

func SolveSudoku(grid [][]int) [][]int {
	solve_sudoku(grid, 0, 0) // calls helper function for the entire grid
	return grid
}

// - This is solve_sudoku
// - This is a recursive function that uses backtracking.
// - The first three if and else if statements check the general location of the cell
// - The input is a 9x9 sudoku grid and the exact row and column to start at
// - It solves from left the right hgoing down the grid.
// - It returns true if the row is equal to 9. This is because the indexing starts at
// 0 so if the pointer reaches row 9 it means it has already checked all of row 8 which is
// the last row.
// - The same idea applies to columns, if it reaches column 9 then we know that that entire row is checked
// and we can go to start the next row which is grid[row + 1][0] and recursively call the function at that location
// - If the row and column inputs are valid it then checks if it has been filled (i.e doesn't contain 0)
// - If it has been filled it moves on until it finds a cell with a zero by recursively calling the function on the adjacent cell.

// - The last else statement is where the backtracking happens
// - If we reach the else block, this means that at this location there is a zero
// - The function then iterates through all possible numbers 1-9 then calls isLegal using the the coordinates and
// a number from 1-9 to see if you can put the number there without breaking the rules
//- If it can't find a suitable number it returns false
//- If it can find one, it fills that cell with that number and recursively calls the function again
// in the adjacent block. It then recurses through the grid, if it runs into a case where it can't find a suitable
// number it would backtrack and set the spot to 0.

func solve_sudoku(grid [][]int, r, c int) bool {
	if r == 9 {
		return true // function has checked entire grid and can return true

	} else if c == 9 {
		return solve_sudoku(grid, r+1, 0) //  function has checked the entire row and can go to the next row

	} else if grid[r][c] != 0 {
		return solve_sudoku(grid, r, c+1) // if the cell is already filled, move to the next cell

	} else {
		for k := 1; k <= 9; k++ {
			if is_legal(grid, r, c, k) { // checks if the selected number in the cell would be allowed. If, yes place number in the cell
				grid[r][c] = k
				if solve_sudoku(grid, r, c+1) { // recursive step, calls function on next cell
					return true
				}
				grid[r][c] = 0 // if further moves are false, stop backtrack and make cell 0
			}
		}
		return false // no possible moves
	}

}

// This is the is_legal function
// - It returns a boolean value on whether adding a number to a cell at the given coordinates is legal.
// - This means that each row, column, and sub-grid contains all the numbers from 1-9 without repetition.
// - This function first checks if the given number is in the same row (all values in grid[row]).
// - It then checks if the given number is in the same column (all values in grid[i][col] for a 0<= i <9).
// - Lastly, it checks if the number is in the same 3 by 3 grid we do this by calucating the start and end coordinates of the grid
// They are 3*(row/ floor(3)) and 3*(col/floor(3)
// - We then check all values in range 3 * (row/ floor(3)) to (3 * (row/ floor(3))) +3
// and 3 * col/floor(3) to (3 * (col/floor(3))) + 3, we must check all cells in that range. There are 9 in total.

func is_legal(grid [][]int, row, col, value int) bool {

	for _, num := range grid[row] {
		if num == value { // checks if the given number is in the same row
			return false // The attempted move is not allowed
		}
	}

	for i := 0; i < 9; i++ {
		if grid[i][col] == value { // checks if the given number is in the same column
			return false // the attempted move is not allowed
		}
	}

	startRow, startCol := 3*(row/3), 3*(col/3) // get start coordinates of the 3x3 grid
	for i := 0; i < 3; i++ {                   // use nested for loops to iterate within the 3x3 grid
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == value { // checks if the given number is in the 3x3 grid
				return false // the attempted move is not allowed
			}
		}
	}

	return true // the cell with selected number is allowed
}
