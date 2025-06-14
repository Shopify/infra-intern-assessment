package main

import "fmt"

// Solves the a sudoku
// Uses a set approach as the possible solutions for a square is the intersection of the sets of possible numbers for the column, row, and box the number is in
// Finds the next number to solve using a stack
func SolveSudoku(sudoku [][]int) [][]int {
	// prepares all the sets for the rows, columns, and boxes as well as the stack which stores all empty slots
	rowSets, colSets, boxSets, nextS := initSudoku(sudoku)
	// stores the first value to solve for in the sudoku
	first := nextS.pop()
	return dfsSudoku(sudoku, first[0], first[1], rowSets, colSets, boxSets, nextS)
}

// initializes necessary variables to solve the sudoku
func initSudoku(sudoku [][]int) ([]*Set, []*Set, []*Set, Stack) {
	// create the next stack
	nextS := Stack{}

	// Create the sets for the rows
	rowSets := make([]*Set, 9)
	for i, row := range sudoku {
		// Create the set for a specific row
		rowSets[i] = &Set{arr: make([]bool, 9)}

		// add all numbers to set and remove those that are found in the row
		for j := 0; j < 9; j++ {
			rowSets[i].add(j)
		}
		for _, val := range row {
			if val != 0 {
				rowSets[i].remove(val - 1)
			}
		}
	}

	//Create the sets for the columns
	colSets := make([]*Set, 9)
	for i := 0; i < 9; i++ {
		// Create the set for the column
		colSets[i] = &Set{arr: make([]bool, 9)}

		// add all numbers to the set and remove those that are found in the column
		for j := 0; j < 9; j++ {
			colSets[i].add(j)
		}
		for j := 0; j < 9; j++ {
			if sudoku[j][i] != 0 {
				colSets[i].remove(sudoku[j][i] - 1)
			} else {
				// push all empty squares in the sudoku into the next stack
				nextS.push([]int{j, i})
			}
		}
	}

	//Create the sets for the boxes
	boxSets := make([]*Set, 9)
	for i := 0; i < 9; i++ {
		// the row and column for the top left of a box
		base := []int{(i / 3) * 3, (i % 3) * 3}

		// create the set for the box
		boxSets[i] = &Set{arr: make([]bool, 9)}

		// add all numbers to the box and remove those found
		for j := 0; j < 9; j++ {
			boxSets[i].add(j)
		}
		for j := 0; j < 9; j++ {
			if sudoku[base[0]+(j/3)][base[1]+(j%3)] != 0 {
				boxSets[i].remove(sudoku[base[0]+(j/3)][base[1]+(j%3)] - 1)
			}
		}
	}

	return rowSets, colSets, boxSets, nextS
}

// Calculates the number of the box, 0 is top left, 8 is bottom right in a zig zagging pattern
func calcBox(row int, col int) int {
	return ((row / 3) * 3) + (col / 3)
}

// recursive funtion to solve a square in the sudoku
func dfsSudoku(sudoku [][]int, row int, col int, rowSets []*Set, colSets []*Set, boxSets []*Set, nextS Stack) [][]int {
	// finds the set of possible numbers for a square
	possSet := intersection(rowSets[row], colSets[col], boxSets[calcBox(row, col)])

	// if there are no possible numbers, return to a previous square
	if possSet.isEmpty() {
		return nil
	}

	// an integer slice of all possible values of a square
	poss := possSet.possValues()

	// tests every possible value for a square
	for _, val := range poss {

		// remove value from all sets
		rowSets[row].remove(val)
		colSets[col].remove(val)
		boxSets[calcBox(row, col)].remove(val)

		// put value on sudoku board
		sudoku[row][col] = val + 1

		// if there is no next empty square, then the sudoku has been solved
		if nextS.isEmpty() {
			return sudoku
		}

		// get the next square and solve it
		next := nextS.pop()
		temp := dfsSudoku(sudoku, next[0], next[1], rowSets, colSets, boxSets, nextS)

		// if the next square cannot be solved, try another possible value
		if temp != nil {
			return temp
		}

		// return all sets and the sudoku to their original values
		nextS.push(next)
		rowSets[row].add(val)
		colSets[col].add(val)
		boxSets[calcBox(row, col)].add(val)
		sudoku[row][col] = 0
	}

	// if no possible value works for this square, change a previous square and return to this one later
	return nil
}

func main() {
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
	fmt.Println(SolveSudoku(input))
}

// A stack that stores the next positions
type Stack struct {
	top *node
}

// Node that stores a row and column in its integer array
type node struct {
	val  []int
	past *node
}

// returns the string representation of a node
func (n *node) String() string {
	return fmt.Sprint(n.val)
}

// pops the top value off of the stack
func (s *Stack) pop() []int {
	temp := s.top
	s.top = s.top.past
	return temp.val
}

// pushes a value to the top of the stack
func (s *Stack) push(val []int) {
	n := node{val: val, past: s.top}
	s.top = &n
}

// returns if a stack is empty
func (s *Stack) isEmpty() bool {
	return s.top == nil
}

// returns the string representation of a stack
func (s *Stack) String() string {
	out := ""
	cur := s.top
	for cur != nil {
		out += cur.String()
		cur = cur.past
	}
	return out
}

// A Set which stores whether a number is within the set
type Set struct {
	arr []bool
	len int
}

// Returns the intersection of three sets
func intersection(s1 *Set, s2 *Set, s3 *Set) *Set {
	out := Set{arr: make([]bool, 9), len: 0}

	// checks if a value is if all three sets, if so then it is in the intersection
	for i := 0; i < 9; i++ {
		if s1.arr[i] && s2.arr[i] && s3.arr[i] {
			out.arr[i] = true
			out.len++
		}
	}
	return &out
}

// Removes a value from the set
func (s *Set) remove(val int) {
	if s.arr[val] {
		s.len--
	}
	s.arr[val] = false
}

// Adds a value to the set
func (s *Set) add(val int) {
	if s.arr[val] {
		s.len++
	}
	s.arr[val] = true
}

// Returns the possible values based on the set
func (s *Set) possValues() []int {
	out := make([]int, 0)

	// puts all values that are true in the array of the set in a slice
	for i := 0; i < 9; i++ {
		if s.arr[i] {
			out = append(out, i)
		}
	}
	return out
}

// returns string representation of the set
func (s *Set) String() string {
	return fmt.Sprintf("%v", s.possValues())
}

// returns if a set is empty
func (s *Set) isEmpty() bool {
	return s.len == 0
}
