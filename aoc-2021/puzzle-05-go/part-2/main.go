package main

import (
	"fmt"
	"math"
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
// type diagram [10][10]int

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

// filter out the horizontal, vertical and diagonal lines
func filterLineSegm(x1, y1, x2, y2 int) bool {
	if x1 == x2 || y1 == y2 {
		fmt.Println("Hor or vert")
		return true
	}

	// compute deltaX/deltaY, convert to float64 and get the absolute values
	dX := float64(x1 - x2)
	dY := float64(y1 - y2)
	if math.Abs(dX) == math.Abs(dY) {
		fmt.Println("Diagonal")
		return true
	} else {
		return false
		// make sure to add and else return statement
		// otherwise Go complains about missing return values
	}
}

// compute all the line segment points
// and add them to the array/diagram
func addPointsToDiagram(x1, y1, x2, y2 int, d *diagram) {

	// a multiplier for the direction // -1 or +1 // default 0
	var xDir int
	var yDir int

	if x1 > x2 {
		xDir = -1
	} else if x1 == x2 {
		xDir = 0
	} else {
		xDir = 1
	}

	if y1 > y2 {
		yDir = -1
	} else if y1 == y2 {
		yDir = 0
	} else {
		yDir = 1
	}

	// fmt.Println("Points:", x1, y1, x2, y2)
	// fmt.Println("Directions:", xDir, yDir)

	for curX, curY := x1, y1; curX != x2 || curY != y2; {
		// fmt.Println("current: ", curX, curY)
		d[curY][curX] += 1
		curX += xDir
		curY += yDir

		// last point edge case
		if curX == x2 && curY == y2 {
			d[y2][x2] += 1
		}
	}
}

// analyse the ventLines and update the diagram
// loop over all the filtered input lines
// and add 1 to the diagram for all the individual line points
func createDiagram(l ventLines, d *diagram) {
	// [[0 9 5 9] [9 4 3 4] [2 2 2 1] [7 0 7 4] [0 9 2 9] [3 4 1 4]]

	// loop over all the lineSegments in the ventLines
	for _, lineSegm := range l {
		var x1, y1, x2, y2 int

		x1 = lineSegm[0]
		y1 = lineSegm[1]
		x2 = lineSegm[2]
		y2 = lineSegm[3]

		// check if lineSegm meets the criterium
		if filterLineSegm(x1, y1, x2, y2) {
			// fmt.Println(x1, y1, x2, y2)
			addPointsToDiagram(x1, y1, x2, y2, d)
		}
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

func printDiagram(d *diagram, enable bool) {
	// check for enable // so that the Go compiler doesnt warn about unused functions
	if enable {
		// print the vents on different lines
		for i, v := range d {
			if i == 0 {
				fmt.Println() // add empty line
			}
			fmt.Println(v)
		}
	}
}

func main() {
	// store all the parsed input strings
	filteredLines := parseInput()
	// fmt.Println(filteredLines)

	// call diagramSize to compute the size of diagram type
	// set the diagram type to the corresponding x and y values
	diagramSize(filteredLines)

	// create a new *diagram and initialize with zero values
	var vents = new(diagram)
	// analyse the line segments and add the points to the diagram
	createDiagram(filteredLines, vents)

	// print diagram
	// printDiagram(vents, false) // set to true to enable the function
	printDiagram(vents, true) // set to true to enable the function

	overlapPoints := computeOverlap(vents)
	fmt.Println("Overlapping points", overlapPoints)

}
