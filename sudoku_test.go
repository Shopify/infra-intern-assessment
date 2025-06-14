package main

import (
    "reflect"
    "testing"
)

// TestSolveSudoku tests the SolveSudoku function with different scenarios.
func TestSolveSudoku(t *testing.T) {
    tests := []struct {
        name     string
        input    [][]int
        expected [][]int
        valid    bool // Indicates if the puzzle is solvable
    }{
        {
            name: "Standard Puzzle",
            input: [][]int{
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
            expected: [][]int{
                {5, 3, 4, 6, 7, 8, 9, 1, 2},
                {6, 7, 2, 1, 9, 5, 3, 4, 8},
                {1, 9, 8, 3, 4, 2, 5, 6, 7},
                {8, 5, 9, 7, 6, 1, 4, 2, 3},
                {4, 2, 6, 8, 5, 3, 7, 9, 1},
                {7, 1, 3, 9, 2, 4, 8, 5, 6},
                {9, 6, 1, 5, 3, 7, 2, 8, 4},
                {2, 8, 7, 4, 1, 9, 6, 3, 5},
                {3, 4, 5, 2, 8, 6, 1, 7, 9},
            },
            valid: true,
        },
        {
            name: "Almost Solved Puzzle",
            input: [][]int{
                {1, 2, 3, 4, 5, 6, 7, 8, 9},
                {4, 5, 6, 7, 8, 9, 1, 2, 3},
                {7, 8, 9, 1, 2, 3, 4, 5, 6},
                {2, 1, 4, 3, 6, 5, 8, 9, 7},
                {3, 6, 5, 8, 9, 7, 2, 1, 4},
                {8, 9, 7, 2, 1, 4, 5, 6, 0}, // One cell (last one) is empty
                {5, 3, 1, 6, 4, 2, 9, 7, 8},
                {6, 4, 2, 9, 7, 8, 3, 0, 5}, // One cell (8th one) is empty
                {9, 7, 8, 5, 3, 1, 6, 4, 2},
            },
            expected: [][]int{
                {1, 2, 3, 4, 5, 6, 7, 8, 9},
                {4, 5, 6, 7, 8, 9, 1, 2, 3},
                {7, 8, 9, 1, 2, 3, 4, 5, 6},
                {2, 1, 4, 3, 6, 5, 8, 9, 7},
                {3, 6, 5, 8, 9, 7, 2, 1, 4},
                {8, 9, 7, 2, 1, 4, 5, 6, 3},
                {5, 3, 1, 6, 4, 2, 9, 7, 8},
                {6, 4, 2, 9, 7, 8, 3, 5, 1},
                {9, 7, 8, 5, 3, 1, 6, 4, 2},
            },
            valid: true,
        },
        {
            name: "Invalid Puzzle",
            input: [][]int{
                {1, 2, 3, 4, 5, 6, 7, 8, 9},
                {1, 2, 3, 4, 5, 6, 7, 8, 9}, // Duplicate row
                {7, 8, 9, 1, 2, 3, 4, 5, 6},
                {2, 1, 4, 3, 6, 5, 8, 9, 7},
                {3, 6, 5, 8, 9, 7, 2, 1, 4},
                {8, 9, 7, 2, 1, 4, 5, 6, 3},
                {5, 3, 1, 6, 4, 2, 9, 7, 8},
                {6, 4, 2, 9, 7, 8, 3, 5, 1},
                {9, 7, 8, 5, 3, 1, 6, 4, 2},
            },
            valid: false, // This puzzle should not be solvable
        },
        // Additional test cases can be added here
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            solved := SolveSudoku(tt.input)
            if tt.valid && !reflect.DeepEqual(tt.input, tt.expected) {
                t.Errorf("%s: Sudoku puzzle was not solved correctly.\nExpected:\n%v\nGot:\n%v", tt.name, tt.expected, tt.input)
            }
            if !tt.valid && solved {
                t.Errorf("%s: Invalid puzzle was incorrectly marked as solvable.", tt.name)
            }
        })
    }
}
