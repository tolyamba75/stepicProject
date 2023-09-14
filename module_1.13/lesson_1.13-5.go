package main

import "fmt"

func main() {
	/* Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами. */

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b > c && a+c > b && b+c > a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
