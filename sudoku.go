package main

import (
	"container/list"
)

func dfs(board *[][]int, queue *list.List, rows *[9][9]int, columns *[9][9]int, groups *[9][9]int) bool {
	// if the length of the process queue is empty then all empty nodes
	// have been filled without ever violating the frequency constraint
	if queue.Len() == 0 {
		return true // return here, the current state of the board is the solution
	}

	// grab current node for processing
	elem := queue.Front()
	cur := elem.Value.([]int)
	x, y := cur[0], cur[1]

	for i := 0; i < 9; i++ {
		// iterate through 9 possible values for this node, determine if a solution exists
		if (*rows)[y][i] == 0 && (*columns)[x][i] == 0 && (*groups)[(y/3)*3+(x/3)][i] == 0 { // check for violations
			// add this number and remove it from the queue, then check the next node in queue
			queue.Remove(queue.Front())
			(*rows)[y][i] = 1
			(*columns)[x][i] = 1
			(*groups)[(y/3)*3+(x/3)][i] = 1
			(*board)[y][x] = i + 1

			if dfs(board, queue, rows, columns, groups) {
				return true
			} else {
				// the recursive call has not found a solution, so this value cannot work, move on
				(*rows)[y][i] = 0
				(*columns)[x][i] = 0
				(*groups)[(y/3)*3+(x/3)][i] = 0
				(*board)[y][x] = 0
				queue.PushFront([]int{x, y})
			}
		}
	}

	return false
}

func SolveSudoku(board [][]int) [][]int {
	// Fill in adjacency matrices mapping which numbers are
	// present in rows/column/groups

	rows := [9][9]int{}
	columns := [9][9]int{}
	groups := [9][9]int{}
	queue := list.New() // use Linked List for O(1) popping & left append

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if board[y][x] == 0 { // if a node is empty, mark it for processing
				queue.PushBack([]int{x, y})
			} else { // else document it in our matrices
				rows[y][board[y][x]-1] = 1
				columns[x][board[y][x]-1] = 1
				groups[(y/3)*3+(x/3)][board[y][x]-1] = 1
			}
		}
	}

	// call the dfs function to compute a solution of the board
	dfs(&board, queue, &rows, &columns, &groups)

	return board
}
