package main

import (
	"reflect"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
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

	expected := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}

func TestSolveSudoku2(t *testing.T) {
	input := [][]int{
		{0, 2, 0, 6, 0, 8, 0, 0, 0},
		{5, 8, 0, 0, 0, 9, 7, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{3, 7, 0, 0, 0, 0, 5, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 8, 0, 0, 0, 0, 1, 3},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 9, 8, 0, 0, 0, 3, 6},
		{0, 0, 0, 3, 0, 6, 0, 9, 0},
	}

	expected := [][]int{
		{1, 2, 3, 6, 7, 8, 9, 4, 5},
		{5, 8, 4, 2, 3, 9, 7, 6, 1},
		{9, 6, 7, 1, 4, 5, 3, 2, 8},
		{3, 7, 2, 4, 6, 1, 5, 8, 9},
		{6, 9, 1, 5, 8, 3, 2, 7, 4},
		{4, 5, 8, 7, 9, 2, 6, 1, 3},
		{8, 3, 6, 9, 2, 4, 1, 5, 7},
		{2, 1, 9, 8, 5, 7, 4, 3, 6},
		{7, 4, 5, 3, 1, 6, 8, 9, 2},
	}

	solved := SolveSudoku(input)

	if !reflect.DeepEqual(solved, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, solved)
	}
}

func TestSolveSudokuInvalid(t *testing.T) {
	input := [][]int{
		{5, 3, 3, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solved := SolveSudoku(input)

	if solved != nil {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", nil, solved)
	}
}

func TestSolveSudokuInputMutation(t *testing.T) {
	input := [][]int{
		{0, 2, 0, 6, 0, 8, 0, 0, 0},
		{5, 8, 0, 0, 0, 9, 7, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{3, 7, 0, 0, 0, 0, 5, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 8, 0, 0, 0, 0, 1, 3},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 9, 8, 0, 0, 0, 3, 6},
		{0, 0, 0, 3, 0, 6, 0, 9, 0},
	}

	expected := [][]int{
		{0, 2, 0, 6, 0, 8, 0, 0, 0},
		{5, 8, 0, 0, 0, 9, 7, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{3, 7, 0, 0, 0, 0, 5, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 8, 0, 0, 0, 0, 1, 3},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 9, 8, 0, 0, 0, 3, 6},
		{0, 0, 0, 3, 0, 6, 0, 9, 0},
	}

	SolveSudoku(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", expected, input)
	}
}

func TestSolveSudokuInvalidGrid(t *testing.T) {
	input := [][]int{
		{0, 2, 0, 6, 0, 8, 0, 0, 0},
		{5, 8, 0, 0, 0, 9, 7, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{3, 7, 0, 0, 0, 0, 5, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 4},
		{0, 0, 8, 0, 0, 0, 0, 1, 3},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 9, 8, 0, 0, 0, 3, 6},
		{0, 0, 0, 3, 0, 6, 1, 0, 9, 0},
	}

	solved := SolveSudoku(input)

	if solved != nil {
		t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", nil, input)
	}
}
