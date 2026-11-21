## twomissingnumbers

### Instructions

Write a function that takes a slice of distinct `int` values and returns the two missing numbers in ascending order.

- The input contains distinct integers from `1` through `len(nums)+2`, with exactly two numbers missing.
- The inputs avoid integer overflow.

### Expected function

```go
func TwoMissingNumbers(nums []int) (int, int) {

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
	fmt.Println(piscine.TwoMissingNumbers([]int{1, 3, 6, 2}))
	fmt.Println(piscine.TwoMissingNumbers(nil))
	fmt.Println(piscine.TwoMissingNumbers([]int{2, 4, 5, 6}))
	fmt.Println(piscine.TwoMissingNumbers([]int{2, 3}))
}
```

And its output:

```console
$ go run . | cat -e
4 5$
1 2$
1 3$
1 4$
$
```
