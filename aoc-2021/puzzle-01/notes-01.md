## Solution notes

- There are 2000 lines in the input.txt file
- For each line
- Compare the value with the value of the previous line
- If the value is lower > "decreased"
- If the value is higher > "increased"

Count the number of times the value "increased"

How to loop over a text file?
https://www.geeksforgeeks.org/how-to-read-a-file-line-by-line-to-string-in-golang/

```go
package main

// store count of "increased"
var incCounter int
// read the values from the .txt file

func main() {
    var prev string
    // parse the lines
    // read line 1
    // store value of line in prev
    if currentLine > prev {
        incCounter++
    }

    fmt.Println(incCounter)

}

```