```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m[0]) //This will cause a runtime panic
}
```