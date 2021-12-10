# Notes part 2

AIM
HORIZONTAL
DEPTH

Down X
- increase AIM by X units

Up X
- decrease AIM by X units

Forward X
- increase HORIZONTAL by X
- increases DEPTH by current AIM*X


--------------------------------

X
AIM = 0
HORIZONTAL = 0
DEPTH = AIM * X

up/down >> only changes AIM
down >> ADDS
up >> SUBTRACTS

```
forward 5 // X=5 >> HOR=5 DEPTH=0*5
down 5 // X=5 >> AIM=5
forward 8 // X=8 >> HOR=13 DEPTH=5*8=40
up 3 // X=3 >> AIM=2
down 8 // X=8 >> AIM=10
forward 2 // X=2 >> HOR=15 DEPTH=2*10=20 >> 60
```


----------------------------------

```go
AIM
HORIZ
DEPTH

if "up" {
    AIM -= i // subtract
}

if "down" {
    AIM += i // add
}

if "forward" {
    HORIZ += i
    DEPTH += AIM * i
}

fmt.Println(HORIZ * DEPTH)
```
