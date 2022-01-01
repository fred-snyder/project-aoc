package main

import (
	"bufio"
	"fmt"
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

// commonBitAtIndex loops over lines and
// calculates the most common bit
// for a specific index
func commonBitAtIndex(data []string, index int, most bool) int {

	// count the number of bits
	var bit0, bit1 int

	// loop over all values in the file
	for _, val := range data {
		// first convert to a character
		// then convert to an int
		i, _ := strconv.Atoi(string(val[index]))
		if i == 0 {
			bit0++
		} else if i == 1 {
			bit1++
		}
	}

	// if most common bit
	if most {
		if bit0 < bit1 {
			return 1
		} else {
			return 0
		}

	} else { // if least common bit
		if bit0 < bit1 {
			return 0
		} else {
			return 1
		}
	}
}

// convert binary [ 1 0 1 1 0] to decimal
func intSliceToDec(sl []int) int64 {

	var bin []string

	// val is an int
	for _, val := range sl {
		// create a slice of string so that we can convert it decimal
		bin = append(bin, strconv.Itoa(val))
	}

	// create a string from the slice of string
	binString := strings.Join(bin, "")

	// convert base2 to int64
	i, _ := strconv.ParseInt(binString, 2, 64)
	return i
}

func main() {
	// read data from file
	data, err := fileToLines(FILE)
	check(err)

	var gammaBin, epsilonBin []int

	// create a loop of the length of 1 line
	for i := 0; i < len(data[0]); i++ {
		// calculate the most common bit
		gammaBin = append(gammaBin, commonBitAtIndex(data, i, true))
		// calculate the least common bit
		epsilonBin = append(epsilonBin, commonBitAtIndex(data, i, false))
	}

	// convert binary to decimal
	gamma := intSliceToDec(gammaBin)
	epsilon := intSliceToDec(epsilonBin)

	// output power consumption
	fmt.Println("binary gamma: ", gammaBin, "binary epsilon: ", epsilonBin)
	fmt.Println("power consumption", gamma*epsilon)
}
