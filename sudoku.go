package main

const BoardSize = 9
const BoxSize = 3

// State
// 9 x 9 array of length 9 bitsets
// Each bit represents whether (index + 1) is still in the domain (list of possible numbers) for that sudoku cell

// ex: Number               123456789
//     BoardDomains[0][0] = 010101010
//     Domain                2 4 6 8

// Bitset representing whether a number is still possible for that cell
type CellDomain uint16

// Array of CellDomains representing the entire board
var BoardDomains [BoardSize][BoardSize]CellDomain

// allowNumber adds k to the cell's domain
func allowNumber(c CellDomain, k int) CellDomain {
	return c | (1 << (k - 1))
}

// disallowNumber removes k from the cell's domain
func disallowNumber(c CellDomain, k int) CellDomain {
	return c & ^(1 << (k - 1))
}

// allowAllNumbers allows every possible number in the cell's domain
func allowAllNumbers() CellDomain {
	return (1 << BoardSize) - 1
}

// disallowAllNumbers removes every number from the cell's domain
func disallowAllNumbers() CellDomain {
	return 0
}

// isAllowedNumber reports whether the number is allowed in the cell's domain
func isAllowedNumber(c CellDomain, k int) bool {
	return (c & (1 << (k - 1))) != 0
}

// getAllowedNumbers returns every number still allowed in the cell's domain
func getAllowedNumbers(c CellDomain) []int {
	allowedNumbers := make([]int, 0, BoardSize)

	for i := 1; i < BoardSize+1; i++ {
		if isAllowedNumber(c, i) {
			allowedNumbers = append(allowedNumbers, i)
		}
	}
	return allowedNumbers
}

// initBoard initializes the BoardDomains into bitsets representing the remaining domain (ie: possible numbers) per cell
func initBoard(input [][]int) {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			BoardDomains[i][j] = allowAllNumbers()

			// If cell is filled, then set domain to just that number
			if input[i][j] != 0 {
				BoardDomains[i][j] = disallowAllNumbers()
				BoardDomains[i][j] = allowNumber(BoardDomains[i][j], input[i][j])
			}
		}
	}
}

// finalizeBoard returns the properly formatted 2D array from the BoardDomains
func finalizeBoard() (output [][]int) {
	output = make([][]int, BoardSize)
	for i := 0; i < BoardSize; i++ {
		output[i] = make([]int, BoardSize)
		for j := 0; j < BoardSize; j++ {

			allowedNumbers := getAllowedNumbers(BoardDomains[i][j])

			if len(allowedNumbers) == 1 {
				output[i][j] = allowedNumbers[0]
			} else {
				output[i][j] = 0
			}
		}
	}
	return output
}

// validate reports whether the numbers 1-9 can be recursively selected from the remaining bitset cells
func validate(remainingCells []CellDomain) bool {
	// Base case means a path forming 1-9 is possible
	if len(remainingCells) == 0 {
		return true
	}

	curNum := BoardSize - len(remainingCells) + 1

	for i, cell := range remainingCells {
		if isAllowedNumber(cell, curNum) {
			// Copies bitset array and removes the selected element
			newRemainingCells := append([]CellDomain{}, remainingCells[:i]...)
			newRemainingCells = append(newRemainingCells, remainingCells[i+1:]...)

			// Checks if the next remaining numbers can be selected from the remaining bitsets
			if validate(newRemainingCells) {
				return true
			}
		}
	}
	return false
}

