package main

import (
	"fmt"
	//"bufio"
	//"os"
	//"strings"
	// "log"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func (t float64) float64 {
		return 0.5 * a * math.Pow(t, 2) + v0 * t + s0
	}
}

func GetFloat64(msg string) float64 {
	var v float64
	for {
		fmt.Printf(msg);
		if _, err := fmt.Scan(&v); err == nil {
			return v
		}
	}
}

func main() {
	a := GetFloat64("Enter a: ")
	v0 := GetFloat64("Enter v0: ")
	s0 := GetFloat64("Enter s0: ")

	fn := GenDisplaceFn(a, v0, s0)

	t := GetFloat64("Enter t: ")

	fmt.Println(fn(t))
}
