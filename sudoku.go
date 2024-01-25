package main

const BoardSize = 9
const BoxSize = 3

var CellDomains [BoardSize][BoardSize]uint16

// STATE
// 9 x 9 array of 9 bit vectors
// Each bit represents whether (index + 1) is still in the domain for that sudoku cell

// ex: Number              123456789
//     CellDomains[0][0] = 010101010
//     Domain               2 4 6 8

// MAIN DFS Function O()
// Recursively iterates over the board (row wise)
// Everytime an slot is assigned (backtracking possible), a generalized arc consistency is run to verify that the every cell's domain is still valid

// GAC/Domain TESTER
// Checks that every row, col, and box is valid
// Will not check that they will necessarily work together
// Runs DFS on the 9 9-bit vectors to ensure a path is possible (the numbers 1-9 can be selected from each domain)
// O(N^N)

// 123   456   789
//
// 100   000   000
// 010   000   000
// 001   000   000
//
// 010   100   000
// 010   010   000
// 010   001   000
//
// 010   000   100
// 010   000   010
// 010   000   001

// initBoard initializes the CellDomains into bitsets representing the remaining domain per cell
func initBoard(input [][]int) {
	// Loops through input board
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if input[i][j] != 0 {
				// If cell is filled, then collapse domain (0b000_010_000)
				CellDomains[i][j] = 1 << (input[i][j] - 1)
			} else {
				// if cell is open, then open domain (0b111_111_111)
				CellDomains[i][j] = (1 << BoardSize) - 1
			}
		}
	}
}

// finalizeBoard returns the properly formatted 2D array from the CellDomains
func finalizeBoard() (output [][]int) {
	// Initializes and loops through first dimension
	output = make([][]int, BoardSize)
	for i := 0; i < BoardSize; i++ {

		// Initializes and loops through the 2nd dimension row
		output[i] = make([]int, BoardSize)
		for j := 0; j < BoardSize; j++ {

			// Tracks how many true bits have been seen
			trueBits := 0

			// Loops through the bitset
			for k := 0; k < BoardSize; k++ {
				// Adds the index + 1 to the element if the bit is True
				if CellDomains[i][j]&(1<<k) != 0 {
					output[i][j] = k + 1
					trueBits += 1
				}
			}

			// If the domain is still flexible, set to zero (indicates a problem)
			if trueBits != 1 {
				output[i][j] = 0
			}
		}
	}
	return output
}

// validate reports whether the numbers 1-9 can be recursively selected from the remaining bitset cells (house is valid)
func validate(arr []uint16) bool {
	// Base case means a path forming 1-9 is possible
	if len(arr) == 0 {
		return true
	}

	// represents the bitset index of number that's being searched (ex: starts with looking for True at index 0)
	k := BoardSize - len(arr)

	// Loops through remaining bitsets
	for i := 0; i < len(arr); i++ {
		// Checks if bitset is 1 at index k
		if (arr[i] & (1 << k)) != 0 {
			// Copies bitset array and removes the selected element
			newArr := append([]uint16{}, arr[:i]...)
			newArr = append(newArr, arr[i+1:]...)

			// Checks if the next remaining numbers can be selected from the remaining bitsets
			if validate(newArr) {
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
		res := validate([]uint16{
			CellDomains[row][0],
			CellDomains[row][1],
			CellDomains[row][2],
			CellDomains[row][3],
			CellDomains[row][4],
			CellDomains[row][5],
			CellDomains[row][6],
			CellDomains[row][7],
			CellDomains[row][8],
		})
		if !res {
			return false
		}
	}
	// Verify all columns
	for col := 0; col < BoardSize; col++ {
		res := validate([]uint16{
			CellDomains[0][col],
			CellDomains[1][col],
			CellDomains[2][col],
			CellDomains[3][col],
			CellDomains[4][col],
			CellDomains[5][col],
			CellDomains[6][col],
			CellDomains[7][col],
			CellDomains[8][col],
		})
		if !res {
			return false
		}
	}
	// Checks all boxes
	// Loops through box offsets 3,6,9 in both dimensions
	for h_offset := 0; h_offset < BoardSize; h_offset += BoxSize {
		for v_offset := 0; v_offset < BoardSize; v_offset += BoxSize {
			res := validate([]uint16{
				CellDomains[0+h_offset][0+v_offset],
				CellDomains[1+h_offset][0+v_offset],
				CellDomains[2+h_offset][0+v_offset],
				CellDomains[0+h_offset][1+v_offset],
				CellDomains[1+h_offset][1+v_offset],
				CellDomains[2+h_offset][1+v_offset],
				CellDomains[0+h_offset][2+v_offset],
				CellDomains[1+h_offset][2+v_offset],
				CellDomains[2+h_offset][2+v_offset],
			})
			if !res {
				return false
			}
		}
	}
	return true
}

// DFS recursively runs depth first search throughout the board to try all numbers and backtracks with boolean representing whether a solution has been found
func DFS(cell int) bool {
	// if cell number past board, then report if valid board
	if cell >= BoardSize*BoardSize {
		return validateAllDomains()
	}

	// Get row & col from cell number
	row := cell % BoardSize
	col := cell / BoardSize

	// Saves old domain in case of backtrack
	oldDomain := CellDomains[row][col]

	// Loops through all possible numbers in domain/bitset indexes
	for k := 0; k < BoardSize; k++ {

		// Checks if number is in domain
		if oldDomain&(1<<k) != 0 {

			// Collapses domain to just that number
			CellDomains[row][col] = 1 << k

			// Creates a slice to store the adjacent cells that have their domains modified (pruned from tree)
			pruned := []int{}

			// loops through the same column as the selected cell
			for adjRow := 0; adjRow != row && adjRow < BoardSize; adjRow++ {
				// Checks if the number is still in the cell's domain
				if CellDomains[adjRow][col]&(1<<k) != 0 {
					// Flips the bit to off
					CellDomains[adjRow][col] ^= 1 << k
					// Saves the cell number for backtracking
					pruned = append(pruned, adjRow+col*BoardSize)
				}
			}

			// loops through the same row as the selected cell
			for adjCol := 0; adjCol != col && adjCol < BoardSize; adjCol++ {
				// Checks if the number is still in the cell's domain
				if CellDomains[row][adjCol]&(1<<k) != 0 {
					// Flips the bit to off
					CellDomains[row][adjCol] ^= 1 << k
					// Saves the cell number for backtracking
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

					// Checks if the number is still in the cell's domain
					if CellDomains[i+h_offset][j+v_offset]&(1<<k) != 0 {
						// Flips the bit to off
						CellDomains[i+h_offset][j+v_offset] ^= 1 << k
						// Saves the cell number for backtracking
						pruned = append(pruned, i+h_offset+(j+v_offset)*BoardSize)
					}
				}
			}

			// If passes full domain check, then recurse
			if validateAllDomains() {
				// Returns true if solution found
				if DFS(cell + 1) {
					return true
				}
			}

			// Backtrack and allow the pruned numbers back into the domains
			for _, v := range pruned {
				row_i := v % BoardSize
				col_i := v / BoardSize
				// Flips the bit to on
				CellDomains[row_i][col_i] |= 1 << k
			}

			// Sets selected cell to its old domain
			CellDomains[row][col] = oldDomain

		}
	}
	return false
}

func SolveSudoku(input [][]int) [][]int {
	// Initializes cell domains
	initBoard(input)

	// Starts at cell 0
	DFS(0)

	// Returns final grid
	return finalizeBoard()
}
