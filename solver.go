package main

import (
	"container/heap"
)

// Solves the sudoku using an optimized backtracking approach which assigns values at positions with the lowest number of possible values.
// Optimizations:
//   - Positions are tracked using a priority queue (min heap) based of number of possible values
//   - All positions with 1 possible value are assigned before solving with backtracking
//   - When a position is assigned, all positions affected by assignment (row, column and box) have their priorities updated for quicker access.
//   - Only begin recursive calls when the position with least number of possible values is more than 1.
//   - Recursive call invokes all optimizations again (At the cost of extra storage of board, possibilities and priority queue at each recursion)
func SudokuSolver(board [][]int) (bool, [][]int) {

	allPossibilities := getAllPossibilities(board)
	pq := initializePriorityQueue(allPossibilities)

	var position Position
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		position = item.position

		//If position is in possibilities, it needs to be assigned
		// This check is important as priority queue may have positions which have already been assigned but haven't be removed from queue from older priority.
		if possibilities, ok := allPossibilities[position]; ok {

			if possibilities.Size() == 0 {
				//If position has no possibilities, this isn't the correct solution and we backtrack
				return false, nil

			} else if possibilities.Size() == 1 {
				//Assign only possible value at position

				value := possibilities.Ints()[0]
				board[position.row][position.column] = value

				//Remove positions from possibilities
				delete(allPossibilities, position)

				//Update the possibilities map and Priority queue for positions in same row, column and box.
				updatePossibilitiesAndPriorityQueue(&allPossibilities, &pq, position, value)
			} else {
				//Break to use backtrack approach when there are more than one possible
				break
			}
		}
	}

	//Check and return solved board back up the recursion stack.
	if isBoardSolved(board) {
		return true, board
	}

	//Recursive check and bactrack over possible values at position with more than one value.
	possibilities := allPossibilities[position]
	possibleValues := possibilities.Ints()

	//TODO: Invoke SudokuSolver simultaneously with Goroutines to improve performance
	//REQUIREMENT: Check if code is concurrency safe and aggregate results from all channels.

	for i := 0; i < len(possibleValues); i++ {
		//Create a deep copy of board so it isn't affected by changes in recursive calls
		boardCopy := deepCopy(board)
		//Assign and check possible value
		boardCopy[position.row][position.column] = possibleValues[i]
		isSolved, solution := SudokuSolver(boardCopy)
		//Return solution up recursive stack
		if isSolved {
			return true, solution
		}
		//Backtrack and use next possible value
	}
	//If no value can be assigned then the current board is incorrect.
	return false, nil
}
