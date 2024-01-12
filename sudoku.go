package main

import "fmt"

type cell struct {
	i int
	j int
	value int
	shiftedPossibleSol []bool
}

func (c *cell) addSol(v int) {
	c.shiftedPossibleSol[v - 1] = true
}

func (c *cell) removeSol(v int) {
	c.shiftedPossibleSol[v - 1] = false
}

func (c *cell) removeSolSlice(v []int) {
	for _, v := range v{
		c.removeSol(v)
	}
}

func (c *cell) tryOneSolution() int {
	for i, v := range c.shiftedPossibleSol {
		if v {
			return i + 1
		}
	}
	return -1
}

func (c *cell) possibleSolution() []int {
	r := []int{}
	for i, v := range c.shiftedPossibleSol {
		if v {
			r = append(r, i + 1)
		}
	}
	return r
}

func newCell(i int, j int, val int) cell {
	set := make([]bool, 9)

	for i := range set {
		set[i] = true
	}

	return cell{
		i: i,
		j: j,
		value: val,
		shiftedPossibleSol: set,
	}
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

func solver(workingBoard [][]cell, missingCell []cell, current int) bool {
	if current >= len(missingCell) {
		return true
	}
	c := &missingCell[current]

	s := c.possibleSolution()
	for solCount := len(s) - 1; solCount >= 0; solCount-- {
		c.value = s[solCount]

		if !changeIsValid(workingBoard, *c) {
			continue
		}

		if solver(workingBoard, missingCell , current + 1) {
			return true
		}
	}
	return false
}

func solve(in [][]int) {
	workingBoard := transformToCells(in)
	prunePossibleSol(workingBoard)
	allMissingCells := []cell{}

	var a *cell
	for i := range workingBoard {
		for j := range workingBoard[i] {
			a = &workingBoard[i][j]
			if a.value == 0 {
				allMissingCells = append(allMissingCells, *a)
			}
		}
	}
	fmt.Println(allMissingCells)
	if solver(workingBoard, allMissingCells, 0) {
		for i := range workingBoard {
			for j, c := range workingBoard[i] {
				in[i][j] = c.value
			}
		}
	}
}

func changeIsValid(workingBoard [][]cell, change cell) bool {
	return rowIsValid(workingBoard, change) && columnIsValid(workingBoard, change) && blockIsValid(workingBoard, change)
}

func rowIsValid(workingBoard [][]cell, change cell) bool {
	for j := range workingBoard {
		if change.value == workingBoard[change.i][j].value {
			return false
		}
	}
	return true
}

func columnIsValid(workingBoard [][]cell, change cell) bool {
	for i := range workingBoard[change.j] {
		if change.value == workingBoard[i][change.j].value {
			return false
		}
	}
	return true
}

func blockIsValid(workingBoard[][]cell, change cell) bool {
	for i := change.i / 3; i < change.i/3 + 3; i++ {
		for j := change.j / 3; j < change.j/3 + 3; j++ {
			if workingBoard[i][j].value == change.value {
				return false
			}
		}
	}
	return true
}


func transformToCells(in [][]int) [][]cell {
	board := make([][]cell, len(in))
	for i := range in {
		for j := range in[i] {
			board[i] = append(board[i], newCell(i, j, in[i][j]))
		}
	}
	return board
}



func copySlice[T any](src [][]T) (dest [][]T) {
	dest = make([][]T, len(src))
	for i := range src{
		dest[i] = make([]T, len(src[i]))
		copy(dest[i], src[i])
	}

	return dest
}

func main() {
	input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solve(input)
	fmt.Println(input)
}
