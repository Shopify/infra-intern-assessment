package main

import (
	"fmt"
	"math"
	"reflect"
)

// The approach:
//
// This is simply just a constraint satisfaction problem. We can first parse
// the input and create a CSP model to model the scenario. We can then use
// backtracking + AC3 to try and search for a solution to the Sudoku puzzle.

// NOTE:
// I'd like to preface this by saying I have not coded in Go before and had to quickly
// teach myself the basics of programming in Golang to attempt this challenge. This is why
// the naming conventions used may be inconsistent as I was focused on solving the problem
// in a timely manner. I hope that, at the very least, this demonstrates my adaptibility and
// my eagerness to learn. Thanks to my strong foundational knowledge, the main difficulty of
// this challenge didn't lie in coming up with a solution, but just translating what I know
// into a new programming language.
//
// As for my implementation, my approach to this problem is motivated by a similar problem
// I encountered when taking CSC384: Intro to Artificial Intelligence at the University of
// Toronto which was solving a Kropki Sudoku board for one of my assignments.

// An object to represent a two-tuple of integers.
//
// Attributes:
//   - a: The first component of the two-tuple.
//   - b: The second component of the two-tuple.
type twoTup struct {
	a int
	b int
}

// An object to represent a cell in the Sudoku puzzle.
//
// Attributes:
//   - name: A name for this variable.
//   - row: The row location of the cell.
//   - col: The column location of the cell.
//   - dom: The permanent domain of this variable.
//   - curdom: A mapping of values still in the domain. Maps to true if the value
//     is still in the domain.
//   - assignedVal: The value assigned to the variable or 0 if it is unassigned.
type variable struct {
	name        string
	row         int
	col         int
	dom         []int
	curdom      map[int]bool
	assignedVal int
}

// An object to represent a constraint between two cells.
//
// Attributes:
//   - name: A name for this constraint.
//   - cell1: A cell in the Sudoku puzzle.
//   - cell2: Another distinct cell in the Sudoku puzzle.
//   - sat_tuples: A collection of tuples that satisfy the constraint. They always map to true.
type binConstraint struct {
	name       string
	cell1      variable
	cell2      variable
	sat_tuples map[twoTup]bool
}

// An object to represent a CSP.
//
// Attributes:
//   - vars: A list of variable objects representing variables in the CSP.
//   - cons: A list of constraint objects representing constraints in the CSP.
//   - vars_to_cons: A mapping of the name representation of a variable object to a list of constriants
//     containing that variable.
type CSP struct {
	vars         []variable
	cons         []binConstraint
	vars_to_cons map[string]([]binConstraint)
}

// A class to perform backtracking search.
type BT struct {
	csp             CSP
	unassigned_vars []variable
}

// Returns whether the variable has an assigned value.
func isAssigned(curr_var variable) bool {
	return curr_var.assignedVal != -1
}

// Returns a list of values in the given variable's current domain.
func getCurrentDomain(curr_var variable) []int {
	values := []int{}

	if isAssigned(curr_var) {
		values = append(values, curr_var.assignedVal)
	} else {
		for _, val := range curr_var.dom {
			if curr_var.curdom[val] {
				values = append(values, val)
			}
		}
	}

	return values
}

// Returns the size of the variable's current domain.
func getCurrentDomainSize(curr_var variable) int {
	count := 0

	for _, v := range curr_var.curdom {
		if v {
			count++
		}
	}

	return count
}

// Returns whether the given value is in the given variable's current domain.
func inCurrentDomain(curr_var variable, value int) bool {
	_, ok := curr_var.curdom[value]
	if !ok {
		return false
	}

	if isAssigned(curr_var) {
		return value == curr_var.assignedVal
	} else {
		return curr_var.curdom[value]
	}
}

// Returns a list of unassigned variables for the given binary constraint
func getUnassignedVarsFromCons(cons binConstraint) []variable {
	vars := []variable{}

	if !isAssigned(cons.cell1) {
		vars = append(vars, cons.cell1)
	}

	if !isAssigned(cons.cell2) {
		vars = append(vars, cons.cell2)
	}

	return vars
}

