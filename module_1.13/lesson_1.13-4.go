package main

import (
	"fmt"
	"math"
)

func main() {

	/* Заданы три числа - a,b,c. Нужно проверить, является ли треугольник прямоугольным.*/

	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if math.Pow(c, 2) == math.Pow(b, 2)+math.Pow(a, 2) {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

// math.Pow(a, 2) == math.Pow(b, 2)+math.Pow(c, 2) || math.Pow(b, 2) == math.Pow(a, 2)+math.Pow(c, 2) ||
