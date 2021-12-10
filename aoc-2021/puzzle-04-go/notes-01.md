## Notes part 1

Create a structure to represent a bingo board
Analyse the file get the first line and store as the bingoNumbers

```go
// Perhaps ??
type cell struct {
    id: int
    value: int
    marked: bool // ???

    // perhaps
    row: int
    column: int
}

type board struct {
    // or type board []cells ??
    column []cell //??
    row [] cell // ??
    // properties for every cell?
}

// Some methods ??
func (b board) checkRow() {
    // the check if any of the rows has bingo?
    // or checkCell???
}

```