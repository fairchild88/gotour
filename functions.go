package main

import (
    "fmt"
)

// type come after the variable name
func add( x int, y int) int {
    return x + y
}

// two or more consecutive named function parameters share a type, you can omit the type from all but the last
func add2(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
    fmt.Println(add2(42, 13))
}
