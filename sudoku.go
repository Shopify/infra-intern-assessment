// File: sudoku.go
// Created Date: Jan 29 2024
// Author: Indranil Palit
// Description: This file contains the code to solve the Sudoku puzzle

package main

import "fmt"

const N = 9 // Size of the Sudoku grid

// I am using Backtracking to solve the Sudoku puzzle.
// This function takes a partially filled-in grid and attempts to assign values to all the empty cells in such a way to meet the requirements for Sudoku solution.
// This function returns true if a solution is found, otherwise returns false.
func solveSudoku(grid [][]int) bool {
	var row, col int

	// Find an empty cell. Basically iterate through the grid and find the first empty cell
	found := false
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				row, col = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	// Base case for recursion: If no empty cell is found, then the Sudoku puzzle is solved
	if !found {
		return true
	}

	// Once we find a cell to be filled, we try possible digits.
	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num

			// Recursively call the function to solve the puzzle
			if solveSudoku(grid) {
				return true
			}

			// Backtrack if after putting the digit, we are not able to solve the puzzle. Put 0 in the cell and try the next digit
			grid[row][col] = 0
		}
	}

	return false
}

// Function to solve the Sudoku puzzle
func SolveSudoku(grid [][]int) [][]int {

	// Create a deep-copy of the grid so that we can modify it without affecting the original grid
	newGrid := make([][]int, N)
	for i := range grid {
		newGrid[i] = make([]int, N)
		copy(newGrid[i], grid[i])
	}

	solveSudoku(newGrid) // I am not checking the return value here since the problem statement guarantees that there will be a solution.

	return newGrid
}

// Function to print the Sudoku grid. I am not sure if it's required but I am doing it since in the problem description it says:
// "Your program should implement an efficient algorithm to solve the Sudoku puzzle and print the solved grid to the console."
func printGrid(grid [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	// Define the Sudoku grid
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
	solution := SolveSudoku(grid)

	printGrid(solution)
}
