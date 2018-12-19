package main

import (
	"fmt"
	"log"
)

func main() {
	var fpNumber float64

	fmt.Printf("enter a floating point number: ")
	_, err := fmt.Scanln(&fpNumber)
	if err != nil {
		log.Fatalf("can't read floating point number, reason: %s", err)
	}

	fmt.Printf("%d", int64(fpNumber))
}
