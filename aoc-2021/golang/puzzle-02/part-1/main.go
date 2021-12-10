package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const FILE = "example_input.txt"

const FILE = "input.txt"

// helper function for error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput() []string {
	// open file and check for errors
	file, err := os.Open(FILE)
	check(err)
	defer file.Close() // defer closing file until function returns

	// store for the lines in a slice
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	return lines
}

func main() {
	data := parseInput()

	var forw, downw, upw int

	for _, val := range data {
		// split the line
		split := strings.Fields(val)

		switch split[0] {
		case "forward":
			int, _ := strconv.Atoi(split[1])
			forw += int
		case "down":
			int, _ := strconv.Atoi(split[1])
			downw += int
		case "up":
			int, _ := strconv.Atoi(split[1])
			upw += int
		default:
			fmt.Println("Default case")
		}
	}

	fmt.Println("Forward: ", forw)
	fmt.Println("Downward: ", downw)
	fmt.Println("Upward: ", upw)
	fmt.Println()
	fmt.Println("Horizontal: ", forw)
	fmt.Println("Depth: ", downw-upw)
	fmt.Println("Multiplied: ", forw*(downw-upw))
}
