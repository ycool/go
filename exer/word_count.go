package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    r := make(map[string]int)
    for _, w := range words {
        r[w]++
    }
    return r
}

func main() {
    wc.Test(WordCount)
}