## passwordstrength

### Instructions

Write a function that takes a password as a `string` and returns its strength as an `int` from `0` to `4`.

Score one point for each of the following:

- The length is greater than or equal to `8`.
- The password contains an uppercase ASCII letter from `A` to `Z`.
- The password contains a digit from `0` to `9`.
- The password contains a special character from `!@#$%`.

### Expected function

```go
func PasswordStrength(s string) int {

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
	fmt.Println(piscine.PasswordStrength("Go2026!"))
	fmt.Println(piscine.PasswordStrength(""))
	fmt.Println(piscine.PasswordStrength("abc"))
	fmt.Println(piscine.PasswordStrength("GoLang2026!"))
}
```

And its output:

```console
$ go run . | cat -e
3$
0$
0$
4$
$
```
