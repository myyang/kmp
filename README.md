# kmp
Simple KMP implement

Example:

```
package main

import (
    "github.com/myyang/kmp"
)

func main() {
    s := kmp.KMP("ABCC", "ABCCCC")
    println(s)
}
```

Output:

```
[0]
```
