package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which prompts the user to first enter a name, and then enter an
address. Your program should create a map and add the name and address to the
map using the keys â€œnameâ€ and â€œaddressâ€, respectively.
Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.
*/

func main() {
	// Make reader object to get user input
	reader := bufio.NewReader(os.Stdin)
	m := make(map[string]string, 2)
	var s string
	// Read user input
	fmt.Printf("Enter yout name: ")
	s, _ = reader.ReadString('\n')
	m["name"] = strings.TrimSuffix(s, "\n")
	fmt.Printf("Enter you address: ")
	s, _ = reader.ReadString('\n')
	m["address"] = strings.TrimSuffix(s, "\n")

	// Try to Marshal object
	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error marshalling the json object")
		os.Exit(1)
	}
	fmt.Println(string(j))
}
