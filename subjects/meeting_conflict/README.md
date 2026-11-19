# Level 9: MeetingConflict

**Difficulty:** Medium

## Instructions

Meetings are valid, unsorted half-open intervals [start,end), 0 <= start < end <= 1440. Return true if any two overlap; touching endpoints are not conflicts.

## Expected function

```go
func MeetingConflict(meetings [][2]int) bool {
    // Implement your solution.
}
```

## Example

```text
MeetingConflict([][2]int{{540,600},{570,630}}) -> true
```

## Rules

Use Go standard library only. The evaluator calls the function directly.
Do not print the result instead of returning it. Keep the function name and signature unchanged.
