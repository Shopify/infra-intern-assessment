package main

import (
	"errors"
	"slices"
)

const MATRIX_N = 9
const EMPTY_VALUE = 0

var ErrorInvalidMatrix = errors.New("invalid matrix length")
var ErrorUnsolvableMatrix = errors.New("unsolvable matrix")

// SolveSudoku recursively solves the matrix input and returns the solved matrix with
// any errors that may have occured.
func SolveSudoku(matrix [][]int) ([][]int, error) {
	if len(matrix) != MATRIX_N {
		return matrix, ErrorInvalidMatrix
	}
	for _, row := range matrix {
		if len(row) != MATRIX_N {
			return matrix, ErrorInvalidMatrix
		}
	}

	solved_matrix, is_solved := solve(matrix)
	if is_solved {
		return solved_matrix, nil
	} else {
		return matrix, ErrorUnsolvableMatrix
	}
}

type placement struct {
	i   int
	j   int
	val int
}

// We use bactracking in order to exhaust all possible values (brute force).
func solve(matrix [][]int) ([][]int, bool) {
	for i := 0; i < MATRIX_N; i++ {
		for j := 0; j < MATRIX_N; j++ {
			if matrix[i][j] == EMPTY_VALUE {
				for val := 1; val <= MATRIX_N; val++ {
					if checkPlacement(matrix, placement{i: i, j: j, val: val}) {
						matrix[i][j] = val
						if solved_matrix, is_solved := solve(matrix); is_solved {
							return solved_matrix, is_solved
						} else {
							matrix[i][j] = EMPTY_VALUE
						}
					}
				}

				// exhausted all possible values
				return matrix, false
			}
		}
	}
	// no empty values
	return matrix, true
}

const SQUARE_LEN = 3

func checkPlacement(matrix [][]int, p placement) bool {
	inSquare := checkPlacementInSquare(matrix, p)
	inCol := checkPlacementInCol(matrix, p)
	inRow := checkPlacementInRow(matrix, p)
	return inSquare && inCol && inRow
}

func checkPlacementInSquare(matrix [][]int, p placement) bool {
	square := make([]int, 0)
	vertical_offset := SQUARE_LEN * (p.i / SQUARE_LEN)
	horizontal_offset := SQUARE_LEN * (p.j / SQUARE_LEN)

	for i := 0; i < SQUARE_LEN; i++ {
		for j := 0; j < SQUARE_LEN; j++ {
			square = append(square, matrix[i+vertical_offset][j+horizontal_offset])
		}
	}
	return !slices.Contains(square, p.val)
}

func checkPlacementInCol(matrix [][]int, p placement) bool {
	column := make([]int, 0)
	for _, row := range matrix {
		column = append(column, row[p.j])
	}
	return !slices.Contains(column, p.val)
}

func checkPlacementInRow(matrix [][]int, p placement) bool {
	row := matrix[p.i]
	return !slices.Contains(row, p.val)
}
