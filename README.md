# utils
Utilities

## Usage

```go
package main

import (
	"fmt"
	"github.com/yzinkovets/utils/env"
)

func main() {
    fmt.Println(env.Must("HOME"))
}
```