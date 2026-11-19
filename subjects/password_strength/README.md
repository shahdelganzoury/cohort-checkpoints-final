# Level 3: PasswordStrength

**Difficulty:** Medium

## Instructions

Score one point each for length >= 8, presence of uppercase ASCII A-Z, digit 0-9, and a special character from !@#$%. Return 0..4.

## Expected function

```go
func PasswordStrength(s string) int {
    // Implement your solution.
}
```

## Example

```text
PasswordStrength("Go2026!") -> 3
```

## Rules

Use Go standard library only. The evaluator calls the function directly.
Do not print the result instead of returning it. Keep the function name and signature unchanged.
