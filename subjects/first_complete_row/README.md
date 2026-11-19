# Level 10: FirstCompleteRow

**Difficulty:** Medium

## Instructions

Given a nonempty rectangular grid of distinct integers and a sequence of distinct called integers, return the 1-based call index when the first row becomes fully called, or -1. Called values may be absent from the grid.

## Expected function

```go
func FirstCompleteRow(grid [][]int, called []int) int {
    // Implement your solution.
}
```

## Example

```text
FirstCompleteRow([][]int{{4,1,7},{3,9,2},{8,5,6}}, []int{9,4,1,3,7,2}) -> 5
```

## Rules

Use Go standard library only. The evaluator calls the function directly.
Do not print the result instead of returning it. Keep the function name and signature unchanged.
