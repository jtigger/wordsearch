package grid

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"strings"
)

type Grid struct {
	Letters   [][]rune
	Solutions map[string]*Solution
}

func (g *Grid) Width() int {
	return len(g.Letters[0])
}
func (g *Grid) Height() int {
	return len(g.Letters)
}

func (g *Grid) writeWord(x int, y int, word string) {
	for i, char := range word {
		g.Letters[y][x+i] = char
	}
	g.Solutions[word] = &Solution{
		Word:      word,
		Row:       y,
		Column:    x,
		Direction: Directions[2],
	}
}

type Solution struct {
	Word        string
	Row, Column int
	Direction   Direction
}

type Direction struct {
	Name string
}

var Directions = []Direction{
	Direction{"up"},
	Direction{"diagonally up and right"},
	Direction{"right"},
	Direction{"diagonally up and left"},
	Direction{"down"},
	Direction{"diagonally down and left"},
	Direction{"left"},
	Direction{"diagonally up and left"},
}

func NewGrid(letters [][]rune) Grid {
	newGrid := Grid{
		Letters: letters,
	}
	newGrid.Solutions = make(map[string]*Solution)
	return newGrid
}

func Generate(wordstream io.Reader) (Grid, error) {
	words, err := parseWords(wordstream)
	if err != nil {
		return NewGrid(nil), err
	}
	size := calcSmallestFittingSquare(words)
	letters := buildLetters(size)
	newGrid := NewGrid(letters)
	for _, word := range words {
		placeWord(&newGrid, word)
	}
	fillEmptyLetters(&newGrid)
	return newGrid, nil
}

func fillEmptyLetters(grid *Grid) {
	for _, row := range grid.Letters {
		for i := range row {
			if row[i] == 0 {
				row[i] = 'a' + rune(rand.Intn(25))
			}
		}
	}
}

func placeWord(g *Grid, word string) {
	var x, y int

	for i := 0; i < 500; i++ {
		x = rand.Intn(g.Width())
		y = rand.Intn(g.Height())

		if collides(g, word, x, y) {
			continue
		}
		g.writeWord(x, y, word)
		return
	}
	panic(fmt.Sprintf("Failed to place \"%s\" in %+v after 500 attempts", word, g))
}

func collides(g *Grid, word string, x, y int) bool {
	if x+len(word) > g.Width() {
		return true
	}
	for idx := x; idx < x+len(word); idx++ {
		if g.Letters[y][idx] != 0 {
			return true
		}
	}
	return false
}

func calcSmallestFittingSquare(words []string) int {
	var longest int
	for _, word := range words {
		longest = max(longest, len(word))
	}

	// ensure total number of letters will fit in the space
	numRunes := countRunes(words)
	minimumSquare := int(math.Ceil(math.Sqrt(float64(numRunes))))

	return max(longest, minimumSquare)
}

func countRunes(words []string) (size int) {
	for _, word := range words {
		size += len(word)
	}
	return
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func parseWords(wordstream io.Reader) ([]string, error) {
	bytes, err := ioutil.ReadAll(wordstream)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes), ","), nil
}

func buildLetters(size int) [][]rune {
	letters := make([][]rune, size)

	for i := range letters {
		letters[i] = make([]rune, size)
	}
	return letters
}
