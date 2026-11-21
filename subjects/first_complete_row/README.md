## firstcompleterow

### Instructions

Write a function that takes a nonempty rectangular grid of distinct `int` values and a sequence of distinct called `int` values, and returns the 1-based index of the call that first completes a row.

- If no row is fully called, the function returns `-1`.
- A called value may be absent from the grid.
- The index of the first call is `1`.
- After each call, check every row. Return the index of the call that completes one. If several rows finish on that same call, return that index.
- A row is complete only when every value in it has been called.
- If `called` is empty or `nil`, the function returns `-1`.
- A call that is not in the grid fills no cell, but it still counts toward the index.

### Expected function

```go
func FirstCompleteRow(grid [][]int, called []int) int {

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
	fmt.Println(piscine.FirstCompleteRow([][]int{{4, 1, 7}, {3, 9, 2}, {8, 5, 6}}, []int{9, 4, 1, 3, 7, 2}))
	fmt.Println(piscine.FirstCompleteRow([][]int{{1}}, []int{1}))
	fmt.Println(piscine.FirstCompleteRow([][]int{{1}}, []int{2}))
	fmt.Println(piscine.FirstCompleteRow([][]int{{1, 2}, {3, 4}}, []int{1, 3, 2}))
}
```

And its output:

```console
$ go run . | cat -e
5$
1$
-1$
3$
$
```
