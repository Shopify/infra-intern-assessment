package main

import (
	"container/heap"
)

func SudokuSolver(board [][]int) (bool, [][]int) {
	allPossibilities := getAllPossibilities(board)
	// for position := range allPossibilities {
	// 	possibilities := allPossibilities[position]
	// 	fmt.Println(position, possibilities.String())
	// }
	pq := initializePriorityQueue(allPossibilities)
	// for pq.Len() > 0 {
	// 	item := heap.Pop(&pq).(*Item)
	// 	fmt.Printf("%.2d:%d,%d\n", item.priority, item.position.row, item.position.column)
	// }

	pq_ctr := 0
	var position Position
	// fmt.Println(board)
	for pq_ctr < pq.Len() {
		pq_ctr++
		item := heap.Pop(&pq).(*Item)
		position = item.position
		// fmt.Println(position)
		//If position is in possibilities, it needs to be assigned
		if possibilities, ok := allPossibilities[position]; ok {
			// fmt.Println(possibilities.Len())
			//If position has no possibilities, this isn't the correct solution and we backtrack
			if possibilities.Size() == 0 {
				return false, nil
			} else if possibilities.Size() == 1 {
				//Assign single possible value
				value := possibilities.Ints()[0]
				board[position.row][position.column] = value
				//Remove positions from possibilities
				// printAllPositionPossibilities(&allPossibilities)
				// fmt.Println("Delete")
				delete(allPossibilities, position)
				// printAllPositionPossibilities(&allPossibilities)
				// fmt.Println("Update Possibilities")
				updatedPositions := updatePositionPosibilities(allPossibilities, position, value)
				// fmt.Println(updatedPosition)
				// fmt.Println(pq.Len())
				// printAllPositionPossibilities(&allPossibilities)
				// fmt.Println("Update PQ")
				updatePriorityQueue(allPossibilities, &pq, updatedPositions)
				// fmt.Println(pq.Len())
				// printAllPositionPossibilities(&allPossibilities)
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
	// fmt.Println(position, possibleValues)

	for i := 0; i < len(possibleValues); i++ {
		boardCopy := deepCopy(board)
		boardCopy[position.row][position.column] = possibleValues[i]
		isSolved, solution := SudokuSolver(boardCopy)
		if isSolved {
			return true, solution
		}
	}
	// fmt.Println("Updated single values")
	// fmt.Println(board)
	return false, nil
}
