/*
Let us assume the following formula for displacement s as a function of time t,
 acceleration a, initial velocity vo, and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration,
initial velocity, and initial displacement. Then the program should prompt the
user to enter a value for time and the program should compute the displacement
after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes
three float64 arguments, acceleration a, initial velocity vo, and initial
displacement so.
GenDisplaceFn() should return a function which computes
displacement as a function of time, assuming the given values acceleration,
initial velocity, and initial displacement. The function returned by GenDisplaceFn()
 should take one float64 argument t, representing time, and return one float64
 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration,
 initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use
 the following statement to call GenDisplaceFn() to generate a function fn which
 will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

*/

package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
import "math"
//import "log"

//stores acceleration, initial velocity, initial displacement
var slice []float64 = make([]float64, 3)


//generates func based on the values stored in slice
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
  fn := func(time float64) float64 {
    calc := 0.5 * a * math.Pow(time, 2) + vo*time + so
    return calc
  }
  return fn
}

func GetValue() float64 {
  inputReader := bufio.NewReader(os.Stdin)
  value, _ := inputReader.ReadString('\n')
  value = strings.TrimSuffix(value, "\n")
  value_float, _ := strconv.ParseFloat(value, 64)
  return value_float
}

func main() {
  inputReader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter values for acceleration, initial velocity, and initial displacement. Press enter after each value!!")
  //asks for 3 float inputs, one per line
  for x:=0; x<3; x++ {
    input, _ := inputReader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")
    input_float, _ := strconv.ParseFloat(input, 64)
    slice[x] = input_float
  }

  fmt.Println("Your input: ", slice)


  time := GetValue()

  funct := GenDisplaceFn(slice[0], slice[1], slice[2])
  fmt.Println(funct(time))

  }
