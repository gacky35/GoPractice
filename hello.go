package main

import (
    "fmt"
)

var (
    msg1 = "Hello"
    msg2 string
)

func main() {
    msg2 = "world"
    fmt.Println(msg1, msg2)
}
