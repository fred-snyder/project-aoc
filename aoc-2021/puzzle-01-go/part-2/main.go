package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const FILE_PATH = "input.txt"

// const FILE_PATH = "example_input.txt"

func numbers() []int {
	file, err := os.Open(FILE_PATH)

	if err != nil {
		log.Fatalf("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var nums []int

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, v)
	}

	return nums
}

func main() {
	nums := numbers()

	// new summed numbers in windows of 3
	var windows []int

	for i, val := range nums {
		if i == 0 {
			windows = append(windows, val)
			continue
		}

		if i == 1 {
			windows = append(windows, val)
			windows[i-1] += val
		}

		if i >= 2 {
			windows = append(windows, val)
			windows[i-1] += val
			windows[i-2] += val
		}
	}

	// store new slice // remove last two elements
	l := len(windows) - 2
	summedWindows := windows[0:l]

	// storage variables // declare and initialize to zero
	var prev, incCounter int

	for i, val := range summedWindows {
		// catch edge case
		if i == 0 {
			prev = val // assign the first value to prev
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
