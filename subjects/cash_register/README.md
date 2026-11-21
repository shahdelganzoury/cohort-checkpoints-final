## cashregister

### Instructions

Write a function that takes a nonnegative amount as an `int` and returns the minimum number of coins needed to make that amount as an `int`.

- The denominations are `25`, `10`, `5`, and `1`, and each may be used any number of times.
- The amount is greater than or equal to `0`.
- An amount of `0` returns `0`.
- An amount equal to one denomination returns `1`.
- Use as many `25` coins as possible, then `10`, then `5`, then `1`.

### Expected function

```go
func CashRegister(amount int) int {

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
	fmt.Println(piscine.CashRegister(68))
	fmt.Println(piscine.CashRegister(0))
	fmt.Println(piscine.CashRegister(1))
	fmt.Println(piscine.CashRegister(5))
	fmt.Println(piscine.CashRegister(10))
}
```

And its output:

```console
$ go run . | cat -e
7$
0$
1$
1$
1$
$
```
