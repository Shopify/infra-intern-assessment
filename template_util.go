package main

import (
	"bufio"
	"embed"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
)

// GenerateSudokuTemplates returns an int which is number of elements appended to the templates slice.
// It generates all possible patterns a digit from 1 to 9 can be placed in a traditional Sudoku grid, which is 46656.
// Each template is a bit vector representation of a Sudoku grid where 1 represents a valid position and 0 represents an empty cell.
func GenerateSudokuTemplates(currGrid *big.Int, freeGrid *big.Int, templates *[]big.Int, row int, count int) int {
	if row >= 9 {
		*templates = append(*templates, *new(big.Int).Set(currGrid))
		return count + 1
	}

	for i := 0; i < 9; i++ {
		cellPos := row*9 + i
		if freeGrid.Bit(cellPos) != 0 {
			continue
		}
		prevFreeGrid := new(big.Int).Set(freeGrid)
		setCell(currGrid, freeGrid, cellPos)
		count = GenerateSudokuTemplates(currGrid, freeGrid, templates, row+1, count)
		currGrid.SetBit(currGrid, cellPos, 0)
		freeGrid.Set(prevFreeGrid)
	}

	return count
}

// setCell modifies a bit vector representation of a Sudoku grid by setting the pos (from 0-80) and then
// updates the freeGrid bit vector to reflect the new conflicts introduced by setting the cell.
func setCell(grid *big.Int, freeGrid *big.Int, pos int) {
	// Set cell
	grid.SetBit(grid, pos, 1)

	// Set row conflicts
	row := pos / 9
	for i := 0; i < 9; i++ {
		freeGrid.SetBit(freeGrid, row*9+i, 1)
	}

	// Set column conflicts
	col := pos % 9
	for i := 0; i < 9; i++ {
		freeGrid.SetBit(freeGrid, i*9+col, 1)
	}

	// Set box conflicts
	boxRow := row / 3
	boxCol := col / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			freeGrid.SetBit(freeGrid, (boxRow*3+i)*9+(boxCol*3+j), 1)
		}
	}
}

// SaveTemplatesToFile saves the value of each bit vector template
// encoded in base64 into a text file specified by filename.
func SaveTemplatesToFile(templates []big.Int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, template := range templates {
		templateBytes := template.Bytes()
		encodedTemplate := base64.StdEncoding.EncodeToString(templateBytes)

		_, err := fmt.Fprintln(writer, encodedTemplate)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadTemplatesFromFile reads from an embed.FS file and returns a slice of big.Int's
// which represents a template of a Sudoku grid It reads every base64 value as a line
// which is then converted into a bit vector.
func ReadTemplatesFromEmbed(filename embed.FS) ([]big.Int, error) {
	file, err := filename.Open("templates.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var templates []big.Int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var template big.Int
		decodedTemplate, err := base64.StdEncoding.DecodeString(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("Error parsing template from file")
		}
		template.SetBytes(decodedTemplate)
		templates = append(templates, template)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return templates, nil
}
