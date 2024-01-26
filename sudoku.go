package main

import (
	"math/big"
	"sort"
)

// Using big.Int as a bit vector since it isn't the worst
// performance wise and is availble through the standard library
// https://medium.com/@val_deleplace/7-ways-to-implement-a-bit-set-in-go-91650229b386
type SudokuNum struct {
	occurs            big.Int
	conflicts         big.Int
	possibleTemplates []big.Int
	symbol            uint8
}

func SolveSudoku(input [][]int) [][]int {
	var baseTemplate big.Int
	var freeGrid big.Int
	var templates []big.Int

	// Generate the 46656 possible sudoku templates
	generateSudokuTemplates(&baseTemplate, &freeGrid, &templates, 0, 0)

	sudokuVec := vectorize(input)
	inputNums := vecToBits(sudokuVec)

	for i := range inputNums {
		findValidTemplates(templates, &inputNums[i])
	}

	sort.Slice(inputNums, func(i, j int) bool {
		return len(inputNums[i].possibleTemplates) < len(inputNums[j].possibleTemplates)
	})

	freeGrid.SetInt64(0)
	dfsBacktrack(inputNums, &freeGrid, &sudokuVec, 0)

	return sudokurize(sudokuVec)
}

func dfsBacktrack(input []SudokuNum, freeGrid *big.Int, solvedVec *[]uint8, index int) bool {
	if index >= 9 {
		return true
	}

	num := input[index]

	for _, template := range num.possibleTemplates {
		t := new(big.Int).And(&template, freeGrid)

		if len(t.Bits()) == 0 {
			freeGrid.Xor(freeGrid, &template)
			rtn := dfsBacktrack(input, freeGrid, solvedVec, index+1)
			if rtn {
				insertBitsToVec(template, num.symbol, solvedVec)
				return true
			} else {
				freeGrid.Xor(freeGrid, &template)
			}
		}
	}

	return false
}

// https://stackoverflow.com/questions/64257065/is-there-another-way-of-testing-if-a-big-int-is-0
func findValidTemplates(templates []big.Int, input *SudokuNum) {
	for _, template := range templates {
		t1 := new(big.Int).And(&template, &input.conflicts)
		if len(t1.Bits()) == 0 {
			t2 := new(big.Int).And(&template, &input.occurs)
			if t2.Cmp(&input.occurs) == 0 {
				input.possibleTemplates = append(input.possibleTemplates, template)
			}
		}
	}
}

func vecToBits(vec []uint8) []SudokuNum {
	result := make([]SudokuNum, 9)

	for i := range result {
		result[i].symbol = uint8(i + 1)
	}

	for i, n := range vec {
		if n != 0 {
			for j := 0; j < 9; j++ {
				result[j].conflicts.SetBit(&result[j].conflicts, i, 1)
			}
			result[n-1].occurs.SetBit(&result[n-1].occurs, i, 1)
			result[n-1].conflicts.SetBit(&result[n-1].conflicts, i, 0)
		}
	}
	return result
}

func insertBitsToVec(bits big.Int, num uint8, input *[]uint8) {
	for i := 0; i < 81; i++ {
		if bits.Bit(i) == 1 {
			(*input)[i] = num
		}
	}
}

func vectorize(input [][]int) []uint8 {
	var result []uint8
	for _, arr := range input {
		for _, n := range arr {
			result = append(result, uint8(n))
		}
	}
	return result
}

func sudokurize(input []uint8) [][]int {
	var result [][]int
	for i := 0; i < 9; i++ {
		var row []int
		for j := 0; j < 9; j++ {
			row = append(row, int(input[i*9+j]))
		}
		result = append(result, row)
	}
	return result
}

// TODO: cache this
func generateSudokuTemplates(currGrid *big.Int, freeGrid *big.Int, templates *[]big.Int, row int, count int) int {
	if row >= 9 {
		newTemplate := new(big.Int).Set(currGrid)
		*templates = append(*templates, *newTemplate)
		return count + 1
	}

	for i := 0; i < 9; i++ {
		cellPos := row*9 + i
		if freeGrid.Bit(cellPos) != 0 {
			continue
		}
		prevFreeGrid := new(big.Int).Set(freeGrid)
		setCell(currGrid, freeGrid, cellPos)
		count = generateSudokuTemplates(currGrid, freeGrid, templates, row+1, count)
		currGrid.SetBit(currGrid, cellPos, 0)
		freeGrid.Set(prevFreeGrid)
	}

	return count
}

func setCell(grid *big.Int, freeGrid *big.Int, pos int) {
	// Set cell
	grid.SetBit(grid, pos, 1)

	// Set row conflicts
	row := pos / 9
	for i := 0; i < 9; i++ {
		freeGrid.SetBit(freeGrid, row*9+i, 1)
	}

	// Set column conflicts
	col := pos % 9
	for i := 0; i < 9; i++ {
		freeGrid.SetBit(freeGrid, i*9+col, 1)
	}

	// Set box conflicts
	boxRow := row / 3
	boxCol := col / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			freeGrid.SetBit(freeGrid, (boxRow*3+i)*9+(boxCol*3+j), 1)
		}
	}
}
