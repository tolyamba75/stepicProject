package main

import "fmt"

func main() {
	var number1, number2 string
	fmt.Scan(&number1, &number2)
	for i := 0; i < len(number1); i++ {
		for k := 0; k < len(number2); k++ {
			if number1[i] == number2[k] {
				fmt.Print(string(number1[i]), " ")
				break
			}
		}
	}
}
