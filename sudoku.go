package main

import "fmt" 

const SIZE = 9

// We assume that a puzzle is actually solveable and also assume that a full grid implies it has been solve
// Implements backtracking algorithm in order to solve the problem
func SolveSudoku(input [][]int) [][]int {

	//make a copy of the grid
	grid := copyGrid(input)

	//iterate through all the cells in the grid
	for i := 0; i < SIZE; i++ {
		for k := 0; k < SIZE; k++ {

			//if the current cell needs to be filled
			if grid[i][k] == 0 {

				//iterate through all possible values
				for val := 1; val <= SIZE; val++ {

					//check if it is valid to place the value in that cell
					if isValid(i, k, val, input) {

						//place the value in the cell
						grid[i][k] = val

						//store the result grid from continuing the solve process
						resultGrid := SolveSudoku(grid)

						//if the grid is not nil that means we have a solved grid. return it
						if resultGrid != nil {
							return resultGrid
						}

						//if this value didn't work then reset the cell back to zero
						grid[i][k] = 0
					}
				}
				//return nil indicating the grid is unsolvable in its current state; backtrack to previous state
				return nil
			}
		}
	}

	//print out the grid once it has been solved
	printGrid(grid)

	//return the solved grid
	return grid
}

// checks if a specific value can be placed within a particular cell
func isValid(row, col, value int, grid [][]int) bool {

	//ensure the value is not contained in the row or column
	for i := 0; i < SIZE; i++ {
		if grid[row][i] == value || grid[i][col] == value {
			return false
		}
	}

	//calculate the top x coordinate
	topX := 3 * (row / 3)

	//calculate the top y coordinate
	topY := 3 * (col / 3)

	//iterate through the x and y coordinates and ensure value is not present
	for i := topX; i < topX+3; i++ {
		for k := topY; k < topY+3; k++ {

			//if the value is already present then return false
			if grid[i][k] == value {
				return false
			}
		}
	}

	return true
}

// helper function to copy the grid so original is not passed around allows for more dynamic programming
func copyGrid(input [][]int) [][]int {

	grid := make([][]int, SIZE)

	for i := range grid {
		grid[i] = make([]int, SIZE)
		copy(grid[i], input[i][:])
	}

	return grid
}

// Prints a Sudoku grid to the terminal
// It takes in a grid paramater that is assumed to be 9 columns and 9 rows
func printGrid(grid [][]int) {
	fmt.Println("[")
	for i := 0; i < SIZE; i++ {
		fmt.Print("[")
		for k := 0; k < SIZE-1; k++ {
			fmt.Printf("%d, ", grid[i][k])
		}
		fmt.Printf("%d]\n", grid[i][8])
	}
	fmt.Println("]")
}
