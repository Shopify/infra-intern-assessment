//Go program to solve a 9x9 sudoku puzzle
package main

//program is using the packages with import path "fmt"
import (
	"fmt"
)

// SolveSudoku function solves the given Sudoku grid and returns the solved grid.
func SolveSudoku(grid [][]int) [][]int {
	// If the sudoku grid is solvable, return the solved grid
	if solve(&grid) {
		return grid
	}
	// If no solution is found, return nil
	return nil
}

// solve function recursively solves the grid using backtracking algorithm.
func solve(grid *[][]int) bool {
	// Find the first unassigned location in the grid
	row, col := findUnassignedLocation(grid)
	// If no unassigned location is found, return true indicating all cells are assigned
	if row == -1 && col == -1 {
		return true
	}

	// Try assigning numbers from 1 to 9 to the current position
	for num := 1; num <= 9; num++ {
		// Check if it's safe to place the number at the current position
		if isSafe(grid, row, col, num) {
			// Assign the number
			(*grid)[row][col] = num

			// If the assignment is successful, recursively solve
			if solve(grid) {
				return true
			}

			// If the assigned number does not result in a solution, reset it
			(*grid)[row][col] = 0
		}
	}
	// No valid number found for the current position
	return false
}

// findUnassignedLocation function finds the first unassigned location in the grid.
func findUnassignedLocation(grid *[][]int) (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if (*grid)[row][col] == 0 {
				return row, col
			}
		}
	}
	// If no unassigned location is found, return (-1, -1)
	return -1, -1
}

// isSafe function checks if it is safe to place a number in a specific position.
func isSafe(grid *[][]int, row, col, num int) bool {
	return !usedInRow(grid, row, num) &&
		!usedInCol(grid, col, num) &&
		!usedInBox(grid, row-row%3, col-col%3, num)
}

// usedInRow function checks if the given number is present in the specified row.
func usedInRow(grid *[][]int, row, num int) bool {
	for col := 0; col < 9; col++ {
		if (*grid)[row][col] == num {
			return true
		}
	}
	return false
}

// usedInCol function checks if the given number is present in the specified column.
func usedInCol(grid *[][]int, col, num int) bool {
	for row := 0; row < 9; row++ {
		if (*grid)[row][col] == num {
			return true
		}
	}
	return false
}

// usedInBox function checks if the given number is present in the specified 3x3 box.
func usedInBox(grid *[][]int, boxStartRow, boxStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if (*grid)[row+boxStartRow][col+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}

func main() {
	// Example usage:
	grid := [][]int{
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

	// Print the unsolved Sudoku puzzle (input)
	fmt.Println("Sudoku puzzle:")
	printGrid(grid)

	// Solve the Sudoku puzzle
	solvedGrid := SolveSudoku(grid)

	// Check if a solution exists
	if solvedGrid != nil {
		// Print the solved Sudoku puzzle
		fmt.Println("\nSolved Sudoku:")
		printGrid(solvedGrid)
	} else {
		// Print a message indicating that no solution exists
		fmt.Println("\nNo solution exists.")
	}
}

// printGrid prints the Sudoku grid in the desired structure.
func printGrid(grid [][]int) {
	fmt.Println("[") // Print the starting bracket for the grid
	for _, row := range grid {
		fmt.Print("  [") // Print the starting bracket for each row
		for i, val := range row {
			fmt.Printf("%d", val) // Print the value of the cell
			// Add a comma if it's not the last element in the row
			if i < len(row)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("],") // Print the ending bracket for the row and a comma
	}
	fmt.Println("]") // Print the ending bracket for the grid
}
