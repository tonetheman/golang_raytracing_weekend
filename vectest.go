package main

import "fmt"
import "math"

type float float64

type vec struct {
  e [3]float64
}

func (v *vec) length() float64 {
  return math.Sqrt(v.squared_length())
}

func (v* vec) squared_length() float64 {
  return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v* vec) make_unit_vector() {
  k := 1.0 / v.length()
  v.e[0] *= k
  v.e[1] *= k
  v.e[2] *= k
}

func add(v1 vec, v2 vec) vec {
  return vec{[3]float64 {
    v1.e[0] + v2.e[0],v1.e[1] + v2.e[1],
    v1.e[2] + v2.e[2]}}
}

func sub(v1 vec, v2 vec) vec {
  return vec{[3]float64 {
    v1.e[0] - v2.e[0],v1.e[1] - v2.e[1],
    v1.e[2] - v2.e[2]}}
}

func mul(v1, v2 vec) vec {
  return vec{[3]float64 {
    v1.e[0] * v2.e[0],v1.e[1] * v2.e[1],
    v1.e[2] * v2.e[2]}}
}

func div(v1, v2 vec) vec {
  return vec{[3]float64 {
    v1.e[0] / v2.e[0],v1.e[1] / v2.e[1],
    v1.e[2] / v2.e[2]}}
}

func const_mul(t float64, v vec) vec {
    return vec{[3]float64{t*v.e[0],t*v.e[1],t*v.e[2]}}
}

func main() {

  v := vec{[3]float64{1.0,2.0,3.0}}

  fmt.Println(v, v.length())
  fmt.Println("calling make unit vector")
  v.make_unit_vector()
  fmt.Println(v)

}
