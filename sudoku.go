package main

const SIZE = 9

func SolveSudoku(in [][]int) [][]int {
	workingCopy := copySlice(in)
	workingCells := toCells(workingCopy)

	prune(workingCells)
	allMissingCells := []cell{}

	for i := range workingCells {
		for _, c := range workingCells[i] {
			if c.value == 0 {
				allMissingCells = append(allMissingCells, c)
			}
		}
	}
	search(workingCopy, allMissingCells, 0)
	return workingCopy
}

func search(workingCells [][]int, missingCell []cell, current int) bool {
	if current >= len(missingCell) {
		return true
	}

	c := &missingCell[current]
	sol := c.PossibleSol()
	for count := len(sol) - 1; count >= 0; count-- {
		c.value = sol[count]
		workingCells[c.i][c.j] = sol[count]

		if !changeIsValid(workingCells, *c) {
			continue
		}

		if search(workingCells, missingCell, current+1) {
			return true
		}
	}
	workingCells[c.i][c.j] = 0
	return false
}


func prune(board [][]cell) {
	var rowValues, columnValues []int
	var cRow, cColumn cell
	var rowMissingCell, columnMissingCell []cell
	for i := range board {
		rowValues = []int{}
		columnValues = []int{}
		rowMissingCell = []cell{}
		columnMissingCell = []cell{}

		for j := range board[i] {
			cRow = board[i][j]

			if cRow.value != 0 {
				rowValues = append(rowValues, cRow.value)
			} else {
				rowMissingCell = append(rowMissingCell, cRow)
			}

			cColumn = board[j][i]
			if cColumn.value != 0 {
				columnValues = append(columnValues, cColumn.value)
			} else {
				columnMissingCell = append(columnMissingCell, cColumn)
			}
		}

		for _, v := range rowMissingCell {
			v.removeSolSlice(rowValues)
		}

		for _, v := range columnMissingCell {
			v.removeSolSlice(columnValues)
		}
	}
}


func changeIsValid(workingCells [][]int, changed cell) bool {
	return rowIsValid(workingCells, changed) && columnIsValid(workingCells, changed) && blockIsValid(workingCells, changed)
}

func rowIsValid(workingCells [][]int, changed cell) bool {
	for j := range workingCells {
		if changed.j == j {
			continue
		}
		if changed.value == workingCells[changed.i][j] {
			return false
		}
	}
	return true
}

func columnIsValid(workingCells [][]int, changed cell) bool {
	for i := range workingCells[changed.j] {
		if changed.i == i {
			continue
		}
		if changed.value == workingCells[i][changed.j] {
			return false
		}
	}
	return true
}

func blockIsValid(workingCells [][]int, changed cell) bool {
	// Divide since both type is int, its floor division
	blockStartRow := (changed.i / 3) * 3
	blockStartColumn := (changed.j / 3) * 3
	for i := blockStartRow; i < blockStartRow+3; i++ {
		for j := blockStartColumn; j < blockStartColumn+3; j++ {
			if changed.j == j && changed.i == i {
				continue
			}
			if workingCells[i][j] == changed.value {
				return false
			}
		}
	}
	return true
}


func copySlice[T any](src [][]T) (dest [][]T) {
	dest = make([][]T, len(src))
	for i := range src {
		dest[i] = make([]T, len(src[i]))
		copy(dest[i], src[i])
	}

	return dest
}
