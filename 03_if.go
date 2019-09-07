package main

import (
    "fmt"
)

var (
    example = []string {
        "golang",
        "hands-on",
        "in",
        "kagawa",
    }
)

func main() {
    var data []string
    data = example
    for _, v := range data {
        fmt.Println(v)
        if v == "in" {
            fmt.Println("OK")
        } else {
            fmt.Println("NG")
        }
    }
}
