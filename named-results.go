package main

import "fmt"

//named return values
func split(sum int) (x, y, z int) {
    x = sum * 4 / 9
    y = sum - x
    z = 2
    return
}

func main() {
    fmt.Println(split(17))
}
