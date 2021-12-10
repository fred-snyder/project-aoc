package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const FILE = "input.txt"

// const FILE = "example_input.txt"

// global variable for lines
// var lines []string
var lines = make(map[int]string)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToLines(path string) (err error) {
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lineCounter := 0
	for scanner.Scan() {
		lines[lineCounter] = scanner.Text()
		lineCounter++
	}

	err = scanner.Err()
	return // same applies for err // we can just return
}

// commonBitAtIndex loops over lines and
// calculates the most common bit
// for a specific line[index]
func commonBitAtIndex(index int, most bool) int {

	// count the number of bits
	var bit0, bit1 int

	// loop over all values in the file
	// and count the number of 1s and 0s
	for _, line := range lines {
		// first convert to a character then convert to an int
		i, _ := strconv.Atoi(string(line[index]))
		if i == 0 {
			bit0++
		} else if i == 1 {
			bit1++
		}
	}

	// if most common bit
	if most {
		if bit1 == bit0 { // if there are the same
			return 1
		} else if bit0 < bit1 {
			return 1
		} else {
			return 0
		}
		// if least common bit
	} else {
		if bit0 == bit1 { // if there are the same
			return 0
		} else if bit0 < bit1 {
			return 0
		} else {
			return 1
		}
	}
}

// convert binary "10110" to decimal
func stringToDec(sl string) int64 {
	// convert base2 to int64
	i, _ := strconv.ParseInt(sl, 2, 64)
	return i
}

// loop over all the lines and check which lines meet the criteria
func processLines(most bool) {
	ln := len(lines[0]) // lenght of 1 line

	// create a loop with iteration count of the length of 1 line
	for i := 0; i < ln; i++ {
		// find the most/least common bit at a specific index
		commbit := commonBitAtIndex(i, most)

		/*
			remove all the lines that do not meet the criterium
			most: keep lines that start with the most
			least: keep starting with the lease
		*/

		// loop over all the items in the map
		for lineI, val := range lines {
			if most {
				// if the the string at the current index is the same as the most common bit
				// remove that line
				if len(lines) > 1 && string(val[i]) == strconv.Itoa(commbit) {
					delete(lines, lineI)
				}
			} else {
				if len(lines) > 1 && string(val[i]) == strconv.Itoa(commbit) {
					delete(lines, lineI)
				}
			}
		}
	}

	fmt.Println("lines after processLines(): ", lines)
}

func lastLineItem(m map[int]string) string {
	var v string
	for _, val := range m {
		v = val
	}
	return v
}

func main() {
	// call the readfile function // store in global variable
	fileToLines(FILE)
	processLines(false) // calculate the least common bit
	ox := lastLineItem(lines)
	oxDec := stringToDec(ox) // convert binary to decimal

	// call the readfile function again to reset the the lines map[]
	fileToLines(FILE)
	processLines(true) // calculate the most common bit
	co2 := lastLineItem(lines)
	co2Dec := stringToDec(co2) // convert binary to decimal

	// output power consumption
	fmt.Println("binary ox: ", ox, "binary co2: ", co2)
	fmt.Println("life support rating", oxDec*co2Dec)
}
