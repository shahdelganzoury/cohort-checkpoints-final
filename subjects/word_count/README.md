## wordcount

### Instructions

Write a function that takes a slice of `string` and returns a `map[string]int` where each key is a word and each value is the number of times that word appears in the slice.

- Matching is case-sensitive: `go` and `Go` are different words.
- If the slice is empty or `nil`, the function returns an empty map (not `nil`).

### Expected function

```go
func WordCount(words []string) map[string]int {

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
	fmt.Println(piscine.WordCount([]string{"a", "b", "a"}))
	fmt.Println(piscine.WordCount([]string{"go", "Go", "go"}))
	fmt.Println(piscine.WordCount([]string{"x"}))
	fmt.Println(piscine.WordCount([]string{}))
	fmt.Println(piscine.WordCount(nil))
}
```

And its output:

```console
$ go run . | cat -e
map[a:2 b:1]$
map[Go:1 go:2]$
map[x:1]$
map[]$
map[]$
$
```
