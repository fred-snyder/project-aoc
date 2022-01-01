package main

import (
	"strconv"
	"strings"
)

// create a new board from a slice of input strings
func createBoard(bNum []string) *board {
	// create a new board
	b := new(board)
	// create the row index counter
	var rowIndex, colIndex int

	// loop over all the board input numbers
	for i, val := range bNum {
		// create a new cell
		cell := new(cell)
		// set the correct properties of the cell
		cell.rowIndex = rowIndex
		cell.colIndex = colIndex
		cell.value, _ = strconv.Atoi(val) // convert string to int
		cell.marked = false

		// add the cell to the board
		b.cells = append(b.cells, cell)

		// increase column index
		// and after 5 indexes go to the next row
		if i > 0 && (i+1)%5 == 0 {
			rowIndex++
			colIndex = 0
		} else {
			colIndex++
		}
	}

	return b // return the board
}

// create a map of all the boards
func createAllBoards(f string) map[int]*board {
	var allBoards = make(map[int]*board)
	var bCount int
	var b []string

	// separate all the numbers in f into a []string (strings.Fields)
	// create multiple separate board slices
	for i, val := range strings.Fields(f) {

		b = append(b, val)

		// there are 25 numbers in a board
		if i > 0 && (i+1)%25 == 0 {
			// append the newboard to the allBoards map
			allBoards[bCount] = createBoard(b)
			b = nil
			bCount++
		}
	}

	return allBoards
}
