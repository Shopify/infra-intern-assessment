package main

import (
	"embed"
	"math/big"
	"sync"
)

//go:embed templates.txt
var templateFile embed.FS

// SudokuNum is a struct that represents the positions of a single digit
// within a Sudoku puzzle. Thus, there are 9 SudokuNum's, one for each digit.
// It uses big.Int as a bit vector representation of a Sudoku grid.
type SudokuNum struct {
	occurs            big.Int
	conflicts         big.Int
	possibleTemplates []big.Int
	symbol            uint8
}

// SolveSudoku solves a Sudoku puzzle by using templating, backtracking,
// and a little bit of parallelization. Can panic if the input is invalid.
func SolveSudoku(input [][]int) [][]int {
	var wg sync.WaitGroup
	var freeGrid big.Int
	var templates []big.Int

	// ~25000000 ns/op faster to read from file (on average)
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
	if len(sudokuVec) != 81 {
		panic("Invalid Sudoku puzzle")
	}
	inputNums := vecToBits(sudokuVec)

	// - Since every goroutine is exclusively responsible
	//   for modifying one SudokuNum, atomic access is not necessary
	// - Parallelizing with Goroutines saves ~5000000 ns/op tested on M1-pro
	wg.Add(len(inputNums))
	for i := range inputNums {
		go findValidTemplates(templates, &inputNums[i], &wg)
	}
	wg.Wait()

	valid := dfsBacktrack(inputNums, &freeGrid, sudokuVec, 0)
	if !valid {
		panic("Invalid Sudoku puzzle")
	}

	return sudokurize(sudokuVec)
}

// dfsBacktrack is a recursive function that uses backtracking to solve a Sudoku puzzle by
// trying every possible template from a filtered set for each digit. It returns true
// upon finding the first set of valid templates (i.e. a solution). It may be modified
// to check if a Sudoku puzzle only has one solution by creating >1 set of templates for a given grid.
func dfsBacktrack(input []SudokuNum, freeGrid *big.Int, solvedVec []uint8, index int) bool {
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

// findValidTemplates is a helper function for SolveSudoku that filters out invalid templates
// for a given digit (i.e. SudokuNum). It uses bitwise operations to encourage vectorization.
// It is parallelized by using goroutines and blocking until all goroutines in the WaitGroup are finished.
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

// Below are helper functions for converting from 2d representation of Sudoku puzzle to a bit vector

// vectorize returns a slice of uint8's which represents a 1d vector representation of a Sudoku puzzle.
// It takes a 2d representation of a Sudoku puzzle and converts it to a 1d vector representation.
func vectorize(input [][]int) []uint8 {
	var result []uint8
	for _, arr := range input {
		for _, n := range arr {
			result = append(result, uint8(n))
		}
	}
	return result
}

// vecToBits returns a slice of 9 SudokuNum's. one for each digit of a 9x9 Sudoku puzzle.
// It coverts a 1d vector representation of a Sudoku puzzle to a bit vector representation.
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

// insertBitsToVec inserts the positions of the 1's in bits into their
// respective position of the input slice representation of a Sudoku puzzle.
func insertBitsToVec(bits big.Int, num uint8, input []uint8) {
	for i := 0; i < 81; i++ {
		if bits.Bit(i) == 1 {
			input[i] = num
		}
	}
}

// sudokurize returns a 2d array representation of a Sudoku puzzle.
// Converts an input vector representation of a Sudoku puzzle to a 2d representation.
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
