package main

import "fmt"

// test input
var inputFish = []uint8{3, 4, 3, 1, 2}

const day int = 256

// var inputFish = []uint8{1,1,3,5,1,3,2,1,5,3,1,4,4,4,1,1,1,3,1,4,3,1,2,2,2,4,1,1,5,5,4,3,1,1,1,1,1,1,3,4,1,2,2,5,1,3,5,1,3,2,5,2,2,4,1,1,1,4,3,3,3,1,1,1,1,3,1,3,3,4,4,1,1,5,4,2,2,5,4,5,2,5,1,4,2,1,5,5,5,4,3,1,1,4,1,1,3,1,3,4,1,1,2,4,2,1,1,2,3,1,1,1,4,1,3,5,5,5,5,1,2,2,1,3,1,2,5,1,4,4,5,5,4,1,1,3,3,1,5,1,1,4,1,3,3,2,4,2,4,1,5,5,1,2,5,1,5,4,3,1,1,1,5,4,1,1,4,1,2,3,1,3,5,1,1,1,2,4,5,5,5,4,1,4,1,4,1,1,1,1,1,5,2,1,1,1,1,2,3,1,4,5,5,2,4,1,5,1,3,1,4,1,1,1,4,2,3,2,3,1,5,2,1,1,4,2,1,1,5,1,4,1,1,5,5,4,3,5,1,4,3,4,4,5,1,1,1,2,1,1,2,1,1,3,2,4,5,3,5,1,2,2,2,5,1,2,5,3,5,1,1,4,5,2,1,4,1,5,2,1,1,2,5,4,1,3,5,3,1,1,3,1,4,4,2,2,4,3,1,1}

// instead of creating 1 giant array
// and modifying that array
// you can also process the offspring of 1 fish at the time
// calculate the total offspring of that generations
// and then sum everything together

func calcOffspring(c uint8) int {
	// store first gen and all offspring
	var offspring []uint8
	offspring = append(offspring, c)

	// start the time
	for i := 0; i < day; i++ {
		for index, cyc := range offspring {
			if cyc != 0 {
				offspring[index]--
			} else if cyc == 0 {
				offspring[index] = 6
				offspring = append(offspring, 8)
			}
		}
	}

	return len(offspring)
}

func main() {

	var totalOffspring int
	for _, cycle := range inputFish {
		offspr := calcOffspring(cycle)
		totalOffspring += offspr
		// fmt.Println(day)
		// fmt.Printf("Cycle: %d, Offspring: %d \n", cycle, totalOffspring)
	}

	// test 256 days with just 1 generation
	//! fatal error: out of memory
	fmt.Println(calcOffspring(3))

	fmt.Println(totalOffspring)
}
