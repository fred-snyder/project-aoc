# Notes part 2

Add another function that specifically checks for diagonal lines.
Or intergrate the logic in the addPoints function

And modify the filter so that diagonal lines do not get filtered out.

First just process diagonal lines. So that you know what the logic should look like.

```go
// Diagonal is deltaX == deltaY
x1 - x2 == y1 - y2
// also consider that the x1 > x2 and y1 > y2
// if not swap the values
```


```
// x1,y1       x2,y2
// (9,7) (8,8) (7,9)

// d[y][x]
// d[7][9]
// d[8][8]
// d[9][7]
// x1 is bigger then x2
```

## Problem while solving part 2

*The loops that add points to the diagram are to complicated.*
- Possible solution: create a function that adds the correct points for every linesegment.
- Instead of 4 separate loops: try to create 1 less complex loop that catches all the different line directions.

Old code:

```go
    // diagonal is deltaX == deltaY
    // filter input, only keep horizontal, vertical or diagonal lines
    if (X1 > X2 && (X1-X2 == Y1-Y2) && X1-X2 > 0) || (X2 < X1 && X2-X1 == Y2-Y1 && X2-X1 > 0) || X1 == X2 || Y1 == Y2 {
        l = append(l, X1, Y1, X2, Y2)
        allLines = append(allLines, l)
    }
```

```go
	// example: [0 9 5 9] [x1 y1 x2 y2]
	// (0,9) >>> (5,9)
	// (0,9) (1,9) (2,9) (3,9) (4,9) (5,9)

	// swap values so that x1 or y1 is always the lowest value
	// otherwise the loop below doesnt process: [9 4 3 4]
	// first check if the lines are not diagonal (dX=dY)
	// then swap the values // because we only want to swap the values for straight lines
	if x1-x2 != y1-y2 || x2-x1 != y2-y1 {
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
	}

	// add the diagonal lines from top right to bottom left
	if x1-x2 == y1-y2 {
		fmt.Println("Diagonal lines from left to right??")
		fmt.Println(x1, y1, x2, y2)
		_y2 := y2 // temp store y

		for i := x1; i >= x2; i-- {
			d[_y2][i] += 1
			_y2++ //
		}
	}

	// add the inverse direction diagonal lines
	//TODO: finish these points
	//TODO: currently doesnt work
	//OPTIONALLY: rethink the way these loops work
	//FOR EXAMPLE: create a new function that calculates all inbetween points for every situation
	if x2-x1 == y2-y1 {
		fmt.Println("Diagonal reverse lines")
		fmt.Println(x1, y1, x2, y2)
		_y1 := y1 // temp store y

		for i := x2; i >= x1; i++ {
			d[_y1][i] += 1
			_y1++ //
		}
	}

	// add the y values for hor/vert lines
	if x1 == x2 {
		for i := y1; i <= y2; i++ {
			// fmt.Println("y", i)
			d[i][x1] += 1
		}
	}

	// add the x values for hor/vert lines
	if y1 == y2 {
		for i := x1; i <= x2; i++ {
			// fmt.Println("x", i)
			d[y1][i] += 1
		}
	}
```