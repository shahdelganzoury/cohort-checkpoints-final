## celsiustofahrenheit

### Instructions

Write a function that takes a temperature in degrees Celsius as a `float64` and returns the same temperature in degrees Fahrenheit as a `float64`.

- Use the formula: `F = C * 9 / 5 + 32`.
- Do not round the result.

### Expected function

```go
func CelsiusToFahrenheit(c float64) float64 {

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
	fmt.Println(piscine.CelsiusToFahrenheit(0))
	fmt.Println(piscine.CelsiusToFahrenheit(100))
	fmt.Println(piscine.CelsiusToFahrenheit(-40))
	fmt.Println(piscine.CelsiusToFahrenheit(25))
	fmt.Println(piscine.CelsiusToFahrenheit(37.5))
}
```

And its output:

```console
$ go run . | cat -e
32$
212$
-40$
77$
99.5$
$
```
