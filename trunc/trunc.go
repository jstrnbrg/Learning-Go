package main

import "fmt"


func main() {
  fmt.Print("Enter text: ")
  var input float32
  fmt.Scanln(&input)
  fmt.Print(int(input), "\n")

}
