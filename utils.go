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
	var possibilities IntSet
	for i := 1; i <= SudokuSize; i++ {
		possibilities.Add(i)
	}
	// Remove possibilities present in row and column
	for i := 0; i < SudokuSize; i++ {
		// fmt.Println(possibilities.String(), board[position.row][i])
		if possibilities.Has(board[position.row][i]) {
			possibilities.Delete((board[position.row][i]))
		}
		// fmt.Println(possibilities.String(), board[i][position.column])
		if possibilities.Has(board[i][position.column]) {
			possibilities.Delete(board[i][position.column])
		}
		// fmt.Println(possibilities.String())
	}

	boxRowIndex, boxColumnIndex := (position.row/3)*3, (position.column/3)*3
	// fmt.Println(possibilities.String())
	for row := boxRowIndex; row < boxRowIndex+SudokuBoxSize; row++ {
		for column := boxColumnIndex; column < boxColumnIndex+SudokuBoxSize; column++ {
			// fmt.Println(row, column, board[row][column])
			if !(position.row == row || position.column == column) && possibilities.Has(board[row][column]) {
				possibilities.Delete(board[row][column])
			}
		}
	}
	// fmt.Println(possibilities.String())
	// fmt.Println(possibilities.String())
	return possibilities
}

func initializePriorityQueue(allPossibilities map[Position]IntSet) PriorityQueue {
	pq := make(PriorityQueue, len(allPossibilities))
	i := 0
	for position := range allPossibilities {
		// fmt.Println(position, allPossibilities[position].Len())
		possibilities := allPossibilities[position]
		pq[i] = &Item{
			position: position,
			priority: possibilities.Size(),
		}
		i++
	}
	heap.Init(&pq)
	// for pq.Len() > 0 {
	// 	item := heap.Pop(&pq).(*Item)
	// 	fmt.Printf("%.2d:%d,%d\n", item.priority, item.position.row, item.position.column)
	// }
	return pq
}

func updatePositionPosibilities(allPossibilities map[Position]IntSet, position Position, value int) []Position {
	var updatedPositions []Position

	//Update all possibilities in the row and column
	for k := 0; k < SudokuSize; k++ {
		//Remove possibiliy of value in column
		if k != position.column {
			updateRowPosition := Position{row: position.row, column: k}
			if possibilities, ok := allPossibilities[updateRowPosition]; ok {
				if possibilities.Has(value) {
					possibilities.Delete(value)
					updatedPositions = append(updatedPositions, updateRowPosition)
				}
			}
		}
		//Remove possibiliy of value in row
		if k != position.row {
			updateColumnPosition := Position{row: k, column: position.column}
			if possibilities, ok := allPossibilities[updateColumnPosition]; ok {
				if possibilities.Has(value) {
					possibilities.Delete(value)
					updatedPositions = append(updatedPositions, updateColumnPosition)
				}
			}
		}
	}

	boxRowIndex, boxColumnIndex := (position.row/3)*3, (position.column/3)*3

	//Update possibilities of value within the box
	for row := boxRowIndex; row < boxRowIndex+SudokuBoxSize; row++ {
		for column := boxColumnIndex; column < boxColumnIndex+SudokuBoxSize; column++ {
			if !(position.row == row || position.column == column) {
				updateBoxPosition := Position{row: row, column: column}
				if possibilities, ok := allPossibilities[updateBoxPosition]; ok {
					if possibilities.Has(value) {
						possibilities.Delete(value)
						updatedPositions = append(updatedPositions, updateBoxPosition)
					}
				}
			}
		}
	}

	return updatedPositions
}

func updatePriorityQueue(allPossibilities map[Position]IntSet, pq *PriorityQueue, updatedPositions []Position) {
	for i := 0; i < len(updatedPositions); i++ {
		possibilities := allPossibilities[updatedPositions[i]]
		item := &Item{
			position: updatedPositions[i],
			priority: possibilities.Size(),
		}
		// fmt.Println(item.position, item.priority)
		heap.Push(pq, item)
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
