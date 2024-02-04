package main

//Returns a set of all possible solutions for the specified position
func GetSolns(input [][]int, pos_x int, pos_y int) map[int]bool {
	var possible_solns map[int]bool = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true} // I wish there was a proper Set type

	// Remove numbers in same row and column from possible solutions
	for i := 0; i < 9; i++ {
		delete(possible_solns, input[pos_y][i])
		delete(possible_solns, input[i][pos_x])
	}

	box_x := (pos_x / 3) * 3
	box_y := (pos_y / 3) * 3 // Rounds to the upper left corner of 3x3 box

	// Remove numbers in same box from possible solutions
	for dx := 0; dx < 3; dx++ {
		for dy := 0; dy < 3; dy++ {
			delete(possible_solns, input[box_y+dy][box_x+dx])
		}
	}

	return possible_solns
}

// Solves the puzzle (if solvable) starting from the given location and
// returns true. Otherwise, returns false and the puzzle is unmodified.
func RunSolve(input [][]int, pos_x int, pox_y int) bool {

	//move x, y to the first blank
	for input[pox_y][pos_x] != 0 {
		pos_x++
		if pos_x == 9 {
			pos_x = 0
			pox_y++
			if pox_y == 9 {
				return true // Base case 1: there are no more blank values to solve for
			}
		}
	}

	possible_solns := GetSolns(input, pos_x, pox_y)

	// Reccursive step: try every solution generated. If one solves it, good. If not, continue.
	// Because the reccursive step reduceses the blank areas by one, the reccursive call must
	// return at some point because of base case #1.
	for soln := range possible_solns {
		input[pox_y][pos_x] = soln
		if RunSolve(input, pos_x, pox_y) {
			return true 
		}
	}

	input[pox_y][pos_x] = 0
	return false // Base case 2: Either the input puzzle has no solution (len(possible_solns) == 0) or all the possible solns failed.
}

func SolveSudoku(input [][]int) [][]int {
	RunSolve(input, 0, 0)
	return input
}
