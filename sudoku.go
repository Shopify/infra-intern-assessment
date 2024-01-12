package main

import (
	"fmt"
)

/**
Do not return anything, modify board in-place instead.

The first thing to do is to maintain 9 cols plus 9 rows set so 18
Then maintain 9 more sets for all blocks

As you traverse, you also want to maintain a row of coordinates you need to fill

Then start filling out the solution simply by doing this

use dfs to fill out the solution, and when you backtrack dlete.
if you manage to get to the end of the array, return,
*/

func SolveSudoku(board [][]int) [][]int {
	// Create sets as needed
	rowSet := make([]map[int]bool, 9)
	colSet := make([]map[int]bool, 9)
	blockSet := make(map[string]map[int]bool)
	var toTraverse [][]int

	for i := 0; i < 9; i++ {
		rowSet[i] = make(map[int]bool)
		colSet[i] = make(map[int]bool)
	}

	indexes := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
	for _, rowBlock := range indexes {
		for _, colBlock := range indexes {
			key := fmt.Sprintf("%d-%d", rowBlock[0], colBlock[0])
			if _, exists := blockSet[key]; !exists {
				blockSet[key] = createBlockSetInt(rowBlock, colBlock, board)
			}
		}
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			number := board[row][col]
			if number == 0 {
				toTraverse = append(toTraverse, []int{row, col})
			} else {
				rowSet[row][number] = true
				colSet[col][number] = true
			}
		}
	}

	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index >= len(toTraverse) {
			return true
		}

		row, col := toTraverse[index][0], toTraverse[index][1]
		key := getBoxHash(row, col)
		sets := []map[int]bool{rowSet[row], colSet[col], blockSet[key]}

		for i := 1; i <= 9; i++ {
			if !sets[0][i] && !sets[1][i] && !sets[2][i] {
				for _, set := range sets {
					set[i] = true
				}
				board[row][col] = i
				if dfs(index + 1) {
					return true
				} else {
					board[row][col] = 0
					for _, set := range sets {
						delete(set, i)
					}
				}
			}
		}
		return false
	}

	dfs(0)
	// print solution
	//fmt.Println(board)
	return board
}

func getBoxHash(r, c int) string {
	var row, col int

	if r >= 6 {
		row = 6
	} else if r >= 3 {
		row = 3
	} else {
		row = 0
	}

	if c >= 6 {
		col = 6
	} else if c >= 3 {
		col = 3
	} else {
		col = 0
	}

	return fmt.Sprintf("%d-%d", row, col)
}

func createBlockSetInt(rowIndexes, colIndexes []int, board [][]int) map[int]bool {
	result := make(map[int]bool)
	for _, r := range rowIndexes {
		for _, c := range colIndexes {
			if board[r][c] != 0 {
				result[board[r][c]] = true
			}
		}
	}
	return result
}
