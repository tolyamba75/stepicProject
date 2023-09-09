package main

import (
	"fmt"
)

func main() {

	/* Дано неотрицательное целое число.
	Найдите и выведите первую цифру числа. */

	var number string

	fmt.Scan(&number)
	fmt.Print(string(number[0]))

}
