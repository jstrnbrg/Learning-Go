/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
 Each line of the text file has a first name and a last name, in that order,
 separated by a single space on the line.

Your program will define a name struct which has two fields,
 fname for the first name, and lname for the last name.
  Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and
create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice,
 and after all lines have been read from the file,
 your program will have a slice containing one struct for each line in the file.
 After reading all lines from the file, your program should iterate through your
 slice of structs and print the first and last names found in each struct.
*/
package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "log"

func main() {


  type name struct  {
    fname string
    lname string
  }

  slice := make([]name, 0, 3)

  //ask for file name
  fmt.Print("Enter filename: ")
  inputReader := bufio.NewReader(os.Stdin)
  filename, _ := inputReader.ReadString('\n')
  filename = strings.TrimSuffix(filename, "\n")

  //open file or print error
  file, err := os.Open(filename) // For read access.
  if err != nil {
	   fmt.Println(err)
   }

  //create scanner object
  scanner := bufio.NewScanner(file)

  //loop through lines, split lines into words, create new instance of name struct
  //and add the words at index 0 1nd 1 to fname and lname
  //then append the slice with that new instance
  for scanner.Scan() {
      line := scanner.Text()
      splitted := strings.Split(line, " ")
      n1 := name{splitted[0], splitted[1]}
      slice = append(slice, n1)
    }
    if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  //loop through slice, ignore counter, print fname and lname of each name instance
  for _, name := range slice {
     fmt.Printf("Name: %s Age: %s\n", name.fname, name.lname)
  }
}
