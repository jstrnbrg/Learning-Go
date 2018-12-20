package main

import "fmt"

func main() {
  x := [...]int {4, 8, 5}
  y := x[0:2] //[4,8]
  z := x[1:3] //[8,5]
  y[0] = 1    //y=[1,8], x=[1,8,5]
  z[1] = 3    //z=[8,3], x=[1,8,3]
  fmt.Print(x) //[1,8,3]
}
