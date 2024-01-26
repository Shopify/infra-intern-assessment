package main

import (
	"container/heap"
)

func SudokuSolver(board [][]int) (bool, [][]int) {
	allPossibilities := getAllPossibilities(board)
	pq := initializePriorityQueue(allPossibilities)
	var position Position
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		position = item.position
		//If position is in possibilities, it needs to be assigned
		if possibilities, ok := allPossibilities[position]; ok {
			//If position has no possibilities, this isn't the correct solution and we backtrack
			if possibilities.Size() == 0 {
				return false, nil
			} else if possibilities.Size() == 1 {
				//Assign single possible value
				value := possibilities.Ints()[0]
				board[position.row][position.column] = value
				//Remove positions from possibilities
				delete(allPossibilities, position)
				updatedPositions := updatePositionPosibilities(allPossibilities, position, value)
				updatePriorityQueue(allPossibilities, &pq, updatedPositions)
			} else {
				break
			}
		}
	}

	if isBoardSolved(board) {
		return true, board
	}

	possibilities := allPossibilities[position]
	possibleValues := possibilities.Ints()

	for i := 0; i < len(possibleValues); i++ {
		boardCopy := deepCopy(board)
		boardCopy[position.row][position.column] = possibleValues[i]
		isSolved, solution := SudokuSolver(boardCopy)
		if isSolved {
			return true, solution
		}
	}
	return false, nil
}
