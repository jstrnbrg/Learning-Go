/*Write a program which prompts the user to first enter a name, and then enter an address.
 Your program should create a map and add the name and address to the map using
  the keys “name” and “address”, respectively. Your program should use Marshal()
  to create a JSON object from the map, and then your program should
  print the JSON object.
*/


package main

import "fmt"
import "bufio"
import "os"
import "encoding/json"
import "strings"
//

func main() {

  infoMap := make(map[string]string)
  fmt.Print("Enter name: ")
  inputReader := bufio.NewReader(os.Stdin)
  name, _ := inputReader.ReadString('\n')
  fmt.Println(name)
  if name != "" {
    infoMap["name"] = strings.TrimSuffix(name, "\n")
  }

  fmt.Print("Enter address: ")
  address, _ := inputReader.ReadString('\n')
  if address != "" {
    infoMap["address"] = strings.TrimSuffix(address, "\n")
  }

  barr, err := json.Marshal(infoMap)
  if err != nil {
		fmt.Println("Error marshalling the json object")
		os.Exit(1)
	}
  fmt.Println(string(barr))



}
