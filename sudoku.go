/*
Go Sudoku solver with a stochastic simulated-annealing algorithm.

I wanted to challenge myself with something other than backtracking :P.
Implementation of the algorithm is influenced by Rhyd Lewis' paper:
(https://link.springer.com/article/10.1007/s10732-007-9012-8)

*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Returns a solved Sudoku puzzle. Also works with impossible sudoku puzzles,
// in which case would return the "best" solution with the least errors.
func SolveSudoku(board [][]int) [][]int {

	solutionFound := false
	newBoard := make([][]int, len(board))
	decreaseFactor := 0.99

	for numTries := 0; !solutionFound && numTries < 10; numTries++ {
		stuckCount := 0

		newBoard = make([][]int, len(board))
		for i := range board {
			newBoard[i] = make([]int, len(board[i]))
			copy(newBoard[i], board[i])
		}

		blocks := Create3x3Blocks()
		sigma := CalcInitSigma(newBoard, board, blocks)
		iters := CalcNumberOfIters(board)
		RandomlyFill3x3Blocks(newBoard, blocks)
		cost := float64(CalculateCost(newBoard))

		if cost <= 0 {
			solutionFound = true
		}

		for !solutionFound {
			prevCost := cost
			for i := 0; i < iters; i++ {
				deltaCost := ChooseNewState(newBoard, board, blocks, sigma)
				cost += deltaCost
				// fmt.Println(cost)
				if cost <= 0 {
					solutionFound = true
					return newBoard
				}
			}
			sigma *= decreaseFactor
			if cost <= 0 {
				solutionFound = true
				return newBoard
			}

			if cost >= prevCost {
				stuckCount++
			} else {
				stuckCount = 0
			}

			if stuckCount > 77 {
				break
			} else if stuckCount > 45 {
				sigma += 2
			}
		}
	}
	return newBoard
}

// Calculate the "cost" of the current board state (Number of errors).
// This is the function we try to minimize
func CalculateCost(board [][]int) int {
	numberOfErrors := 0
	for i := 0; i < 9; i++ {
		numberOfErrors += CalculateNumberOfErrorsRowColumn(i, i, board)
	}
	return numberOfErrors
}

// Helper function to calculate the errors in the (col, row) cross
func CalculateNumberOfErrorsRowColumn(row int, column int, sudoku [][]int) int {
	rowElem := make(map[int]int)
	colElem := make(map[int]int)
	numberOfErrors := 0
	for i := 0; i < 9; i++ {
		_, r := rowElem[sudoku[row][i]]
		_, c := colElem[sudoku[i][column]]
		if sudoku[row][i] != 0 {
			if r {
				numberOfErrors++
			} else {
				rowElem[sudoku[row][i]] += 1
			}
		}
		if sudoku[i][column] != 0 {
			if c {
				numberOfErrors++
			} else {
				colElem[sudoku[i][column]] += 1
			}
		}
	}
	return numberOfErrors
}

// Creates a slice of blocks, where each block holds 9 pairs of coordinates to
// the corresponding (col, row) location on the board.
func Create3x3Blocks() [][][2]int {
	blocks := [][][2]int{}
	for bCol := 0; bCol < 9; bCol += 3 {
		for bRow := 0; bRow < 9; bRow += 3 {
			block := [][2]int{}
			for y := 0; y < 3; y++ {
				for x := 0; x < 3; x++ {
					block = append(block, [2]int{bCol + x, bRow + y})
				}
			}
			blocks = append(blocks, block)
		}
	}
	return blocks
}

// Randomly fills a 3x3 block with valid numbers.
// Ensures each block has unique numbers from 1-9
func RandomlyFill3x3Blocks(board [][]int, blocks [][][2]int) {
	for _, block := range blocks {
		exists := []int{}
		for _, box := range block {
			curr := board[box[1]][box[0]]
			if curr != 0 {
				exists = append(exists, curr)
			}
		}

		choices := []int{}
		for i := 1; i <= 9; i++ {
			if !contains(exists, i) {
				choices = append(choices, i)
			}
		}

		for _, box := range block {
			curr := board[box[1]][box[0]]
			if curr == 0 && len(choices) > 0 {
				index := rand.Intn(len(choices))
				choice := choices[index]
				for choice == curr {
					index = rand.Intn(len(choices))
					choice = choices[index]
				}
				board[box[1]][box[0]] = choice
				choices = remove(choices, index)
			}
		}
	}
}

// Picks two random **different** boxes in a block, each box is a coordinate pair (col, row)
func PickTwoRandomBoxInBlock(sudoku [][]int, block [][2]int) [2][2]int {
	for true {
		firstBox := block[rand.Intn(len(block))]
		secondBox := block[rand.Intn(len(block))]

		if firstBox != secondBox && sudoku[firstBox[1]][firstBox[0]] == 0 && sudoku[secondBox[1]][secondBox[0]] == 0 {
			return [2][2]int{firstBox, secondBox}
		}
	}
	return [2][2]int{}
}

// Swaps the values of two boxes in place on a board
func FlipBoxes(board [][]int, blocks [2][2]int) {
	board[blocks[0][1]][blocks[0][0]], board[blocks[1][1]][blocks[1][0]] = board[blocks[1][1]][blocks[1][0]], board[blocks[0][1]][blocks[0][0]]
}

func ChooseNewState(board [][]int, ogBoard [][]int, blocks [][][2]int, s float64) float64 {
	block := blocks[rand.Intn(len(blocks))]

	flip := PickTwoRandomBoxInBlock(ogBoard, block)

	currentCost := CalculateNumberOfErrorsRowColumn(flip[0][1], flip[0][0], board) + CalculateNumberOfErrorsRowColumn(flip[1][1], flip[1][0], board)
	FlipBoxes(board, flip)
	newCost := CalculateNumberOfErrorsRowColumn(flip[0][1], flip[0][0], board) + CalculateNumberOfErrorsRowColumn(flip[1][1], flip[1][0], board)

	costDifference := float64(newCost - currentCost)
	rho := math.Exp(-costDifference / s)

	if rand.Float64() < rho {
		return costDifference
	} else {
		FlipBoxes(board, flip)
		return 0
	}
}

// Calculates number of iteration to run annealing process per cycle
func CalcNumberOfIters(ogBoard [][]int) int {
	iter := 0
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if ogBoard[y][x] != 0 {
				iter++
			}
		}
	}

	return iter
}

// Calculate initial sigma value based on the cost standard deviation of flips.
func CalcInitSigma(board [][]int, ogBoard [][]int, blocks [][][2]int) float64 {
	list := []int{}
	for i := 0; i < 9; i++ {
		block := blocks[rand.Intn(len(blocks))]
		flip := PickTwoRandomBoxInBlock(ogBoard, block)
		FlipBoxes(board, flip)
		list = append(list, CalculateCost(board))
		FlipBoxes(board, flip)
	}

	var sum, mean, sd float64
	for _, val := range list {
		sum += float64(val)
	}
	mean = sum / float64(len(list))

	for _, val := range list {
		sd += math.Pow(float64(val)-mean, 2)
	}

	sd = math.Sqrt(sd / float64(len(list)))

	return sd
}

/********** START: Utility Helper Functions   ************/

// Checks if an integer is in a int slice
func contains(arr []int, item int) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}
	return false
}

// Removes item from int slice, this WILL change the order of the slice
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

/********** END: Utility Helper Functions   ************/

func main() {
	// input := [][]int{
	// 	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	// 	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	// 	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	// 	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	// 	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	// 	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	// 	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	// 	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	// 	{0, 0, 0, 0, 8, 0, 0, 7, 9},
	// }
	input := [][]int{
		{0, 0, 9, 0, 0, 5, 0, 2, 0},
		{2, 4, 0, 7, 0, 0, 0, 0, 1},
		{0, 0, 6, 0, 4, 0, 0, 0, 0},
		{0, 6, 0, 0, 0, 0, 0, 0, 0},
		{4, 1, 0, 0, 3, 0, 0, 5, 0},
		{0, 0, 0, 9, 0, 0, 3, 0, 0},
		{0, 0, 2, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 8, 0, 0, 7},
		{5, 3, 0, 0, 9, 0, 0, 1, 0},
	}

	fmt.Println(SolveSudoku(input))
}
