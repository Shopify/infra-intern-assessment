# Shopify Infrastructure Intern Assessment - Sudoku Solver 
By Rick Zhang

This Go program provides a backtracking algorithm-based Sudoku solver.
The 'SolveSudoku' function attempts to solve a 9x9 Sudoku puzzle. 
We assume that there exists a solution to the sudoku puzzle. When successful, 
we will print the board out.

This program creates constraints as the board is filled and uses 
backtracking in order to reduce the number of combinations to consider.

Preconditions:
1. The input grid will be a 9x9 two-dimensional array of integers.
2. The input grid will have exactly one solution.
3. The input grid may contain zeros (0) to represent empty cells.

Postconditions:
1. The solved sudoku puzzle will be printed.
2. The solved sudoku puzzle will be returned.


Here is a general overview of the algorithm:
1. Create a counter for the numbers used in each row, column and box.
2. Starting from (0,0), we will iterate through each row from left to right, 
	as if we were reading the board like a book.
3. If the cell is not already set, try placing every feasible number (a feasible number is one that does not 
	violate the rules of sudoku) in the cell. 
	Placing a number in an empty cell includes:
	 - adding that number to the corresponding row, column and box counters
	 - adding the number to the board
4. After placing a number, go back to step 2, starting at the newly filled cell.
5. If there are no feasible numbers, we have made a wrong move. We will backtrack.
	Backtracking includes:
	 - removing the number from the cell
	 - removing the number from its corresponding row, column and box counters.
	Now, go back to step 3 to try a different number.