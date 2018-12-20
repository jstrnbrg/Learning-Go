package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	count, str, sli := 0, "", make([]int, 3)
	for {
		fmt.Print("Add an integer to the slice: ")
		fmt.Scan(&str)
		if str == "X" {
			break
		}
		num, _ := strconv.Atoi(str)
		if count < 3 {
			sli[0] = num
			count++
		} else {
			sli = append(sli, num)
		}
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
