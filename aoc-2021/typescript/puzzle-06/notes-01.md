# Notes part 1

lanternfish
must spawn quickly to reach such large numbers
You should model their growth rate to be sure.
maybe exponentially quickly? 
each lanternfish creates a new lanternfish once every 7 days.
**cycles are not synchronized between every lanternfish**

you can model each fish as a single number that represents the number of days until it creates a new lanternfish.

a new lanternfish would surely needs slightly longer before it's capable of producing more lanternfish: two more days for its first cycle.

example with a val of 3:

```plain
increase day +1 > internal timer -1 >>> 2
increase day +1 > internal timer -1 >>> 1
increase day +1 > internal timer -1 >>> 0
if timer == 0 {
    FISH 1: increase day +1 > internal timer -1 >>> 6
    createFish(8)
    FISH 2: add new fish internal timer of 8.
}

FISH 1    
    increase day +1 > internal timer -1 >>> 5
FISH 2
    increase day +1 > internal timer -1 >>> 7
```

## Example input

```plain
Initial state: 3,4,3,1,2
After  1 day:  2,3,2,0,1
After  2 days: 1,2,1,6,0,8
After 17 days: 0,1,0,5,6,0,1,2,2,3,0,1,2,2,2,3,3,4,4,5,7,8
After 18 days: 6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8

In this example, after 18 days, there are a total of 26 fish. After 80 days, there would be a total of 5934.
```

## Logic sketch

```js
// starting amount
let allFish = [];

let day = 0

const fish1 = {
    internalTimer: 3

    decreaseTimer() {
        if (this.internalTimer === 0) {
            this.internalTimer = 6
            allFish.append(newFish(8))
        } else {
            this.internalTimer = this.internalTimer - 1
        }
    }
}

function startCycles(80) {
    // loop
}

len(allFish)
```
