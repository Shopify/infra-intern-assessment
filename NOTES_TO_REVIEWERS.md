# Notes to Reviewers

Hi there! My name is Justin Jao, and here's my solution for the Shopify Infrastructure Intern Assessment challenge. I've written up some further explanations to outline my approach, and address any possible questions reviewers may have.

Happy reading!

## Overview of Approach

This approach is a standard backtracking algorithm, that essentially performs a "depth-first-search" over the solution space of the Sudoku Board. The general overview of the solution is as follows:

* We scan through the board, looking for a square that has yet to be filled
* Once an empty square has been found, we try to fill in a number from 1-9, filling in the smallest number that does not have a conflict
* For each number, we scan the row, column and box outline (which is calculated via modulo and taking into account the row and column). If the current number is not present in any of these 3 places, we choose that number
* We then recursively try to find solutions by picking another empty square and trying to fill it in.
* if at any point a "dead end" is reached (i.e. no solution can be found with the current chosen numbers), we "backtrack", changing previously filled numbers into the next largest number (in the exact order filled out) and continuing the search until a correct solution is found 

## Other Notes
* This backtracking implementation was based on the explanations outlined from [GeeksForGeeks](https://www.geeksforgeeks.org/sudoku-backtracking-7/).
* This solution only works given the constraints (in particular, that each input grid has exactly one solution)
* I considered other implementations, such as rule based or probability based algorithms, but chose backtracking as this was more efficient, as well as to optimize for easier code reviewing by others.
* To further optimize this solution, further work could involve replacing the recursive solution with an iterative implementation, by utilizing a stack to avoid the recursive calls.
* I added extra tests in the `sudoku_extra_test.go` file, because the instructions did not stipulate that we could modify `sudoku_test.go`, but I wanted to add a couple more tests to be certain that my solution was working.
* Extra puzzles were generated via [sudokuweb.org](https://www.sudokuweb.org/)
* For the sake of easier reviewing, I have added more comments than necessary. But if I were to productionize this, I would not use this many comments, only adding them if they were necessary to explain why a certain thing was done, rather than what a line of code is doing.
