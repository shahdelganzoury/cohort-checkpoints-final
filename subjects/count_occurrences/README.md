## countoccurrences

### Instructions

Write a function that takes a slice of `int` and a `target` value, and returns how many times `target` appears in the slice.

- If the slice is empty or `nil`, the function returns `0`.
- If `target` does not appear, the function returns `0`.

### Expected function

```go
func CountOccurrences(nums []int, target int) int {

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
	fmt.Println(piscine.CountOccurrences([]int{1, 2, 3, 2, 2}, 2))
	fmt.Println(piscine.CountOccurrences([]int{1, 2, 3}, 4))
	fmt.Println(piscine.CountOccurrences([]int{5, 5, 5}, 5))
	fmt.Println(piscine.CountOccurrences([]int{-1, 0, -1}, -1))
	fmt.Println(piscine.CountOccurrences(nil, 1))
}
```

And its output:

```console
$ go run . | cat -e
3$
0$
3$
2$
0$
$
```
