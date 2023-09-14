package main

import "fmt"

func main() {
	/* Даны два числа. Найти их среднее арифметическое. */

	var number1, number2 float32
	fmt.Scan(&number1, &number2)

	fmt.Print((number1 + number2) / 2)

}
