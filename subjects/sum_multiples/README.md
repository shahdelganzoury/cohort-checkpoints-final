## summultiples

### Instructions

Write a function that takes two `int` values, `n` and `factor`, and returns the sum of all the multiples of `factor` from `1` up to and including `n`.

- If `factor` is `0` or negative, the function returns `0`.
- If `n` is `0` or negative, the function returns `0`.
- If `factor` is bigger than `n`, the function returns `0`.

### Expected function

```go
func SumMultiples(n int, factor int) int {

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
	fmt.Println(piscine.SumMultiples(10, 3))
	fmt.Println(piscine.SumMultiples(10, 5))
	fmt.Println(piscine.SumMultiples(20, 4))
	fmt.Println(piscine.SumMultiples(5, 10))
	fmt.Println(piscine.SumMultiples(10, 0))
	fmt.Println(piscine.SumMultiples(0, 3))
}
```

And its output:

```console
$ go run . | cat -e
18$
15$
60$
0$
0$
0$
$
```
