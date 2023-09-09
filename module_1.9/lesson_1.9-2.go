package main

import (
	"fmt"
)

func main() {

	// По данному трехзначному числу определите, все ли его цифры различны.

	var number int

	fmt.Scan(&number)

	if (number%10) == (number%100/10) || (number%10) == (number/100) || (number/100) == (number%100/10) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
