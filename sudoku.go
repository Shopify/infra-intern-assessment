package main

import (
	"fmt"
)


// Check if the input sudoku is valid
// Params:
// grid: This is the grid that we are trying to solve
// Returns: True if the sudoku is of valid size else false
func validateSudoku(grid [][]int) bool{
	if len(grid) != 9 {
		return false
	}
	for _, row := range grid{
		if len(row) != 9{
			return false
		}
	}
	return true

}

// Function to deep copy the grid to the destination (newgrid)
// Params:
// grid: This is the grid that we will be copying
// newgrid: This is the grid that stores the copied grid
func deepcopy(grid [][]int, newgrid[][]int){
	for i := 0; i < 9; i++{
		for j :=0; j < 9; j++{
			newgrid[i][j] = grid[i][j]
		}
	}
}

// Check if the num is in the row
// Params:
// grid: This is the grid that we are trying to solve
// row: This is index of the row
// num: This is the number that we are checking
func numInRow(grid [][]int, row int, num int) bool{
	// Loop to check the col
	for i := 0; i < 9; i++{
		if grid[row][i] == num{
			return true
		}
	}
	return false
}

// Check if the num is in the col
// Params:
// grid: This is the grid that we are trying to solve
// col: This is index of the col
// num: This is the number that we are checking
func numInCol(grid [][]int, col int, num int) bool{
	// Loop to check the col
	for i := 0; i < 9; i++{
		if grid[i][col] == num{
			return true
		}
	}
	return false
}

// Check if the num is in the box
// Params:
// grid: This is the grid that we are trying to solve
// row: This is index of the row
// col: This is index of the col
// num: This is the number that we are checking
func numInBox(grid [][]int, row int, col int, num int) bool{
	// Change the row and col to the row and col index of the starting of the box that we wanna check
	row = row - row % 3
	col = col - col % 3

	// Loop to check the box
	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			if grid[i+row][j+col] == num{
				return true
			}
		}
	}
	return false
}

// Check if the num is valid
// Params:
// grid: This is the grid that we are trying to solve
// row: This is index of the row
// col: This is index of the col
// num: This is the number that we are checking
func checkIfValid(grid [][]int, row int, col int, num int) bool{
	return !(numInRow(grid,row,num) || numInCol(grid,col,num) || numInBox(grid, row, col, num))
}

// Function to print the sudoku grid
func printSudoku(grid [][]int){
	for _, row := range grid{
		fmt.Println(row)
	}
}

// Function to create an empty slice of size [9][9]
// Returns: A slice of size 9x9 initialised with 0 
func createEmptySlice() [][]int{
	emptySlice := make([][]int, 9)
	for i := 0; i < 9; i++{
		emptySlice[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			emptySlice[i][j] = 0
		}
	}
	return emptySlice
}

// Function to find the next location to fill
// Params:
// grid: This is the grid that we are trying to solve
// Returns: 
// Tuple containing the row and col index
func findNextLocation(grid [][]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}

	return -1, -1
}

// Function to implement the backtracking algorithm and get the solution if it exists
// The solution gets stored in the grid that is passed into the function
// Returns: A bool value indicating if the solution was found or not
func backtrack(grid [][]int) bool {
	x, y := findNextLocation(grid)

	// Check if the sudoku is solved
	if x == -1 || y == -1 {
		fmt.Println("Solved Sudoku: ")
		printSudoku(grid)
		return true
	}

	// Solve the sudoku
	for i := 1; i <= 9; i++ {
		if checkIfValid(grid, x, y, i) {
			grid[x][y] = i
			if backtrack(grid){
				return true
			}
		}
		grid[x][y] = 0
	}
	return false
}


// Function to solve the sudoku
// Returns: A slice of size 9x9 containing the solved sudoku if the solution exists else returns nil 
func SolveSudoku(grid [][]int) [][]int{

	// Validate the input
	if !validateSudoku(grid){
		return nil
	}
	
	// Create an empty slice to store the solved grid
	solved := createEmptySlice()

	// Deep copy grid to avoid changing it
	deepcopy(grid, solved)

	// Backtrack to solve it and check if a solution is found
	if !backtrack(solved){
		// Return nil if no solution found
		return nil
	}

	// Return the solved grid 
	return solved
}