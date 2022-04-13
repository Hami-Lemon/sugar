# Sugar

一些golang的help func。

## 已提供的func

### Conditional Operator(三目运算符)

```go
func Condition[T any](con bool, v1, v2 T) T
```

当条件con为true时，返回v1，否则返回v2

```go
package example

import "github.com/Hami-Lemon/sugar"

func example() {
	sugar.Condition(1 < 2, 1, 2) // 1
}
```