// Creates satisfying tuples for binary difference constraints.
// Returns a slice of twoTup objects.
func createDifferenceConstraintSatTuples() []twoTup {
	var tups_so_far []twoTup

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			// A satisfying tuple is a pair of two integers such that they are not equal.
			if i != j {
				tups_so_far = append(tups_so_far, twoTup{i, j})
				tups_so_far = append(tups_so_far, twoTup{j, i})
			}
		}
	}

	return tups_so_far
}

// Helper method to add satisfying tuples to a constraint.
func addSatisfyingTuples(constraint binConstraint, sat_tuples []twoTup) {
	// Iterate over the satisfying tuples and add it to the map. Each value of a satisfying
	// tuple should be true
	for _, tup := range sat_tuples {
		_, ok := constraint.sat_tuples[tup]
		// Add the satisfying tuple to the map if it is not already in it.
		if !ok {
			constraint.sat_tuples[tup] = true
		}
	}
}

// Create the variables for this Sudoku puzzle. Each variable represents a cell.
// Returns a list of variables representing cells in the Sudoku puzzle as a 1D and 2D list.
func createVariables(input [][]int) ([81]variable, [9][9]variable) {
	var variables [81]variable
	var variables2D [9][9]variable

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var entry = input[i][j]
			// Default assigned value of -1 to each cell
			var curr_var = variable{fmt.Sprintf("Row%d,Col%d", i, j), i, j, []int{}, make(map[int]bool), -1}
			if 1 <= entry && entry <= 9 {
				curr_var.dom = append(curr_var.dom, entry)
				curr_var.curdom[entry] = true
			} else {
				for i := 1; i < 10; i++ {
					curr_var.dom = append(curr_var.dom, i)
					curr_var.curdom[i] = true
				}
			}

			variables[9*i+j] = curr_var
			variables2D[i][j] = curr_var
		}
	}

	return variables, variables2D
}

// Returns a list of binary row and column constraints for this Sudoku puzzle.
func createRowAndColConstraints(sat_tuples []twoTup, variables [81]variable) []binConstraint {
	var constraints []binConstraint
	n, dim := 81, 9

	for i, curr_var := range variables {
		// Create binary constraints between curr_var and the other cells in the column.
		for idx := i + dim; idx < n; idx = idx + dim {
			var constraint = binConstraint{
				fmt.Sprintf("ColCons(%s,%s)", curr_var.name, variables[idx].name),
				curr_var,
				variables[idx],
				make(map[twoTup]bool),
			}
			addSatisfyingTuples(constraint, sat_tuples)
			constraints = append(constraints, constraint)
		}

		// Create binary constraints between curr_var and the other cells in the row.
		for idx := i + 1; idx < dim*(i/dim+1); idx++ {
			var constraint = binConstraint{
				fmt.Sprintf("RowCons(%s,%s)", curr_var.name, variables[idx].name),
				curr_var,
				variables[idx],
				make(map[twoTup]bool),
			}
			addSatisfyingTuples(constraint, sat_tuples)
			constraints = append(constraints, constraint)
		}
	}

	return constraints
}

// Returns a list of binary constraints for each 3x3 cage in the 9x9 Sudoku board.
func createCageConstraints(sat_tuples []twoTup, variables [81]variable) []binConstraint {
	var constraints []binConstraint
	n, dim := 81, 9
	var row_limit int

	for i, curr_var := range variables {
		if i < 27 {
			row_limit = int(math.Floor(float64(n) / float64(3)))
		} else if i < 54 {
			row_limit = 2 * int(math.Floor(float64(n)/float64(3)))
		} else {
			row_limit = n
		}
		var col_lower = i - (i % 3)
		var col_upper = i + 3 - (i % 3)

		for row := i; row < row_limit; row = row + dim {
			// If we are in the top-right cell of a 3x3 cage, there are no pairings in the first row.
			if row == i && (i%3 == 2) {
				continue
			}

			for col := col_lower; col < col_upper; col++ {
				// We don't want to duplicate constraints.
				if row == i && col <= i {
					continue
				}
				other_var := variables[(int(math.Floor(float64(row)/float64(dim))))*dim+(col%dim)]
				var constraint = binConstraint{
					fmt.Sprintf("CageCons(%s,%s)", curr_var.name, other_var.name),
					curr_var,
					other_var,
					make(map[twoTup]bool),
				}
				addSatisfyingTuples(constraint, sat_tuples)
				constraints = append(constraints, constraint)
			}
		}
	}

	return constraints
}

