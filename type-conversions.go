package main
import (
    "fmt"
    "math"
)

func main() {
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)
    fmt.Printf("%v\n", u)

    j := 42
    f1 := float64(j)
    u1 := uint(f1)

    //explicit conversion
    fmt.Printf("%v\n", math.Sqrt(float64(u1)))
}
