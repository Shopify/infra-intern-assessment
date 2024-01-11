package main

import (
	"reflect"
	"testing"
)

func TestNextUnfilled(t *testing.T) {
	runTest := func(testName string, input [][]int, startPos Position, expectedPos Position, expectedError bool) {
		t.Run(testName, func(t *testing.T) {
			solved, err := nextUnfilled(input, startPos)
			if (err != nil) != expectedError || (err == nil && solved != expectedPos) {
				t.Errorf("%s: Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", testName, expectedPos, solved)
			}
		})
	}

	input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 9},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	runTest("Typical use case", input, Position{0, 0}, Position{0, 2}, false)
	runTest("Starting on an unfilled tile", input, Position{0, 3}, Position{0, 3}, false)
	runTest("Next unfilled tile is on the next line", input, Position{0, 8}, Position{1, 1}, false)
	runTest("No next tile, should error", input, Position{8, 8}, Position{0, 0}, true)
}

func TestIsValuePlacable(t *testing.T) {
	runTest := func(testName string, board [][]int, testingPos Position, value int, expected bool) {
		t.Run(testName, func(t *testing.T) {
			if result := isValuePlaceable(board, testingPos, value); result != expected {
				t.Errorf("%s: Expected placement to be %t, but got %t.", testName, expected, result)
			}
		})
	}
	board := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 6, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
	}
	runTest("Invalid placement in the same column", board, Position{3, 1}, 1, true)
	runTest("Invalid placement in the same row", board, Position{0, 3}, 1, false)
	runTest("Invalid placement in the same small square", board, Position{2, 1}, 1, false)
	runTest("valid placement", board, Position{7, 8}, 1, true)
}

type BoardTest struct {
	input    [][]int
	expected [][]int
}

func TestSolveSudoku(t *testing.T) {
	boards := []BoardTest{
		{[][]int{
			{5, 3, 0, 0, 7, 0, 0, 0, 0},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 6, 0},
			{8, 0, 0, 0, 6, 0, 0, 0, 3},
			{4, 0, 0, 8, 0, 3, 0, 0, 1},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 6, 0, 0, 0, 0, 2, 8, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 5},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		},
			[][]int{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9},
			}},
	}
	for _, test := range boards {
		solved := SolveSudoku(test.input)
		if !reflect.DeepEqual(solved, test.expected) {
			t.Errorf("Sudoku puzzle was not solved correctly. Expected:\n%v\n\nGot:\n%v", test.expected, solved)
		}
	}

}
