package main

import "fmt"

//default false
var c, python, java bool

//declaration with initializers
var j, k int = 1, 2

//!constants cannot be declared using the := syntax
const Pi = 3.14

//numeric constants are high-precision values
//untyped constant takes the type needed by its context
const (
    Big = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int {return x * 10 + 1}
func needFloat(x float64) float64 {return x * 0.1}
func main() {
    //default 0
    var i int
    fmt.Println(i, c, python, java)

    //with initializer present, type omitted
    var flutter, kotlin, csharp  = true, false, "no!"
    fmt.Println(j, k, flutter, kotlin, csharp)

    //short variables declarations
    //only inside function works
    m := 3

    dart, erlang :=  false, "yes!"
    fmt.Println(m, dart, erlang)

    const world = "世界"
    fmt.Println("Hello", world)
    fmt.Println("Happy", Pi, "day")

    const truth = true
    fmt.Println("Go rules?", truth)

    fmt.Printf("%T type %v", Small, Small)
    //overflows
    //fmt.Printf("%T type %v", Big, Big)
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
