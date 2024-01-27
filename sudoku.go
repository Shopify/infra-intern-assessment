package main

import (
	"embed"
	"math/big"
	"sync"
)

//go:embed templates.txt
var templateFile embed.FS

// SudokuNum is a struct that represents the positions of a single
// digit within Sudoku puzzle. Thus, there are 9 SudokuNum's for each digit
// It uses big.Int as a bit vector representation of a Sudoku grid
// It stores the following information:
type SudokuNum struct {
	occurs            big.Int
	conflicts         big.Int
	possibleTemplates []big.Int
	symbol            uint8
}

func SolveSudoku(input [][]int) [][]int {
	var wg sync.WaitGroup
	var freeGrid big.Int
	var templates []big.Int

	// ~13000000 ns/op faster to read from file
	templates, err := ReadTemplatesFromEmbed(templateFile)
	if err != nil {
		panic(err)
	}
	// Uncomment the following to generate the 46656 possible templates instead
	/*
		// var baseTemplate big.Int
		// GenerateSudokuTemplates(&baseTemplate, &freeGrid, &templates, 0, 0)
		// SaveTemplatesToFile(templates, "templates.txt")
		// freeGrid.SetInt64(0)
	*/

	sudokuVec := vectorize(input)
	inputNums := vecToBits(sudokuVec)

	// - Since every goroutine is exclusively responsible
	//   for modifying one SudokuNum, atomic access is not necessary
	// - Parallelizing with Goroutines saves ~5000000 ns/op tested on M1-pro
	wg.Add(len(inputNums))
	for i := range inputNums {
		go findValidTemplates(templates, &inputNums[i], &wg)
	}
	wg.Wait()

	dfsBacktrack(inputNums, &freeGrid, &sudokuVec, 0)

	return sudokurize(sudokuVec)
}

func dfsBacktrack(input []SudokuNum, freeGrid *big.Int, solvedVec *[]uint8, index int) bool {
	if index >= 9 {
		return true
	}

	for _, template := range input[index].possibleTemplates {
		t := new(big.Int).And(freeGrid, &template)

		if len(t.Bits()) == 0 {
			freeGrid.Xor(freeGrid, &template)
			found := dfsBacktrack(input, freeGrid, solvedVec, index+1)
			if found {
				insertBitsToVec(template, input[index].symbol, solvedVec)
				return true
			} else {
				freeGrid.Xor(freeGrid, &template)
			}
		}
	}

	return false
}

func findValidTemplates(templates []big.Int, input *SudokuNum, wg *sync.WaitGroup) {
	defer wg.Done()

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
