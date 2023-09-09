package main

import (
	"fmt"
)

func main() {

	// Требуется определить, является ли данный год високосным.

	var year int

	fmt.Scan(&year)

	if year%400 == 0 {
		fmt.Print("YES")
	} else if year%4 == 0 && year%100 != 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
