## buscapacity

### Instructions

Write a function that takes a bus `capacity` as an `int` and the stops as a slice of `[2]int`, and returns the maximum occupancy as an `int`.

- Each stop is a pair `{boarding, leaving}`. Passengers leave before boarding, and the bus starts empty.
- If `capacity` is negative, the function returns `-1`.
- The function returns `-1` when a count is negative, more passengers leave than are on the bus, or boarding would exceed `capacity`.
- Filling the bus exactly to `capacity`, or leaving every passenger currently on board, is valid.
- If there are no stops, including a `nil` slice, and `capacity` is greater than or equal to `0`, the function returns `0`.
- The result is the highest number of passengers on the bus after any stop.

### Expected function

```go
func BusCapacity(capacity int, stops [][2]int) int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BusCapacity(20, [][2]int{{10, 0}, {8, 3}, {2, 5}, {4, 1}}))
	fmt.Println(piscine.BusCapacity(10, [][2]int{{10, 0}}))
	fmt.Println(piscine.BusCapacity(10, [][2]int{{11, 0}}))
	fmt.Println(piscine.BusCapacity(10, [][2]int{{0, 1}}))
}
```

And its output:

```console
$ go run . | cat -e
15$
10$
-1$
-1$
$
```
