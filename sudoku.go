package main

import (
	"math/big"
)

// Using big.Int as a bit vector since it isn't the worst
// performance wise and is availble through the standard library
// https://medium.com/@val_deleplace/7-ways-to-implement-a-bit-set-in-go-91650229b386
type SudokuNum struct {
	occurs    big.Int
	conflicts big.Int
	templates []big.Int
}

func vecToBits(vec []int8) []SudokuNum {
	result := make([]SudokuNum, 9)
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

func insertBitsToVec(bits big.Int, num int8, input []int8) {
	for i := 0; i < 81; i++ {
		if bits.Bit(i) == 1 {
			input[i] = num
		}
	}
}

func vectorize(input [][]int) []int8 {
	var result []int8
	for _, arr := range input {
		for _, n := range arr {
			result = append(result, int8(n))
		}
	}
	return result
}

func sudokurize(input []int8) [][]int {
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

func SolveSudoku(input [][]int) [][]int {

	return input
}

// func printBits(input []SudokuNum) {
// 	for i, n := range input {
// 		fmt.Printf("\n%d: %b\n   %b\n", i+1, &n.occurs, &n.conflicts)
// 		t := make([]int8, 81)
// 		insertBitsToVec(n.occurs, 9, t)
// 		fmt.Println(t)
// 	}
// }
