## classifystatuscode

### Instructions

Write a function that takes an HTTP status code as an `int` and returns a `string` describing its category.

- `100` to `199` returns `Informational`.
- `200` to `299` returns `Success`.
- `300` to `399` returns `Redirection`.
- `400` to `499` returns `Client Error`.
- `500` to `599` returns `Server Error`.
- Any other value returns `Unknown`.

### Expected function

```go
func ClassifyStatusCode(code int) string {

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
	fmt.Println(piscine.ClassifyStatusCode(200))
	fmt.Println(piscine.ClassifyStatusCode(301))
	fmt.Println(piscine.ClassifyStatusCode(404))
	fmt.Println(piscine.ClassifyStatusCode(503))
	fmt.Println(piscine.ClassifyStatusCode(99))
	fmt.Println(piscine.ClassifyStatusCode(600))
}
```

And its output:

```console
$ go run . | cat -e
Success$
Redirection$
Client Error$
Server Error$
Unknown$
Unknown$
$
```
