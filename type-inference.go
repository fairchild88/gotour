package main

import "fmt"

//the variables's type is infered from the value on the right hand side
func main() {
    v := 42
    fmt.Printf("%T type %v\n", v, v)
    v1 := 42.1
    fmt.Printf("%T type %v\n", v1, v1)
    v2 := 42 + 2i
    fmt.Printf("%T type %v\n", v2, v2)
}
