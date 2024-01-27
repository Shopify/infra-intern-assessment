package main

import (
	"container/heap"
)

// Gets map of posibilities of all positions.
// Map is accessed using the position and it stores a integer set (IntSet) of all possible values which can be assigned at a given position
func getAllPossibilities(board [][]int) map[Position]IntSet {
	allPossibilities := make(map[Position]IntSet)
	for row := 0; row < SudokuSize; row++ {
		for column := 0; column < SudokuSize; column++ {
			if board[row][column] == 0 {
				position := Position{
					row: row, column: column,
				}
				allPossibilities[position] = getPossibilitiesAt(board, position)
			}
		}
	}

	return allPossibilities
}

// Gets all possible values of position in the given matrix.
// It initializes a set of all possible value 1-9 and then removes any value found in the row, column or box.
func getPossibilitiesAt(board [][]int, position Position) IntSet {
	possibilities := make(IntSet)

	for i := 1; i <= SudokuSize; i++ {
		possibilities.Add(i)
	}
	// Remove possibilities present in row and column
	for i := 0; i < SudokuSize; i++ {
		if possibilities.Has(board[position.row][i]) {
			possibilities.Delete((board[position.row][i]))
		}
		if possibilities.Has(board[i][position.column]) {
			possibilities.Delete(board[i][position.column])
		}
	}

	// Remove possibilities present in box
	boxRowIndex, boxColumnIndex := (position.row/3)*3, (position.column/3)*3
	for row := boxRowIndex; row < boxRowIndex+SudokuBoxSize; row++ {
		for column := boxColumnIndex; column < boxColumnIndex+SudokuBoxSize; column++ {
			if !(position.row == row || position.column == column) && possibilities.Has(board[row][column]) {
				possibilities.Delete(board[row][column])
			}
		}
	}
	return possibilities
}

// Initializes Priority Queue (min heap) using all position possibilities.
// The priority queue stores the position and prioritizes it based on number of possible values which could be assigned in the current state.
// The lower the number of values, the higher its priority would be as its built on a min heap.
func initializePriorityQueue(allPossibilities map[Position]IntSet) PriorityQueue {
	pq := make(PriorityQueue, len(allPossibilities))
	i := 0
	for position := range allPossibilities {
		possibilities := allPossibilities[position]
		pq[i] = &Item{
			position: position,
			priority: possibilities.Size(),
		}
		i++
	}
	heap.Init(&pq)
	return pq
}

// Removes possible value at position and adds item with new priority to priority queue.
func removePossibleValueAtPosition(allPossibilities *map[Position]IntSet, pq *PriorityQueue, position Position, value int) {
	if possibilities, ok := (*allPossibilities)[position]; ok && possibilities.Has(value) {
		possibilities.Delete(value)
		item := &Item{
			position: position,
			priority: possibilities.Size(),
		}
		heap.Push(pq, item)
	}
}

// This function removes all possibilities of given value for all unsassigned position in the row, column and box.
// By reducing possibilities, their priority goes up. Hence we add item with new priority to priority queue.
func updatePossibilitiesAndPriorityQueue(allPossibilities *map[Position]IntSet, pq *PriorityQueue, position Position, value int) {
	//Update all possibilities in the given row and column
	for k := 0; k < SudokuSize; k++ {
		//Remove possibility of value in entire row
		updateRowPosition := Position{row: position.row, column: k}
		removePossibleValueAtPosition(allPossibilities, pq, updateRowPosition, value)

		//Remove possibility of value in entire column
		updateColumnPosition := Position{row: k, column: position.column}
		removePossibleValueAtPosition(allPossibilities, pq, updateColumnPosition, value)
	}

	boxRowIndex, boxColumnIndex := (position.row/3)*3, (position.column/3)*3
	//Remove possibility of value within entire box
	for row := boxRowIndex; row < boxRowIndex+SudokuBoxSize; row++ {
		for column := boxColumnIndex; column < boxColumnIndex+SudokuBoxSize; column++ {
			updateBoxPosition := Position{row: row, column: column}
			removePossibleValueAtPosition(allPossibilities, pq, updateBoxPosition, value)
		}
	}
}

// Checks if boards is solved by looking for unassigned positions
func isBoardSolved(board [][]int) bool {
	for row := 0; row < SudokuSize; row++ {
		for column := 0; column < SudokuSize; column++ {
			if board[row][column] == 0 {
				return false
			}
		}
	}
	return true
}

// Creates deep copy of board
func deepCopy(board [][]int) [][]int {
	duplicate := make([][]int, SudokuSize)
	data := make([]int, SudokuSize*SudokuSize)
	for i := range board {
		start := i * SudokuSize
		end := start + SudokuSize
		duplicate[i] = data[start:end:end]
		copy(duplicate[i], board[i])
	}
	return duplicate
}
