package main

import "fmt"

func main() {
	var count, number, sum int
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		fmt.Scan(&number)
		if number%8 == 0 && (number/10 != 0) && (number/100 == 0) {
			sum += number
		}
	}
	fmt.Print(sum)
}
