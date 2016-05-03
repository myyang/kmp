# kmp
Simple KMP implement

Example:

```go
package main

import (
    "github.com/myyang/kmp"
    "fmt"
)

func main() {
    s := kmp.KMP("ABCC", "ABCCCC")
    fmt.Printf("%v\n", s)
}
```

Output:

```
[0]
```
