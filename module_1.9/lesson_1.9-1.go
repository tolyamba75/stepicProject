package main

import "fmt"

func main() {

	/*  На ввод подается целое число.
	Если число положительное - вывести сообщение "Число положительное",
	если число отрицательное - "Число отрицательное".
	Если подается ноль - вывести сообщение "Ноль". */

	var num int

	fmt.Scan(&num)
	fmt.Println(num)

	switch {
	case num < 0:
		fmt.Print("Число отрицательное")
	case num > 0:
		fmt.Print("Число положительное")
	case num == 0:
		fmt.Print("Ноль")
	}
}