// validateAllDomains reports whether all rows, cols, and boxes individually have possible solutions
func validateAllDomains() bool {
	// Check all rows
	for row := 0; row < BoardSize; row++ {
		res := validate([]CellDomain{
			BoardDomains[row][0],
			BoardDomains[row][1],
			BoardDomains[row][2],
			BoardDomains[row][3],
			BoardDomains[row][4],
			BoardDomains[row][5],
			BoardDomains[row][6],
			BoardDomains[row][7],
			BoardDomains[row][8],
		})
		if !res {
			return false
		}
	}
	// Verify all columns
	for col := 0; col < BoardSize; col++ {
		res := validate([]CellDomain{
			BoardDomains[0][col],
			BoardDomains[1][col],
			BoardDomains[2][col],
			BoardDomains[3][col],
			BoardDomains[4][col],
			BoardDomains[5][col],
			BoardDomains[6][col],
			BoardDomains[7][col],
			BoardDomains[8][col],
		})
		if !res {
			return false
		}
	}
	// Checks all boxes
	// Loops through box offsets 3,6,9 in both dimensions
	for h_offset := 0; h_offset < BoardSize; h_offset += BoxSize {
		for v_offset := 0; v_offset < BoardSize; v_offset += BoxSize {
			res := validate([]CellDomain{
				BoardDomains[0+h_offset][0+v_offset],
				BoardDomains[1+h_offset][0+v_offset],
				BoardDomains[2+h_offset][0+v_offset],
				BoardDomains[0+h_offset][1+v_offset],
				BoardDomains[1+h_offset][1+v_offset],
				BoardDomains[2+h_offset][1+v_offset],
				BoardDomains[0+h_offset][2+v_offset],
				BoardDomains[1+h_offset][2+v_offset],
				BoardDomains[2+h_offset][2+v_offset],
			})
			if !res {
				return false
			}
		}
	}
	return true
}

// DFS recursively searches through the board to try all numbers and backtracks with boolean representing whether a solution has been found
func DFS(cell int) bool {
	// if cell number past board, then report if valid board
	if cell >= BoardSize*BoardSize {
		return validateAllDomains()
	}

	// Get row & col from cell number
	row := cell % BoardSize
	col := cell / BoardSize

	// Saves old domain in case of backtrack
	oldDomain := BoardDomains[row][col]

	// Loops through all possible numbers in domain
	for k := 1; k < BoardSize+1; k++ {

		if isAllowedNumber(oldDomain, k) {

			// Collapses domain to just that number
			BoardDomains[row][col] = disallowAllNumbers()
			BoardDomains[row][col] = allowNumber(BoardDomains[row][col], k)

			// Creates a slice to store the adjacent cells that had their domains modified (pruned from tree)
			pruned := []int{}

			// loops through the same column as the selected cell
			for adjRow := 0; adjRow != row && adjRow < BoardSize; adjRow++ {
				if isAllowedNumber(BoardDomains[adjRow][col], k) {
					BoardDomains[adjRow][col] = disallowNumber(BoardDomains[adjRow][col], k)
					pruned = append(pruned, adjRow+col*BoardSize)
				}
			}

			// loops through the same row as the selected cell
			for adjCol := 0; adjCol != col && adjCol < BoardSize; adjCol++ {
				if isAllowedNumber(BoardDomains[row][adjCol], k) {
					BoardDomains[row][adjCol] = disallowNumber(BoardDomains[row][adjCol], k)
					pruned = append(pruned, row+adjCol*BoardSize)
				}
			}

			// loops through the same box as the selected cell
			for i := 0; i < BoxSize; i++ {
				for j := 0; j < BoxSize; j++ {

					// Calculates the index offset of the box
					h_offset := (row / BoxSize) * BoxSize
					v_offset := (col / BoxSize) * BoxSize

					// Skips if cell has already been collapsed or been pruned
					if i+h_offset == row && j+v_offset == col {
						continue
					}

					// Skip if number not in domain
					if !isAllowedNumber(BoardDomains[i+h_offset][j+v_offset], k) {
						continue
					}

					BoardDomains[i+h_offset][j+v_offset] = disallowNumber(BoardDomains[i+h_offset][j+v_offset], k)
					pruned = append(pruned, i+h_offset+(j+v_offset)*BoardSize)
				}
			}

			// If passes full domain check, then recurse
			if validateAllDomains() {
				if DFS(cell + 1) {
					return true
				}
			}

			// Backtrack and allow the pruned numbers back into the domains
			for _, v := range pruned {
				row_i := v % BoardSize
				col_i := v / BoardSize
				BoardDomains[row_i][col_i] = allowNumber(BoardDomains[row_i][col_i], k)
			}

			// Sets selected cell to its old domain
			BoardDomains[row][col] = oldDomain

		}
	}
	return false
}

func SolveSudoku(input [][]int) [][]int {
	initBoard(input)

	DFS(0)

	return finalizeBoard()
}
