# Notes part 1

create a graphlike data structure?

```go
// 
type graph struct {
    lines map[int]int
}
```


```go
// or
type lines map[int]int //?
type graph []lines //?
```

1. Parse all lines and check the biggest X and Y values.
   So that you know how big the data structure should be
2. Creata a data structure that lets you mark points in a sort of graph or array/table structure
3. Count the points that have a value higher then 2

## Step 1

```go
// parse the input.txt

// from every line
// create a []]int
// call that a line type
// [X1,Y1,X2,Y2]
// [0,9,5,9]
// [8,0,0,8]

// skip the lines that do not meet the horizontal/vertical criterium
// so: X1 = X2 or Y1 = Y2

// append the lines to a []line
```

## Step 2

Create a multidimensional array.

```go
var maxX, maxY int // the multi-dimensional array size // find the largest values first
var diagram [maxX][maxY]int
// initialize the array with zero values

// loop over all the lines
// and add the points to the diagram array from the starting value up to and including the end value.
// increment the index for every count

```

## Stap 3

Count the number of items in the array that have a value higher then 2.



```go

	// [x1, y1, x2, y2]
	// [0 9 5 9]
	// x1 = 0
	// y1 = 9
	// x2 = 5
	// y2 = 9
	// (0,9) >>> (5,9)
	// (0,9) (1,9) (2,9) (3,9) (4,9) (5,9)

	// [9 4 3 4]
	// loop over the x the values up to and including x2

	//TODO: how to "get" the points between the line-segments?
	// calculate the difference between the points
	// so loop over the numbers in between??
```