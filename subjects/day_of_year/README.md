## dayofyear

### Instructions

Write a function that takes a `day`, a `month`, and a `year` as `int` values and returns the one-based day of that Gregorian year as an `int`.

- `year` is greater than or equal to `1`.
- If the date is invalid, the function returns `-1`.
- A year is a leap year when it is divisible by `400`, or divisible by `4` but not by `100`. February then has `29` days. Otherwise February has `28` days.
- The month lengths are `31`, `28` or `29`, `31`, `30`, `31`, `30`, `31`, `31`, `30`, `31`, `30`, and `31`.
- The date is invalid when `year` is less than `1`, `month` is outside `1` to `12`, `day` is less than `1`, or `day` is past the last day of that month.
- January `1` is day `1`. The result counts every day from the start of the year through the given day.

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
