package main

import (
    "fmt"
    "math/rand"
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

func init() {
    rand.Seed(time.Now().UnixNano())
}

func shuffle(src []string) []string {
    idxs := rand.Perm(len(src))
    // dst := make([]string, 0)
    // for _, v := range idxs {
    //     dst = append(dst, src[v])
    // }
    dst := make([]string, len(src))
    for i, v := range idxs {
        dst[i] = src[v]
    }
    return dst
}

func main() {
    var data []string
    var ccnt, wcnt int
    data = shuffle(example)
    ln := len(example)
    for i := 3; i >= 1; i-- {
        fmt.Printf("%d\n", i)
        time.Sleep(time.Second)
    }
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
