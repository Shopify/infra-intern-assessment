#!/bin/bash

# Usage: ./solve_sudoku.sh '[[5, 3, 0, 0, 7, 0, 0, 0, 0], [6, 0, 0, 1, 9, 5, 0, 0, 0], ...]'

BOARD=$1   #cmd line arg

curl -X POST https://hhqspow89c.execute-api.us-east-1.amazonaws.com/prod/ShopifySudoku \
     -H "Content-Type: application/json" \
     -d "{\"board\": $BOARD}"

