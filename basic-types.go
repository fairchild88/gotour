package main

import (
    "fmt"
    "math/cmplx"
)

var (
    //bool
    ToBe bool = false

    //int
    One int = 1

    //int8-64
    Two int8 = 2
    Three int16 = 3
    Four uint32 = 1 << 32 - 1
    MaxInt uint64 = 1 << 64 - 1

    //uint8-64
    UOne uint = 4
    UTwo uint8 = 5
    UThree uint16 = 5
    UFour uint32 = 6
    UFive uint64 = 7

    //str
    str string = "str"

    //alias
    b byte = 1 //alias for uint8
    u rune = 2 //alias for int32

    //float
    f1 float32 = 1.1
    f2 float64 = 2.2 

    //complex
    y complex64 = -12i
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Printf("#bool\nType %T Value %v\n", ToBe, ToBe)

    fmt.Printf("\n#int\nType %T Value %v\n", One, One)

    fmt.Printf("\n#int8-64\nType %T Value %v\n", Two, Two)
    fmt.Printf("Type %T Value %v\n", Three, Three)
    fmt.Printf("Type %T Value %v\n", Four, Four)
    fmt.Printf("Type %T Value %v\n", MaxInt, MaxInt)

    fmt.Printf("\n#uint8-64\nType %T Value %v\n", UOne, UOne)
    fmt.Printf("Type %T Value %v\n", UTwo, UTwo)
    fmt.Printf("Type %T Value %v\n", UThree, UThree)
    fmt.Printf("Type %T Value %v\n", UFour, UFour)
    fmt.Printf("Type %T Value %v\n", UFive, UFive)

    fmt.Printf("\n#alias\nType %T (byte) Value %v\n", b, b)
    fmt.Printf("Type %T (rune) Value %v\n", u, u)

    fmt.Printf("\n#str\nType %T Value %v\n", str, str)

    fmt.Printf("\n#float32-64\nType %T Value %v\n", f1, f1)
    fmt.Printf("Type %T Value %v\n", f2, f2)

    fmt.Printf("\n#complex64-128\nType %T Value %v\n", y, y)
    fmt.Printf("Type %T Value %v\n", z, z)

    //zero value
    var i int //0
    var f float64//
    var b bool//false
    var s string//""
    var s1 string = "12" //""
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
    fmt.Printf("%v %v %v %v\n", i, f, b, s)
    fmt.Printf("%v %v %v %q\n", i, f, b, s1)
}
