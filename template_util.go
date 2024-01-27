package main

import (
	"bufio"
	"embed"
	"fmt"
	"math/big"
	"os"
)

func GenerateSudokuTemplates(currGrid *big.Int, freeGrid *big.Int, templates *[]big.Int, row int, count int) int {
	if row >= 9 {
		newTemplate := new(big.Int).Set(currGrid)
		*templates = append(*templates, *newTemplate)
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

func SaveTemplatesToFile(templates []big.Int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, template := range templates {
		_, err := fmt.Fprintln(writer, &template)
		if err != nil {
			return err
		}
	}

	return nil
}

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
		_, success := template.SetString(scanner.Text(), 10)
		if !success {
			return nil, fmt.Errorf("Error parsing template from file")
		}
		templates = append(templates, template)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return templates, nil
}
