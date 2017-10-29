package main

import "fmt"

type V struct {
  X, Y int
}

func main() {

  fmt.Println("hey")

  x := V{0,1}
  fmt.Println(x)

  y := V{Y:2}
  fmt.Println(y)

}
