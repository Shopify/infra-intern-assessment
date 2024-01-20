package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	builder := bytes.NewBuffer(nil)

	// Read all texts from standard input and write to the buffer.
	for scanner.Scan() {
		builder.Write(scanner.Bytes())
	}

	board := make(Sudoku, 0, 9)

	// Parse the JSON-style Sudoku data into the Sudoku board.
	if err := json.Unmarshal(builder.Bytes(), &board); err != nil {
		fmt.Printf("Parse sudoku from stdin failed: %v", err)
		return
	}

	// Print the solved Sudoku board to the stdout.
	fmt.Printf("%s\n", SolveSudoku(board))
}
