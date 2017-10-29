package main

import "fmt"
import  "reflect"

func main() {

  a := []bool{true,true}
  fmt.Println(a,&a)

  fmt.Println(reflect.TypeOf(a))
  fmt.Println("len and cap", len(a), cap(a))

  junk := []int{1,2,3,4,5}
  fmt.Println(junk,len(junk),cap(junk))

  junk2 := junk[1:2]
  fmt.Println("junk2", junk2, len(junk2), cap(junk2))
}
