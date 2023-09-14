package main

import "fmt"

func main() {

	/* Дано трехзначное число. Найдите сумму его цифр. */

	var number, sum int

	fmt.Scan(&number)
	sum = 0

	for number%10 > 0 {
		sum += number % 10
		number /= 10
	}
	fmt.Println(sum)
}
