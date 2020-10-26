package grid_test

import (
	"strings"
	"testing"
	"wordsearch/pkg/grid"
)

func assertGridSize(t *testing.T, input string, grid *grid.Grid, size int) {
	if len(grid.Letters) != size {
		t.Fatalf("Expected \"%s\" to be placed in a %dx%d grid, but the returned grid was: %v", input, size, size, grid.Letters)
	}
	for _, row := range grid.Letters {
		if len(row) != size {
			t.Fatalf("Expected \"%s\" to be placed in a %dx%d grid, but the returned grid was: %v", input, size, size, grid.Letters)
		}
	}
}

func TestGenerateCorrectSizeGrid(t *testing.T) {
	input := "12"
	newGrid, _ := grid.Generate(strings.NewReader(input))
	assertGridSize(t, input, &newGrid, 2)

	input = "12,34"
	newGrid, _ = grid.Generate(strings.NewReader(input))
	assertGridSize(t, input, &newGrid, 2)

	input = "12345,12345,12345,12345,12345,12345"
	newGrid, _ = grid.Generate(strings.NewReader(input))
	assertGridSize(t, input, &newGrid, 6)
}

func TestGeneratePlacesAllWords(t *testing.T) {
	words := []string{"some", "word", "are", "hard"}
	input := strings.Join(words,",")

	newGrid, _ := grid.Generate(strings.NewReader(input))

	var rows []string
	for _, row := range newGrid.Letters {
	   rows = append(rows, string(row))
	}
	gridWords := strings.Join(rows, "")

	for _, word := range words {
		if !strings.Contains(gridWords, word) {
			t.Fatalf("\"%s\" was not placed in %+v", word, newGrid.Letters)
		}
	}
}
