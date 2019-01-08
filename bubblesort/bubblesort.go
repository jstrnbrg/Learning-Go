package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
//import "log"



func swap(input []int, i int) {

  var temp int
  if input[i] > input[i+1] {
    temp = input[i]
    input[i] = input[i+1]
    input[i+1] = temp
  }
  fmt.Println(input)

}


func BubbleSort(input []int) {
    //loop through new int slice
    until := len(input)
    for x := 0; x < len(input) -1; x++ {
      fmt.Println("Sort number: ", x +1)

      for index, _ := range input[:until - 1] {
          swap(input[:until], index)
      }
      until--
    }
}



func main() {

  slice := make([]string, 0)

  //ask for file name
  fmt.Print("Enter up to 10 numbers, seperated by comma: ")
  inputReader := bufio.NewReader(os.Stdin)
  string, _ := inputReader.ReadString('\n')
  read_line := strings.TrimSuffix(string, "\n")
  slice = strings.Split(read_line, ",")

  //fill int slice with values from string slice after converting them into ints
  fib := make([]int, len(slice))
  for i, s := range slice {
    fib[i], _ = strconv.Atoi(s)
  }

  BubbleSort(fib)
  fmt.Println("Sorted: ", fib)





}
