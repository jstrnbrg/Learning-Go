package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Name struct {
	fname, lname string
}

func truncate(s string, maxLength int) string {
	if len(s) > maxLength {
		s = s[:maxLength]
	}
	return s
}



func main() {
	var filename string
	var names []Name

	fmt.Print("Enter a filename: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error scanning file: " + err.Error())
			break
		}

		text := scanner.Text()
		tokens := strings.SplitN(text, " ", 2)

		if len(tokens) == 2 {
			fname := truncate(tokens[0], 20)
			lname := truncate(tokens[1], 20)
			names = append(names, Name{fname, lname})
		}
	}

	format := "%-20s  %-20s\n"
	fmt.Println()
	fmt.Printf(format, "First Name", "Last Name")
	fmt.Println()
	for _, n := range names {
		fmt.Printf(format, n.fname, n.lname)
	}
}
