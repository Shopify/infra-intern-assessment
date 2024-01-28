# Shopify Infrastructure Engineering Intern Assessment
### Submission by Jessica Li
## About This Repository
This repo has been forked from Shopify's Infra Intern 2024 Assessment. 
The task is to create a Go program to solve a Sudoku puzzle. More information can be found [here](instructions.md).
### Constraints
- The input grid will be a 9x9 two-dimensional array of integers.
- The input grid will have exactly one solution.
- The input grid may contain zeros (0) to represent empty cells.

### Repo Contents
- `sudoku.go` contains the main `SolveSudoku` function as well as additional helper functions to solve a given 9x9 sudoku grid
- `sudoku_test.go` contains the given test case and additional test cases to run `sudoku.go` against
- `go.mod` specifies the project dependency and version required to run the code
- `.github\workflows\test.yml` contains GitHub Actions

## How to Run
There is no main function within `sudoku.go` as it was not specified by the requirements. However, the validity can be tested against `sudoku_test.go`. I have added additional test cases as seen in the next section. 

Within the repo root directory:
```bash
go test -v . # runs all tests with verbose output 

go test -run TestSolveSudoku -v # runs the specified (original) test
```

### Test Cases
I have created additional test cases in `sudoku_test.go` to validate my code by using varying levels of difficulty sudoku boards and solutions. The test cases were obtained from [QQWing Sudoku](https://qqwing.com/generate.html) and were formatted with a small Python script below. 
```Python
def print_sudoku(board):
    # Replace '.' with '0'
    board = board.replace(".", "0")

    # Print the Sudoku board in the required format
    for i in range(9):
        row = ", ".join(board[i * 9 : i * 9 + 9])
        print(f'{{{row}}},')
```
## Implementation
The algorithm that I used in `sudoku.go` follows the standard brute force and backtracking method that most Sudoku players would typically use. Essentially, the implementation uses recursive backtracking to explore different possibilities, attempting to fill each empty spot with numbers 1 to 9 and backtracking if a valid solution is not found (i.e. resets the current spot to 0). The process continues until a solution is discovered or all possibilities are exhausted.
