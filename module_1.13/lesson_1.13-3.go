package main

import "fmt"

func main() {

	/* Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.*/

	var hours, minutes, k int
	fmt.Scan(&k)

	hours = k / 60 / 60
	minutes = k / 60 % 60

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
