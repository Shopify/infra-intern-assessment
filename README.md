# Shopify Intern Assessment Production Engineering

## Context

We're structuring this problem as a [Constraint Satisfaction Problem](https://en.wikipedia.org/wiki/Constraint_satisfaction_problem), using this as the foundation we will be using backtracking search with degree heuristics as the algorithm

Peers of a cell  represent all cells that directly depend on this cell to satisfy the constraints, aka all cells in the same row, column and sub square.

## Algorithm

- Optimized Backtracking Search with Degree Heuristics
- We perform domain pruning upon all peers of a cell when that cell's values are updated to reduce the search space
- When we first iterate through the board, we set each cell with the domain of [1,..,9], if we run into a cell with a value that's already defined we can prune each of it's peer's domain
- If we run into a cell that isn't assigned, we create a variable for it to assign to during backtracking search2
- We use a simple degree heuristic by selecting the variables with the fewest number of values in it's domain first

# References

https://en.wikipedia.org/wiki/Backtracking
https://en.wikipedia.org/wiki/Constraint_satisfaction_problem
https://norvig.com/sudoku.html
