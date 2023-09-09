package main

import "fmt"

func main() {
	var x, p, y, year int

	fmt.Scan(&x * 100) // начальный вклад, с копейками
	fmt.Scan(&p)       // процентная ставка
	fmt.Scan(&y * 100) // конечная сумма, с копейками

	for x <= y {
		year += 1
		x += x / 100 * p
	}

	fmt.Println(year)

}
