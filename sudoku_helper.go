// Package main provides a Sudoku solver that uses backtracking and recursion to find a solution.
package main

import "fmt"

// EmptyCell represents the location of an empty cell in the Sudoku grid.
type EmptyCell struct {
	Row, Col int
}

const (
	gridSize    = 9
	subgridSize = 3
)

// Stack is a simple stack data structure for EmptyCell.
type Stack []EmptyCell

// Push adds an empty cell to the stack.
func (s *Stack) Push(cell EmptyCell) {
	*s = append(*s, cell)
}

// Pop removes and returns the top empty cell from the stack.
func (s *Stack) Pop() (EmptyCell, bool) {
	if len(*s) == 0 {
		return EmptyCell{}, false
	}
	index := len(*s) - 1
	cell := (*s)[index]
	*s = (*s)[:index]
	return cell, true
}

// PrintGrid prints the Sudoku grid to the console.
func PrintGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%2d ", cell)
		}
		fmt.Println()
	}
}

// IsValidInRow checks if placing a number in a row is valid.
func IsValidInRow(grid [][]int, row, num int) bool {
	for _, cell := range grid[row] {
		if cell == num {
			return false
		}
	}
	return true
}

// IsValidInCol checks if placing a number in a column is valid.
func IsValidInCol(grid [][]int, col, num int) bool {
	for _, row := range grid {
		if row[col] == num {
			return false
		}
	}
	return true
}

// IsValidInSubgrid checks if placing a number in a 3x3 subgrid is valid.
func IsValidInSubgrid(grid [][]int, startRow, startCol, num int) bool {
	for i := 0; i < subgridSize; i++ {
		for j := 0; j < subgridSize; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

// IsValidPlacement checks if placing a number in a specific position is valid.
func IsValidPlacement(grid [][]int, row, col, num int) bool {
	return IsValidInRow(grid, row, num) && IsValidInCol(grid, col, num) &&
		IsValidInSubgrid(grid, row-row%subgridSize, col-col%subgridSize, num)
}

// findEmptyCells finds and returns a stack of empty cells in the Sudoku grid.
func findEmptyCells(grid [][]int) Stack {
	var emptyCells Stack
	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				emptyCells.Push(EmptyCell{Row: i, Col: j})
			}
		}
	}
	return emptyCells
}