// Create the CSP Model for the Sudoku puzzle.
func sudokuModel(input [][]int) (CSP, [9][9]variable) {
	variables, variables2D := createVariables(input)
	csp := CSP{[]variable{}, []binConstraint{}, make(map[string][]binConstraint)}

	// Add the variables to the CSP
	for _, curr_var := range variables {
		csp.vars = append(csp.vars, curr_var)
		csp.vars_to_cons[curr_var.name] = []binConstraint{}
	}
	// Create the satisfying tuples and the constraints
	difference_sat_tuples := createDifferenceConstraintSatTuples()
	cage_constraints := createCageConstraints(difference_sat_tuples, variables)
	row_and_col_constraints := createRowAndColConstraints(difference_sat_tuples, variables)

	// Add the constraints to the CSP
	for _, cons := range cage_constraints {
		var1 := cons.cell1
		var2 := cons.cell2
		csp.vars_to_cons[var1.name] = append(csp.vars_to_cons[var1.name], cons)
		csp.vars_to_cons[var2.name] = append(csp.vars_to_cons[var2.name], cons)
		csp.cons = append(csp.cons, cons)
	}

	for _, cons := range row_and_col_constraints {
		var1 := cons.cell1
		var2 := cons.cell2
		csp.vars_to_cons[var1.name] = append(csp.vars_to_cons[var1.name], cons)
		csp.vars_to_cons[var2.name] = append(csp.vars_to_cons[var2.name], cons)
		csp.cons = append(csp.cons, cons)
	}

	return csp, variables2D
}

// A wrapper function to run backtracking search.
func btSearch(solver BT) {
	// Reset the variable domains to their permanent domains.
	restoreVariableDomains(solver)

	// Gather the unassigned variables.
	solver.unassigned_vars = []variable{}
	for _, curr_var := range solver.csp.vars {
		if !isAssigned(curr_var) {
			solver.unassigned_vars = append(solver.unassigned_vars, curr_var)
		}
	}

	// Checking arc consistency before assigning any variables.
	status, _ := AC3(solver.csp)

	// Execute the recursive helper.
	if status {
		status = btHelper(solver)
	} else {
		fmt.Print("Funky stuff happened (bad) :(")
	}

	if status {
		fmt.Print("Sudoku solved.")
	} else {
		fmt.Print("No solution :(")
	}
}

// A helper method to perform the recursion for backtracking search.
func btHelper(solver BT) bool {
	// All variables have been assigned so a solution has been found.
	if len(solver.unassigned_vars) == 0 {
		return true
	} else {
		// Choose the next variable to assign.
		next_assigned_var := solver.unassigned_vars[0]
		solver.unassigned_vars = solver.unassigned_vars[1:]
		// Choose the order in which to assign values.
		value_order := getCurrentDomain(next_assigned_var)

		for _, val := range value_order {
			// Assign a value
			if !isAssigned(next_assigned_var) {
				next_assigned_var.assignedVal = val
			}

			// Restore arc consistency after the assignment.
			status, prunings := AC3(solver.csp, next_assigned_var)

			// Recurse if arc-consistency was restored.
			if status {
				if btHelper(solver) {
					return true
				}
			}

			// Unsuccessful assignment so restore the pruned values.
			restoreValues(solver, prunings)
			if isAssigned(next_assigned_var) {
				next_assigned_var.assignedVal = -1
			}
		}
		// Backtrack.
		solver.unassigned_vars = append(solver.unassigned_vars, next_assigned_var)
		return false
	}
}

