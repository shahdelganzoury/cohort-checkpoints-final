## dayofyear

### Instructions

Write a function that takes a `day`, a `month`, and a `year` as `int` values and returns the one-based day of that Gregorian year as an `int`.

- `year` is greater than or equal to `1`.
- If the date is invalid, the function returns `-1`.

### Expected function

```go
func DayOfYear(day, month, year int) int {

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
	fmt.Println(piscine.DayOfYear(1, 3, 2024))
	fmt.Println(piscine.DayOfYear(1, 3, 2025))
	fmt.Println(piscine.DayOfYear(29, 2, 2025))
	fmt.Println(piscine.DayOfYear(29, 2, 2024))
}
```

And its output:

```console
$ go run . | cat -e
61$
60$
-1$
60$
$
```
