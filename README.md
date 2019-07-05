# zdecorator
###### 装饰器

## 获得zdecorator
`go get -u github.com/zlyuancn/zdecorator`

## 使用zdecorator

```go
package main

import (
    "fmt"
    "github.com/zlyuancn/zdecorator"
)

func main() {
    t2 := func(fn *zdecorator.SrcFunction) zdecorator.Function {
        return func() {
            fmt.Println("t2")
            fn.Call()
        }
    }
    t1 := func(fn *zdecorator.SrcFunction) zdecorator.Function {
        return func() {
            fmt.Println("t1")
        }
    }
    t := func() {
        fmt.Println("t")
    }

    d := zdecorator.New(t)
    d.AddDecorator(t1, t2)
    d.Call()
}
```
