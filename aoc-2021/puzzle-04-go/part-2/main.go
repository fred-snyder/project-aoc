package main

import (
	"fmt"
)

// define the properties of a board cell
type cell struct {
	rowIndex int
	colIndex int
	value    int
	marked   bool
}

// all the cells of a board
// plus properties about the state of a board
type board struct {
	cells []*cell
	win   bool
}

// check all the cells of all the rows
// and see if a board has won
func (b *board) checkWin(num int) {
	// count the row indexes that have a marked cell
	var row, column [5]int

	// for every cell
	// if marked store the row and column number
	for _, cell := range b.cells {
		if cell.marked {
			row[cell.rowIndex]++
			column[cell.colIndex]++
		}
	}

	// check if any of the rows has 5 marked cells
	for _, val := range row {
		if val == 5 {
			b.win = true
		}
	}

	// check if any of the columns has 5 marked cells
	for _, val := range column {
		if val == 5 {
			b.win = true
		}
	}
}

// sum all the unmarked values/cells of a board
func (b *board) sumUnmarkedCells() {
	var sum int
	for _, cell := range b.cells {
		if !cell.marked {
			sum += cell.value
		}
	}
	fmt.Println("Sum of b *board:", sum, &b)
}

// for the current game num
// check if a board has a row or a column with 5 marked cells
func (b *board) checkBoard(num int) bool {
	var winningBoard bool

	// update the marked status of the cells
	b.markCells(num)
	b.checkWin(num)

	if b.win {
		winningBoard = true
		b.sumUnmarkedCells()
		fmt.Println("GameNum: ", num)
	}
	return winningBoard
}

// check every cell if equel to input num (gameNumber)
func (b *board) markCells(num int) {
	for _, cell := range b.cells {
		if cell.value == num {
			cell.marked = true
		}
	}
}

/*
func (b *board) printBoard() {
	for _, cell := range b.cells {
		// for readability add + 1 to the row/col indexes
		fmt.Println("Cell row: ", cell.rowIndex+1, "Cell column: ", cell.colIndex+1)
		fmt.Println("Cell value: ", cell.value, "Cell state: ", cell.marked)
		fmt.Println()
	}
}
*/

// a function that checks all the boards
// for every gameNumber
func checkAllBoards(num int, index int, b map[int]*board) {
	for _, board := range b {
		if !board.win {
			board.checkBoard(num)
		}
	}
}

func main() {
	// create all the boards from the giant raw string literal
	boards := numbers()
	allBoards := createAllBoards(boards)

	// var gameNumbers = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	var gameNumbers = []int{0, 56, 39, 4, 52, 7, 73, 57, 65, 13, 3, 72, 69, 96, 18, 9, 49, 83, 24, 31, 12, 64, 29, 21, 80, 71, 66, 95, 2, 62, 68, 46, 11, 33, 74, 88, 17, 15, 5, 6, 98, 30, 51, 78, 76, 75, 28, 53, 87, 48, 20, 22, 55, 86, 82, 90, 47, 19, 25, 1, 27, 60, 94, 38, 97, 58, 70, 10, 43, 40, 89, 26, 34, 32, 23, 45, 50, 91, 61, 44, 35, 85, 63, 16, 99, 92, 8, 36, 81, 84, 79, 37, 93, 67, 59, 54, 41, 77, 42, 14}

	// start the game
	for i, num := range gameNumbers {
		checkAllBoards(num, i, allBoards)
	}

	// manually multiply the gameNum with the Println output from the call to board.sumUnmarkedCells()
}

//TODO: Figure out which board will win last. Once it wins, what would its final score be?
