# Level 4: BusCapacity

**Difficulty:** Medium

## Instructions

Each stop is {boarding, leaving}. Passengers leave before boarding. Starting empty, return maximum occupancy, or -1 for invalid negative counts, impossible leaving, or capacity exceeded. Capacity >= 0.

## Expected function

```go
func BusCapacity(capacity int, stops [][2]int) int {
    // Implement your solution.
}
```

## Example

```text
BusCapacity(20, [][2]int{{10,0},{8,3},{2,5},{4,1}}) -> 15
```

## Rules

Use Go standard library only. The evaluator calls the function directly.
Do not print the result instead of returning it. Keep the function name and signature unchanged.
