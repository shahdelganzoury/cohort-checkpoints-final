## classifystatuscode

### Instructions

Write a function that takes an HTTP status code as an `int` and returns a `string` describing its category.

- `200` to `299` returns `SUCCESS`.
- `400` to `499` returns `CLIENT_ERROR`.
- `500` to `599` returns `SERVER_ERROR`.
- Any other value returns `UNKNOWN`. This includes `1xx`, `3xx` redirects, negative numbers, and `600` or above.

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
SUCCESS$
UNKNOWN$
CLIENT_ERROR$
SERVER_ERROR$
UNKNOWN$
UNKNOWN$
```