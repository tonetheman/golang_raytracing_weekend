
package main

import "fmt"

type float float64

func main() {

  const nx int =  200
  const ny int = 100

  fmt.Println("P3")
  fmt.Println(nx,ny,255)

  for j:=ny-1;j>=0; j-- {
    for i := 0;i<nx;i++ {
      r := float(i)/float(nx)
      g := float(j)/float(ny)
      b := 0.2

      ir := int(255.9 * r)
      ig := int(255.9 * g)
      ib := int(255.9 * b)

      fmt.Println(ir,ig,ib)


    }
  }

}
