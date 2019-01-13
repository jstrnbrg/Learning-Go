/*
Define an interface type called Animal which describes the methods of an animal.
 Specifically, the Animal interface should contain the methods Eat(), Move(),
 and Speak(), which take no arguments and return no values.
 The Eat() method should print the animal’s food, the Move() method should print the animal’s
 locomotion, and the Speak() method should print the animal’s spoken sound.
 Define three types Cow, Bird, and Snake. For each of these three types,
 define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and
 Snake all satisfy the Animal interface. When the user creates an animal,
 create an object of the appropriate type. Your program should call the
 appropriate method when the user issues a query command.
*/



package main

import "fmt"
import "bufio"
import "os"
import "strings"

var slice []Animal

type Animal interface {
  Name() string
  Eat()
  Move()
  Speak()
}

type Cow struct {
  name string
  food string
  locomotion string
  sound string
}

type Bird struct {
  name string
  food string
  locomotion string
  sound string
}

type Snake struct {
  name string
  food string
  locomotion string
  sound string
}

func (x Cow) Name() string {
  return x.name
}

func (x Bird) Name() string {
  return x.name
}

func (x Snake) Name() string {
  return x.name
}

func (x Cow) Eat() {
  fmt.Println(x.food)
}

func (x Bird) Eat() {
  fmt.Println(x.food)
}

func (x Snake) Eat() {
  fmt.Println(x.food)
}

func (x Cow) Move() {
  fmt.Println(x.locomotion)
}

func (x Bird) Move() {
  fmt.Println(x.locomotion)
}

func (x Snake) Move() {
  fmt.Println(x.locomotion)
}

func (x Cow) Speak() {
  fmt.Println(x.sound)
}

func (x Bird) Speak() {
  fmt.Println(x.sound)
}

func (x Snake) Speak() {
  fmt.Println(x.sound)
}

func CreateAnimal(name string, type1 string) Animal {
  var a1 Animal
  switch type1 {
  case "cow":
    cow := Cow{name, "grass", "walk", "moo"}
    a1 = cow
    return a1
  case "bird":
    bird := Bird{name, "worms", "fly", "peep"}
    a1 = bird
    return a1
  case "snake":
    snake := Snake{name, "mice", "slither", "hss"}
    a1 = snake
    return a1
  default:
      return nil
    }
}

func PerformAction(name string, info string) {
  for _ , v  := range slice {
    if (name == v.Name()) {
      switch info {
      case "eat":
        v.Eat()
        continue
      case "move":
        v.Move()
        continue
      case "speak":
        v.Speak()
        continue
      default:
          fmt.Println("No such action defined. Use eat, speak or move.")
        }
    }


  }
}


func main() {
  inputReader := bufio.NewReader(os.Stdin)

  for {
    fmt.Printf(">")
    input, _ := inputReader.ReadString('\n')
    input = strings.ToLower(strings.TrimSuffix(input, "\n"))
    input_words := strings.Split(input, " ")

    if (input_words[0] == "newanimal") {

      animal := CreateAnimal(input_words[1], input_words[2])
      fmt.Println(animal)
      if animal != nil {
        fmt.Println("Created it!")
        slice = append(slice, animal)
        fmt.Println(slice)

      }


    } else if (input_words[0] == "query") {
        PerformAction(input_words[1], input_words[2])


    } else {
      fmt.Println("First word needs to be either 'newanimal' or 'query'")
      continue
    }




}
}
