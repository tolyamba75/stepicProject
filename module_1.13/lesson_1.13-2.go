package main

import "fmt"

func main() {

	/* Дано трехзначное число. Переверните его, а затем выведите. */

	var number, renumber int

	fmt.Scan(&number)
	renumber = 0

	for number%10 > 0 {
		renumber = renumber*10 + (number % 10)
		number /= 10
	}
	fmt.Println(renumber)
}
