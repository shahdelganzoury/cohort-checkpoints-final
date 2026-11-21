## meetingconflict

### Instructions

Write a function that takes the meetings as a slice of `[2]int` and returns `true` when any two meetings overlap.

- Each meeting is a valid half-open interval `[start, end)`, with `0 <= start < end <= 1440`.
- The meetings are not necessarily sorted.
- Intervals that only touch at an endpoint are not a conflict.
- Two intervals `[start1, end1)` and `[start2, end2)` overlap when `start1 < end2` and `start2 < end1`.
- An empty slice, a `nil` slice, or a single meeting returns `false`.
- A meeting contained inside another meeting is a conflict.
- Compare every pair of different meetings. A meeting does not conflict with itself.

### Expected function

```go
func MeetingConflict(meetings [][2]int) bool {

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
	fmt.Println(piscine.MeetingConflict([][2]int{{540, 600}, {570, 630}}))
	fmt.Println(piscine.MeetingConflict(nil))
	fmt.Println(piscine.MeetingConflict([][2]int{{540, 600}, {600, 660}}))
	fmt.Println(piscine.MeetingConflict([][2]int{{0, 1440}, {1, 2}}))
}
```

And its output:

```console
$ go run . | cat -e
true$
false$
false$
true$
$
```
