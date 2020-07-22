# Soft (Short Format)

Help to format string using `{}` like `python` or `rust`. 

---

## Install

```shell
go get -u github.com/yam8511/soft
```

## Quick Start


```go
package main

import (
    "fmt"
    "github.com/yam8511/soft"
)

func main() {
    fmt.Println(
        soft.SFormat(
            "My name is {0}, {1} years old. Again, My Name is {0}",
            "Zuolar", 9999,
        ),
    )

    fmt.Println(
        soft.MFormat(
            "Hello {a}, my name is {b}", 
            map[string]interface{}{
                "a": "World", 
                "b": "Zuolar",
            },
        ),
    )
    fmt.Println(
        soft.TFormat("My name is {.Name}, {.Age} years old. Again, My Name is {.Name}", struct {
            Name string
            Age  int
        }{
            Name: "Zuolar",
            Age:  9999,
        }),
    )
}
```
