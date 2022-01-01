package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const FILE_PATH = "input.txt"

func main() {
	// os.Open() opens specific file in read-only mode
	// and this return a pointer of type os.
	file, err := os.Open(FILE_PATH)

	if err != nil {
		log.Fatalf("failed to open file")
	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan() method.
	scanner.Split(bufio.ScanLines)
	var nums []int

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, v)
	}

	// storage variables // declare and initialize to zero
	var prev, incCounter int

	for i, val := range nums {
		if i == 0 {
			prev = val // assign the first value to the first prev
			continue   // and jump to the next iteration of the loop
		}

		if val > prev {
			incCounter++
		}

		prev = val
	}

	fmt.Println(nums)
	fmt.Println("Number of times increased: ", incCounter)
}
