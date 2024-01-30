## Soduku Solver (For Shopify Infra Intern Assessment)

The following is the algorithm used to implement the sudoku solver. A box/row/column will be called a group for convenience.

1. Determine the set of numbers possible for each cell
2. For each cell whose set contains only one number, set the cell's value to be that number. Remove this number from all other cell's possible numbers set in its group.
3. If there is a number that can only go in one cell in a group, assign that number to that cell. Remove this number from all other cell's possible numbers set.
4. Box Row Column Elimination. If in a box, there are numbers only possible to be in cells for a single row/column. Then that number is not possible in any other cells in that row/column. An example is in the illustration below. 1 can only go in the middle column for the bottom-middle box. Thus, 1 cannot go anywhere in column e.
<p align="center"><image src="sudokuBoxRowElimination.png"/></p>

5. N-Tuple group elimination. If in any group, there are N cells that can only take on the same N values, then these N values must be restricted to that set and can be remove from the set of possible numbers for all other cells in the group. 

    For example, if in a row, two cells have possible values (1,2) and (1,2). Then 1 and 2 must go in these two cells. No other cell in that row can be 1 or 2. The same applies for a column or box

6. Hidden Tuple Elimination. If, in a group, there are N numbers that can only belong in the same N cells, then no other numbers can belong in the cell. Remove those numbers from the cells' set of possible values. In the image below, 1 and 4 can only go into the two red cells We would thus be able to remove 3,7 and 3,5 from the cell's possible values. (Note that this would let us solve that square d7 is 7).
<p align="center"><image src="sudokuHiddenTupleElimination.png"/></p>

7. Repeat steps 2-6 until no further simplifications can be done.

8. If the sudoku board is still not solved, then a guess must be made. Make guesses in order from smallest set of possible values to largest set of possible values. This way, a guess is more likely to be correct.

### Representing Set of Numbers Available

In Soduku, the numbers available to a cell can only range from 0 to 9. Since this is small, the numbers can be conveniently represented using a bitset. If a number $n$ is possible to be in a cell, then the bit $(1 << n)$ will be set for that cell.

To check if only one element is available for a cell, take ``x & (x-1)`` where x is the bitset of available numbers for the cell. If there is only one bit set in x, then the result of that computation will be 0 since subtracting 1 will turn the first set bit into a 0 and every less significant bit to a 1.