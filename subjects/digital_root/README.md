## digitalroot

### Instructions

Write a function that takes a nonnegative `int` and returns its digital root as an `int`.

- Add the decimal digits repeatedly until one digit remains.
- `n` is greater than or equal to `0`.
- If `n` is already a single digit, including `0`, the function returns `n` unchanged.
- A sum of `10` or more is summed again, until the total is from `0` through `9`.

### Expected function

```go
func DigitalRoot(n int) int {

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
	fmt.Println(piscine.DigitalRoot(9875))
	fmt.Println(piscine.DigitalRoot(0))
	fmt.Println(piscine.DigitalRoot(9))
	fmt.Println(piscine.DigitalRoot(10))
	fmt.Println(piscine.DigitalRoot(99))
}
```

And its output:

```console
$ go run . | cat -e
2$
0$
9$
1$
9$
$
```
