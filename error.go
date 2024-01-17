package main

import (
	"fmt"
)

const (
	ERR_INVALID_SUDOKU_SIZE = -000001 // sudoku size is invalid
	ERR_INVALID_SUDOKU      = -000002 // sudoku cannot be solved
)

func reportErr(errCode int) {
	var errInfo string
	switch errCode {
	case ERR_INVALID_SUDOKU_SIZE:
		errInfo = "Sudoku size is invalid"
	case ERR_INVALID_SUDOKU:
		errInfo = "Sudoku is unsolvable"
	}

	fmt.Printf("%v with Error code:%v", errInfo, errCode)
}
