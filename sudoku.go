package main

import "fmt"

// struct for storing cell location
type loc struct {
	row int
	col int
}

// Finding all available numbers for the cell at location
func AvailableNums(input [][]int, location loc) []int {
	nums := [9]bool{false, false, false, false, false, false, false, false, false}
	for i := 0; i < 9; i++ {
		// check available nums by row
		if input[location.row][i] != 0 {
			nums[input[location.row][i]-1] = true
		}
		// check available nums by column
		if input[i][location.col] != 0 {
			nums[input[i][location.col]-1] = true
		}
	}
	// check available nums by grid
	var grid_x, grid_y int
	grid_x = (location.row / 3) * 3
	grid_y = (location.col / 3) * 3
	for i := grid_x; i < grid_x+3; i++ {
		for j := grid_y; j < grid_y+3; j++ {
			if input[i][j] != 0 {
				nums[input[i][j]-1] = true
			}
		}
	}
	// unseen numbers are available nums for this cell
	output := []int{}
	for i := 0; i < 9; i++ {
		if !nums[i] {
			output = append(output, i+1)
		}
	}
	return output
}

// check if all cells are filled, find the next empty cell if not
func IsDone(input [][]int, pos *loc) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if input[i][j] == 0 {
				(*pos).row = i
				(*pos).col = j
				return false
			}
		}
	}
	return true
}

// helper function to solve the puzzle
func _solve(input [][]int) bool {
	pos := loc{-1, -1}
	// check if puzzle is solved and update next empty cell
	if IsDone(input, &pos) {
		return true
	}
	nums := AvailableNums(input, pos)
	row := pos.row
	col := pos.col
	// try each possible number for current cell, dfs
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		input[row][col] = num
		if _solve(input) {
			return true
		} else {
			input[row][col] = 0
		}
	}
	return false
}

func SolveSudoku(input [][]int) [][]int {
	if !_solve(input) {
		fmt.Print("unsolvable!\n")
	}
	return input
}
