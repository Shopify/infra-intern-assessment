package main

import (
	"fmt"
	"log"
	"strings"
)

type Position struct {
	Row, Column int
}

// ! Returns grid with solved sudoku without modifying input grid
func SolveSudoku(inputGrid [][]int) [][]int {
	const EMPTY = 0

	// Check that the input grid is 9x9
	if len(inputGrid) != 9 || len(inputGrid[0]) != 9 {
		log.Printf("Invalid input grid dimensions")
		return inputGrid
	}

	// Copy input grid over to be used in search
	board := make([][]int, 9)
	for i := range board {
		board[i] = make([]int, 9)
		copy(board[i], inputGrid[i])
	}

	//! Helper function to find all possible values for the tile at (row, col)
	getPossibilities := func(square Position) []int {
		row := square.Row
		col := square.Column

		// Check if already assigned
		if board[row][col] != EMPTY {
			return make([]int, 0)
		}

		possibilities := make(map[int]bool)
		for i := 1; i <= 9; i++ {
			possibilities[i] = true
		}

		// Check row
		for _, existingVal := range board[row] {
			possibilities[existingVal] = false
		}
		// Check columns
		for i := 0; i < 9; i++ {
			possibilities[board[i][col]] = false
		}
		// Check subsquare
		subsquareStartR := row - (row % 3)
		subsquareStartC := col - (col % 3)
		for i := subsquareStartR; i < subsquareStartR+3; i++ {
			for j := subsquareStartC; j < subsquareStartC+3; j++ {
				possibilities[board[i][j]] = false
			}
		}

		var ret []int
		for possibility, valid := range possibilities {
			if valid {
				ret = append(ret, possibility)
			}
		}

		return ret
	}

	getUnassignedTiles := func() []Position {
		var ret []Position
		for r, row := range board {
			for c, val := range row {
				if val == EMPTY {
					ret = append(ret, Position{r, c})
				}
			}
		}
		return ret
	}

	//! Modifies board in place
	var recursiveBacktrack func([]Position) bool
	recursiveBacktrack = func(unassignedTiles []Position) bool {
		if len(unassignedTiles) == 0 {
			return true
		}

		var validUnassignedTiles []Position
		for _, unassignedTile := range unassignedTiles {
			if len(getPossibilities(unassignedTile)) > 0 {
				validUnassignedTiles = append(validUnassignedTiles, unassignedTile)
			}
		}

		// If any unassigned tiles have no valid possible values, we can stop early
		if len(validUnassignedTiles) < len(unassignedTiles) {
			return false
		}

		assigningTile := unassignedTiles[len(unassignedTiles)-1]
		possibilities := getPossibilities(assigningTile)

		for _, assigningValue := range possibilities {
			board[assigningTile.Row][assigningTile.Column] = assigningValue
			unassignedTiles = unassignedTiles[:len(unassignedTiles)-1]
			if recursiveBacktrack(unassignedTiles) {
				return true
			}
			board[assigningTile.Row][assigningTile.Column] = EMPTY
			unassignedTiles = append(unassignedTiles, assigningTile)
		}

		return false
	}

	if res := recursiveBacktrack(getUnassignedTiles()); res {
		log.Printf("Successfully solved puzzle!")

		boardStr := "Solution: \n"
		boardStr += "+-----------------+\n"
		for _, row := range board {
			boardStr += strings.Join(strings.Fields(fmt.Sprint(row)), " ")
			boardStr += "\n"
		}
		boardStr += "+-----------------+"
		log.Printf("%s", boardStr)
	} else {
		log.Printf("Failed to solve puzzle...")
	}
	return board
}
