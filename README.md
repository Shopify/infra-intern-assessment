# Introduction
Hello! My name is Karthik and I'm a 3a Computer Science student at the University of Waterloo. 

# Shopify Infra Intern Assessment

## Solution
This Sudoku solver works using a search and backtracking solution. In short, it tries to fill each empty cell on the board in order starting from the top left, and checks if the number we filled in is valid or not.

Here's how it works:

1. It starts at the top-left cell (0,0) and traverses through each cell in the grid row by row ignoring cells that are already filled.
2. If a cell is not filled, it tries all numbers from 1-9 until it finds one that is valid to enter then recursively tries to fill in the next
3. If it manages to fill all the cells after the current cell without violating the Sudoku rules, we have found a solution.
4. If it becomes impossible to fill a cell with a valid number, it clears that cell and backtracks

The helper function `isValid()` does the validation check. It accepts a number and a cell position and checks if it's valid (no repetition on row, col or block) to put in the cell.

## Conclusion
Thanks for taking a look! Hope to hear back from you guys soon!
