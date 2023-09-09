package main

import "fmt"

func main() {
	var x, p, y, year int

	fmt.Scan(&x) // начальный вклад в копейках
	fmt.Scan(&p) // процентная ставка
	fmt.Scan(&y) // конечная сумма в копейках

	x *= 100
	y *= 100
	for x <= y {
		year++
		x += x / 100 * p
	}

	fmt.Println(year)

}
