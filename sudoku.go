package main

const SIZE = 9

func SolveSudoku(in [][]int) [][]int {
	workingCopy := copySlice(in)
	workingCells := transformToCells(workingCopy)

	prunePossibleSol(workingCells)
	allMissingCells := []cell{}

	for i := range workingCells {
		for _, c := range workingCells[i] {
			if c.value == 0 {
				allMissingCells = append(allMissingCells, c)
			}
		}
	}
	solver(workingCopy, allMissingCells, 0)
	return workingCopy
}


func prunePossibleSol(board [][]cell) {
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

func solver(workingCells [][]int, missingCell []cell, current int) bool {
	if current >= len(missingCell) {
		return true
	}
	c := &missingCell[current]

	s := c.PossibleSol()
	for solCount := len(s) - 1; solCount >= 0; solCount-- {
		c.value = s[solCount]
		workingCells[c.i][c.j] = s[solCount]

		if !changeIsValid(workingCells, *c) {
			continue
		}

		if solver(workingCells, missingCell, current+1) {
			return true
		}
	}
	workingCells[c.i][c.j] = 0
	return false
}

func changeIsValid(workingCells [][]int, change cell) bool {
	return rowIsValid(workingCells, change) && columnIsValid(workingCells, change) && blockIsValid(workingCells, change)
}

func rowIsValid(workingCells [][]int, change cell) bool {
	for j := range workingCells {
		if change.j == j {
			continue
		}
		if change.value == workingCells[change.i][j] {
			return false
		}
	}
	return true
}

func columnIsValid(workingCells [][]int, change cell) bool {
	for i := range workingCells[change.j] {
		if change.i == i {
			continue
		}
		if change.value == workingCells[i][change.j] {
			return false
		}
	}
	return true
}

func blockIsValid(workingCells [][]int, change cell) bool {
	// Divide since both type is int, its floor division
	blockStartRow := (change.i / 3) * 3
	blockStartColumn := (change.j / 3) * 3
	for i := blockStartRow; i < blockStartRow+3; i++ {
		for j := blockStartColumn; j < blockStartColumn+3; j++ {
			if change.j == j && change.i == i {
				continue
			}
			if workingCells[i][j] == change.value {
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
