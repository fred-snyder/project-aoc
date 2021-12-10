package main

import (
	"fmt"
)

// define a line segment type
// [x1, y1, x2, y2]
type lineSegm []int

// define a structure for the vent lines
// an array of line segments
type ventLines []lineSegm

// define a type for the diagram
// a multidimensional array or basically a huge table
// the first dimension of the diagram corresponds with the y-value
// the second (inner) dimension corresponds with the x-value
type diagram [991][991]int

// compute the maximum size for the diagram
func diagramSize(l ventLines) (int, int) {
	var x int
	var y int

	// loop over all the ventLines
	for _, v := range l {
		// then loop over all the line segments
		for i, val := range v {
			// first compute max X value
			if i == 0 || i == 2 {
				if val > x {
					x = val
				}
			}
			// then compute max Y value
			if i == 1 || i == 3 {
				if val > y {
					y = val
				}
			}
		}
	}

	// add 1 to the size to accomodate for index 0
	fmt.Println("Array dimension 1:", x+1, "Array dimension 2:", y+1)

	return x, y
}

// compute all the line segment points
// and add them to the array/diagram
func addPointsToDiagram(x1, y1, x2, y2 int, p *diagram) {
	// example: [0 9 5 9] [x1 y1 x2 y2]
	// (0,9) >>> (5,9)
	// (0,9) (1,9) (2,9) (3,9) (4,9) (5,9)

	// swap values so that x1 or y1 is always the lowest value
	// otherwise the loop below doesnt process: [9 4 3 4]
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	// add the y values
	if x1 == x2 {
		for i := y1; i <= y2; i++ {
			// fmt.Println("y", i)
			p[i][x1] += 1
		}
	}

	// add the x values
	if y1 == y2 {
		for i := x1; i <= x2; i++ {
			// fmt.Println("x", i)
			p[y1][i] += 1
		}
	}
}

// analyse the ventLines and update the diagram
// loop over all the filtered input lines
// and add 1 to the diagram for all the individual line points
func createDiagram(l ventLines, p *diagram) {
	// [[0 9 5 9] [9 4 3 4] [2 2 2 1] [7 0 7 4] [0 9 2 9] [3 4 1 4]]

	// loop over all the lineSegments in the ventLines
	for _, lineSegm := range l {
		var x1, y1, x2, y2 int

		x1 = lineSegm[0]
		y1 = lineSegm[1]
		x2 = lineSegm[2]
		y2 = lineSegm[3]

		// fmt.Println(x1, y1, x2, y2)
		addPointsToDiagram(x1, y1, x2, y2, p)
	}
}

// compute the number of points that overlap
func computeOverlap(d *diagram) int {
	var overlap int
	for _, y := range d {
		for _, x := range y {
			if x > 1 {
				overlap++
			}
		}
	}
	return overlap
}

func printDiagram(d *diagram) {

	// print the vents on different lines
	for i, v := range d {
		if i == 0 {
			fmt.Println() // add empty line
		}
		fmt.Println(v)
	}
}

func main() {
	// store all the parsed input strings
	filteredLines := parseInput()
	fmt.Println(filteredLines)

	// call diagramSize to compute the size of diagram type
	// set the diagram type to the corresponding x and y values
	diagramSize(filteredLines)

	// create a new *diagram and initialize with zero values
	var vents = new(diagram)
	// analyse the line segments and add the points to the diagram
	createDiagram(filteredLines, vents)

	// print diagram
	// printDiagram(vents)

	overlapPoints := computeOverlap(vents)
	fmt.Println("Overlapping points", overlapPoints)

}
