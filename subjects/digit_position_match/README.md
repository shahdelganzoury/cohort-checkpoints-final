## digitpositionmatch

### Instructions

Write a function that takes two equal-length strings of decimal digits and returns, as an `int`, how many indexes hold the same character.

- Leading zeroes are meaningful.
- Compare the characters at each index. Do not convert the strings to numbers.
- Two empty strings return `0`.
- Every matching index counts, including zeroes.

### Expected function

```go
func DigitPositionMatch(a, b string) int {

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
	fmt.Println(piscine.DigitPositionMatch("58342", "57382"))
	fmt.Println(piscine.DigitPositionMatch("", ""))
	fmt.Println(piscine.DigitPositionMatch("000", "000"))
	fmt.Println(piscine.DigitPositionMatch("123", "456"))
}
```

And its output:

```console
$ go run . | cat -e
3$
0$
3$
0$
$
```
