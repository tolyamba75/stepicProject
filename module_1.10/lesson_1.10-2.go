package main

import "fmt"

func main() {
	var a, b, sum int
	sum = 0

	fmt.Scan(&a, &b)

	for ; a <= b; a++ {
		sum += a
	}

	fmt.Print(sum)
}
