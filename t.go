package main

import "fmt"

func main() {

    i,j := 42, 2701
    p := &i

    fmt.Println("i and j",i,j)
    fmt.Println(*p)

    fmt.Println("the pointer is", p)
    p = &j
    fmt.Println(*p)
    fmt.Println("the pointer is",p)
}
