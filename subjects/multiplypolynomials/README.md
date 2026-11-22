## multiplypolynomials

### Instructions

A polynomial is a linked list of terms sorted by `Exp` in **descending** order, with no duplicate exponents and no zero coefficients. Given two polynomials `a` and `b`, return their product as a **new** list in the same form.

Terms whose coefficients cancel to `0` **must not appear**. If either input is `nil`, the product is zero: return `nil`. Do not modify the input lists.

Example 1:

```
Input: a = [(2,2),(3,0)], b = [(1,1),(1,0)]
Output: [(2,3),(2,2),(3,1),(3,0)]
Explanation: (2x^2 + 3)(x + 1) = 2x^3 + 2x^2 + 3x + 3. Each term is written as (Coef,Exp).
```

Example 2:

```
Input: a = [(1,1),(1,0)], b = [(1,1),(-1,0)]
Output: [(1,2),(-1,0)]
Explanation: (x + 1)(x - 1): the x terms cancel, so they are left out.
```

Example 3:

```
Input: a = nil, b = [(5,3)]
Output: nil
Explanation: Anything times the zero polynomial is zero.
```

Constraints:

* Each polynomial has at most `200` terms.
* `0 <= Exp <= 5000`, `1 <= |Coef| <= 100`
* Follow up: `O(|a| * |b| + E)` time, where `E` is the largest result exponent.

### Allowed structure

You are allowed to use the following structure. Declare it in your solution file and use it as is:

```go
type ListNode struct {
	Coef int
	Exp  int
	Next *ListNode
}
```

### Expected function

```go
func MultiplyPoly(a, b *ListNode) *ListNode {

}
```

### Usage

Here is a possible program to test your function (it assumes the structure above is declared in your solution file):

```go
package main

import (
	"fmt"
)

func main() {
	a := buildPoly(2, 2, 3, 0)
	b := buildPoly(1, 1, 1, 0)
	fmt.Println(show(MultiplyPoly(a, b)))
	fmt.Println(show(a), show(b))
	fmt.Println(show(MultiplyPoly(buildPoly(1, 1, 1, 0), buildPoly(1, 1, -1, 0))))
	fmt.Println(show(MultiplyPoly(nil, buildPoly(5, 3))))
	fmt.Println(show(MultiplyPoly(buildPoly(3, 5000), buildPoly(-2, 5000, 4, 0))))
	fmt.Println(show(MultiplyPoly(buildPoly(7, 0), buildPoly(-1, 0))))
	fmt.Println(show(MultiplyPoly(buildPoly(1, 2, 1, 1), buildPoly(1, 1, -1, 0))))
}

// buildPoly takes coef, exp pairs in descending exponent order.
func buildPoly(pairs ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for i := 0; i+1 < len(pairs); i += 2 {
		tail.Next = &ListNode{Coef: pairs[i], Exp: pairs[i+1]}
		tail = tail.Next
	}
	return dummy.Next
}

func show(p *ListNode) string {
	if p == nil {
		return "nil"
	}
	s := "["
	for ; p != nil; p = p.Next {
		s += fmt.Sprintf("(%d,%d)", p.Coef, p.Exp)
	}
	return s + "]"
}
```

And its output:

```console
$ go run . | cat -e
[(2,3)(2,2)(3,1)(3,0)]$
[(2,2)(3,0)] [(1,1)(1,0)]$
[(1,2)(-1,0)]$
nil$
[(-6,10000)(12,5000)]$
[(-7,0)]$
[(1,3)(-1,1)]$
```
