/*
Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’,
 ends with the character ‘n’, and contains the character ‘a’.
 The program should print “Not Found!” otherwise.
 The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/


package main

import "fmt"
import "strings"
import "bufio"
import "os"
//

func main() {
  fmt.Print("Enter text: ")
  inputReader := bufio.NewReader(os.Stdin)
  input, _ := inputReader.ReadString('\n')
  input =  strings.ToLower(input)
  response := strings.Join(strings.Fields(input),"")
  response = strings.TrimSpace(response)
  fmt.Printf("Stripped input: %q\n", response)

  if strings.Contains(response, "a") && strings.HasPrefix(response, "i") && strings.HasSuffix(response, "n") {
    fmt.Println("Found!")
  } else {
      fmt.Println("Not Found!")
  }
}
