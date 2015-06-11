package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"flag"
)

var (
	n          = flag.Int("n", 8, "n queens, in an nxn grid")
	CantBeDone = errors.New("Can't be done")
)

type Board struct {
	grid   [][]bool
	queens []Location
}

type Location struct {
	X int
	Y int
}

func main() {
	flag.Parse()

	board := createBoard(*n)

	err := board.PlaceQueens(*n)
	fmt.Println(board.String())
	if err != nil {
		log.Fatalln("What the actual fuck:", err)
	}

}

// Simple, but used here
func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

// Returns an nxn grid off booleans
// The boolean is true if there's a queen placed there
func createBoard(n int) Board {
	grid := make([][]bool, n)
	for i, _ := range grid {
		row := make([]bool, n)
		grid[i] = row
	}

	locations := make([]Location, 0)

	return Board{
		grid:   grid,
		queens: locations,
	}
}

// Attempts to place n queens on the board.
// Returns an error if it can't be done
func (b *Board) PlaceQueens(n int) error {
	// check for freeness
	if n <= 0 {
		return nil
	}

	// Since len(queens) == len(grid), and they can't share rows,
	// we'll just assign it one by it's number
	y := n - 1
	row := b.grid[y]

placements:
	for x := range row {

		// Already a queen here
		if row[x] {
			continue placements
		}

		// Check if this location violates the rules
		for _, queen := range b.queens {
			// See if the x or y match (same row/column), or diagonals (differences of x and y match)
			if x == queen.X || y == queen.Y || abs(x-queen.X) == abs(y-queen.Y) {
				continue placements
			}
		}

		// Place the queen
		row[x] = true
		b.queens = append(b.queens, Location{X: x, Y: y})

		// Recurse!
		err := b.PlaceQueens(n - 1)
		// If we couldn't place the queens
		if err != nil {
			// Remove our queen
			row[x] = false
			b.queens = b.queens[:len(b.queens)-1]

			// Continue to next option
			continue placements
		}

		// We did place the queens, we did!
		return nil
	}

	// Couldn't find a place for the queen :(
	return CantBeDone
}

// For pretty printing the board
func (b *Board) String() string {
	toPrint := make([]string, len(b.grid)+2)
	toPrint[0] = strings.Repeat("_", len(b.grid)+2)
	toPrint[len(toPrint)-1] = strings.Repeat("-", len(b.grid)+2)
	for y, row := range b.grid {
		rowText := make([]string, len(row)+2)
		rowText[0] = "|"
		rowText[len(rowText)-1] = "|"

		for x, element := range row {
			if element {
				rowText[x+1] = "Q"
			} else {
				rowText[x+1] = "x"
			}
		}

		toPrint[y+1] = strings.Join(rowText, "")
	}
	return strings.Join(toPrint, "\n")
}
