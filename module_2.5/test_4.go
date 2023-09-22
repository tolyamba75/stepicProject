package main

import "fmt"

func main() {
	var str string

	fmt.Scan(&str)
	result := ""
	for i, lit := range str {
		if i%2 == 1 {
			result += string(lit)
		}
	}

	fmt.Print(result)
}
