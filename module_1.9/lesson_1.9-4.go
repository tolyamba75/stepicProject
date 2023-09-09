package main

import (
	"fmt"
)

func main() {

	/*
		Определите является ли билет счастливым.
		Счастливым считается билет,
		в шестизначном номере
		которого сумма первых трёх цифр совпадает с суммой трёх последних.
	*/

	var (
		number, number1, number2, number3, number4, number5, number6 int
	)

	fmt.Scan(&number)

	number1 = number / 100000
	number2 = number / 10000 % 10
	number3 = number / 1000 % 10
	number4 = number / 100 % 10
	number5 = number / 10 % 10
	number6 = number % 10

	if (number1 + number2 + number3) == (number4 + number5 + number6) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
