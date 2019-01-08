package main
import (
  "fmt"
)

func Swap(slice []int, p int) {
  var temp = slice[p]
  slice[p] = slice[p + 1]
  slice[p + 1] = temp
}

func BubbleSort(slice []int) {
  for i := len(slice) - 1; i > 0; i-- {
    for j := 1; j <=i; j++ {
      if slice[j] < slice[j - 1] {
        Swap(slice, j - 1)
      }
    }
  }
}

func main() {
  // your code goes here

  fmt.Println("Enter 10 numbers to sort")
  slice := []int{}
  num := 0

  for i := 0; i < 10; i++ {
    fmt.Scanf("%d", &num)
    slice = append(slice, num)
  }

  BubbleSort(slice)

  fmt.Printf("Sorted numbers are: %v", slice)
 }
