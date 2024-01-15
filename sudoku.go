package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Check whether a sequence of 9 numbers is "valid":
// - The numbers 1 to 9 must be contained at most 1 time
// - There may not be any numbers greater than 9 in the sequence
// - Note that empty cells (containing 0) are ignored
func checkSeq(seq []int) bool {
	freq := make([]int, 9) // Frequency array for 1 to 9
	for _, v := range seq {
		if v == 0 { // Ignore empty cells
			continue
		}
		if v > 9 { // Cells only contain 1 to 9
			return false
		}
		freq[v-1]++ // Record occurence
	}
	for _, v := range freq {
		if v > 1 {
			return false // Maximum 1 occurence per number
		}
	}
	return true
}

// Check whether the 9x9 grid is "valid":
// - All rows must be valid as per 'checkSeq'
// - All columns must be valid as per 'checkSeq'
// - All 3x3 subgrids must be valid as per 'checkSeq'
func validGrid(grid [][]int) bool {
	// Check all rows
	for _, row := range grid {
		if !checkSeq(row) {
			return false
		}
	}

	// Check all columns
	for c := 0; c < 9; c++ {
		seq := make([]int, 0)
		for j := 0; j < 9; j++ {
			seq = append(seq, grid[j][c])
		}
		if !checkSeq(seq) {
			return false
		}
	}

	// Check all 3x3 subgrids
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			// i and j denote the top left "corner" of the subgroup
			subgrid := make([]int, 0)
			for a := 0; a < 3; a++ {
				for b := 0; b < 3; b++ {
					subgrid = append(subgrid, grid[i+a][j+b])
				}
			}
			if !checkSeq(subgrid) {
				return false
			}
		}
	}

	return true
}

// Move a pointer (i, j) "forward"
func forward(i, j int) (int, int) {
	if j == 8 { // About to "overflow" so wrap around
		return i + 1, 0
	}
	return i, j + 1
}

func backward(i, j int) (int, int) {
	if j == 0 {
		return i - 1, 8
	}
	return i, j - 1
}

func printGrid(grid [][]int) {
	fmt.Println(strings.Repeat("=", 17))
	for _, row := range grid {
		for _, v := range row {
			fmt.Print(strconv.Itoa(v) + " ")
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("=", 17))
}

func SolveSudoku(grid [][]int) [][]int {
	// Store where the pre-initialized cells are
	fixed := make([][]int, 9)
	for i := 0; i < 9; i++ {
		row := make([]int, 9)
		for j := 0; j < 9; j++ {
			row[j] = grid[i][j]
		}
		fixed[i] = row
	}

	i, j := 0, 0 // Start the pointer at top left
	for !(i == 8 && j == 8) || !validGrid(grid) {
		if validGrid(grid) {
			// If current position is unchangeable, skip it
			if fixed[i][j] != 0 {
				i, j = forward(i, j)
				continue
			}
			grid[i][j]++
			i, j = forward(i, j)
		} else {
			i, j = backward(i, j) // Move back to fix the error
			if fixed[i][j] != 0 {
				continue // Keep moving until we get a changeable cell
			}
			done := false
			// Try all numbers greater than the current cell value
			for grid[i][j] < 10 {
				if validGrid(grid) { // Problem fixed, move on
					i, j = forward(i, j)
					done = true
					break
				}
				grid[i][j]++
			}
			if done {
				continue
			}

			grid[i][j] = 0       // Problem not solved, move back even more
			if validGrid(grid) { // Invalid grids will move back automatically
				for { // Find the next changeable cell (going backwards)
					i, j = backward(i, j)
					if fixed[i][j] == 0 {
						break
					}
				}
			}
		}
	}

	printGrid(grid)
	return grid
}
