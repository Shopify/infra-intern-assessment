# Sudoku Solver - Shopify-Infra-Intern

## Built With

* [![Go][Go-logo]][Go-url]
* [![APIGateway][AWS-API-Gateway-logo]][AWS-API-Gateway-url]
* [![AWSLambda][AWS-Lambda-logo]][AWS-Lambda-URL]

A sudoku solution must satisfy **all of the following rules**:

1.  Each of the digits `1-9` must occur exactly once in each row.
2.  Each of the digits `1-9` must occur exactly once in each column.
3.  Each of the digits `1-9` must occur exactly once in each of the 9 `3x3` sub-boxes of the grid.

The `'0'` character indicates empty cells.

## Pseudocode

Backtracking:
Backtracking Algorithm to Solve Sudoku:
  Part 1: Check Placement
    - Define a function `CanPlace(board, row, col, num)` that returns `True` if it's legal to place `num` at `board[row][col]`, considering Sudoku rules. Otherwise, return `False`.

  Part 2: Solve Sudoku
    - Define a recursive function `Solve(board)` that attempts to fill the board with valid numbers.
    - Start with the first cell (row=0, col=0).
    - Iterate over each cell in the board row by row and column by column.
      - If the current cell is empty (denoted by 0):
        - Loop through possible numbers (1 to 9) and place each number in the cell.
        - Call `CanPlace` to check if the current number can be legally placed.
        - If `CanPlace` returns `True`, place the number and call `Solve` recursively for the next cell.
        - If the recursive call returns `True`, the board is solved; return `True`.
        - If placing the number does not lead to a solution, reset the cell (backtrack) and try the next number.
      - If no number can be placed, return `False` (triggering further backtracking).
    - If the end of a row is reached (col=9), move to the next row (row+1) and reset col to 0.
    - The base condition is when the last row is reached (row=9), meaning the board is solved; return `True`.
    <br>
`Time Complexity - O(9^(n*n))`
     
![Image](https://i.imgur.com/jXDkaEX.jpg)
![image](https://i.imgur.com/cqbF8rV.jpg)


## Deployment
The Go code has been deployed as an AWS Lambda function, with Amazon API Gateway handling the incoming requests.
  ![Image](https://i.imgur.com/iT2rHE8.png)

### Accessing Sudoku Solver via RESTAPI through script

Before executing the `solve_sudoku.sh` script, it's important to ensure that the file has the appropriate permissions set. You can update the file permissions to allow the owner to read and execute the script by using the `chmod` command. Here's how you can do it:

1. Open your terminal.
2. Navigate to the directory where the `solve_sudoku.sh` file is located.
3. Run the following command to set the permissions:

```bash
chmod 500 solve_sudoku.sh
  ```
* Usage
  ```sh
  ./solve_sudoku.sh '[[5, 3, 0, 0, 7, 0, 0, 0, 0], [6, 0, 0, 1, 9, 5, 0, 0, 0], ...]
  ```
  
[AWS-logo]: https://img.shields.io/badge/AWS-232F3E?style=for-the-badge&logo=amazon-aws&logoColor=white
[AWS-url]: https://aws.amazon.com/

[Go-logo]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://www.terraform.io/](https://go.dev/)https://go.dev/

[AWS-API-Gateway-logo]: https://img.shields.io/badge/AWS_API_Gateway-232F3E?style=for-the-badge&logo=amazon-aws&logoColor=white
[AWS-API-Gateway-url]: https://aws.amazon.com/api-gateway/

[AWS-Lambda-logo]: https://img.shields.io/badge/AWS_Lambda-232F3E?style=for-the-badge&logo=amazon-aws&logoColor=white
[AWS-Lambda-url]: https://aws.amazon.com/lambda/
