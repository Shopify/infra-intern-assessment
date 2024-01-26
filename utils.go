package main

import (
	"container/heap"
)

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

func removeValueFromPossibilities(allPossibilities *map[Position]IntSet, pq *PriorityQueue, position Position, value int) {
	if possibilities, ok := (*allPossibilities)[position]; ok {
		if possibilities.Has(value) {
			possibilities.Delete(value)
			item := &Item{
				position: position,
				priority: possibilities.Size(),
			}
			heap.Push(pq, item)
		}
	}
}

func updatePosibilitiesAndPriorityQueue(allPossibilities *map[Position]IntSet, pq *PriorityQueue, position Position, value int) {
	//Update all possibilities in the row and column
	for k := 0; k < SudokuSize; k++ {
		//Remove possibiliy of value in column
		if k != position.column {
			updateRowPosition := Position{row: position.row, column: k}
			removeValueFromPossibilities(allPossibilities, pq, updateRowPosition, value)
		}
		//Remove possibiliy of value in row
		if k != position.row {
			updateColumnPosition := Position{row: k, column: position.column}
			removeValueFromPossibilities(allPossibilities, pq, updateColumnPosition, value)
		}
	}

	boxRowIndex, boxColumnIndex := (position.row/3)*3, (position.column/3)*3
	//Update possibilities of value within the box
	for row := boxRowIndex; row < boxRowIndex+SudokuBoxSize; row++ {
		for column := boxColumnIndex; column < boxColumnIndex+SudokuBoxSize; column++ {
			if !(position.row == row || position.column == column) {
				updateBoxPosition := Position{row: row, column: column}
				removeValueFromPossibilities(allPossibilities, pq, updateBoxPosition, value)
			}
		}
	}
}

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