// Restore the pruned values.
func restoreValues(solver BT, prunings map[*variable]int) {
	for var_pt, val := range prunings {
		(*var_pt).curdom[val] = true
	}
}

// Unassign all variables in the CSP and restore the current domains to their permanent domains.
func restoreVariableDomains(solver BT) {
	for _, curr_var := range solver.csp.vars {
		if isAssigned(curr_var) {
			curr_var.assignedVal = -1
		}
		for val := range curr_var.curdom {
			curr_var.curdom[val] = true
		}
	}
}

// Return an item from the "queue"
func pop(queue map[string]binConstraint) binConstraint {
	for k, v := range queue {
		delete(queue, k)
		return v
	}
	return binConstraint{}
}

// Execute the AC3 Algorithm.
func AC3(csp CSP, last_assigned_var ...variable) (bool, map[*variable]int) {
	var pruned map[*variable]int = make(map[*variable]int)
	var constraints []binConstraint
	queue := make(map[string]binConstraint)
	if len(last_assigned_var) > 0 {
		constraints = csp.vars_to_cons[last_assigned_var[0].name]
	} else {
		constraints = csp.cons
	}
	// Adding constraints to the queue
	for _, c := range constraints {
		queue[c.name] = c
	}

	for len(queue) > 0 {
		curr_constraint := pop(queue)

		// Restore arc consistency for this constraint.
		for _, curr_var := range getUnassignedVarsFromCons(curr_constraint) {
			if revise(curr_constraint, curr_var, pruned) {
				// If values were pruned and led to an empty domain, return as
				// the current assignment will not lead to a solution.
				if getCurrentDomainSize(curr_var) == 0 {
					return false, pruned
				}
				// Otherwise, add back relevant constraints back to the queue.
				for _, cons := range csp.vars_to_cons[curr_var.name] {
					if !reflect.DeepEqual(cons, curr_constraint) {
						queue[cons.name] = cons
					}
				}
			}
		}
	}

	return true, pruned
}

// Revise the given variable's domain to restore arc consistency.
// Returns true iff a domain was reduced.
func revise(constraint binConstraint, curr_var variable, pruned map[*variable]int) bool {
	revised := false

	// Try all values in the current domain.
	for _, val := range getCurrentDomain(curr_var) {
		// If a value is not part of a satisfying assignment, prune it, and update the flag.
		if !(isPartOfSatTuple(constraint, curr_var, val)) {
			curr_var.curdom[val] = false
			pruned[&curr_var] = val
			revised = true
		}
	}

	return revised
}

// Returns whether the current variable-value assignment is part of a satisfying tuple
func isPartOfSatTuple(constraint binConstraint, curr_var variable, value int) bool {
	for tup := range constraint.sat_tuples {
		var1, var2 := constraint.cell1, constraint.cell2
		val1, val2 := tup.a, tup.b

		if (reflect.DeepEqual(curr_var, var1) && val1 == value) || (reflect.DeepEqual(curr_var, var2) && val2 == value) {
			if (inCurrentDomain(var1, val1) && (!isAssigned(var1) || val1 == var1.assignedVal)) &&
				(inCurrentDomain(var2, val2) && (!isAssigned(var2) || val2 == var2.assignedVal)) {
				return true
			}
		}
	}

	return false
}

// SolveSudoku takes a 9x9 Sudoku board represented as a 2D array of integers as input
// and returns the solved board as a new 2D array of integers.
func SolveSudoku(input [][]int) [][]int {
	csp, variables2D := sudokuModel(input)
	var output [][]int
	var temp []int

	solver := BT{csp, []variable{}}
	btSearch(solver)

	// Constructing the solution.
	// Note: this is a bit hack-y. The domains are correctly reduced but they aren't being
	// assigned properly so this work is to reconstruct the solution.
	for row := 0; row < 9; row++ {
		temp = []int{}
		for col := 0; col < 9; col++ {
			curr_var := variables2D[row][col]
			for k, v := range curr_var.curdom {
				if v {
					temp = append(temp, k)
					break
				}
			}
		}
		output = append(output, temp)
	}

	return output
}
