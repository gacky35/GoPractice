package main

import (
    "fmt"
    "time"
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
    var ccnt, wcnt int
    data = example
    ln := len(example)
    for i, v := range data {
        fmt.Printf("[%d/%d] %s\n", i+1, ln, v)
        fmt.Print("input : ")
        var ans string
        stm := time.Now()
        fmt.Scanln(&ans)
        ptm := time.Since(stm)
        fmt.Printf("%.3f(sec)\n", ptm.Seconds())
        if v == ans {
            fmt.Println("OK")
            ccnt++
        } else {
            fmt.Println("NG")
            wcnt++
        }
    }
    fmt.Printf("[Errata] Correct:%d, Wrong:%d\n", ccnt, wcnt)
}
