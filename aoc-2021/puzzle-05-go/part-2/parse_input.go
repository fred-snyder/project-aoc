package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const FILE = "input.txt"

// const FILE = "example_input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToLines(path string) (lines []string, err error) {
	// open file
	f, err := os.Open(path)
	check(err)
	defer f.Close() // close after function returns

	// we want a *Scanner so that we can scan the content
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// we don't have to explicitly declare lines because
		// we declared lines in the function return blueprint
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	return // same applies for err // we can just return
}

// loop over the string input
// and create the lines
func parseInput() ventLines {
	inputLines, _ := fileToLines(FILE)
	var allLines ventLines

	for _, line := range inputLines {
		var l lineSegm
		var X1, Y1, X2, Y2 int

		// split the line into different parts
		spl1 := strings.Split(line, " -> ")
		spl2_1 := strings.Split(spl1[0], ",")
		spl2_2 := strings.Split(spl1[1], ",")

		// convert byte to string // convert string to int
		X1, _ = strconv.Atoi(spl2_1[0])
		Y1, _ = strconv.Atoi(spl2_1[1])
		X2, _ = strconv.Atoi(spl2_2[0])
		Y2, _ = strconv.Atoi(spl2_2[1])

		// append points to lines
		l = append(l, X1, Y1, X2, Y2)
		allLines = append(allLines, l)
	}

	return allLines
}
