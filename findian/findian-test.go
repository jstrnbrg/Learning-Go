package main
import "fmt"

func main(){
	var input_str string
	fmt.Scanf("%s", &input_str)

	for _, alpha := range(input_str) {
		if alpha == 'i' || alpha == 'a' || alpha == 'n' {
			fmt.Println("Found!")
			return
		}
	}

	fmt.Println("Not Found!")
}
