```go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int)
    m["0"] = 10
    fmt.Println(m["0"]) //Correct way to access map element 
}
```