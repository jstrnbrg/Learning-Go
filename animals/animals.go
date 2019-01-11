package main

import "fmt"
import "bufio"
import "os"
import "strings"
//import "strconv"
//import "math"
//import "log"


type Animal struct {
  food string
  locomotion string
  noise string
}

func (x Animal) Eat() {
  fmt.Println(x.food)
}

func (x Animal) Move() {
  fmt.Println(x.locomotion)
}

func (x Animal) Speak() {
  fmt.Println(x.noise)
}


//based on input[0] this func returns a pointer to a newly created animal object
func CreateAnimal(input string) *Animal {
  switch input {
  case "cow":
    cow := &Animal{"grass", "walk", "moo"}
    return cow
  case "bird":
    bird := &Animal{"worms", "fly", "peep"}
    return bird
  case "snake":
    snake := &Animal{"mice", "slither", "hss"}
    return snake
  default:
      return nil
    }
}

func PerformAction(input string, animal Animal) {
  switch input {
  case "eat":
    animal.Eat()
  case "move":
    animal.Move()
  case "speak":
    animal.Speak()
  }
}


func main() {
  inputReader := bufio.NewReader(os.Stdin)

  for {
    fmt.Printf(">")
    input, _ := inputReader.ReadString('\n')
    input = strings.ToLower(strings.TrimSuffix(input, "\n"))
    input_words := strings.Split(input, " ")


    animal := CreateAnimal(input_words[0])
    PerformAction(input_words[1], *animal)



    }
  }
