package main

import "fmt"

func main() {
    var v int
    fmt.Scanf("%X", &v)
    fmt.Printf("%[1]d\n%[1]c\n", v)
}