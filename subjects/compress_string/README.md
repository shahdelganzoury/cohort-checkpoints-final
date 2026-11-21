## compressstring

### Instructions

Write a function that takes a `string` and returns the same text with each run of consecutive lowercase ASCII letters encoded.

- Write each run as the letter followed by its decimal count, including a count of `1`.
- If the input is empty, the function returns an empty string.

### Expected function

```go
func CompressString(s string) string {

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
	fmt.Println(piscine.CompressString("aaabbcccc"))
	fmt.Println(piscine.CompressString(""))
	fmt.Println(piscine.CompressString("a"))
	fmt.Println(piscine.CompressString("aabbaa"))
}
```

And its output:

```console
$ go run . | cat -e
a3b2c4$
$
a1$
a2b2a2$
$
```
